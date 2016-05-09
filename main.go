package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jasonlvhit/gocron"
	"github.com/sclevine/agouti"
)

var agoutiDriver *agouti.WebDriver
var page *agouti.Page

var storage = ""

func main() {

	agoutiDriver = agouti.PhantomJS()
	agoutiDriver.Start()

	page, _ = agoutiDriver.NewPage(agouti.Browser("firefox"))
	storage = getChange("trumpcoin")
	gocron.Every(10).Minute().Do(task)
	gocron.Start()
	fmt.Println(storage)

	n := negroni.Classic()
	router := mux.NewRouter()

	router.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(storage))
	})
	n.UseHandler(router)
	n.Run(":5000")

}

func task() {
	fmt.Println("I am runnning task.")
	storage = getChange("trumpcoin")
}

func getChange(coin string) string {
	page.Navigate("http://coinmarketcap.com/currencies/" + coin)

	change := page.FirstByClass("negative_change")

	isFound, err := change.Visible()
	checkErr(err)

	if isFound {
		txt, err := change.Text()
		checkErr(err)
		return formatString(txt)
	}
	change = page.FirstByClass("positive_change")
	isFound, err = change.Visible()
	checkErr(err)
	if isFound {
		txt, err := change.Text()
		checkErr(err)
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

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
