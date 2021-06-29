package fmpgo

type stockTimeSeries struct {}

// RealTimeQuote allows you to batch real time quotes for tickers.
// You may also supply a single ticker as well.
func (sts stockTimeSeries) RealTimeQuote(t... Ticker) []Quote {
    // TODO
    return []Quote{}
}