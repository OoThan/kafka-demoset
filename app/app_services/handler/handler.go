package handler

import (
	"github.com/gin-gonic/gin"
	"kafka-demoset/app/internal/middlerware"
)

type Handler struct {
	R *gin.Engine
}

type HConfig struct {
	R *gin.Engine
}

func NewHandler(c *HConfig) *Handler {
	return &Handler{
		R: c.R,
	}
}

func (h *Handler) Register() {

	h.R.Use(middlerware.Cors())

	// home handler
	homeHandler := NewHomeHandler(h)
	homeHandler.Register()
}
