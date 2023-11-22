package helper

import (
	"io"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
)

type fileSpec struct {
	PodName      string
	PodNamespace string
	File         pathSpec
}

type pathSpec interface {
	String() string
}

type CopyOptions struct {
	Container  string
	Namespace  string
	NoPreserve bool
	MaxTries   int

	ClientConfig      *restclient.Config
	ClientSet         kubernetes.Interface
	ExecParentCmdName string

	args []string
}

type TarPipe struct {
	src       fileSpec
	o         *CopyOptions
	reader    *io.PipeReader
	outStream *io.PipeWriter
	bytesRead uint64
	retries   int
}
