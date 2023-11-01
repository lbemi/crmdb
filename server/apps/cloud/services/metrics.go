package services

import "sync"

type ResourceMetrics struct {
	data sync.Map
}

func (r *ResourceMetrics) GetMetrics() {

}
