import request from '/@/utils/request';

export function useNodeApi() {
	return {
		listNode: (param: any) => {
			return request({
				url: '/node',
				method: 'get',
				params: param,
			});
		},

		getNode: (name: string, param: any) => {
			return request({
				url: `/node/${name}`,
				method: 'get',
				params: param,
			});
		},

		updateLabel: (param: any, data: any) => {
			return request({
				url: '/node',
				method: 'patch',
				params: param,
				data: data,
			});
		},

		schedulable: (param: any, name: string, unschedulable: boolean) => {
			return request({
				url: `/node/${name}/${unschedulable}`,
				method: 'put',
				params: param,
			});
		},
		drainNode: (name: string, param: any) => {
			return request({
				url: `/node/${name}/drain`,
				method: 'post',
				params: param,
			});
		},
	};
}
