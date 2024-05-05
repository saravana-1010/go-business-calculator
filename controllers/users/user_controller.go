package users

import (
	"github.com/gin-gonic/gin"
	"github.com/saravana-1010/go-business-calculator/services/userService"
)

func CreateUser(c *gin.Context)  {
	userService.UserService.CreateUser()
}

func GetUser(c *gin.Context)  {
	userService.UserService.GetUser()
}

func LoginUser(c *gin.Context)  {
	userService.UserService.LoginUser()
}

func DeleteUser(c *gin.Context)  {
	userService.UserService.DeleteUser()
}

func GetAllUsers(c *gin.Context)  {
	userService.UserService.GetAllUsers()
}

func UpdateUser(c *gin.Context)  {
	userService.UserService.UpdateUser()
}