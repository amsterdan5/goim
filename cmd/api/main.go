package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/amsterdan/goim/conf"
	"github.com/amsterdan/goim/httpserv"
	"github.com/amsterdan/goim/httpserv/route"
	"github.com/amsterdan/goim/internal/initial"
	"github.com/amsterdan/goim/pkg/logs"
)

var configFile = flag.String("config", "./config/app.yaml", "配置文件")

func main() {
	flag.Parse()

	ctx := context.WithValue(context.Background(), logs.LogField("tid"), time.Now().UnixMicro())

	// 读取配置
	cfg, err := conf.NewConf(*configFile)
	if err != nil {
		panic(err)
	}

	// 初始化日志
	initial.Init(ctx, cfg)

	// http服务
	app, err := httpserv.New(cfg)
	if err != nil {
		panic(err)
	}

	api := route.NewApi(ctx, app.HTTPServer)

	// 退出信号
	exitChan := make(chan bool, 1)

	logs.Ctx(ctx).Info(fmt.Sprintf("服务【%s】启动成功, :%d", cfg.AppName, cfg.HttpServerConf.Port))
	go func() {
		// 启动服务
		if err := app.HTTPServer.Start(api); err != nil {
			logs.Ctx(ctx).Error(fmt.Sprintf("服务【%s】启动失败, :%d", cfg.AppName, cfg.HttpServerConf.Port), err)
			exitChan <- true
		}
	}()

	safeExit(ctx, exitChan, app.HTTPServer)
}

// 优雅退出
func safeExit(ctx context.Context, exit chan bool, serv *httpserv.Server) {
	c := make(chan os.Signal, 1)

	// 接受信号，kill -9 无法监听
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGTTOU)

	select {
	case <-exit:
		os.Exit(1)
	case cc := <-c:
		log.Println("接收到关闭信号,", cc.String())
	}

	cancelCtx, cancel := context.WithTimeout(ctx, time.Duration(10))
	defer cancel()

	graceStopChan := make(chan bool, 1)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := serv.Stop(cancelCtx); err != nil {
			panic(err)
		}
	}()

	go func() {
		wg.Wait()
		graceStopChan <- true
	}()

	select {
	case <-cancelCtx.Done():
		log.Println("服务正常退出")
		os.Exit(1)
	case <-graceStopChan:
		log.Println("非正常退出")
		os.Exit(0)
	}
}
