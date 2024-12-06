package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Initialize() {

	r := gin.Default()
	initializeRoutes(r)

	log.Println("server starting port:8080")
	r.Run(":8080")
}
