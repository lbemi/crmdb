import request from '@/utils/request';

export function useSecretApi() {
	return {
		createSecret: (param: any, data: any) => {
			return request({
				url: '/secrets',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updateSecret: (param: any, data: any) => {
			return request({
				url: '/secrets',
				method: 'put',
				params: param,
				data: data,
			});
		},
		listSecret: (namespace: string, param: any) => {
			return request({
				url: '/secrets/namespaces/' + namespace,
				method: 'get',
				params: param,
			});
		},
		deleteSecret: (namespace: string, name: string, param: any) => {
			return request({
				url: `/secrets/namespaces/${namespace}/${name}`,
				method: 'delete',
				params: param,
			});
		},
	};
}
