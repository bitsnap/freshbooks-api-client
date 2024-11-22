package enums_manual

import "slices"

type CreditsCreditType string

const (
	CreditGoodwill   CreditsCreditType = "goodwill"
	CreditPrepayment                   = "prepayment"
)

func KnownCreditType(name string) bool {
	return slices.Contains([]CreditsCreditType{
		CreditGoodwill,
		CreditPrepayment,
	}, CreditsCreditType(name))
}
