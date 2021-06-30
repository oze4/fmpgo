package fmpgo

type stockTimeSeries struct {
	apiKey string
}

// RealTimeQuote allows you to batch real time quotes for tickers.
// You may also supply a single ticker as well.
func (sts stockTimeSeries) RealTimeQuote(tickers []string) []Quote {
	// TODO
	return []Quote{}
}

func (sts stockTimeSeries) TickerSearch(ticker string) []Ticker { // Since it is a search, we may need to return more than one result
	return []Ticker{}
}

func (sts stockTimeSeries) CompanyProfile(ticker string) {

}

func (sts stockTimeSeries) KeyExecutives(ticker string) {

}
