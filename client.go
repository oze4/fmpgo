package fmpgo

// Client knows how to interact with the fmpcloud.io API
type Client interface {
    APIKey() string
    StockTimeSeries() StockTimeSeries
}

type client struct {
    apiKey string
}

func (c *client) APIKey() string {
    return c.apiKey;
}

func (c *client) StockTimeSeries() StockTimeSeries {
    return stockTimeSeries{apiKey: c.apiKey}
}
