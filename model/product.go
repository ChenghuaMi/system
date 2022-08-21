package model

import "time"

type Product struct {
	Id int
	title string
	CateId int
	BrandId int
	Price float64
	Thumb string
	Store int
	CreatedAt time.Time
	UpdatedAt time.Time
}
func(*Product) TableName() string {
	return "product"
}