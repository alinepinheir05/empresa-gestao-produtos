package domain

type Product struct {
    ID           uint    `gorm:"primaryKey" json:"id"`
    Name         string  `json:"name"`
    Code         string  `gorm:"unique" json:"code"`
    Description  string  `json:"description"`
    Unit         string  `json:"unit"`
    CostEstimate float64 `json:"costEstimate"`
    Weight       float64 `json:"weight"`
    Color        string  `json:"color"`
    Volume       float64 `json:"volume"`
    Thickness    float64 `json:"thickness"`
    RawMaterial  string  `json:"rawMaterial"`
    ClientBrand  string  `json:"clientBrand"`
    EAN          string  `json:"ean"`
}
