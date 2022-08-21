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
)

type CateService struct {

}
func(s *CateService) List(params form.SearchForm)  ([]*form.Tree,error) {
	var m []model.Cate
	db := global.Db
	if params.S != "" {
		db = db.Where("cate_name like ?","%" +params.S+ "%")
	}
	result := db.Debug().Find(&m)
	if result.Error != nil {
		return nil,result.Error
	}
	rs := s.loopData(m)
	return rs,nil
}
func(s *CateService) loopData(data []model.Cate) []*form.Tree {
	roots := s.getRoot(data)
	for _,root := range roots {
		root.Child = s.makeChild(data,root)
	}
	return roots
}
func(s *CateService) makeChild(data []model.Cate,root *form.Tree) []*form.Tree {
	var children = []*form.Tree{}
	for _,item := range data {
		if item.Pid == root.Id {
			ss := &form.Tree{
				Id:        item.Id,
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
				Pid:       item.Pid,
				CateName:  item.CateName,
				Child:     []*form.Tree{},
			}

			ss.Child = s.makeChild(data,ss)
			children = append(children,ss)
		}
	}
	return children
}
func(s *CateService) getRoot(data []model.Cate) []*form.Tree {
	roots := []*form.Tree{}
	for _,item := range data {
		if item.Pid == 0 {
			child := &form.Tree{
				Id:        item.Id,
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
				Pid:       item.Pid,
				CateName:  item.CateName,
				Child:     []*form.Tree{},
			}
			roots = append(roots,child)
		}
	}
	return roots

}
func (s *CateService) AddCate(params product.CateForm) (*model.Cate,error) {
	cateObj := model.Cate{}
	m,err := s.findOne(params.CateName)
	if err == nil && m.Id == 0  {
		cateObj.CateName = params.CateName
		cateObj.Pid = params.Pid
		cateObj.CreatedAt = time.Now()
		cateObj.UpdatedAt = time.Now()
		global.Db.Create(&cateObj)
		return &cateObj,err
	}
	return nil,errs.ErrorDataExist
}
func(s *CateService) findOne(catename string,id ...int) (*model.Cate,error) {
	m := new(model.Cate)
	db := global.Db.Where("cate_name = ?",catename)
	if len(id) > 0 {
		db = db.Where("id <> ?",id[0])
	}
	err := db.First(&m).Error

	if err == gorm.ErrRecordNotFound {
		return m,nil
	} else if err != nil {
		return m,err
	}
	return m,nil

}
func(s *CateService) EditCate(params product.EditCateForm)(cateObj *model.Cate,err error) {
	m,err := s.findOne(params.CateName,params.Id)
	if m.Id > 0 && err == nil {
		return nil,errs.ErrorDataExist
	}
	cateObj,err = s.findOne(params.CateName)
	cateObj.CateName = params.CateName
	cateObj.Pid = params.Pid
	cateObj.UpdatedAt = time.Now()
	global.Db.Save(cateObj)
	return cateObj,err
}
func(s *CateService) DeleteCate(params product.DeleteCateForm) bool {
	cateIds := strings.Split(params.CateIds,",")
	var m model.Cate
	global.Db.Delete(&m,cateIds)
	return true

}
func(s *CateService) getCateNames(cateIds []int) (map[int]string,error) {
	var m = []model.Cate{}
	result := map[int]string{}
	ss := global.Db.Find(&m,cateIds)
	for _,item := range m {
		result[item.Id]  = item.CateName
	}
	return result,ss.Error
}