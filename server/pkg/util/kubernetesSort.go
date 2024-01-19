package util

//func SortKubernetesResource(podList []runtime.Object) {
//	sort.Slice(podList, func(i, j int) bool {
//		kind := podList[i].GetObjectKind()
//		// sort by creation timestamp in descending order
//		if podList[j].ObjectMeta.GetCreationTimestamp().Time.Before(podList[i].ObjectMeta.GetCreationTimestamp().Time) {
//			return true
//		} else if podList[i].ObjectMeta.GetCreationTimestamp().Time.Before(podList[j].ObjectMeta.GetCreationTimestamp().Time) {
//			return false
//		}
//
//		// if the creation timestamps are equal, sort by name in ascending order
//		return podList[i].ObjectMeta.GetName() < podList[j].ObjectMeta.GetName()
//	})
//}
