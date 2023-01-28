package handler

import (
	"github.com/gin-gonic/gin"
	libkafka "kafka-demoset/app/_lib/kafka"
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

type TestReq struct {
	Message string `json:"message"`
}

func (ctr *homeHandler) welcome(c *gin.Context) {
	res := &dto.ResponseObject{}
	req := &TestReq{}
	if err := c.ShouldBind(&req); err != nil {
		res.ErrCode = 401
		res.ErrMsg = err.Error()
		c.JSON(200, res)
		return
	}

	libkafka.TestKafkaMessageEvent(req.Message)

	res.ErrCode = 0
	res.ErrMsg = "success"
	res.Data = gin.H{
		"data": "welcome from kafka test!",
	}
	c.JSON(200, res)
}
