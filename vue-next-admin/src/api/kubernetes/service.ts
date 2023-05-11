import { V1Service } from '@kubernetes/client-node';
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
		deleteService: (param: any, name: string, namespace: string) => {
			return request({
				url: `/service/${namespace}/${name}`,
				method: 'delete',
				params: param,
			});
		},
		updateService: (param: any, data: V1Service) => {
			return request({
				url: `/service`,
				method: 'put',
				params: param,
				data: data,
			});
		},
	};
}
