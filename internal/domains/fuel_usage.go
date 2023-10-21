package domains

import (
	"time"

	"github.com/shopspring/decimal"
)

type FuelUsage struct {
	ID                 int64           `gorm:"column:id"`
	FuelUseDate        time.Time       `gorm:"column:fuel_use_date"`
	FuelPrice          decimal.Decimal `gorm:"column:fuel_price"`
	KilometerBeforeUse int64           `gorm:"column:kilometer_before_use"`
	KilometerAfterUse  int64           `gorm:"column:kilometer_after_use"`
	Description        string          `gorm:"column:description"`
	TotalMoney         decimal.Decimal `gorm:"column:total_money"`
	CreateTime         time.Time       `gorm:"column:create_time"`
	UpdateTime         time.Time       `gorm:"column:update_time"`
}

func (d FuelUsage) TableName() string {
	return "fuel_usages"
}
