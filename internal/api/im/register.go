package im

import (
	"context"
	"net/http"

	"github.com/amsterdan/goim/httpserv"
)

type Register struct {
	Ctx    context.Context
	Server *httpserv.Server
}

func (res Register) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpserv.ServerHTTP(w, r, res)
}

func (reg Register) GetParams(r *http.Request) (*httpserv.ReqParam, error) {
	return httpserv.GetParamBody(r)
}

func (r Register) ParseBody(ctx context.Context, params *httpserv.ReqParam) (httpserv.RespBody, error) {
	return httpserv.RespBody("test"), nil
}
