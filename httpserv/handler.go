package httpserv

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/amsterdan/goim/pkg/logs"
)

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)                  // 执行服务
	GetParams(r *http.Request) (*ReqParam, error)                      // 获取参数
	ParseBody(ctx context.Context, params *ReqParam) (RespBody, error) // 处理请求
}

type Header struct {
	Header http.Header
}

type ReqParam struct {
	Header   *Header
	queryMap map[string]string // uri请求参数
	body     []byte            // 请求体
	addr     string            // 请求地址
}

type RespBody string

func (r ReqParam) Headers() http.Header {
	return r.Header.Header
}

func (r ReqParam) GetQuery() map[string]string {
	return r.queryMap
}

func (r ReqParam) Body() []byte {
	return r.body
}

func (r ReqParam) Addr() string {
	return r.addr
}

// 处理请求
func ServerHTTP(w http.ResponseWriter, r *http.Request, h Handler) {
	statusCode := http.StatusOK
	ctx := r.Context()

	startime := time.Now()

	var body RespBody
	var header *Header

	params, err := h.GetParams(r)
	if err != nil {
		statusCode = http.StatusBadRequest
	} else {
		header = params.Header
		body, err = h.ParseBody(ctx, params)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
	}

	w.WriteHeader(statusCode)
	w.Write([]byte(body))

	logs.Ctx(ctx).WithField("cost time", time.Since(startime).String()).
		WithFields("header", header.Header, "method", r.Method, "uri", r.URL.String(), "body", body)
}

// 解析参数
func GetParam(r *http.Request, keys ...string) (*ReqParam, error) {
	return checkParams(r, keys...)
}

func checkParams(r *http.Request, keys ...string) (*ReqParam, error) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		return nil, err
	}

	query := make(map[string]string)

	for _, k := range keys {
		query[k] = r.FormValue(k)
	}

	params := &ReqParam{
		Header: &Header{
			Header: r.Header,
		},
		queryMap: query,
		addr:     r.RemoteAddr,
	}
	return params, nil
}

// 解析body体
func GetParamBody(r *http.Request) (*ReqParam, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	params := &ReqParam{
		Header: &Header{
			Header: r.Header,
		},
		body: body,
		addr: r.RemoteAddr,
	}
	return params, nil
}
