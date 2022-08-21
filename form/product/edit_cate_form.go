package product

type EditCateForm struct {
	Id int `form:"id" binding:"required"`
	CateName string `form:"cate_name" binding:"required"`
	Pid 	int `form:"pid"`
}
