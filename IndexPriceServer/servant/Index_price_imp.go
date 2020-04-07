package servant

import (
	"CoinContract/IndexPriceServer/validator"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"sync"

	"CoinContract/IndexPriceServer/models/redis"
	"CoinContract/IndexPriceServer/utils"

	"github.com/go-xweb/log"
)

type IndexPriceObjImp struct {
}

const (
	// 现货指数数据来源
	bitstamp_url = "https://www.bitstamp.net/api/v2/ticker/"
	gdax_url     = "https://api.pro.coinbase.com/products/"
	gemini_url   = "https://api.gemini.com/v2/ticker/"
	kraken_url   = "https://api.kraken.com/0/public/Ticker?pair="
	// 备用地址
	bitstamp_spare_url = "https://api.cryptowat.ch/markets/bitstamp/"
	gdax_spare_url     = "https://api.cryptowat.ch/markets/coinbase-pro/"
	gemini_spare_url   = "https://api.cryptowat.ch/markets/gemini/"
	kraken_spare_url   = "https://api.cryptowat.ch/markets/kraken/"
	// 权重
	bitstamp_weight float64 = 0.25
	gdax_weight     float64 = 0.25
	gemini_weight   float64 = 0.25
	kraken_weight   float64 = 0.25
	// 价格
	price_pre = "index_price"
)

var (
	// 读写锁
	priceRW sync.RWMutex
)

// 计算指数价格
func (imp *IndexPriceObjImp) GetIndexPrice(Symbol string, Value *string) (ret string, err error) {
	defer Recover()
	// 查询redis是否有值
	price, err := redis.IndexPriceAll(price_pre)
	if err != nil {
		priceRW.RLock() // 定义读锁
		// 获取4个交易所的最新价格
		other_price := getOtherPrice(Symbol)
		fmt.Printf("4个交易所的最新价格：%v\n", other_price)
		// 价格进行加权平均，获取指数价格
		price = getWeightPrice(other_price)
		fmt.Printf("指数价格price：%f\n", price)
		// 价格存入redis
		if err = redis.IndexPriceAdd(price_pre, price); err != nil {
			*Value = validator.ErrResponse(validator.RedisErr)
			log.Infof("价格存入redis错误")
			priceRW.RUnlock()
			return ret, err
		}
		priceRW.RUnlock()
	}
	*Value = validator.Response(validator.Succeed, price)
	return ret, nil
}

// 价格加权平均
func getWeightPrice(other_price map[string]float64) float64 {
	var price float64 = 0.0
	// 获取中位数价格
	mid, _ := getMidPrice(other_price)
	fmt.Printf("中位数价格%f", mid)
	// 修正偏离幅度
	other_price = correctDeviate(other_price, mid)
	fmt.Printf("修正后价格：%v\n", other_price)
	// 加权
	for k, v := range other_price {
		if k == "bitstamp_price" {
			price += bitstamp_weight * v
		} else if k == "gdax_price" {
			price += gdax_weight * v
		} else if k == "kraken_price" {
			price += gemini_weight * v
		} else if k == "gemini_price" {
			price += kraken_weight * v
		}
	}
	return price
}

// 偏离幅度
func correctDeviate(other_price map[string]float64, mid float64) map[string]float64 {
	m := make(map[string]float64)
	for k, v := range other_price {
		if (v-mid)/mid > 0.1 || (v-mid)/mid < -0.1 {
			m[k] = mid
		} else {
			m[k] = v
		}
	}
	return m
}

// 获取中位数价格
func getMidPrice(other_price map[string]float64) (float64, []float64) {
	s := sort.Float64Slice{other_price["bitstamp_price"], other_price["gdax_price"], other_price["kraken_price"], other_price["gemini_price"]}
	sort.Float64s(s)
	mid := (s[1] + s[2]) / 2.0
	return mid, s
}

// 获取其他平台最新价
func getOtherPrice(symbol string) map[string]float64 {
	other_price := map[string]float64{
		"bitstamp_price": 0.0,
		"gdax_price":     0.0,
		"kraken_price":   0.0,
		"gemini_price":   0.0,
	}
	wg := sync.WaitGroup{}
	wg.Add(4)
	// bitstamp交易所最新价
	go func() {
		other_price["bitstamp_price"] = utils.StringToFloat64(get_bitstamp_price(symbol))
		wg.Done()
	}()
	// gdax交易所最新价
	go func() {
		other_price["gdax_price"] = utils.StringToFloat64(get_gdax_price(symbol))
		wg.Done()
	}()
	// gemini交易所最新价
	go func() {
		other_price["kraken_price"] = utils.StringToFloat64(get_gemini_price(symbol))
		wg.Done()
	}()
	// kraken交易所最新价
	go func() {
		other_price["gemini_price"] = get_kraken_price(symbol)
		wg.Done()
	}()
	wg.Wait()
	return other_price
}

// 获取bitstamp交易所最新价
func get_bitstamp_price(symbol string) string {
	var b Bitstamp
	url := fmt.Sprintf("%s%susd", bitstamp_url, strings.ToLower(symbol))
	body, err := get(url)
	if err != nil {
		log.Infof("get_bitstamp_price err:%s", err)
		return "0"
	}
	err = json.Unmarshal([]byte(body), &b)
	if err != nil {
		log.Infof("get_bitstamp_price err:%s", err)
		return "0"
	}
	return b.Last
}

// 获取gdax交易所最新价
func get_gdax_price(symbol string) string {
	var b Gdax
	url := fmt.Sprintf("%s%s-USD/ticker", gdax_url, symbol)
	body, err := get(url)
	if err != nil {
		log.Infof("get_gdax_price err:%s", err)
		return "0"
	}
	err = json.Unmarshal([]byte(body), &b)
	if err != nil {
		log.Infof("get_gdax_price err:%s", err)
		return "0"
	}
	return b.Price
}

// 获取gemini交易所最新价
func get_gemini_price(symbol string) string {
	var b Gemini
	url := fmt.Sprintf("%s%susd", gemini_url, strings.ToLower(symbol))
	body, err := get(url)
	if err != nil {
		log.Infof("get_gemini_price err:%s", err)
		return "0"
	}
	err = json.Unmarshal([]byte(body), &b)
	if err != nil {
		log.Infof("get_gemini_price err:%s", err)
		return "0"
	}
	return b.Bid
}

// 获取kraken交易所最新价
func get_kraken_price(symbol string) float64 {
	var b Spare
	url := fmt.Sprintf("%s%susd/price", kraken_spare_url, symbol)
	body, err := get(url)
	if err != nil {
		log.Infof("get_kraken_price err:%s", err)
		return 0.0
	}
	err = json.Unmarshal([]byte(body), &b)
	if err != nil {
		log.Infof("get_kraken_price err:%s", err)
		return 0.0
	}
	return b.Result.Price
}

// Get请求
func get(url string) (string, error) {
	client := &http.Client{
		// Timeout: time.Duration(6 * time.Second),
	}
	// 请求建立连接
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Infof("get %s err:%s", url, err)
		return "", err
	}
	// 提交
	response, err := client.Do(req)
	body, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		log.Infof("read body %s err:%s", url, err)
		return "", err1
	}
	defer response.Body.Close()
	return string(body), nil
}

func Recover() {
	if err := recover(); err != nil {
		fmt.Printf("recover err: %s\n", err)
	}
}

type Bitstamp struct {
	High      string `json:"high"`
	Last      string `json:"last"`
	Timestamp string `json:"timestamp"`
	Bid       string `json:"bid"`
	Vwap      string `json:"vwap"`
	Volume    string `json:"volume"`
	Low       string `json:"low"`
	Ask       string `json:"ask"`
	Open      string `json:"open"`
}

type Gdax struct {
	TradeId int64  `json:"trade_id"`
	Price   string `json:"price"`
	Size    string `json:"size"`
	Time    string `json:"time"`
	Bid     string `json:"bid"`
	Ask     string `json:"ask"`
	Volume  string `json:"volume"`
}

type Gemini struct {
	Symbol  string   `json:"symbol"`
	Open    string   `json:"open"`
	High    string   `json:"high"`
	Low     string   `json:"low"`
	Close   string   `json:"close"`
	Changes []string `json:"changes"`
	Bid     string   `json:"bid"`
	Ask     string   `json:"ask"`
}

type Spare struct {
	Result    Rs `json:"result"`
	Allowance Aw `json:"allowance"`
}

type Rs struct {
	Price float64 `json:"price"`
}

type Aw struct {
	Cost          float64 `json:"cost"`
	Remaining     float64 `json:"remaining"`
	RemainingPaid float64 `json:"remainingPaid"`
	Upgrade       string  `json:"upgrade"`
}
