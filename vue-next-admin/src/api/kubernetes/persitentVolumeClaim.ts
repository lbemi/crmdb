import request from '/@/utils/request';

export function usePVCApi() {
	return {
		listPVC: (namespace: string, param: any) => {
			return request({
				url: '/pvc/' + namespace,
				method: 'get',
				params: param,
			});
		},
	};
}
