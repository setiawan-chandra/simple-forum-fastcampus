package memberships

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/setiawan-chandra/simple-forum-fastcampus/internal/model/memberships"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
}

type Handler struct {
	*gin.Engine

	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, memberbershipSvc membershipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: memberbershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("memberships")
	route.GET("/ping", h.Ping)
	route.POST("/signup", h.SignUp)
}
