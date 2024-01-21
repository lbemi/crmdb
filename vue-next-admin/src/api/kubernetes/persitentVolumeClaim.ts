import request from '@/utils/request';

export function usePersistentVolumeClaimApi() {
	return {
		createPersistentVolumeClaim: (param: any, data: any) => {
			return request({
				url: '/persistentvolumeclaims',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updatePersistentVolumeClaim: (param: any, data: any) => {
			return request({
				url: '/persistentvolumeclaims',
				method: 'put',
				params: param,
				data: data,
			});
		},
		listPersistentVolumeClaim: (namespace: string, param: any) => {
			return request({
				url: '/persistentvolumeclaims/namespaces/' + namespace,
				method: 'get',
				params: param,
			});
		},
		getPersistentVolumeClaimsByStorageClassName: (param: any) => {
			return request({
				url: '/persistentvolumeclaims',
				method: 'get',
				params: param,
			});
		},
		deletePersistentVolumeClaim: (namespace: string, name: string, param: any) => {
			return request({
				url: `/persistentvolumeclaims/namespaces/${namespace}/${name}`,
				method: 'delete',
				params: param,
			});
		},
	};
}
