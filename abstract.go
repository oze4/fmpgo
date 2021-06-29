package fmpgo 

// CompanyValuation knows how to get misc company data in regards to valuation
type CompanyValuation interface {
    /**
     * TODO : ADD RETURN TYPES
     */
    RSSFeed(t Ticker)
    EarningsCalendar(t Ticker) 
    IPOCalendar(t Ticker) 
    StockSplitCalendar(t Ticker)
    DividendCalendar(t Ticker)
    EconomicCalendar(t Ticker)
    FinancialStatementsZIP(t Ticker)
}

// StockTimeSeries knows how to get stock data
type StockTimeSeries interface {
    RealTimeQuote(t... Ticker) []Quote
}
