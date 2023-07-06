import request from '@/utils/request';

export function useAccountApi() {
	return {
		listAccount: (query: any) => {
			return request({
				url: '/accounts',
				method: 'get',
				params: query,
			});
		},
		updateAccount: (id: number, data: any) => {
			return request({
				url: '/accounts/' + id,
				method: 'put',
				data: data,
			});
		},
		addAccount: (data: any) => {
			return request({
				url: '/accounts/register',
				method: 'post',
				data: data,
			});
		},
		deleteAccount: (id: number) => {
			return request({
				url: '/accounts/' + id,
				method: 'delete',
			});
		},
	};
}
