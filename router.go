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

	// e.GET("/account/:account_number", controller.GetAccount)
	e.POST("/login", controllers.GetToken)
	e.GET("/getprofile", controllers.GetProfile)
	
	return e
}