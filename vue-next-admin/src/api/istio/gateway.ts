import request from '@/utils/request';

export function useGatewayApi() {
	return {
		createGateway: (param: any, data: any) => {
			return request({
				url: '/gateways',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updateGateway: (param: any, data: any) => {
			return request({
				url: '/gateways',
				method: 'put',
				params: param,
				data: data,
			});
		},
		listGateway: (namespace: string, param: any) => {
			return request({
				url: '/gateways/namespaces/' + namespace,
				method: 'get',
				params: param,
			});
		},
		deleteGateway: (namespace: string, name: string, param: any) => {
			return request({
				url: `/gateways/namespaces/${namespace}/${name}`,
				method: 'delete',
				params: param,
			});
		},
	};
}
