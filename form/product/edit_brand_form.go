package product

type EditBrandForm struct {
	Id int 				`form:"id" binding:"required"`
	BrandName string 	`form:"brand_name" binding:"required"`
	CateId int 	`form:"cate_id" binding:"required"`
}
