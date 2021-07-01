package fmpgo

import "time"

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

// StockTimeSeries is an endpoint that contains actions and knows how to 
// get time series related data from fmpcloud.
type StockTimeSeries interface {
	/**
	 * TODO : ADD RETURN TYPES
	 */
	RealTimeQuote(tickers []string) []Quote
	TickerSearch(ticker string, limit int) []Ticker
	TickerSearchByExchange(exchange Exchange, ticker string, limit int) Ticker
	CompanyProfile(ticker string)
	KeyExecutives(ticker string)
	HistoricalStockData(ticker string) HistoricalStockData
}

// HistoricalStockData is an action that knows how to get historical stock data from fmpcloud.
type HistoricalStockData interface {
	// Chart returns OHLCV with date. Use for intraday data.
	Chart(from, to time.Time, tf TimeFrame) Chart
	Daily(from, to time.Time) []DailyPriceSummary
	DailyLastXDays(x int) []DailyPriceSummary
	StockDividend() []Dividend
	StockSplit() []Split
}
