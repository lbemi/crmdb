import request from '@/utils/request';

export function usePVCApi() {
	return {
		listPVC: (namespace: string, param: any) => {
			return request({
				url: '/persistentvolumeclaims/namespaces/' + namespace,
				method: 'get',
				params: param,
			});
		},
	};
}
