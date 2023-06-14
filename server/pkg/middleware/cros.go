package middleware

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cross 跨域处理
func Cross() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "think-lang", "server"}
	config.AllowCredentials = true
	config.ExposeHeaders = []string{"New-Token", "New-Expires-In", "Content-Disposition"}

	return cors.New(config)

}

// Cors 处理跨域请求
func Cors(container *restful.Container) *restful.CrossOriginResourceSharing {
	cors := &restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowedHeaders: []string{"Content-Type", "AccessToken", "X-CSRF-Token", "Authorization", "Token", "X-Token", "X-User-Id"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "DELETE", "PUT", "PATCH"},
		CookiesAllowed: false,
		Container:      container}

	return cors
}
