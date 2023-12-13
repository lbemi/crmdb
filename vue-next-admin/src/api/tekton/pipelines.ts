import request from '@/utils/request';

export function useTektonPipelinesApi() {
	return {
		createPipeline: (param: any, data: any) => {
			return request({
				url: '/pipelines',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updatePipeline: (param: any, data: any) => {
			return request({
				url: '/pipelines',
				method: 'put',
				params: param,
				data: data,
			});
		},
		listPipeline: (namespace: string, param: any) => {
			return request({
				url: '/pipelines/namespaces/' + namespace,
				method: 'get',
				params: param,
			});
		},
		deletePipeline: (namespace: string, name: string, param: any) => {
			return request({
				url: `/pipelines/namespaces/${namespace}/${name}`,
				method: 'delete',
				params: param,
			});
		},
	};
}
