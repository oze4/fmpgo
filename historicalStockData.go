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

// Chart allows you to get historical chart data on the following
// time frames: 1min, 5min, 15min, 30min, and 1hour. Returns OHLCV
// with date.
// If you want intraday data for a specific date range, this is what
// you're looking for.
func (hsd historicalStockData) Chart(from, to time.Time, tf TimeFrame) ChartData {
	// TODO : get data from API and return it.
	return ChartData{}
}

// Price allows you to get detailed **daily** historical price info.
// Returns OHLCV with date and much more.
func (hsd historicalStockData) Price() HistoricalPrice {
	return historicalPrice{}
}
