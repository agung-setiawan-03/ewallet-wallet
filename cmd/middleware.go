package cmd

import (
	"ewallet-wallet/helpers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareRefreshToken(c *gin.Context) {

	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("authorization empty")
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
		c.Abort()
		return
	}

	// c.Set("token")

	c.Next()
}
