# fmpgo
 fmpcloud.io client in golang

# WORK IN PROGRESS

This repo is a WIP.

# API Design

Our API is designed to [mirror `fmpcloud.io` documentation](https://fmpcloud.io/documentation). 

Consider the following hierarchy...

![fmpcloud_sidebar](https://raw.githubusercontent.com/oze4/fmpgo/main/.github/docs/img/sidebar.png)

Each 'category' has 'actions'...

![fmpcloud_stockTimeSeries](https://raw.githubusercontent.com/oze4/fmpgo/main/.github/docs/img/stockTimeSeries_Sidebar.png)

The path above will match up with our API, so it's easy to follow the docs and use this library...

```golang
// Build fmpcloud.io client
httpclient := &http.Client{}
fmpclient := &fmpgo.Client{APIKey: "<your_api_key>", HTTPClient: httpclient}
// Build tickers so we can batch real time quotes
abcd := fmpgo.Ticker{Symbol: "ABCD"}
efgh := fmpgo.Ticker{Symbol: "EFGH"}
// Get quotes
quotes, _ := fmpclient.StockTimeSeries.RealTimeQuote([]fmpgo.Ticker{abcd, efgh}) // -> []fmpgo.Quote, error
```
