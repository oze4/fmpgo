package fmpgo

type stockTimeSeries struct {
    ticker Ticker
}

func (sts stockTimeSeries) RealTimeQuote(t... Ticker) []Quote {
    // TODO
    return []Quote{}
}