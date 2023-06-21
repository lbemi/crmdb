import request from '@/utils/request';

export function useConfigMapApi() {
	return {
		createConfigMap: (param: any, data: any) => {
			return request({
				url: '/configmaps',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updateConfigMap: (param: any, data: any) => {
			return request({
				url: '/configmaps',
				method: 'put',
				params: param,
				data: data,
			});
		},
		listConfigMap: (namespace: string, param: any) => {
			return request({
				url: '/configmaps/namespaces/' + namespace,
				method: 'get',
				params: param,
			});
		},
		deleteConfigMap: (namespace: string, name: string, param: any) => {
			return request({
				url: `/configmaps/namespaces/${namespace}/${name}`,
				method: 'delete',
				params: param,
			});
		},
	};
}
