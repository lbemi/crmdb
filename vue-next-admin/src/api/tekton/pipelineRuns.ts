import request from '@/utils/request';

export function useTektonPipelineRunsApi() {
	return {
		createPipelinerun: (param: any, data: any) => {
			return request({
				url: '/pipelineruns',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updatePipelinerun: (param: any, data: any) => {
			return request({
				url: '/pipelineruns',
				method: 'put',
				params: param,
				data: data,
			});
		},
		listPipelinerun: (namespace: string, param: any) => {
			return request({
				url: '/pipelineruns/namespaces/' + namespace,
				method: 'get',
				params: param,
			});
		},
		deletePipelinerun: (namespace: string, name: string, param: any) => {
			return request({
				url: `/pipelineruns/namespaces/${namespace}/${name}`,
				method: 'delete',
				params: param,
			});
		},
	};
}
