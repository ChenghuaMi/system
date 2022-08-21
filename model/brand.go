package model

import "time"

type Brand struct {
	Id int
	BrandName string

	CreatedAt time.Time
	UpdatedAt time.Time
	CateId int
}
func(*Brand) TableName() string {
	return "brand"
}
