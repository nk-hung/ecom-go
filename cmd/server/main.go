package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nk-hung/ecom-go/internal/routes"
)

func main() {
	// cfg := mysql.Config{
	// 	User:                 config.Envs.DBUser,
	// 	Passwd:               config.Envs.DBPassword,
	// 	Addr:                 config.Envs.DBAddress,
	// 	DBName:               config.Envs.DBName,
	// 	Net:                  "tcp",
	// 	AllowNativePasswords: true,
	// 	ParseTime:            true,
	// }
	//
	// db, err := global.NewMySQLStorage(cfg)
	// initStorage(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }
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
	// r := mux.NewRouter()
	// r.Use(gin.Logger())
	//
	// /* Routes */
	// routes.AuthRoute(r)
	// routes.UserRoute(r)
	// routes.RegisterBookStoreRoutes(r)
	//
	// http.Handle("/", r)
	// log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MySQL:: Successfully Connected")
}
