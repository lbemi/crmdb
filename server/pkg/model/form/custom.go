package form

import v1 "k8s.io/api/core/v1"

type PageResult struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}

type PageConfigMap struct {
	Data  []*v1.ConfigMap `json:"data"`
	Total int64           `json:"total"`
}

type PatchNode struct {
	Name   string                 `json:"name"`
	Labels map[string]interface{} `json:"labels"`
}
