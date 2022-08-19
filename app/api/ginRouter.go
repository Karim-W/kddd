package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karim-w/kddd/docs"
	"github.com/karim-w/kddd/infra/middlewares"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
// @schemes   http

//Add your controllers here
func SetupRoutes(handlers ...ApiHandler) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.Cors())
	router.LoadHTMLFiles("index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))
	v1 := router.Group("api/v1")
	v1.Use(middlewares.TransactionIdGenerator())
	for _, handler := range handlers {
		handler.SetupRoutes(v1)
	}
	router.Run()
	return router
}
