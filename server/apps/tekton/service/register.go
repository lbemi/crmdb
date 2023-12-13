package services

import "github.com/lbemi/lbemi/pkg/cache"

type TektonGetter interface {
	Tekton() TektonInterface
}

type TektonInterface interface {
	TaskGetter
}

type Tekton struct {
	clusterName string
	store       *cache.ClientStore
}

func NewTekton(clusterName string, store *cache.ClientStore) *Tekton {
	return &Tekton{clusterName: clusterName, store: store}
}

func (t *Tekton) Tasks(namespace string) ITask {
	if namespace == "all" {
		namespace = ""
	}
	return NewTask(t.store.Get(t.clusterName), namespace)
}
