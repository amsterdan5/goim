package httpserv

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/amsterdan/goim/conf"
	"github.com/amsterdan/goim/internal/connector"
	"github.com/amsterdan/goim/pkg/logs"
)

type Server struct {
	cfg        conf.HttpServerConf
	httpServer *http.Server
	serverMux  *http.ServeMux
	connector  *connector.Connector
}

// 新建http server
func NewServer(cfg conf.HttpServerConf, conn *connector.Connector) (*Server, error) {
	serveMux := http.NewServeMux()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: serveMux,
		ConnContext: func(ctx context.Context, c net.Conn) context.Context {
			return context.WithValue(context.Background(), logs.LogField("tid"), time.Now().UnixMicro())
		},
	}

	return &Server{
		cfg:        cfg,
		httpServer: server,
		serverMux:  serveMux,
		connector:  conn,
	}, nil
}

// 开启服务
func (s *Server) Start(r Router) error {
	// 加载路由
	r.Route()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Port))
	if err != nil {
		return err
	}

	return s.httpServer.Serve(listen)
}

// 关闭服务
func (s *Server) Stop(ctx context.Context) error {
	// 关闭http
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	// 关闭链接
	if err := s.connector.Close(); err != nil {
		return err
	}
	return nil
}

// http配置
func (s *Server) Cfg() conf.HttpServerConf {
	return s.cfg
}

// http路由
func (s *Server) HttpHandler() http.Handler {
	return s.httpServer.Handler
}

// 全局配置
func (s *Server) Connector() *connector.Connector {
	return s.connector
}

func (s *Server) ServerMux() *http.ServeMux {
	return s.serverMux
}
