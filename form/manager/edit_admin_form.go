package manager

type EditAdminForm struct {
	Id int 	`form:"id" binding:"required"`
	Username string `form:"username" binding:"required"`
	Password string	`form:"password" binding:"required"`
	RoleName  string
	RoleId int	`form:"role_id" binding:"required"`
}
