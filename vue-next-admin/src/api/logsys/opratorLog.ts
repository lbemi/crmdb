import request from '@/utils/request';

export function useOperatorLogApi() {
	return {
		listOperatorLog: (query: any) => {
			return request({
				url: '/logs/operator',
				method: 'get',
				params: query,
			});
		},
		getOperatorLog: (id: number) => {
			return request({
				url: '/logs/operator' + id,
				method: 'get',
			});
		},
		deleteOperatorLog: (query: any) => {
			return request({
				url: '/logs/operator',
				method: 'delete',
				params: query,
			});
		},
		deleteAllOperatorLog: () => {
			return request({
				url: '/logs/operator/all',
				method: 'delete',
			});
		},
	};
}
