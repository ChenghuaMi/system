package service

import (
	"gorm.io/gorm"
	"strings"
	"system/errs"
	"system/form"
	"system/form/product"
	"system/global"
	"system/model"
	"time"
	"fmt"
)

type BrandService struct {

}
func(s *BrandService) List(param form.SearchForm) ([]form.Brand,error) {
	var result = []*model.Brand{}
	db := global.Db
	if param.S != "" {
		db = db.Where("brand_name like ?","%" + param.S + "%")
	}
	err := db.Find(&result).Error
	var catids = []int{}
	for _,item := range result {
		catids = append(catids,item.CateId)
	}
	catObj,err := BaseServiceObj.CateServiceObj.getCateNames(catids)
	if err != nil {
		return nil,err
	}
	brands := []form.Brand{}
	for _,item := range result {
		b := model.Brand{
			Id:        item.Id,
			BrandName: item.BrandName,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
			CateId:    item.CateId,
		}
		if v,ok := catObj[item.CateId];ok {
			brands = append(brands,form.Brand{
				Brand:    b,
				CateName: v,
			})
		}else {
			brands = append(brands,form.Brand{
				Brand:    b,
			})
		}
	}
	return brands,err
}
func(s *BrandService) AddBrand(param product.AddBrandForm) (*model.Brand,error) {
	ss,err := s.findOne(param.BrandName)
	if err != nil {
		return nil,err
	}else if ss.Id > 0 {
		return nil,errs.ErrorDataExist
	}
	m := model.Brand{
		BrandName: param.BrandName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CateId: param.CateId,
	}
	result := global.Db.Create(&m)
	return &m,result.Error
}
func(s *BrandService) EditBrand(param product.EditBrandForm) (*model.Brand,error) {
	ss,err := s.findOne(param.BrandName,param.Id)
	result,err := BaseServiceObj.CateServiceObj.getCateNames([]int{param.CateId})
	fmt.Println(result,err)
	if err != nil || len(result) == 0 {
		return nil,errs.ErrorDataNotExist
	}
	if err == nil && ss.Id == 0 {
		obj,err := s.findById(param.Id)
		fmt.Println(obj)
		fmt.Println(err)
		if (err == nil && obj.Id == 0) || err != nil {
			return nil,errs.ErrorDataNotExist
		}
		obj.BrandName = param.BrandName
		obj.UpdatedAt = time.Now()
		obj.CateId = param.CateId
		rs := global.Db.Save(&obj)
		return obj,rs.Error
	}else {
		return nil,err
	}
}
func(s *BrandService) findOne(brandName string,ids ...int) (*model.Brand,error) {
	m := new(model.Brand)
	db := global.Db.Where("brand_name = ?",brandName)
	if len(ids) > 0 {
		db = db.Where("id <> ?",ids[0])
	}
	err := db.First(&m).Error

	if err == gorm.ErrRecordNotFound {
		return m,nil
	} else if err != nil {
		return m,err
	}
	return m,nil
}
func(s *BrandService) findById(id int) (*model.Brand,error) {
	m := &model.Brand{}
	err := global.Db.Where("id = ?",id).First(&m).Error
	if err == gorm.ErrRecordNotFound {
		return m,nil
	}else if err != nil {
		return nil,err
	}else {
		return m,nil
	}

}
func(s *BrandService) DeleteBrand(brandIds string) bool {
	brs := strings.Split(brandIds,",")
	b := model.Brand{}
	global.Db.Delete(&b,brs)
	return true
}
