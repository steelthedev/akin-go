package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) AddProject(ctx *gin.Context) {
	var project Project

	requestBody := AddProjectBody{}

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Request body not valid",
		})

		return
	}

	if result := h.DB.Create(&project); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error occured",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Project created successfully",
	})
}

func (h handler) GetProjects(ctx *gin.Context) {

	var projects []Project

	if result := h.DB.Find(&projects); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Could not load all projects",
		})
		return
	}

	ctx.IndentedJSON(http.StatusAccepted, &projects)
}
