package fmpgo

import (
    "time"
)

type Ticker struct {
    Symbol string //MPAA
    Name string //Motorcar Parts of America Inc
    Currency string //USD
    StockExchange string //Nasdaq Global Select
    ExchangeShortName string //NASDAQ
}

type Quote struct {
    Symbol string
    Name string
    Price float64 // 134.88090000
    ChangesPercentage float64 // 1.33000000
    Change float64 //1.77090000
    DayLow float64 //133.38570000
    DayHigh float64 // 135.22000000
    YearHigh float64 // 145.09000000
    YearLow float64 // 87.82000000
    MarketCap float64 // 2250838638592.00000000
    PriceAvg50 float64 // 127.59588000
    PriceAvg200 float64 // 128.85368000
    Volume float64 //35896256
    AvgVolume float64 //84175388
    Exchange string
    Open float64 //133.41000000
    PreviousClose float64 //133.11000000
    Eps float64 //4.44900000
    PE float64 //30.31713000
    EarningsAnnouncement time.Time //2021-04-28T16:30:00.000+0000
    SharesOutstanding float64 //16687600977
    Timestamp time.Time // unix timestamp aka `time.Now().Unix()`
}
