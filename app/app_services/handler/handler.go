package handler

import "github.com/gin-gonic/gin"

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
