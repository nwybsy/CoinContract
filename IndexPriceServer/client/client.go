package main

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars"

	"CoinContract"
)

func TestGetIndexPrice(Symbol string) {
	comm := tars.NewCommunicator()
	obj := fmt.Sprintf("CoinContract.IndexPriceServer.IndexPriceObjObj@tcp -h 127.0.0.1 -p 10015 -t 60000")
	app := new(CoinContract.IndexPriceObj)
	comm.StringToProxy(obj, app)
	var s1 string
	_, err := app.GetIndexPrice(Symbol, &s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", s1)
}

func main() {
	// comm := tars.NewCommunicator()
	// obj := fmt.Sprintf("CoinContract.IndexPriceServer.IndexPriceObjObj@tcp -h 127.0.0.1 -p 10015 -t 60000")
	// app := new(CoinContract.IndexPriceObj)
	// comm.StringToProxy(obj, app)
	// var out, i int32
	// i = 123
	// ret, err := app.Add(i, i*2, &out)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(ret, out)

	TestGetIndexPrice("BTC")

}
