package pkg

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	routes := router.Group("api")
	routes.POST("/projects/add", h.AddProject)
	routes.GET("/projects", h.GetProjects)
}
