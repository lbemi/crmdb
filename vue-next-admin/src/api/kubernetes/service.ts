import request from '/@/utils/request';

export function useServiceApi() {
	return {
		listService: (namespace: string, param: any) => {
			return request({
				url: '/service/' + namespace,
				method: 'get',
				params: param,
			});
		},
	};
}
