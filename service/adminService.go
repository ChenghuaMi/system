package service

import (
	"errors"
	"gorm.io/gorm"
	"system/errs"
	"system/form"
	manager2 "system/form/manager"
	"system/global"
	"system/model"
	"fmt"
	"system/response/manager"
	"system/util"
)

type AdminService struct {

}
func(s *AdminService) List(param form.Page) (manager.AdminList,error){
	m := []model.Admin{}
	page := param.Page
	size := param.Size
	q := param.Param
	var total int64
	ss := manager.AdminList{}
	obj := global.Db.Preload("AdminRole")
	if q != "" {
		obj = obj.Where("username = ?","%"+ q +"%")
	}
	err  := obj.Offset((page - 1) * size).Limit(size).Find(&m).Offset(-1).Limit(-1).Count(&total).Error
	fmt.Println(err)
	ss.Admins = m
	ss.Size = size
	ss.Total = total
	ss.Page = page
	fmt.Printf("%#v",m)
	return ss,nil
}
func(s *AdminService) AddAdmin(param manager2.AddAdminForm) (model.Admin,error) {
	username := param.Username
	m,err := s.findAdmin(username)
	if err != nil {
		return m,err
	}
	roleObj,err := BaseServiceObj.RoleServiceObj.findById(param.RoleId)
	if err != nil {
		return m,errs.ErrorRoleIdNotExist
	}
	admin := model.Admin{
		Username:  param.Username,
		Password:  util.Md5String(param.Password),
		RoleName:  roleObj.RoleName,
		RoleId:    param.RoleId,
	}
	result := global.Db.Create(&admin)
	return admin,result.Error

}
func(s *AdminService) findAdmin(username string) (model.Admin,error) {
	var m model.Admin
	err := global.Db.Where("username = ?",username).First(&m).Error
	fmt.Println(err,gorm.ErrRecordNotFound)
	if err == gorm.ErrRecordNotFound {
		return m,nil
	}
	fmt.Println("err:",err,m.Id)
	if m.Id > 0 {
		return m,errs.ErrorUserAlreadyExist
	}
	return m,nil
}

func(s *AdminService) EditAdmin(param manager2.EditAdminForm) (model.Admin,error) {
	var m model.Admin
	err := global.Db.Where("id = ?",param.Id).First(&m).Error
	fmt.Println("err::::",err)
	//if err != gorm.ErrRecordNotFound {
	//	return m,errs.ErrorUserNotExist
	//}
	if err != nil {
		return m,errors.New(err.Error())
	}
	m.Username = param.Username
	m.RoleId = param.RoleId
	m.Password = util.Md5String(param.Password)
	global.Db.Save(&m)
	return m,nil
}
func(s *AdminService) DeleteAdmin(param manager2.AdminForm) error {
	id := param.Id
	m := model.Admin{}
	err := global.Db.Where("id = ?",id).First(&m).Error
	if err != nil {
		return err
	}
	global.Db.Delete(&m)
	return nil
}