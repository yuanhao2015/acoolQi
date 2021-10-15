package main

import (
	"acoolqi-admin/config"
	_ "acoolqi-admin/dao/system"
	"acoolqi-admin/pkg/middleware/logger"
	"acoolqi-admin/router"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	port, mode string
)

func init() {
	flag.StringVar(&port, "port", "8080", "server listening on, default 8080")
	flag.StringVar(&mode, "mode", "debug", "server running mode, default debug mode")
}

func main() {
	port := config.GetServerCfg().Port
	flag.Parse()
	gin.SetMode(mode)
	r := router.Init()
	r.Use(logger.LoggerToFile())
	err := r.Run(port)
	if err != nil {
		log.Fatalf("Start server: %+v", err)
	}
}
