package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "asdsadasdas")
	})
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})
	r := router.Group("/user")
	setApiUserGroupRoutes(r)

}

func setApiUserGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "asdsadasdas")
	})
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})
}
