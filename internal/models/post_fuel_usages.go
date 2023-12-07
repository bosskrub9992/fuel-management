package models

import (
	"time"

	"github.com/jinleejun-corp/corelib/validators"
	"github.com/shopspring/decimal"
)

type CreateFuelUsageRequest struct {
	FuelUseDate        time.Time       `json:"fuelUseDate" validate:"max=50"`
	FuelPrice          decimal.Decimal `json:"fuelPrice"`
	KilometerBeforeUse int64           `json:"kilometerBeforeUse"`
	KilometerAfterUse  int64           `json:"kilometerAfterUse"`
	Description        string          `json:"description" validate:"max=500"`
	CarID              int64           `param:"currentCarId" validate:"required"`
	UserIDs            []int64         `json:"userId" validate:"required"`
	IsPaid             bool            `json:"isPaid"`
}

func (req CreateFuelUsageRequest) Validate() error {
	return validators.Validate(req)
}
