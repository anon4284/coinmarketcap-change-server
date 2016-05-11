package main

import (
	"projects/webScraper/logic"
	"projects/webScraper/scraper"
)

func main() {
	scraper.InitWebDriver()
	logic.InitCoins()

	logic.NewCoin("trumpcoin", "trump")
	logic.NewCoin("dogecoin", "doge")
	logic.NewCoin("bitcoin", "btc")
	logic.NewCoin("coexistcoin", "coxst")
	logic.NewCoin("primechain", "prime")
	logic.NewCoin("uncoin", "unc")

	routes()
}
