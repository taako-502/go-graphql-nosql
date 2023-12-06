package config

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS
func SettingCors(hosts []string) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: hosts,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodPut,
			http.MethodOptions,
		},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	})
}
