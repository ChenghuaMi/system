package manager

type RoleForm struct {
	RoleId int `form:"role_id"`
	RoleName string `form:"role_name"`
	Description string 	`form:"description"`
	System byte `form:"system"`
	Disabled byte `form:"disabled"`
}
