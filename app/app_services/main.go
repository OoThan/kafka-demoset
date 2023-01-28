package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	libkafka "kafka-demoset/app/_lib/kafka"
	"kafka-demoset/app/app_services/handler"
	_ "kafka-demoset/app/conf"
	"kafka-demoset/app/internal/logger"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := flag.String("port", "8001", "default port is 8001")
	flag.Parse()

	addr := net.JoinHostPort("", *port)

	libkafka.InitKafkaProducer()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20

	h := handler.NewHandler(&handler.HConfig{
		R: router,
	})
	h.Register()

	server := http.Server{
		Addr:    addr,
		Handler: h.R,
	}

	go func() {
		logger.Sugar.Info("http server started Addr: ", addr)
		if err := server.ListenAndServe(); err != nil {
			logger.Sugar.Error(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c

	if err := server.Shutdown(context.Background()); err != nil {
		logger.Sugar.Error(err)
	}
}
