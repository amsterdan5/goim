package route

import (
	"context"

	"github.com/amsterdan/goim/httpserv"
	"github.com/amsterdan/goim/internal/api/im"
)

type iApi struct {
	ctx    context.Context
	server *httpserv.Server
}

func NewApi(ctx context.Context, s *httpserv.Server) iApi {
	return iApi{
		ctx:    ctx,
		server: s,
	}
}

func (a iApi) Route() {
	mux := a.server.ServerMux()

	httpserv.Route(mux, "GET", "/register", im.Register{Ctx: a.ctx, Server: a.server})
}
