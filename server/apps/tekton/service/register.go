package services

import "github.com/lbemi/lbemi/pkg/cache"

type TektonGetter interface {
	Tekton() TektonInterface
}

type TektonInterface interface {
	TaskGetter
	TaskRunGetter
	PipelineGetter
	PipelineRunGetter
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

func (t *Tekton) TaskRuns(namespace string) ITaskRun {
	if namespace == "all" {
		namespace = ""
	}
	return NewTaskRun(t.store.Get(t.clusterName), namespace)
}

func (t *Tekton) Pipelines(namespace string) IPipeline {
	if namespace == "all" {
		namespace = ""
	}
	return NewPipeline(t.store.Get(t.clusterName), namespace)
}

func (t *Tekton) PipelineRuns(namespace string) IPipelineRun {
	if namespace == "all" {
		namespace = ""
	}
	return NewPipelineRun(t.store.Get(t.clusterName), namespace)
}
