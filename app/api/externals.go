package api

import "github.com/gin-gonic/gin"

type ApiHandler interface {
	SetupRoutes(rg *gin.RouterGroup)
}
