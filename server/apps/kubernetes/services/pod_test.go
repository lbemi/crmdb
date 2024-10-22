package services

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/cmd/app/option"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
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
	_ = NewPod(cliStore, "default", nil)
	// _ = pod.CopyFromPod(context.Background(), "default", "nginx2-7cc8cd4598-f66ws", "nginx", "test.txt", ".")
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

	dirs := pod.GetFileList(context.Background(), "default", "nginx-2-598f88c6dc-f7fb8", "nginx-2", "/")
	//println(dirs)
	f.Log(dirs)
	//println("item:", item[0].Name)
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
	pod.execPodOnce(ctx, "default", "nginx2-7cc8cd4598-f66ws", "nginx", []string{"mv", "test2.txt", "test.txt"})
}
