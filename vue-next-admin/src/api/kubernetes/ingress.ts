import request from '/@/utils/request';

export function useIngressApi() {
	return {
		listIngress: (namespace: string, param: any) => {
			return request({
				url: '/ingress/' + namespace,
				method: 'get',
				params: param,
			});
		},
		deleteIngress: (param: any, name: string, namespace: string) => {
			return request({
				url: `/ingress/${namespace}/${name}`,
				method: 'delete',
				params: param,
			});
		},
	};
}
