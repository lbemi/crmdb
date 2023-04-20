import request from '/@/utils/request';

export function useConfigMapApi() {
	return {
		listConfigMap: (namespace: string, param: any) => {
			return request({
				url: '/configmap/' + namespace,
				method: 'get',
				params: param,
			});
		},
	};
}