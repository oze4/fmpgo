package fmpgo

import "net/http"

// New creates a new client
func New(apiKey string) Client {
	return &client{apiKey: apiKey}
}

// NewWithClient creates a new client using custom transport
func NewWithClient(httpclient *http.Client, apikey string) Client {
	return &client{apiKey: apikey, transport: httpclient}
}

// Client knows how to interact with the fmpcloud.io API
type Client interface {
	APIKey() string
	StockTimeSeries() StockTimeSeries
	CompanyValuation() CompanyValuation
}

// client implements Client
type client struct {
	apiKey    string
	transport *http.Client
}

// WithClient allows the consumer to provide their own *http.Client.
// This will modify the current http client we are using for transport.
func (c *client) WithClient(t *http.Client) {
	c.transport = t
}

func (c *client) APIKey() string {
	return c.apiKey
}

func (c *client) StockTimeSeries() StockTimeSeries {
	return stockTimeSeries{c}
}

func (c *client) CompanyValuation() CompanyValuation {
	return companyValuation{c}
}
