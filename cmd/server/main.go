package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nk-hung/ecom-go/internal/routes"
)

func main1() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoute(router)
	routes.AuthRoute(router)

	router.GET("/apt-v1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-v1"})
	})

	router.GET("api-v2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for apt-v2"})
	})

	router.Run(":" + port)
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MySQL:: Successfully Connected")
}
