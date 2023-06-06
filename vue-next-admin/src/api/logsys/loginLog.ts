import { AnyColumn } from 'element-plus/es/components/table-v2/src/common';
import request from '/@/utils/request';

export function useLoinLogApi() {
	return {
		listLoginLog: (query: any, condition?: any) => {
			return request({
				url: '/logs/login',
				method: 'get',
				params: query,
			});
		},
		getLoginLog: (id: number) => {
			return request({
				url: '/logs/login' + id,
				method: 'get',
			});
		},
		deleteLoginLog: (query: any) => {
			return request({
				url: '/logs/login',
				method: 'delete',
				params: query,
			});
		},
		deleteAllLoginLog: () => {
			return request({
				url: '/logs/login/all',
				method: 'delete',
			});
		},
	};
}
