package fmpgo

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
