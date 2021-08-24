package main

import (
	"github.com/Evolt0/gw-kit/cmd/gw-kit/global"
	"github.com/Evolt0/gw-kit/pkg/apis"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var monitoredSignals = []os.Signal{
	syscall.SIGHUP,
	syscall.SIGINT,
	syscall.SIGTERM,
	syscall.SIGQUIT,
}

func main() {
	config := &global.Config{}
	config.Init()
	logrus.Println(config)
	gin.SetMode(config.App.Mode)
	engine := gin.Default()
	apis.SetRoutes(engine)
	go func() {
		err := engine.Run(config.App.Port)
		logrus.Fatalf("failed to run project! %v", err)
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, monitoredSignals...)
	select {
	case <-quit:
		logrus.Println("exit...")
	}
}
