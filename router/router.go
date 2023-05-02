package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	fmt.Println("Initializer Application !")
	router := gin.Default()
	home(router)
	router.Run(":8888")

}
