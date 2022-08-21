package product

type CateForm struct {
	Id int `form:"id"`
	CateName string `form:"cate_name" binding:"required"`
	Pid 	int `form:"pid"`
}
