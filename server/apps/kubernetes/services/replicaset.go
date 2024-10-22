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

type ReplicasetGetter interface {
	Replicaset(namespace string) ReplicasetImp
}

type ReplicasetImp interface {
	List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult
	Get(ctx context.Context, name string) *appsv1.ReplicaSet
}

type Replicaset struct {
	cli *cache.ClientConfig
	ns  string
}

func NewReplicaset(cli *cache.ClientConfig, namespace string) *Replicaset {
	return &Replicaset{cli: cli, ns: namespace}
}

func (r *Replicaset) List(ctx context.Context, query *entity.PageParam, name string, label string) *entity.PageResult {
	data, err := r.cli.SharedInformerFactory.Apps().V1().ReplicaSets().Lister().ReplicaSets(r.ns).List(labels.Everything())
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	res := &entity.PageResult{}
	var replicasetList = make([]*appsv1.ReplicaSet, 0)
	if name != "" {
		for _, item := range data {
			if strings.Contains(item.Name, name) {
				replicasetList = append(replicasetList, item)
			}
		}
		data = replicasetList
	}

	if label != "" {
		for _, item := range data {
			if strings.Contains(labels.FormatLabels(item.Labels), label) {
				replicasetList = append(replicasetList, item)
			}
		}
		data = replicasetList
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

func (r *Replicaset) Get(ctx context.Context, name string) *appsv1.ReplicaSet {
	replicaSet, err := r.cli.SharedInformerFactory.Apps().V1().ReplicaSets().Lister().ReplicaSets(r.ns).Get(name)
	restfulx.ErrNotNilDebug(err, restfulx.GetResourceErr)
	return replicaSet
}

type ReplicasetHandler struct {
	client      *cache.ClientConfig
	clusterName string
}

func NewReplicasetHandler(client *cache.ClientConfig, clusterName string) *ReplicasetHandler {
	return &ReplicasetHandler{client: client, clusterName: clusterName}
}

func (r *ReplicasetHandler) OnAdd(obj interface{}) {
	r.notifyReplicaset(obj)
}

func (r *ReplicasetHandler) OnUpdate(oldObj, newObj interface{}) {
	r.notifyReplicaset(newObj)
}

func (r *ReplicasetHandler) OnDelete(obj interface{}) {
	r.notifyReplicaset(obj)
}

func (r *ReplicasetHandler) notifyReplicaset(obj interface{}) {
	namespace := obj.(*appsv1.ReplicaSet).Namespace
	replicates, err := r.client.SharedInformerFactory.Apps().V1().ReplicaSets().Lister().ReplicaSets(namespace).List(labels.Everything())
	if err != nil {
		global.Logger.Error(err)
	}

	//按时间排序
	sort.Slice(replicates, func(i, j int) bool {
		return replicates[j].ObjectMeta.GetCreationTimestamp().Time.Before(replicates[i].ObjectMeta.GetCreationTimestamp().Time)
	})
	//fmt.Println(r.clusterName, "-----这个空间-----发生数据变化------------")
	go cache.WebsocketStore.SendClusterResource(r.clusterName, "replicaset", map[string]interface{}{
		"cluster": r.clusterName,
		"type":    "replicaset",
		"result": map[string]interface{}{
			"namespace": namespace,
			"data":      replicates,
		},
	})
}
