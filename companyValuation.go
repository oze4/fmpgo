package fmpgo

// companyValuation implements CompanyValuation
type companyValuation struct {
	ticker string
	client *client
}

// RSSFeed returns RSS data about a ticker
func (c companyValuation) RSSFeed() {

}

// EarningsCalendar returns an earnings calendar for a ticker
func (c companyValuation) EarningsCalendar() {

}

// IPOCalendar returns an IPO calendar for a ticker
func (c companyValuation) IPOCalendar() {

}

// StockSplitCalendar returns stock split data for a ticker
func (c companyValuation) StockSplitCalendar() {

}

// DividendCalendar returns a dividend calendar for a ticker
func (c companyValuation) DividendCalendar() {

}

// EconomicCalendar returns economic events calendar
func (c companyValuation) EconomicCalendar() {

}

// FinancialStatementsZIP returns a .zip (?) containing all financial statements for a ticker
// TODO:(??need to do more research on this endpoint??)
func (c companyValuation) FinancialStatementsZIP() {

}
