package initial

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"

	"github.com/amsterdan/goim/conf"
	"github.com/amsterdan/goim/pkg/logs"
)

func Init(ctx context.Context, c conf.Config) {
	initLog(c.AppName, c.Log.LogDir, c.Log.FileType)

	if c.Pprof.Open {
		go initpprof(ctx, c.Pprof.Port)
	}
}

// 初始化日志
func initLog(app, logDir, filetype string) {
	if filetype == "console" {
		logs.Init(os.Stdout, os.Stderr, os.Stderr, logs.LogField("tid"))
		return
	}

	if filetype == "file" {
		if err := os.MkdirAll(filepath.Dir(logDir), 0777); err != nil {
			panic(err)
		}

		infoLog := fmt.Sprintf("%s%s.log", logDir, app)

		infof, err := os.Create(infoLog)
		if err != nil {
			panic(err)
		}

		logs.Init(infof, infof, infof, logs.LogField("tid"))
		return
	}

	panic("未知日志类型")
}

// 开启性能
func initpprof(ctx context.Context, port int) {
	logs.Ctx(ctx).Info("开启pprof性能检测, ", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		logs.Ctx(ctx).Error("启动失败", err)
	}
}
