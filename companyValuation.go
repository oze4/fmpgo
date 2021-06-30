package fmpgo

type companyValuation struct {
	client *client
}

// RSSFeed returns RSS data about a ticker
func (c companyValuation) RSSFeed(ticker string) {

}

// EarningsCalendar returns an earnings calendar for a ticker
func (c companyValuation) EarningsCalendar(ticker string) {

}

// IPOCalendar returns an IPO calendar for a ticker
func (c companyValuation) IPOCalendar(ticker string) {

}

// StockSplitCalendar returns stock split data for a ticker
func (c companyValuation) StockSplitCalendar(ticker string) {

}

// DividendCalendar returns a dividend calendar for a ticker
func (c companyValuation) DividendCalendar(ticker string) {

}

// EconomicCalendar returns economic events calendar
func (c companyValuation) EconomicCalendar(ticker string) {

}

// FinancialStatementsZIP returns a .zip (?) containing all financial statements for a ticker
// TODO:(??need to do more research on this endpoint??)
func (c companyValuation) FinancialStatementsZIP(ticker string) {

}
