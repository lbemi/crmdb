import request from '/@/utils/request';


export function useUserApi() {
	return {
		listUser: (query: any) => {
			return request({
				url: '/user',
				method: 'get',
                params: query
			});
		},
        updateUser: ( id: number , data :any) => {
			return request({
				url: '/user/' + id ,
				method: 'put',
                data: data
			});
		},
        addUser: (data: any) => {
			return request({
				url: '/user/register',
				method: 'post',
                data: data
			});
		},
        deleteUser: (id: number) => {
			return request({
				url: '/user/' + id,
				method: 'delete',
			});
		},

        updateStatus: (id: number, status: number) => {
            return request({
				url: '/user/' + id + '/status/' + status,
				method: 'put',
			});
        },
	
	};
}
