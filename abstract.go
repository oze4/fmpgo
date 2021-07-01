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
	HistoricalStockData(ticker string) HistoricalStockData
}

// HistoricalStockData gets historical stock data
// This "action" is so large that it has it's own interface & file.
type HistoricalStockData interface {
	// Chart returns OHLCV with date. Use
	// for intraday data.
	Chart(from, to time.Time, tf TimeFrame) ChartData
	// Price returns OHLCV with date and
	// more but only for daily, no
	// intraday time frames.
	Price() HistoricalPrice
}

// HistoricalPrice knows how to get detailed price info via historicalStockData endpoint.
// HistoricalPrice endpoint is a little more flexible than `HistoricalStockData.Chart(...)`
// so it required it's own interface and file.
type HistoricalPrice interface {
	// DailyLine()
	// DailyChangeAndVolume()
	// DailySpecificPeriod(from, to time.Time)
	// DailyLastXDays(x int)
	// DailyStockDividend()
	// DailyStockSplit()
}
