package server

const (
	HTTP Type = "HTTP"
	GRPC Type = "GRPC"
)

type Type string

type Server interface {
	Type() Type
	Start() error
	Stop() error
}
