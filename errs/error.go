package errs

import "errors"

var (
	ErrorUserNotExist = errors.New("用户不存在")
	ErrorUserPassword = errors.New("密码错误")
	ErrorUserAlreadyExist = errors.New("用户已存在")

	ErrorRoleExist = errors.New("角色已存在")
	ErrorRoleIdNotExist = errors.New("id 不存在")
	ErrorRoleNotExist = errors.New("角色不存在")
)
