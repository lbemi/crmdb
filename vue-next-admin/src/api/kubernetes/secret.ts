import request from '/@/utils/request';

export function useSecretApi() {
	return {
		listSecret: (namespace: string, param: any) => {
			return request({
				url: '/secret/' + namespace,
				method: 'get',
				params: param,
			});
		},
		deleteSecret: (namespace: string,name: string, param: any) => {
			return request({
				url: `/secret/${namespace}/${name}`,
				method: 'delete',
				params: param,
			})
	}
}
}
