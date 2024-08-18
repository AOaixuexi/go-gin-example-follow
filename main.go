package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/AOaixuexi/go-gin-example-follow/models"
	"github.com/AOaixuexi/go-gin-example-follow/pkg/logging"

	"github.com/AOaixuexi/go-gin-example-follow/pkg/setting"
	"github.com/AOaixuexi/go-gin-example-follow/routers"
	"github.com/fvbock/endless"
)

// 不关闭现有连接（正在运行中的程序）

// 新的进程启动并替代旧进程

// 新的进程接管新的连接

// 连接要随时响应用户的请求，当用户仍在请求旧进程时要保持连接，新用户应请求新进程，不可以出现拒绝请求的情况

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
