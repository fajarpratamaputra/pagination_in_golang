//Routes/Routes.go
package routes

import (
	"general/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())

	v3 := r.Group("/v3")
	grp1 := v3.Group("/general")
	grp1.GET("/promo", controller.GetPromo)
	return r
}
