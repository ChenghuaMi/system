package service

import (
	"system/errs"
	"system/form/login"
	"system/model"
	"system/global"
	"system/util"
	"fmt"
)

type LoginService struct {

}
//var LoginServiceObj = new(LoginService)

func (s *LoginService) Login(form login.AdminForm) (*model.Admin,error) {
	return s.findOne(form)
}
func(s *LoginService) findOne(form login.AdminForm) (*model.Admin,error){
	var model model.Admin
	username := form.Username
	err := global.Db.Where("username = ?",username).First(&model).Error
	fmt.Println(err)
	if err != nil {
		return nil,errs.ErrorUserNotExist
	}
	fmt.Println(model)
	password := util.Md5String(form.Password)
	if password != model.Password {
		return nil,errs.ErrorUserPassword
	}
	fmt.Println(model)
	return &model,nil
}
