package types

import (
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Secret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Immutable         *bool             `json:"immutable,omitempty"`
	Data              map[string][]byte `json:"data,omitempty"`
	StringData        map[string]string `json:"stringData,omitempty"`
	Type              v12.SecretType    `json:"type,omitempty"`
}
