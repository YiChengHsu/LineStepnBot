package currency

type Currency struct {
	ID         string     `json:"id"`
	Symbol     string     `json:"symbol"`
	Name       string     `json:"name"`
	MarketData MarketData `json:"market_data"`
}

type MarketData struct {
	CurrencyPrice            FiatPrice `json:"current_price"`
	PriceChangePercentage1H  FiatPrice `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentage24H FiatPrice `json:"price_change_percentage_24h_in_currency"`
	PriceChangePercentage7D  FiatPrice `json:"price_change_percentage_7d_in_currency"`
}

type FiatPrice struct {
	TWD float64 `json:"twd"`
	USD float64 `json:"usd"`
}
