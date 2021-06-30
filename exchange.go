package fmpgo

// Exchange represents a stock market exchange (serves as 'enum')
type Exchange string

// Exchanges
const (
	ExchangeNasdaq     Exchange = "nasdaq"
	ExchangeNYSE                = "nyse"
	ExchangeTSX                 = "tsx"
	ExchangeEuroNext            = "euronext"
	ExchangeMutualFund          = "mutual_fund"
	ExchangeETF                 = "etf"
	ExchangeAMEX                = "amex"
	ExchangeIndex               = "index"
	ExchangeCommodity           = "commodity"
	ExchangeForex               = "forex"
	ExchangeCrypto              = "crypto"
)

// IsValid validates provided Exchange
func (e Exchange) IsValid() bool {
	switch e {
	case ExchangeNasdaq,
		ExchangeAMEX,
		ExchangeCommodity,
		ExchangeCrypto,
		ExchangeETF,
		ExchangeEuroNext,
		ExchangeForex,
		ExchangeIndex,
		ExchangeMutualFund,
		ExchangeNYSE,
		ExchangeTSX:
		return true
	default:
		return false
	}
}
