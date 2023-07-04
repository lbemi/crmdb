import request from '@/utils/request';

export function useHostApi() {
	return {
		lisHost: (query: any) => {
			return request({
				url: '/hosts',
				method: 'get',
				params: query,
			});
		},
		updateHost: (id: number, data: any) => {
			return request({
				url: '/hosts/' + id,
				method: 'put',
				data: data,
			});
		},
		addHost: (data: any) => {
			return request({
				url: '/hosts/register',
				method: 'post',
				data: data,
			});
		},
		deleteHost: (id: number) => {
			return request({
				url: '/hosts/' + id,
				method: 'delete',
			});
		},
	};
}
