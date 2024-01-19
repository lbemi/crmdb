package services

import (
	"context"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"sort"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type ControllerRevisionGetter interface {
	ControllerRevision(namespace string) ControllerRevisionImp
}

type ControllerRevisionImp interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult
	Get(ctx context.Context, name string) *appsv1.ControllerRevision
}

type ControllerRevision struct {
	cli *cache.ClientConfig
	ns  string
}

func NewControllerRevision(cli *cache.ClientConfig, namespace string) *ControllerRevision {
	return &ControllerRevision{cli: cli, ns: namespace}
}

func (r *ControllerRevision) List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult {
	data, err := r.cli.SharedInformerFactory.Apps().V1().ControllerRevisions().Lister().ControllerRevisions(r.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &entity.PageResult{}
	var controllerRevisionList = make([]*appsv1.ControllerRevision, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				controllerRevisionList = append(controllerRevisionList, item)
			}
		}
		data = controllerRevisionList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				controllerRevisionList = append(controllerRevisionList, item)
			}
		}
		data = controllerRevisionList
	}
	//按时间排序
	sort.SliceStable(data, func(i, j int) bool {
		return data[j].ObjectMeta.GetCreationTimestamp().Time.Before(data[i].ObjectMeta.GetCreationTimestamp().Time)
	})
	total := len(data)
	// 未传递分页查询参数
	if query.Limit == 0 && query.Page == 0 {
		res.Data = data
	} else {
		if total <= query.Limit {
			res.Data = data
		} else if query.Page*query.Limit >= total {
			res.Data = data[(query.Page-1)*query.Limit : total]
		} else {
			res.Data = data[(query.Page-1)*query.Limit : query.Page*query.Limit]
		}
	}
	res.Total = int64(total)
	return res
}

func (r *ControllerRevision) Get(ctx context.Context, name string) *appsv1.ControllerRevision {
	controllerRevision, err := r.cli.SharedInformerFactory.Apps().V1().ControllerRevisions().Lister().ControllerRevisions(r.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return controllerRevision
}

type ControllerRevisionHandler struct {
	client      *cache.ClientConfig
	clusterName string
}

func NewControllerRevisionHandler(client *cache.ClientConfig, clusterName string) *ControllerRevisionHandler {
	return &ControllerRevisionHandler{client: client, clusterName: clusterName}
}

func (r *ControllerRevisionHandler) OnAdd(obj interface{}) {
	r.notifyControllerRevision(obj)
}

func (r *ControllerRevisionHandler) OnUpdate(oldObj, newObj interface{}) {
	r.notifyControllerRevision(newObj)
}

func (r *ControllerRevisionHandler) OnDelete(obj interface{}) {
	r.notifyControllerRevision(obj)
}

func (r *ControllerRevisionHandler) notifyControllerRevision(obj interface{}) {
	namespace := obj.(*appsv1.ReplicaSet).Namespace
	replicates, err := r.client.SharedInformerFactory.Apps().V1().ControllerRevisions().Lister().ControllerRevisions(namespace).List(labels.Everything())
	if err != nil {
		global.Logger.Error(err)
	}

	//按时间排序
	sort.Slice(replicates, func(i, j int) bool {
		return replicates[j].ObjectMeta.GetCreationTimestamp().Time.Before(replicates[i].ObjectMeta.GetCreationTimestamp().Time)
	})
	//fmt.Println(r.clusterName, "-----这个空间-----发生数据变化------------")
	go cache.WebsocketStore.SendClusterResource(r.clusterName, "controllerRevision", map[string]interface{}{
		"cluster": r.clusterName,
		"type":    "controllerRevision",
		"result": map[string]interface{}{
			"namespace": namespace,
			"data":      replicates,
		},
	})
}
