package userService

import (
	"github.com/saravana-1010/go-business-calculator/domain/users"
	"github.com/saravana-1010/go-business-calculator/utils/errors"
)

var (
	UserService = &userService{}
)

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, errors.RestErr)
	GetUser(int64) (*users.User, errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, errors.RestErr)
	DeleteUser(int64) errors.RestErr
	LoginUser(users.LoginRequest) (*users.User, errors.RestErr)
	GetAllUsers() ([]*users.User, errors.RestErr)
}

type userService struct {}

func (service *userService) CreateUser(user users.User) (*users.User, errors.RestErr) {
	
}

func (service *userService) GetUser(int64) (*users.User, errors.RestErr) {
	
}

func (service *userService) UpdateUser(bool, users.User) (*users.User, errors.RestErr) {

}
	
func (service *userService) DeleteUser(int64) errors.RestErr {

}
	
func (service *userService) LoginUser(users.LoginRequest) (*users.User, errors.RestErr) {

}

func (service *userService) GetAllUsers ([]*users.User, errors.RestErr)  {
	
}