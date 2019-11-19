package model

import (
	database "github.com/Jiahonzheng/CloudGo-IO/db"
	"github.com/Jiahonzheng/CloudGo-IO/dto"
	"github.com/Jiahonzheng/CloudGo-IO/errors"
)

type User struct {
	StuId    string `json:"stu_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (u *User) SignUp(dto *dto.UserSignUp) (error, *User) {
	db := database.Get()
	stuId, email, username, phone := dto.StuId, dto.Email, dto.Username, dto.Phone
	if stuId == "" {
		return errors.ErrInvalidStuId, nil
	}
	if username == "" {
		return errors.ErrInvalidUsername, nil
	}
	if email == "" {
		return errors.ErrInvalidEmail, nil
	}
	if phone == "" {
		return errors.ErrInvalidPhone, nil
	}
	if _, ok := db.Load(stuId); ok {
		return errors.ErrUserExists, nil
	}
	user := &User{
		StuId:    stuId,
		Username: username,
		Email:    email,
		Phone:    phone,
	}
	db.Store(stuId, user)
	return nil, user
}

func (u *User) GetByStuId(stuId string) (error, *User) {
	db := database.Get()
	var user *User
	db.Range(func(key, value interface{}) bool {
		if key != stuId {
			return true
		}
		user = value.(*User)
		return false
	})
	if user != nil {
		return nil, user
	}
	return errors.ErrUserDoesNotExist, user
}
