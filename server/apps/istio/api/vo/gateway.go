package vo

import "istio.io/client-go/pkg/apis/networking/v1beta1"

type PageGateway struct {
	Data  []*v1beta1.Gateway `json:"data"`
	Total int64              `json:"total"`
}
