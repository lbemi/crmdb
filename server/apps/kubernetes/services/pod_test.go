package services

import (
	"context"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/cmd/app/option"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"testing"
)

func TestPod_CopyFromPod(t *testing.T) {
	completedOptions := option.NewOptions().WithConfig("../../../dev.yaml").WithLog().Complete()
	println(completedOptions.Config.Log.Level)
	configFile := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", configFile)
	if err != nil {
		panic(err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	cliStore := &cache.ClientConfig{ClientSet: clientSet, Config: config}
	pod := NewPod(cliStore, "default", nil)
	_ = pod.CopyFromPod(context.Background(), "default", "nginx2-7cc8cd4598-f66ws", "nginx", "test.txt", ".")
}
