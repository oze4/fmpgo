package fmpgo

/**
 * Becuase this endpoint is so massive, it gets it's own file
 */

// historicalStockData implements HistoricalStockData
type historicalStockData struct {
	ticker string
	client *client
}

// HistoricalChart allows you to get historical chart data on the following
// time frames: 1min, 5min, 15min, 30min, and 1hour.
func (hsd historicalStockData) HistoricalChart(tf TimeFrame) {

}

func (hsd historicalStockData) HistoricalPrice() {

}
