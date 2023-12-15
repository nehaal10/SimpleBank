package util

const (
	USD = "USD"
	INR = "INR"
	EUR = "EUD"
	CAD = "CAD"
)

func IsdupportedCurrency(currency string) bool {
	switch currency {
	case USD, INR, EUR, CAD:
		return true
	}
	return false
}
