package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karim-w/kddd/domain/health"
)

type HealthCheckHandler struct {
	svc *health.HealthService
}

func NewHealthCheckHandler(svc *health.HealthService) *HealthCheckHandler {
	return &HealthCheckHandler{
		svc: svc,
	}
}

func (h *HealthCheckHandler) SetupRoutes(rg *gin.RouterGroup) {
	rg.GET("/health", h.HealthCheck)
}

// @BasePath /
// HealthCheck GET
// @Summary Api to check the health of the service
// @Schemes
// @Description Api will return 200 if the service is healthy
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} contracts.HealthResponse
// @Failure 500 {object} contracts.ErrorResponse
// @Router /health [get]
func (h *HealthCheckHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, h.svc.GetHealth())
}
