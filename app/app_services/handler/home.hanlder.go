package handler

import (
	"github.com/gin-gonic/gin"
	"kafka-demoset/app/internal/dto"
)

type homeHandler struct {
	R *gin.Engine
}

func NewHomeHandler(h *Handler) *homeHandler {
	return &homeHandler{
		R: h.R,
	}
}

func (ctr *homeHandler) Register() {
	ctr.R.POST("/", ctr.welcome)
}

func (ctr *homeHandler) welcome(c *gin.Context) {
	res := &dto.ResponseObject{}

	res.ErrCode = 0
	res.ErrMsg = "success"
	res.Data = gin.H{
		"data": "welcome from kafka test!",
	}
	c.JSON(200, res)
}
