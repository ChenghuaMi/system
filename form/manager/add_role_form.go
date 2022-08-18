package manager

type AddRoleForm struct {
	RoleId int `form:"role_id"`
	RoleName string `form:"role_name" binding:"required"`
	Description string 	`form:"description"`
	System byte `form:"system"`
	Disabled byte `form:"disabled"`
}
