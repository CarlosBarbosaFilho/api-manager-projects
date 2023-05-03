package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProject(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Create project"})
}

func GetProject(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Get project"})
}
func Projects(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Get all projects"})
}

func DeleteProject(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete project"})
}

func UpdateProject(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update project"})
}
