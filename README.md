# CoinContract
根据市场上多家交易所的最新成交价格，进行加权平均，计算而得到的指数价格。

# 常用命令
1. 启动
```
go run main.go config.go   --config IndexPriceServer.conf
```

2. 打包
```
make && make install
```

3. 引入tars文件
```
tars2go --outdir=./vendor IndexPriceObj.tars
```