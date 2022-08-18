package service

import (
	"github.com/jinzhu/gorm"
	"system/errs"
	"system/form"
	"system/form/manager"
	"system/global"
	"system/model"
	"fmt"
	manager2 "system/response/manager"
)

type RoleService struct {

}
// 增加角色
func(s *RoleService) AddRole(form manager.AddRoleForm) (*model.AdminRole,error) {
	_,err := s.findOne(form.RoleName)
	if err != nil {
		return nil,errs.ErrorRoleExist
	}
	role := model.AdminRole{
		RoleName:    form.RoleName,
		Description: form.Description,
		System:      form.System,
		Disabled:    form.Disabled,
	}
	global.Db.Create(&role)
	return &role,nil
}
func(s *RoleService) EditRole(param interface{}) (*model.AdminRole,error) {
	form,_ := param.(manager.EditRoleForm)
	m,err := s.findById(form.RoleId)
	if err != nil {
		return nil,err
	}
	m.RoleName=form.RoleName
	m.Disabled = form.Disabled
	m.System = form.System
	global.Db.Save(m)
	return m,nil

}
func(s *RoleService) findById(roleId int) (*model.AdminRole,error) {
	m := model.AdminRole{}

	err := global.Db.Where("role_id = ?",roleId).First(&m).Error
	if err == gorm.ErrRecordNotFound {
		return nil,errs.ErrorRoleIdNotExist
	}
	if err == nil {
		return &m,nil
	}else {
		return nil,errs.ErrorRoleNotExist
	}

}
func(s *RoleService) DeleteRole(form manager.RoleForm) error {
	m,err := s.findById(form.RoleId)
	if err != nil {
		return errs.ErrorRoleIdNotExist
	}
	result := global.Db.Delete(m)
	return result.Error

}
func(s *RoleService) findOne(rolename string) (*model.AdminRole,error) {
	m := model.AdminRole{}
	err := global.Db.Debug().Where("role_name = ?",rolename).First(&m).Error
	if err == nil && m.RoleId > 0   {
		return nil,errs.ErrorRoleExist
	}
	return &m,nil
}

func(s *RoleService) List(param form.Page) (manager2.RoleList,error) {
	ss := []model.AdminRole{}
	m := manager2.RoleList{}
	var total int64
	page := param.Page
	size := param.Size
	q := param.Param
	var err error
	if q != "" {
		err = global.Db.Debug().Where("role_name like ?","%" +q+ "%").Limit(size).Offset((page - 1) * size).Find(&ss).Limit(-1).Offset(-1).Count(&total).Error
	}else {
		err = global.Db.Debug().Limit(size).Offset((page - 1) * size).Find(&ss).Limit(-1).Offset(-1).Count(&total).Error
	}

	fmt.Println(err)
	m.AdminRoles = ss
	m.Size = size
	m.Page = page
	m.Total = total
	return m,nil
}