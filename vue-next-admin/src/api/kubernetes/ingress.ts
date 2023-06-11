import request from '/@/utils/request';

export function useIngressApi() {
	return {
		listIngress: (namespace: string, param: any) => {
			return request({
				url: '/ingress/namespaces/' + namespace,
				method: 'get',
				params: param,
			});
		},
		deleteIngress: (param: any, name: string, namespace: string) => {
			return request({
				url: `/ingress/namespaces/${namespace}/${name}`,
				method: 'delete',
				params: param,
			});
		},
	};
}
