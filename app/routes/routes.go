package routes

import (
	"injar/controllers/categories"
	"injar/controllers/users"
	"injar/controllers/webinars"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware        middleware.JWTConfig
	UserController       users.UserController
	CategoriesController categories.CategoryController
	WebinarController    webinars.WebinarController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	//Auth ...
	auth := e.Group("v1/api/auth")

	auth.POST("/register", cl.UserController.Store)
	auth.POST("/login", cl.UserController.Login)

	//Categories ...
	category := e.Group("v1/api/categories")
	category.Use(middleware.JWTWithConfig(cl.JWTMiddleware))

	category.GET("", cl.CategoriesController.GetAll)
	category.GET("/id/:id", cl.CategoriesController.FindById)
	category.POST("", cl.CategoriesController.Store)
	category.PUT("/id/:id", cl.CategoriesController.Update)
	category.DELETE("/id/:id", cl.CategoriesController.Delete)

	//Webinars ...
	webinar := e.Group("v1/api/webinars")
	webinar.Use(middleware.JWTWithConfig(cl.JWTMiddleware))

	webinar.GET("", cl.WebinarController.GetAll)
	webinar.GET("/id/:id", cl.WebinarController.FindById)
	webinar.POST("", cl.WebinarController.Store)
	webinar.PUT("/id/:id", cl.WebinarController.Update)
	webinar.DELETE("/id/:id", cl.WebinarController.Delete)

}