package core

const (
	HTTP Type = "HTTP"
	GRPC Type = "GRPC"
)

type Type string

type Server interface {
	ServerType() Type
	Start()
	Stop() error
}
