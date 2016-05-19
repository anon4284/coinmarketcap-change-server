package main

import "projects/webScraper/logic"

func main() {
	logic.InitCoins()

	logic.NewCoin("trumpcoin", "trump")
	logic.NewCoin("dogecoin", "doge")
	logic.NewCoin("bitcoin", "btc")
	logic.NewCoin("coexistcoin", "coxst")
	logic.NewCoin("primechain", "prime")
	logic.NewCoin("uncoin", "unc")
	logic.NewCoin("hitcoin", "htc")

	routes()
}
