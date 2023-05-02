package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/projects", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Hello App"})
		})

		v1.GET("/projects/{id}", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Hello App"})
		})

		v1.POST("/projects", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Hello App"})
		})

		v1.PUT("/projects", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Hello App"})
		})

		v1.DELETE("/projects/{id}", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Hello App"})
		})

	}

}
