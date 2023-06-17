import request from '@/utils/request';

export function useDeploymentApi() {
	return {
		listDeployment: (ns: string, param: any) => {
			return request({
				url: '/deployments/namespaces/' + ns,
				method: 'get',
				params: param,
			});
		},
		searchDeployment: (ns: string, param: any) => {
			return request({
				url: `/deployments/namespaces/${ns}/search`,
				method: 'get',
				params: param,
			});
		},
		createDeployment: (param: any, data: any) => {
			return request({
				url: '/deployments',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updateLabel: (param: any, data: any) => {
			return request({
				url: '/deployments',
				method: 'patch',
				params: param,
				data: data,
			});
		},
		detailDeployment: (ns: string, depName: string, param: any) => {
			return request({
				url: '/deployments/namespaces/' + ns + '/' + depName + '/pods',
				method: 'get',
				params: param,
			});
		},
		getDeploymentEvents: (ns: string, depName: string, param: any) => {
			return request({
				url: '/deployments/namespaces/' + ns + '/' + depName + '/events',
				method: 'get',
				params: param,
			});
		},
		deleteDeployment: (ns: string, name: string, param: any) => {
			return request({
				url: `/deployments/namespaces/${ns}/${name}`,
				method: 'delete',
				params: param,
			});
		},
		// "/namespaces/{namespace}/{name}/scale/{replica}").
		scaleDeployment: (ns: string, name: string, n: number, param: any) => {
			return request({
				url: `/deployments/namespaces/${ns}/${name}/scale/${n}`,
				method: 'put',
				params: param,
			});
		},
		updateDeployment: (data: any, param: any) => {
			return request({
				url: `/deployments`,
				method: 'put',
				data: data,
				params: param,
			});
		},
		reDeployDeployment: (ns: string, name: string, param: any) => {
			return request({
				url: `/deployments/namespaces/${ns}/redeploy/${name}`,
				method: 'put',
				params: param,
			});
		},
		// /rollback/:namespace/:name/:reversion
		rollBackDeployment: (ns: string, name: string, reversion: string, param: any) => {
			return request({
				url: `/deployments/namespaces/${ns}/rollback/${name}/${reversion}`,
				method: 'put',
				params: param,
			});
		},
	};
}
