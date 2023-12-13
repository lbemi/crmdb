import request from '@/utils/request';

export function useTektonTasksApi() {
	return {
		createTask: (param: any, data: any) => {
			return request({
				url: '/tasks',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updateTask: (param: any, data: any) => {
			return request({
				url: '/tasks',
				method: 'put',
				params: param,
				data: data,
			});
		},
		listTask: (namespace: string, param: any) => {
			return request({
				url: '/tasks/namespaces/' + namespace,
				method: 'get',
				params: param,
			});
		},
		deleteTask: (namespace: string, name: string, param: any) => {
			return request({
				url: `/tasks/namespaces/${namespace}/${name}`,
				method: 'delete',
				params: param,
			});
		},
	};
}
