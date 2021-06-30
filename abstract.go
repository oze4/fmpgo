package fmpgo

// CompanyValuation knows how to get misc company data in regards to valuation
type CompanyValuation interface {
	/**
	 * TODO : ADD RETURN TYPES
	 */
	RSSFeed(ticker string)
	EarningsCalendar(ticker string)
	IPOCalendar(ticker string)
	StockSplitCalendar(ticker string)
	DividendCalendar(ticker string)
	EconomicCalendar(ticker string)
	FinancialStatementsZIP(ticker string)
}

// StockTimeSeries knows how to get stock data
type StockTimeSeries interface {
	/**
	 * TODO : ADD RETURN TYPES
	 */
	RealTimeQuote(tickers []string) []Quote
	TickerSearch(ticker string) []Ticker // Since it is a search, we may need to return more than one result
	CompanyProfile(ticker string)
	KeyExecutives(ticker string)
}
