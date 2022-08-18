package model

type AdminRole struct {
	RoleId int	`gorm:"primaryKey"`
	RoleName string
	Description string
	System byte
	Disabled byte
}

func(*AdminRole) TableName() string {
	return "admin_role"
}