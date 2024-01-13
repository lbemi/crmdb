import request from '@/utils/request';

export function useDaemonsetApi() {
	return {
		listDaemonset: (ns: string, param: any) => {
			return request({
				url: '/daemonsets/namespaces/' + ns,
				method: 'get',
				params: param,
			});
		},
		searchDaemonset: (ns: string, param: any) => {
			return request({
				url: `/daemonsets/namespaces/${ns}/search`,
				method: 'get',
				params: param,
			});
		},
		createDaemonset: (param: any, data: any) => {
			return request({
				url: '/daemonsets',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updateLabel: (param: any, data: any) => {
			return request({
				url: '/daemonsets',
				method: 'patch',
				params: param,
				data: data,
			});
		},
		detailDaemonset: (ns: string, depName: string, param: any) => {
			return request({
				url: '/daemonsets/namespaces/' + ns + '/' + depName + '/pods',
				method: 'get',
				params: param,
			});
		},
		getDaemonset: (ns: string, depName: string, param: any) => {
			return request({
				url: '/daemonsets/namespaces/' + ns + '/' + depName,
				method: 'get',
				params: param,
			});
		},
		getDaemonsetEvents: (ns: string, depName: string, param: any) => {
			return request({
				url: '/daemonsets/namespaces/' + ns + '/' + depName + '/events',
				method: 'get',
				params: param,
			});
		},
		deleteDaemonset: (ns: string, name: string, param: any) => {
			return request({
				url: `/daemonsets/namespaces/${ns}/${name}`,
				method: 'delete',
				params: param,
			});
		},
		// "/namespaces/{namespace}/{name}/scale/{replica}").

		updateDaemonset: (data: any, param: any) => {
			return request({
				url: `/daemonsets`,
				method: 'put',
				data: data,
				params: param,
			});
		},
		reDeployDaemonset: (ns: string, name: string, param: any) => {
			return request({
				url: `/daemonsets/namespaces/${ns}/redeploy/${name}`,
				method: 'put',
				params: param,
			});
		},
		// /rollback/:namespace/:name/:reversion
		rollBackDaemonset: (ns: string, name: string, reversion: string, param: any) => {
			return request({
				url: `/daemonsets/namespaces/${ns}/rollback/${name}/${reversion}`,
				method: 'put',
				params: param,
			});
		},
	};
}
