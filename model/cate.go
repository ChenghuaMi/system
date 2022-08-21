package model

import "time"

type Cate struct {
	Id int	`gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt 	 time.Time
	Pid		     int
	CateName string
	Brand Brand `gorm:"foreignKey:CateId;references:Id"`
}

func(*Cate) TableName() string {
	return "cate"
}
