package fmpgo

import "net/http"

// Client knows how to interact with the fmpcloud.io API
type Client interface {
	APIKey() string
	StockTimeSeries() StockTimeSeries
}

// client implements Client
type client struct {
	apiKey    string
	transport *http.Client
}

func (c *client) APIKey() string {
	return c.apiKey
}

func (c *client) StockTimeSeries() StockTimeSeries {
	return stockTimeSeries{c}
}
