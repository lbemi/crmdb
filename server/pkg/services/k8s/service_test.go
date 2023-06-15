package k8s

import (
	"context"
	"fmt"
	"github.com/lbemi/lbemi/pkg/cmd"
	"testing"
	"time"
)

func TestService_ListWorkLoad(t *testing.T) {

	const CONFIGURE = "/Users/lei/Documents/GitHub/lbemi/server/dev.yaml"
	factory := cmd.TestInit(CONFIGURE)
	// 初始化已存在的kubernetes集群client
	//go bootstrap.LoadKubernetes(factory)

	//list := factory.Cluster().List()
	//fmt.Println(list)
	time.Sleep(3 * time.Second)

	c := newService(factory.Cluster().GetClient("311"), "mongo")
	load := c.ListWorkLoad(context.Background(), "mongodb-svc")
	fmt.Println(load)
}
