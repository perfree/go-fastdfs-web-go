package form

import "go-fastdfs-web-go/src/models"

type UserForm struct {
	Id              int `binding:"required"`
	Account         string
	Password        string `binding:"required,max=16"`
	Name            string `binding:"required,max=100"`
	Email           string `binding:"required,email"`
	NewPassword	    string `binding:"required,max=16"`
}

func (u *UserForm)GetUser() models.User {
	return models.User{
		Id: u.Id,
		Account: u.Account,
		Password: u.Password,
		Name: u.Name,
		Email: u.Email,
	}
}