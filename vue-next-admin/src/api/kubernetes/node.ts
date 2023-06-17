import request from '@/utils/request';

export function useNodeApi() {
	return {
		listNode: (param: any) => {
			return request({
				url: '/nodes',
				method: 'get',
				params: param,
			});
		},
		listPodByNode: (nodeName: string, param: any) => {
			return request({
				url: `/nodes/${nodeName}/pods`,
				method: 'get',
				params: param,
			});
		},

		getNode: (name: string, param: any) => {
			return request({
				url: `/nodes/${name}`,
				method: 'get',
				params: param,
			});
		},
		getNodeEvents: (name: string, param: any) => {
			return request({
				url: `/nodes/${name}/events`,
				method: 'get',
				params: param
			});
		},
		updateLabel: (param: any, data: any) => {
			return request({
				url: '/nodes/label',
				method: 'patch',
				params: param,
				data: data,
			});
		},

		schedulable: (param: any, name: string, unschedulable: boolean) => {
			return request({
				url: `/nodes/${name}/${unschedulable}`,
				method: 'put',
				params: param,
			});
		},
		drainNode: (name: string, param: any) => {
			return request({
				url: `/nodes/${name}/drain`,
				method: 'post',
				params: param,
			});
		},
	};
}
