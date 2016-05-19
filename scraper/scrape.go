package scraper

import (
	"projects/webScraper/util"

	"github.com/PuerkitoBio/goquery"
)

//PREFIX for webscraping
const PREFIX = "https://coinmarketcap.com/currencies/"

//ScrapeChange get change through webscraping
func ScrapeChange(coin string) (bool, string) {
	url := PREFIX + coin

	doc, _ := goquery.NewDocument(url)
	positive := doc.Find(".positive_change")
	negative := doc.Find(".negative_change")

	if positive.Length() > 0 {
		return true, util.FormatString(positive.First().Text())
	}
	if negative.Length() > 0 {
		return true, util.FormatString(negative.First().Text())
	}
	return false, "Currency not found"
}
