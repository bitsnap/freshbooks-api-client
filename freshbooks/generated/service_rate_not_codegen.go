package generated

import "github.com/shopspring/decimal"

type ServiceRate struct {
	Rate *decimal.Decimal `json:"rate" validate:"required"`
}
