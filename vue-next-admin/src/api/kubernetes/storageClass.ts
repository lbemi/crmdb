import request from '@/utils/request';

export function useStorageClassApi() {
	return {
		createStorageClass: (param: any, data: any) => {
			return request({
				url: '/storageclasses',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updateStorageClass: (param: any, data: any) => {
			return request({
				url: '/storageclasses',
				method: 'put',
				params: param,
				data: data,
			});
		},
		listStorageClass: (param: any) => {
			return request({
				url: '/storageclasses',
				method: 'get',
				params: param,
			});
		},
		deleteStorageClass: (name: string, param: any) => {
			return request({
				url: `/storageclasses/${name}`,
				method: 'delete',
				params: param,
			});
		},
	};
}
