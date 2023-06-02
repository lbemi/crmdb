package server

import (
	"context"
	"fmt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"net"
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
	h.welcomeMsg()
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

func getIP() *[]string {
	ips := make([]string, 0)
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						ips = append(ips, ipnet.IP.String())
					}
				}
			}
		}
	}
	return &ips
}

func (h *HttpSever) welcomeMsg() {
	ips := getIP()
	msg := `
----------------------------------------------------
                  欢迎使用GO-OPS
----------------------------------------------------
服务器监听地址：
`
	for _, i := range *ips {
		msg = msg + "  http://" + i + h.Addr
	}

	msg += `
swagger:
`
	for _, i := range *ips {
		msg = msg + "  http://" + i + h.Addr + "/apidocs.json"
	}
	msg += `
----------------------------------------------------
`
	log.Logger.Infof(msg)
}
