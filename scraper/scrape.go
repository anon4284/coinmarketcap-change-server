package scraper

import (
	"fmt"
	"projects/webScraper/util"

	"github.com/sclevine/agouti"
)

var agoutiDriver *agouti.WebDriver
var page *agouti.Page

func InitWebDriver() {

	agoutiDriver = agouti.PhantomJS()
	agoutiDriver.Start()

	page, _ = agoutiDriver.NewPage(agouti.Browser("firefox"))

}

func ScrapeChange(coin string) (bool, string) {
	fmt.Println("Scrape job")
	page.Navigate("http://coinmarketcap.com/currencies/" + coin)

	change := page.FirstByClass("negative_change")

	isFound, err := change.Visible()
	util.CheckErr(err)

	if isFound {
		txt, _ := change.Text()
		return true, util.FormatString(txt)
	}
	change = page.FirstByClass("positive_change")
	isFound, err = change.Visible()
	util.CheckErr(err)
	if isFound {
		txt, err := change.Text()
		util.CheckErr(err)
		return true, util.FormatString(txt)
	}
	page.Destroy()
	agoutiDriver.Stop()
	return false, "currency not found"
}
