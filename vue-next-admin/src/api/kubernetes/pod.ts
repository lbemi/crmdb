import request from '@/utils/request';

export function usePodApi() {
	return {
		listPods: (ns: string, param: any) => {
			return request({
				url: '/pods/namespaces/' + ns,
				method: 'get',
				params: param,
			});
		},
		getPod: (ns: string, name: string, param: any) => {
			return request({
				url: `/pods/namespaces/${ns}/${name}`,
				method: 'get',
				params: param,
			});
		},
		searchPods: (ns: string, param: any) => {
			return request({
				url: `/pods/namespaces/${ns}/search`,
				method: 'get',
				params: param,
			});
		},
		deletePod: (ns: string | undefined, podName: string | undefined, param: any) => {
			return request({
				url: '/pods/namespaces/' + ns + '/' + podName,
				method: 'delete',
				params: param,
			});
		},
		podLog: (ns: string | undefined, podName: string | undefined, container: string, param: any) => {
			return request({
				url: '/pods/namespaces/' + ns + '/logs/' + podName + '/' + container,
				method: 'get',
				params: param,
				onDownloadProgress: (e) => { },
			});
		},
		podEvents: (ns: string | undefined, podName: string | undefined, param: any) => {
			return request({
				url: '/pods/namespaces/' + ns + '/' + podName + '/events',
				method: 'get',
				params: param,
			});
		},
		getPodFileList: (ns: string | undefined, podName: string | undefined, container: string, param: any) => {
			return request({
				url: '/pods/namespaces/' + ns + '/files/' + podName + '/' + container,
				method: 'get',
				params: param,
			});
		},
	};
}
