package service

import "gf-init/app/dao"

// 中间件管理服务
var User = userService{}

type userService struct{}

// 用户注册
func (s *userService) GetAllUsers() error {
	dao.GetUsers()
	return nil
}
