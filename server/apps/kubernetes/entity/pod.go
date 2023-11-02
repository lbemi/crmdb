package entity

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Pod struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec          corev1.PodSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status        corev1.PodStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
	MemoryUsage   float64          `json:"memory_usage"`
	CpuUsage      float64          `json:"cpu_usage"`
}
