package entity

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Node struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec          corev1.NodeSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status        corev1.NodeStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
	Usage         `json:"usage"`
}

type Usage struct {
	Pod    int     `json:"pod"`
	Cpu    float64 `json:"cpu"`
	Memory float64 `json:"memory"`
}
