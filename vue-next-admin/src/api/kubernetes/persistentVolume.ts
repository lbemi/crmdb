import request from '@/utils/request';

export function usePersistentVolumeApi() {
	return {
		createPersistentVolume: (param: any, data: any) => {
			return request({
				url: '/persistentvolumes',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updatePersistentVolume: (param: any, data: any) => {
			return request({
				url: '/persistentvolumes',
				method: 'put',
				params: param,
				data: data,
			});
		},
		listPersistentVolume: (param: any) => {
			return request({
				url: '/persistentvolumes',
				method: 'get',
				params: param,
			});
		},
		deletePersistentVolume: (name: string, param: any) => {
			return request({
				url: `/persistentvolumes/${name}`,
				method: 'delete',
				params: param,
			});
		},
	};
}
