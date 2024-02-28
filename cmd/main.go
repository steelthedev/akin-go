package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/steelthedev/akin-go/connection"
	"github.com/steelthedev/akin-go/pkg"
)

func main() {
	err := godotenv.Load("../local.env")
	if err != nil {
		log.Fatal("Could not load env")
	}

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome back to go akinwumi",
		})
	})

	dbUrl := string(os.Getenv("dbUrl"))
	dbHandler := connection.Init(dbUrl)

	pkg.RegisterRoutes(router, dbHandler)
	router.Run(":8000")
}
