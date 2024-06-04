package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nk-hung/ecom-go/internal/controllers"
)

// var UserRoute = func(router *mux.Router) {
var UserRoute = func(router *gin.Engine) {
	// router.Use(middleware.Authenticate())
	// router.GET("/users", controllers.GetUsers())
	router.GET("/users/:user_id", controllers.GetUser())
}
