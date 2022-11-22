package form

type ClusterReq struct {
	Name       string `json:"name"`
	KubeConfig string `json:"kube_config"`
}
