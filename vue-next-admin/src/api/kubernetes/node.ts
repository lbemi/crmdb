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
		updateLabel: (param: any, data: any) => {
			return request({
				url: '/node',
				method: 'patch',
				params: param,
				data: data,
			});
		},
	};
}
