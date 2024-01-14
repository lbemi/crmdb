import request from '@/utils/request';

export function useCronJobApi() {
	return {
		listCronJob: (ns: string, param: any) => {
			return request({
				url: '/cronjobs/namespaces/' + ns,
				method: 'get',
				params: param,
			});
		},
		searchCronJob: (ns: string, param: any) => {
			return request({
				url: `/cronjobs/namespaces/${ns}/search`,
				method: 'get',
				params: param,
			});
		},
		createCronJob: (param: any, data: any) => {
			return request({
				url: '/cronjobs',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updateLabel: (param: any, data: any) => {
			return request({
				url: '/cronjobs',
				method: 'patch',
				params: param,
				data: data,
			});
		},
		detailCronJob: (ns: string, depName: string, param: any) => {
			return request({
				url: '/cronjobs/namespaces/' + ns + '/' + depName + '/pods',
				method: 'get',
				params: param,
			});
		},
		getCronJob: (ns: string, depName: string, param: any) => {
			return request({
				url: '/cronjobs/namespaces/' + ns + '/' + depName,
				method: 'get',
				params: param,
			});
		},
		getCronJobEvents: (ns: string, depName: string, param: any) => {
			return request({
				url: '/cronjobs/namespaces/' + ns + '/' + depName + '/events',
				method: 'get',
				params: param,
			});
		},
		deleteCronJob: (ns: string, name: string, param: any) => {
			return request({
				url: `/cronjobs/namespaces/${ns}/${name}`,
				method: 'delete',
				params: param,
			});
		},
		// "/namespaces/{namespace}/{name}/scale/{replica}").
		scaleCronJob: (ns: string, name: string, n: number, param: any) => {
			return request({
				url: `/cronjobs/namespaces/${ns}/${name}/scale/${n}`,
				method: 'put',
				params: param,
			});
		},
		updateCronJob: (data: any, param: any) => {
			return request({
				url: `/cronjobs`,
				method: 'put',
				data: data,
				params: param,
			});
		},
		reDeployCronJob: (ns: string, name: string, param: any) => {
			return request({
				url: `/cronjobs/namespaces/${ns}/redeploy/${name}`,
				method: 'put',
				params: param,
			});
		},
		// /rollback/:namespace/:name/:reversion
		rollBackCronJob: (ns: string, name: string, reversion: string, param: any) => {
			return request({
				url: `/cronjobs/namespaces/${ns}/rollback/${name}/${reversion}`,
				method: 'put',
				params: param,
			});
		},
	};
}
