package handler

import (
	"net/http"

	"github.com/CarlosBarbosaFilho/api-manager-projects/config"
	"github.com/gin-gonic/gin"
)

func getLoggerInternal() *config.Logger {
	logger := *config.GetLogger("Initialize calling to handlers")
	return &logger
}

func CreateProject(ctx *gin.Context) {
	logger := getLoggerInternal()
	logger.Info("Start creating project")
	ctx.JSON(http.StatusOK, gin.H{"message": "Create project"})
	logger.Info("End creating project")
}

func GetProject(ctx *gin.Context) {
	logger := getLoggerInternal()
	logger.Info("Return project")
	ctx.JSON(http.StatusOK, gin.H{"message": "Get project"})
	logger.Infof("Project returned: %v")
}
func Projects(ctx *gin.Context) {
	logger := getLoggerInternal()
	logger.Info("Listing all project")
	ctx.JSON(http.StatusOK, gin.H{"message": "Get all projects"})
}

func DeleteProject(ctx *gin.Context) {
	logger := getLoggerInternal()
	logger.Info("Start deleting project")
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete project"})
	logger.Info("End deleting project")
}

func UpdateProject(ctx *gin.Context) {
	logger := getLoggerInternal()
	logger.Info("Start updating project")
	ctx.JSON(http.StatusOK, gin.H{"message": "Update project"})
	logger.Info("End updating project")
}
