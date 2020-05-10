package dto

import "test.blog.com/testBlog/model"

type UserDto struct {
	Name  string `json:"name"`
	Phone string `json:phone`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:  user.Name,
		Phone: user.Phone,
	}
}
