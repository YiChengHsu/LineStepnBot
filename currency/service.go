package currency

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"line-stepn-bot/config"
	"line-stepn-bot/log"
)

var CurrencyData []Currency
var providerURL string = "https://api.coingecko.com/api/v3/coins"

func SyncCurrency() {
	var err error
	defer func() {
		if err != nil {
			log.Error(err)
		}
	}()

	var CurrentCurrencyData []Currency

	for _, tracking := range config.Global.BlockChain.Currencies {
		data, err := RequestCurrency(tracking)
		if err != nil {
			log.Error(fmt.Errorf("[Sync Error]%w", err))
		}

		CurrentCurrencyData = append(CurrentCurrencyData, data)
	}

	CurrencyData = CurrentCurrencyData
}

func RequestCurrency(currency string) (data Currency, err error) {

	url := fmt.Sprintf("%s/%s", providerURL, currency)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	res, err := new(http.Client).Do(req)
	if err != nil {
		return
	}

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		err = fmt.Errorf("[%s] %s: %d, %s", "GET", url, res.StatusCode, response)
		return
	}

	err = json.Unmarshal(response, &data)
	if err != nil {
		return
	}

	return
}

func LoadCurrency() (CurrencyArr []Currency) {
	CurrencyArr = CurrencyData
	return
}
