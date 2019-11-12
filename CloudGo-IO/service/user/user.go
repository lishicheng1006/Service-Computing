package user

import (
	"github.com/Jiahonzheng/CloudGo-IO/dto"
	"github.com/Jiahonzheng/CloudGo-IO/model"
)

type Service struct{}

var userModel = new(model.User)

func (s *Service) SignUp(dto *dto.UserSignUp) (error, *model.User) {
	return userModel.SignUp(dto)
}

func (s *Service) GetUserInfo(stuId string) (error, *model.User) {
	return userModel.GetByStuId(stuId)
}
