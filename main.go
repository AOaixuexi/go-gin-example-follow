package main

import (
	"fmt"
	"syscall"

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
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		logging.Info("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		logging.Info("Server err: %v", err)
	}
}
