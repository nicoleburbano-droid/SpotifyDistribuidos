package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS: el cliente se abre como archivo (file://), así que todas las peticiones
// son de otro origen. Se devuelve el mismo Origin con credenciales porque
// SockJS (reacciones) hace sus peticiones XHR con credenciales. Incluye las
// cabeceras que necesita gRPC-Web.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		origen := c.Request.Header.Get("Origin")
		if origen == "" {
			origen = "*"
		} else {
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Vary", "Origin")
		}
		c.Header("Access-Control-Allow-Origin", origen)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "content-type, keep-alive, user-agent, cache-control, content-transfer-encoding, x-accept-content-transfer-encoding, x-accept-response-streaming, x-user-agent, x-grpc-web, grpc-timeout")
		c.Header("Access-Control-Expose-Headers", "grpc-status, grpc-message")
		c.Header("Access-Control-Max-Age", "1728000")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
