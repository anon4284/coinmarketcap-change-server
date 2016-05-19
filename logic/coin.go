package logic

import (
	"fmt"
	"projects/webScraper/api"
	"projects/webScraper/scraper"
	"time"
)

type Coin struct {
	API    bool
	Change string
	Name   string
	Code   string
	Ticker *time.Ticker
}

func NewCoin(name string, code string) *Coin {
	c := new(Coin)
	c.Name = name
	c.Code = code
	c.update()
	c.timer()
	CoinStorage[code] = *c
	return c
}

func (c *Coin) update() {
	success, change := api.GetChange(c.Code)
	if success {
		c.API = true
		c.Change = change
		fmt.Println(c.Change)
	} else {
		success, change = scraper.ScrapeChange(c.Name)
		if success {
			c.API = false
			c.Change = change
			fmt.Println(c.Change)
		} else {
			c.API = false
			c.Change = change
		}
	}
	CoinStorage[c.Code] = *c
}

func (c *Coin) timer() {
	var interval time.Duration

	if c.API {
		interval = 2 * time.Minute
	} else {
		interval = 2 * time.Minute
	}

	c.Ticker = time.NewTicker(interval)
	go func() {
		for range c.Ticker.C {
			c.update()
		}
	}()
}

var CoinStorage map[string]Coin

func InitCoins() {
	CoinStorage = make(map[string]Coin)
}
