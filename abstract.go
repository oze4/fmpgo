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
	TickerSearch(ticker string, limit int) []Ticker
	TickerSearchByExchange(ticker string, exchange Exchange, limit int) Ticker
	CompanyProfile(ticker string)
	KeyExecutives(ticker string)
}
