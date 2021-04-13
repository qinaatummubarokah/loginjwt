package main

import (
	"loginjwt/controllers"
	"loginjwt/config"

	"github.com/labstack/echo"

)

func NewRouter() *echo.Echo {
	e := echo.New()

	// Initialize main database
	config.Db = config.Connect()

	e.POST("/login", controllers.GetToken)
	e.GET("/getprofile", controllers.GetProfile)
	e.POST("/register", controllers.Register)
	e.PUT("/updateuser", controllers.UpdateUser)
	
	return e
}