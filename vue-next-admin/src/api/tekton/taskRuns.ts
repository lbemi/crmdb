import request from '@/utils/request';

export function useTektonTaskRunsApi() {
	return {
		createTaskRun: (param: any, data: any) => {
			return request({
				url: '/taskruns',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updateTaskRun: (param: any, data: any) => {
			return request({
				url: '/taskrunss',
				method: 'put',
				params: param,
				data: data,
			});
		},
		listTaskRun: (namespace: string, param: any) => {
			return request({
				url: '/taskruns/namespaces/' + namespace,
				method: 'get',
				params: param,
			});
		},
		deleteTaskRun: (namespace: string, name: string, param: any) => {
			return request({
				url: `/taskruns/namespaces/${namespace}/${name}`,
				method: 'delete',
				params: param,
			});
		},
	};
}
