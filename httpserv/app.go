package httpserv

import (
	"github.com/amsterdan/goim/conf"
	"github.com/amsterdan/goim/internal/connector"
)

type App struct {
	Cfg        conf.Config
	HTTPServer *Server
}

// 新建容器
func New(cfg conf.Config) (*App, error) {
	conn, err := connector.New(cfg)
	if err != nil {
		return nil, err
	}

	server, err := NewServer(cfg.HttpServerConf, conn)
	if err != nil {
		return nil, err
	}

	return &App{
		Cfg:        cfg,
		HTTPServer: server,
	}, nil
}
