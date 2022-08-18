package service

var BaseServiceObj = new(BaseServices)
type BaseServices struct {
	LoginServiceObj *LoginService
	RoleServiceObj *RoleService
	AdminServiceObj *AdminService
}
