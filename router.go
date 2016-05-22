package main

import (
	"net/http"
	"projects/coinmarketcap-change-server/logic"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func routes() {
	n := negroni.Classic()
	router := mux.NewRouter()

	router.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		if val, ok := logic.CoinStorage[code]; ok {
			w.Write([]byte(val.Change))
		} else {
			w.Write([]byte("Coin does not exist"))
		}
	})

	n.UseHandler(router)
	n.Run(":5000")
}
