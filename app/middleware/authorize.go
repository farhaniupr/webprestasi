package middleware

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthorizeRequest is used to authorize a request for a certain end-point group.
func AuthorizeRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		//gin.BasicAuth()
		v := session.Get("user-id")
		if v == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Please login."})
			c.Abort()
		}
		c.Next()
	}
}
