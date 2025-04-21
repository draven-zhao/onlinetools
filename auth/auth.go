package auth

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

const (
	validUser = "admin"
	validPass = "hou147258"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, hasAuth := c.Request.BasicAuth()
		if !hasAuth || user != validUser || pass != validPass {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}