package main

import (
	"CoinContract/IndexPriceServer/servant"

	"github.com/TarsCloud/TarsGo/tars"

	"CoinContract"
)

func main() { //Init servant
	//imp := new(IndexPriceObjImp) //New Imp
	imp := new(servant.IndexPriceObjImp)
	app := new(CoinContract.IndexPriceObj)                          //New init the A Tars
	cfg := tars.GetServerConfig()                                   //Get Config File Object
	app.AddServant(imp, cfg.App+"."+cfg.Server+".IndexPriceObjObj") //Register Servant
	tars.Run()
}
