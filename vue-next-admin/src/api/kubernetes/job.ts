import request from '@/utils/request';

export function useJobApi() {
	return {
		listJob: (ns: string, param: any) => {
			return request({
				url: '/jobs/namespaces/' + ns,
				method: 'get',
				params: param,
			});
		},
		searchJob: (ns: string, param: any) => {
			return request({
				url: `/jobs/namespaces/${ns}/search`,
				method: 'get',
				params: param,
			});
		},
		createJob: (param: any, data: any) => {
			return request({
				url: '/jobs',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updateLabel: (param: any, data: any) => {
			return request({
				url: '/jobs',
				method: 'patch',
				params: param,
				data: data,
			});
		},
		detailJob: (ns: string, depName: string, param: any) => {
			return request({
				url: '/jobs/namespaces/' + ns + '/' + depName + '/pods',
				method: 'get',
				params: param,
			});
		},
		getJob: (ns: string, depName: string, param: any) => {
			return request({
				url: '/jobs/namespaces/' + ns + '/' + depName,
				method: 'get',
				params: param,
			});
		},
		getJobEvents: (ns: string, depName: string, param: any) => {
			return request({
				url: '/jobs/namespaces/' + ns + '/' + depName + '/events',
				method: 'get',
				params: param,
			});
		},
		deleteJob: (ns: string, name: string, param: any) => {
			return request({
				url: `/jobs/namespaces/${ns}/${name}`,
				method: 'delete',
				params: param,
			});
		},
		// "/namespaces/{namespace}/{name}/scale/{replica}").
		scaleJob: (ns: string, name: string, n: number, param: any) => {
			return request({
				url: `/jobs/namespaces/${ns}/${name}/scale/${n}`,
				method: 'put',
				params: param,
			});
		},
		updateJob: (data: any, param: any) => {
			return request({
				url: `/jobs`,
				method: 'put',
				data: data,
				params: param,
			});
		},
		reDeployJob: (ns: string, name: string, param: any) => {
			return request({
				url: `/jobs/namespaces/${ns}/redeploy/${name}`,
				method: 'put',
				params: param,
			});
		},
		// /rollback/:namespace/:name/:reversion
		rollBackJob: (ns: string, name: string, reversion: string, param: any) => {
			return request({
				url: `/jobs/namespaces/${ns}/rollback/${name}/${reversion}`,
				method: 'put',
				params: param,
			});
		},
	};
}
