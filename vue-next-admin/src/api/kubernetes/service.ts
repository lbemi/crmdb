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
		deleteService: ( param: any,name: string, namespace: string) => {
			return request({
				url: `/service/${namespace}/${name}`,
				method: 'delete',
				params: param,
			});
		},
	};
}
