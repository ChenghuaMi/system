package model

type Admin struct {
	Id int
	Username string
	Password string
	RoleName  string
	RoleId int
	AdminRole AdminRole `gorm:"foreignKey:RoleId;references:RoleId"`
}

func(*Admin) TableName() string {
	return "admin"
}
