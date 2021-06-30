package fmpgo

//
//  This file contains an enum like type used for historical chart action under StockTimeSeries
//

// TimeFrame represents a chart/candlestick time frame
type TimeFrame string

// TimeFrames
const (
	TimeFrameOneMin     TimeFrame = "1min"
	TimeFrameFiveMin              = "5min"
	TimeFrameFifteenMin           = "15min"
	TimeFrameThirtyMin            = "30min"
	TimeFrameOneHour              = "1hour"
)

// IsValid validates whether the provided timeframe is valid
func (tf TimeFrame) IsValid() bool {
	switch tf {
	case TimeFrameOneMin,
		TimeFrameFiveMin,
		TimeFrameFifteenMin,
		TimeFrameThirtyMin,
		TimeFrameOneHour:
		return true
	default:
		return false
	}
}
