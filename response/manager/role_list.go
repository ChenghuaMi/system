package manager

import "system/model"

type RoleList struct {
	AdminRoles []model.AdminRole
	Page int
	Size int
	Total int64
}

type AdminList struct {
	Admins []model.Admin
	Page int
	Size int
	Total int64
}
