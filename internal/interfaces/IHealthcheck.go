package interfaces

import "github.com/gin-gonic/gin"

type IHealthcheckServices interface {
	HealthcheckServices() (string, error)
}

type IHealthcheckAPI interface {
	HelathcheckHandlerHTTP(c *gin.Context)
}

type IHealthcheckRepo interface {
}
