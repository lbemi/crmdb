package types

import (
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Secret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Immutable         *bool             `json:"immutable,omitempty"`
	Data              map[string][]byte `json:"data,omitempty"`
	StringData        map[string]string `json:"stringData,omitempty"`
	Type              corev1.SecretType `json:"type,omitempty"`
}

type ServiceWorkLoad struct {
	Deployments  []*v1.Deployment  `json:"deployments"`
	DaemonSets   []*v1.DaemonSet   `json:"daemonSets"`
	StatefulSets []*v1.StatefulSet `json:"statefulSets"`
	EndPoints    *corev1.Endpoints `json:"endPoints"`
	Events       []*corev1.Event   `json:"events"`
}
