package server

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful/v3"
)

type HttpSever struct {
	Addr string
	srv  *http.Server

	Container *restful.Container
}

func NewHttpSever(addr string) *HttpSever {
	container := restful.NewContainer()
	container.EnableContentEncoding(true)
	//restful.TraceLogger(&httpLog{})
	//restful.SetLogger(&httpLog{})
	return &HttpSever{Addr: addr, Container: container, srv: &http.Server{
		Addr:    addr,
		Handler: container,
	}}
}

func (h *HttpSever) Type() Type {
	return HTTP
}

func (h *HttpSever) Start() error {
	log.Logger.Infof("HTTP Server listen: %s", h.Addr)
	go func() {
		if err := h.srv.ListenAndServe(); err != nil {
			log.Logger.Infof("error http serve: %s", err)
		}
	}()

	return nil
}

func (h *HttpSever) Stop() error {
	// 延迟5s停止
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return h.srv.Shutdown(ctx)
}

func (h *HttpSever) RegisterRoutes(routes ...*restful.WebService) {
	for _, route := range routes {
		h.Container.Add(route)
	}
}

type httpLog struct{}

func (t *httpLog) Print(v ...any) {
	log.Logger.Info(v...)
}

func (t *httpLog) Printf(format string, v ...any) {
	log.Logger.Infof(format, v...)
}
