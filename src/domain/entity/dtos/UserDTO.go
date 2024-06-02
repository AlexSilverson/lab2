package dtos

import (
	"AlexSilverson/lab2/src/domain/entity"
	"fmt"

	"gorm.io/gorm"
)

type UserDto struct {
	Login    string `json:"login" validate:"required,email"`
	Password string `json:"password" validate:"required" `
}

func (u UserDto) MapNewUserDtoToUser() entity.User {
	var ans entity.User
	ans.Model = gorm.Model{}
	ans.Login = u.Login

	ans.HashPas = u.Password
	ans.Role = 0
	fmt.Println(ans)
	return ans
}
