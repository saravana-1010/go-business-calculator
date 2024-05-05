package app

import (
	"github.com/saravana-1010/go-business-calculator/controllers/ping"
	"github.com/saravana-1010/go-business-calculator/controllers/users"
)

func mapUrls()  {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.POST("/users/login", users.LoginUser)
}