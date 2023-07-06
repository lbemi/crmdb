import request from '@/utils/request';

export function useHBAApi() {
	return {
		lisGroup: (query: any) => {
			return request({
				url: '/groups',
				method: 'get',
				params: query,
			});
		},
		updateGroup: (id: number, data: any) => {
			return request({
				url: '/groups/' + id,
				method: 'put',
				data: data,
			});
		},
		addGroup: (data: any) => {
			return request({
				url: '/groups/register',
				method: 'post',
				data: data,
			});
		},
		deleteGroup: (id: number) => {
			return request({
				url: '/groups/' + id,
				method: 'delete',
			});
		},
	};
}
