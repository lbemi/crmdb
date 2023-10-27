package form

import (
	"istio.io/client-go/pkg/apis/networking/v1beta1"
	v1 "k8s.io/api/core/v1"
)

type PageResult struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}

type PageConfigMap struct {
	Data  []*v1.ConfigMap `json:"data"`
	Total int64           `json:"total"`
}

type PatchNode struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
}
type PageVirtualService struct {
	Data  []*v1beta1.VirtualService `json:"data"`
	Total int64                     `json:"total"`
}
