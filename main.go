package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/sclevine/agouti"
)

var agoutiDriver *agouti.WebDriver

func main() {

	n := negroni.Classic()
	router := mux.NewRouter()

	router.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {

		currency := r.URL.Query().Get("coin")
		w.Write([]byte(getChange(currency)))
	})
	n.UseHandler(router)
	n.Run(":5000")
}

func getChange(coin string) string {
	agoutiDriver = agouti.PhantomJS()
	agoutiDriver.Start()
	page, err := agoutiDriver.NewPage(agouti.Browser("firefox"))
	if err != nil {
		fmt.Println(err)
	}

	page.Navigate("http://coinmarketcap.com/currencies/" + coin)

	change := page.FirstByClass("negative_change")

	isFound, _ := change.Visible()

	if isFound {
		txt, _ := change.Text()
		return formatString(txt)
	}
	change = page.FirstByClass("positive_change")
	isFound, _ = change.Visible()
	if isFound {
		txt, _ := change.Text()
		return formatString(txt)
	}
	page.Destroy()
	agoutiDriver.Stop()
	return "false"
}

func formatString(str string) string {
	str = strings.Replace(str, "(", "", -1)
	str = strings.Replace(str, "%", "", -1)
	str = strings.Replace(str, ")", "", -1)
	return strings.Replace(str, " ", "", -1)
}
