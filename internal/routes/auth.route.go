package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/nk-hung/ecom-go/internal/controllers"
)

// var AuthRoute = func(router *mux.Router) {
var AuthRoute = func(router *gin.Engine) {
	// router.HandleFunc("/auth", controllers.Login)
	router.POST("users/signup", controller.Signup())
	router.POST("users/login", controller.Login())
}
