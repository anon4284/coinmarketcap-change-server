package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"projects/webScraper/util"
	"strings"
)

//CMCPREFIX prefix for cmc api
const CMCPREFIX = "http://coinmarketcap-nexuist.rhcloud.com/api/"

//CMCResponse CMC Response
type CMCResponse struct {
	Change string
}

//GetChange get current course
func GetChange(coin string) (bool, string) {

	resp, err := http.Get(CMCPREFIX + coin)
	util.CheckErr(err)

	buff := new(bytes.Buffer)
	buff.ReadFrom(resp.Body)

	data := buff.String()

	if strings.Contains(data, "error") {
		return false, data
	}

	response := CMCResponse{}
	json.Unmarshal([]byte(data), &response)

	return true, response.Change

}
