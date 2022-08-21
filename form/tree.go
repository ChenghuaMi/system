package form

import (
	"time"
)

type Tree struct {
	Id int
	CreatedAt    time.Time
	UpdatedAt 	 time.Time
	Pid		     int
	CateName string
	Child []*Tree
}
