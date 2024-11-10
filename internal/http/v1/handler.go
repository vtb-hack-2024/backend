package v1

import (
	"backend-vtb/internal/service"
	"backend-vtb/pkg/auth"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service      service.Base
	tokenManager auth.TokenManager
}

func NewHandler(service service.Base, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		service:      service,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initInfoRouter(v1)
	}
}
