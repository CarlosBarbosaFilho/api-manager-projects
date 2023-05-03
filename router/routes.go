package router

import (
	"github.com/CarlosBarbosaFilho/api-manager-projects/handler"
	"github.com/gin-gonic/gin"
)

func home(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/projects", handler.Projects)
		v1.GET("/projects/{id}", handler.GetProject)
		v1.POST("/projects", handler.CreateProject)
		v1.PUT("/projects", handler.UpdateProject)
		v1.DELETE("/projects/{id}", handler.DeleteProject)
	}

}
