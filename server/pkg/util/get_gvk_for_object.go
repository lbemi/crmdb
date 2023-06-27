package util

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func GVKForObject(obj runtime.Object, scheme *runtime.Scheme) (schema.GroupVersionKind, error) {
	_, ok := obj.(*metav1.PartialObjectMetadata)
	_, ok2 := obj.(*metav1.PartialObjectMetadataList)

	// 如果存在部分metadata 信息会尝试通过, runtime.Object 自带的接口去获取GVK
	if ok || ok2 {
		gvk := obj.GetObjectKind().GroupVersionKind()

		if len(gvk.Kind) == 0 {
			return schema.GroupVersionKind{}, runtime.NewMissingKindErr("unstructured object has no kind")
		}
		if len(gvk.Version) == 0 {
			return schema.GroupVersionKind{}, runtime.NewMissingVersionErr("unstructured object has no version")
		}
		return gvk, nil
	}

	gvks, isUnversioned, err := scheme.ObjectKinds(obj)
	if err != nil {
		return schema.GroupVersionKind{}, err
	}
	if isUnversioned {
		return schema.GroupVersionKind{}, fmt.Errorf("cannot create group-version-kind for unversioned type %T", obj)
	}

	if len(gvks) < 1 {
		return schema.GroupVersionKind{}, fmt.Errorf("no group-version-kinds associated with type %T", obj)
	}
	if len(gvks) > 1 {
		return schema.GroupVersionKind{}, fmt.Errorf(
			"multiple group-version-kinds associated with type %T, refusing to guess at one", obj)
	}
	return gvks[0], nil
}

//func RestoreGVKForList(objects []runtime.Object) error {
//	if len(objects) != 0 {
//		gvk, err := GVKForObject(objects[0], scheme.Scheme)
//		if err != nil {
//			return err
//		}
//		for _, obj := range objects {
//			obj.GetObjectKind()
//		}
//	}
//}
