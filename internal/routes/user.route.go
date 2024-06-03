package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/nk-hung/ecom-go/internal/controllers"
	"github.com/nk-hung/ecom-go/internal/middleware"
)

// var UserRoute = func(router *mux.Router) {
var UserRoute = func(router *gin.Engine) {
	router.Use(middleware.Authentication())
	router.GET("/users", controller.GetUsers())
}
