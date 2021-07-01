package fmpgo

import "time"

/**
 * Becuase this endpoint is so massive, it gets it's own file
 */

// historicalStockData implements HistoricalStockData
type historicalStockData struct {
	ticker string
	client *client
}

// Chart allows you to get historical chart data on the following time frames: 
//    1min, 5min, 15min, 30min, and 1hour 
// Returns OHLCV with date. If you want intraday data for a specific date range, 
// this is what you're looking for.
func (hsd historicalStockData) Chart(from, to time.Time, tf TimeFrame) Chart {
	// TODO : get data from API and return it.
	return Chart{}
}

// Daily allows you to get detailed **daily** historical price data. No intraday time frames.
func (hsd historicalStockData) Daily(from, to time.Time) []DailyPriceSummary {
	return []DailyPriceSummary{}
}

// DailyLastXDays allows you to easily get daily data instead of having to provide
// a range for recent dates.
func (hsd historicalStockData) DailyLastXDays(x int) []DailyPriceSummary {
	return []DailyPriceSummary{}
}

// Get historical dividend data
func (hsd historicalStockData) StockDividend() []Dividend {
	return []Dividend{}
}

// Get historical stock split data
func (hsd historicalStockData) StockSplit() []Split {
	return []Split{}
}
