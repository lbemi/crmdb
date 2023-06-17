import request from '@/utils/request';

export function useUserApi() {
	return {
		listUser: (query: any) => {
			return request({
				url: '/users',
				method: 'get',
				params: query,
			});
		},
		updateUser: (id: number, data: any) => {
			return request({
				url: '/users/' + id,
				method: 'put',
				data: data,
			});
		},
		addUser: (data: any) => {
			return request({
				url: '/users/register',
				method: 'post',
				data: data,
			});
		},
		deleteUser: (id: number) => {
			return request({
				url: '/users/' + id,
				method: 'delete',
			});
		},

		updateStatus: (id: number, status: number) => {
			return request({
				url: '/users/' + id + '/status/' + status,
				method: 'put',
			});
		},

		setUserRole: (id: number, data: any) => {
			return request({
				url: '/users/' + id + '/roles',
				method: 'post',
				data: data,
			});
		},

		getUserRole: (id: number) => {
			return request({
				url: '/users/' + id + '/roles',
				method: 'get',
			});
		},
	};
}
