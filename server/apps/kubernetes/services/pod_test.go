package services

import (
	"context"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/cmd/app/option"
	"github.com/lbemi/lbemi/pkg/util"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"testing"
	"time"
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

func TestPod_ExecPodReadString(f *testing.T) {
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
	pod := NewPod(cliStore, "", nil)

	dirs := pod.ExecPodReadString(context.Background(), "default", "nginx-2-598f88c6dc-f7fb8", "nginx-2")
	item := util.GetDirAndFiles(dirs)
	println("item:", item[0].Name)
	//for _, dir := range pod.ExecPodReadString(context.Background(), "default", "nginx-2-598f88c6dc-f7fb8", "nginx-2", []string{"ls", "-l"}) {
	//	println(string(dir))
	//}
}

// TestPod_ExecPodOnce is a test function that executes a pod once.
//
// It takes a testing.T object as a parameter.
// It does not return anything.
func TestPod_ExecPodOnce(f *testing.T) {
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
	pod := NewPod(cliStore, "", nil)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFunc()
	pod.ExecPodOnce(ctx, "default", "nginx2-7cc8cd4598-f66ws", "nginx", []string{"mv", "test2.txt", "test.txt"})
}
