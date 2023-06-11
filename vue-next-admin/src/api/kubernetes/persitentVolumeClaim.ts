import request from '/@/utils/request';

export function usePVCApi() {
	return {
		listPVC: (namespace: string, param: any) => {
			return request({
				url: '/pvcs/namespaces/' + namespace,
				method: 'get',
				params: param,
			});
		},
	};
}
