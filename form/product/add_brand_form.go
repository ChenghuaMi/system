package product

type AddBrandForm struct {
	BrandName string 	`form:"brand_name" binding:"required"`
	CateId int 	`form:"cate_id" binding:"required"`
}
