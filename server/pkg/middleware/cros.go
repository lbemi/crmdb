package middleware

import (
	"github.com/emicklei/go-restful/v3"
)

// Cors 处理跨域请求
func Cors(container *restful.Container) *restful.CrossOriginResourceSharing {
	return &restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowedHeaders: []string{"Content-Type", "AccessToken", "X-CSRF-Token", "Authorization", "Token", "X-Token", "X-User-Id"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "DELETE", "PUT", "PATCH"},
		CookiesAllowed: false,
		Container:      container}
}
