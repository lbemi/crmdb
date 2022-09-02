package models

type HttpClient struct {
	TimeOut int
	MaxIdle int
	Error   func(error)
}
type ClientOption func(client *HttpClient)

type ClientOptions []ClientOption

func (opts ClientOptions) apply(c *HttpClient) {
	for _, opt := range opts {
		opt(c)
	}
}

func NewHttpClient(opts ...ClientOption) *HttpClient {
	c := &HttpClient{}
	ClientOptions(opts).apply(c)
	return c
}

func (h *HttpClient) DO(url string) {
	println("do")
}

func WithTimeOut(timeout int) ClientOption {
	return func(client *HttpClient) {
		client.TimeOut = timeout
	}
}
func WithMaxIdle(maxidle int) ClientOption {
	return func(client *HttpClient) {
		client.MaxIdle = maxidle
	}
}
