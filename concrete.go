package fmpgo

/**
 * Concrete types for data returned by the fmpcloud API.
 */

import (
	"time"
)

// Ticker represents a stock ticker
type Ticker struct {
	Symbol            string
	Name              string
	Currency          string
	StockExchange     string
	ExchangeShortName string
}

// OHLCV represents open, high, low, close, volume data.
type OHLCV struct {
	open   float64 // 136.94000000,
	low    float64 // 136.91000000,
	high   float64 // 137.20000000,
	close  float64 // 137.01000000,
	volume float64 // 776219
}

// Chart represents chart data for a ticker. Returns OHLCV with date.
type Chart struct {
	OHLCV
	date time.Time
}

// DailyPriceSummary only return daily price data.
type DailyPriceSummary struct {
	OHLCV
	date             time.Time
	adjClose         float64
	unadjustedVolume float64
	change           float64
	changePercent    float64
	vwap             float64
	label            time.Time // | *may need a string here?* | March 12, 19,
	changeOverTime   float64
}

// HistoricalDailyPriceData represents historical daily price data returned
// from fmpcloud API.
type HistoricalDailyPriceData struct {
	Symbol     string
	Historical []DailyPriceSummary
}

// Dividend represents dividend metadata
type Dividend struct {
}

// Split represents stock split metadata (when, how much, etc..)
type Split struct {
}

// Quote represents a stock quote
type Quote struct {
	Symbol               string
	Name                 string
	Price                float64
	ChangesPercentage    float64
	Change               float64
	DayLow               float64
	DayHigh              float64
	YearHigh             float64
	YearLow              float64
	MarketCap            float64
	PriceAvg50           float64
	PriceAvg200          float64
	Volume               float64
	AvgVolume            float64
	Exchange             string
	Open                 float64
	PreviousClose        float64
	Eps                  float64
	PE                   float64
	EarningsAnnouncement time.Time
	SharesOutstanding    float64
	Timestamp            time.Time
}
