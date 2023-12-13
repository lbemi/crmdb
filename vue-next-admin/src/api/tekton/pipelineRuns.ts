import request from '@/utils/request';

export function useTektonPipelineRunsApi() {
	return {
		createPipeline: (param: any, data: any) => {
			return request({
				url: '/pipelineruns',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updatePipeline: (param: any, data: any) => {
			return request({
				url: '/pipelineruns',
				method: 'put',
				params: param,
				data: data,
			});
		},
		listPipeline: (namespace: string, param: any) => {
			return request({
				url: '/pipelineruns/namespaces/' + namespace,
				method: 'get',
				params: param,
			});
		},
		deletePipeline: (namespace: string, name: string, param: any) => {
			return request({
				url: `/pipelineruns/namespaces/${namespace}/${name}`,
				method: 'delete',
				params: param,
			});
		},
	};
}
