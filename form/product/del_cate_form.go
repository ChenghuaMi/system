package product
type DeleteCateForm struct {
	CateIds string	`form:"cate_ids" binding:"required"`
}
