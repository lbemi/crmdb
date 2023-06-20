import { Service } from 'kubernetes-types/core/v1';
import request from '@/utils/request';

export function useServiceApi() {
	return {
		listService: (namespace: string, param: any) => {
			return request({
				url: '/services/namespaces/' + namespace,
				method: 'get',
				params: param,
			});
		},
		getService: (namespace: string, name: string, param: any) => {
			return request({
				url: `/services/namespaces/${namespace}/${name}`,
				method: 'get',
				params: param,
			});
		},
		listServiceWorkLoad: (namespace: string, name: string, param: any) => {
			return request({
				url: `/services/namespaces/${namespace}/workload/${name}`,
				method: 'get',
				params: param,
			});
		},
		deleteService: (param: any, name: string, namespace: string) => {
			return request({
				url: `/services/namespaces/${namespace}/${name}`,
				method: 'delete',
				params: param,
			});
		},
		updateService: (param: any, data: Service) => {
			return request({
				url: `/services`,
				method: 'put',
				params: param,
				data: data,
			});
		},
		createService: (param: any, data: Service) => {
			return request({
				url: `/services`,
				method: 'post',
				params: param,
				data: data,
			});
		},
	};
}
