import request from '@/utils/request';

export function useStatefulSetApi() {
	return {
		listStatefulSet: (ns: string, param: any) => {
			return request({
				url: '/statefulsets/namespaces/' + ns,
				method: 'get',
				params: param,
			});
		},
		searchStatefulSet: (ns: string, param: any) => {
			return request({
				url: `/statefulsets/namespaces/${ns}/search`,
				method: 'get',
				params: param,
			});
		},
		createStatefulSet: (param: any, data: any) => {
			return request({
				url: '/statefulsets',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updateLabel: (param: any, data: any) => {
			return request({
				url: '/statefulsets',
				method: 'patch',
				params: param,
				data: data,
			});
		},
		detailStatefulSet: (ns: string, depName: string, param: any) => {
			return request({
				url: '/statefulsets/namespaces/' + ns + '/' + depName + '/pods',
				method: 'get',
				params: param,
			});
		},
		getStatefulSet: (ns: string, depName: string, param: any) => {
			return request({
				url: '/statefulsets/namespaces/' + ns + '/' + depName,
				method: 'get',
				params: param,
			});
		},
		getStatefulSetEvents: (ns: string, depName: string, param: any) => {
			return request({
				url: '/statefulsets/namespaces/' + ns + '/' + depName + '/events',
				method: 'get',
				params: param,
			});
		},
		deleteStatefulSet: (ns: string, name: string, param: any) => {
			return request({
				url: `/statefulsets/namespaces/${ns}/${name}`,
				method: 'delete',
				params: param,
			});
		},
		// "/namespaces/{namespace}/{name}/scale/{replica}").
		scaleStatefulSet: (ns: string, name: string, n: number, param: any) => {
			return request({
				url: `/statefulsets/namespaces/${ns}/${name}/scale/${n}`,
				method: 'put',
				params: param,
			});
		},
		updateStatefulSet: (data: any, param: any) => {
			return request({
				url: `/statefulsets`,
				method: 'put',
				data: data,
				params: param,
			});
		},
		reDeployStatefulSet: (ns: string, name: string, param: any) => {
			return request({
				url: `/statefulsets/namespaces/${ns}/redeploy/${name}`,
				method: 'put',
				params: param,
			});
		},
		// /rollback/:namespace/:name/:reversion
		rollBackStatefulSet: (ns: string, name: string, reversion: string, param: any) => {
			return request({
				url: `/statefulsets/namespaces/${ns}/rollback/${name}/${reversion}`,
				method: 'put',
				params: param,
			});
		},
	};
}
