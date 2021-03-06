package fmpgo

// stockTimeSeries implements StockTimeSeries
type stockTimeSeries struct {
	ticker string
	client *client
}

func (sts stockTimeSeries) HistoricalStockData(ticker string) HistoricalStockData {
	return historicalStockData{client: sts.client, ticker: ticker}
}

// RealTimeQuote allows you to batch real time quotes for tickers.
// For single ticker use like : `RealTimeQuote([]string{"TCKR"})`
func (sts stockTimeSeries) RealTimeQuote(tickers []string) []Quote {
	// TODO
	return []Quote{}
}

// TickerSearch allows you to search all exchanges for a ticker symbol.
func (sts stockTimeSeries) TickerSearch(ticker string, limit int) []Ticker {
	return []Ticker{}
}

// TickerSearchByExchange allows you to search a specific exchange for a ticker symbol.
func (sts stockTimeSeries) TickerSearchByExchange(exchange Exchange, ticker string, limit int) Ticker {
	return Ticker{}
}

// CompanyProfile returns profile data about a company.
func (sts stockTimeSeries) CompanyProfile(ticker string) {

}

// KeyExecutives returns key "c-suite" executives.
func (sts stockTimeSeries) KeyExecutives(ticker string) {

}
