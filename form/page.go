package form

type Page struct {
	Page int `form:"page" binding:"required"`
	Size int `form:"size" binding:"required"`
	Param string `form:"param"`
}
