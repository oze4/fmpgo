package fmpgo

// CompanyValuation knows how to get misc company data in regards to valuation
type CompanyValuation interface {
	/**
	 * TODO : ADD RETURN TYPES
	 */
	RSSFeed()
	EarningsCalendar()
	IPOCalendar()
	StockSplitCalendar()
	DividendCalendar()
	EconomicCalendar()
	FinancialStatementsZIP()
}

// StockTimeSeries knows how to get stock data
type StockTimeSeries interface {
	/**
	 * TODO : ADD RETURN TYPES
	 */
	RealTimeQuote(tickers []string) []Quote
	TickerSearch(ticker string, limit int) []Ticker
	TickerSearchByExchange(exchange Exchange, ticker string, limit int) Ticker
	CompanyProfile(ticker string)
	KeyExecutives(ticker string)
}
