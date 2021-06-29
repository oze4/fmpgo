package fmpgo

import (
    "net/http"
)

type FMPClient struct {
    APIKey string
    HTTPClient *http.Client
    StockTimeSeries StockTimeSeries
}
