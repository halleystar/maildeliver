package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitApi() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	initEmail(r)
	fmt.Println("hahah")
	r.Run(":8080")
}
