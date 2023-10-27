import request from '@/utils/request';

export function useVirtualServiceApi() {
	return {
		createVirtualService: (param: any, data: any) => {
			return request({
				url: '/virtualservices',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updateVirtualService: (param: any, data: any) => {
			return request({
				url: '/virtualservices',
				method: 'put',
				params: param,
				data: data,
			});
		},
		listVirtualService: (namespace: string, param: any) => {
			return request({
				url: '/virtualservices/namespaces/' + namespace,
				method: 'get',
				params: param,
			});
		},
		deleteVirtualService: (namespace: string, name: string, param: any) => {
			return request({
				url: `/virtualservices/namespaces/${namespace}/${name}`,
				method: 'delete',
				params: param,
			});
		},
	};
}
