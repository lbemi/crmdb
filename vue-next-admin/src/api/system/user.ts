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
		getTestMenu: (params?: object) => {
			return request({
				url: '/gitee/lyt-top/vue-next-admin-images/raw/master/menu/testMenu.json',
				method: 'get',
				params,
			});
		},
	
	};
}
