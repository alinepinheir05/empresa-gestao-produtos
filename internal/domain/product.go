package domain

type Product struct {
	ID           uint    `gorm:"primaryKey;autoIncrement"`
	Name         string  `gorm:"size:100;not null"`
	Code         string  `gorm:"size:50;unique;not null"`
	Description  string  `gorm:"size:255"`
	Unit         string  `gorm:"size:20"`
	CostEstimate float64 `gorm:"type:decimal(10,2)"`
}
