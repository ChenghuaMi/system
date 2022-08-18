package manager

type AdminForm struct {
	Id string `form:"id" binding:"required"`
}
