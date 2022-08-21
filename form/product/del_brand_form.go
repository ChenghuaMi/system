package product

type DeleteBrandForm struct {
	BrandIds string `form:"brand_ids" binding:"required"`
}
