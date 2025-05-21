package services

import (
	"net/http"
	"net/url"

	"github.com/lbemi/lbemi/pkg/cache"
	"k8s.io/apimachinery/pkg/util/proxy"
	"k8s.io/client-go/rest"
)

const (
	proxyBaseURL = "/lbemi/proxy"
)

type ProxyGetter interface {
	Proxy() IProxy
}

type Proxy struct {
	config      *cache.ClientConfig
	clusterName string
}
type IProxy interface {
	GET(path url.URL, w http.ResponseWriter, r *http.Request) error
}

func NewProxy(config *cache.ClientConfig, clusterName string) IProxy {
	return &Proxy{config: config, clusterName: clusterName}
}

func (p *Proxy) GET(path url.URL, w http.ResponseWriter, r *http.Request) error {

	transport, err := rest.TransportFor(p.config.Config)
	if err != nil {
		return err
	}

	target, err := p.parseTarget(path, p.config.Config.Host, p.clusterName)
	if err != nil {
		return err
	}

	httpPorxy := proxy.NewUpgradeAwareHandler(target, transport, true, false, nil)
	httpPorxy.UpgradeTransport = proxy.NewUpgradeRequestRoundTripper(transport, transport)
	httpPorxy.ServeHTTP(w, r)
	return nil

}

func (p *Proxy) parseTarget(target url.URL, host string, name string) (*url.URL, error) {
	kubeURL, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	target.Path = target.Path[len(proxyBaseURL+"/"+name):]

	target.Host = kubeURL.Host
	target.Scheme = kubeURL.Scheme
	return &target, nil
}
