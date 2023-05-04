import request from '/@/utils/request';

export function usePodApi() {
	return {
		listPods: (ns: string, param: any) => {
			return request({
				url: '/pod/' + ns,
				method: 'get',
				params: param,
			});
		},
		deletePod: (ns: string | undefined, podName: string | undefined, param: any) => {
			return request({
				url: '/pod/' + ns + '/' + podName,
				method: 'delete',
				params: param,
			});
		},
		podLog: (ns: string | undefined, podName: string | undefined, container: string, param: any) => {
			return request({
				url: '/pod/log/' + ns + '/' + podName + '/' + container,
				method: 'get',
				params: param,
				onDownloadProgress: (e) => {},
			});
		},
		podEvents: (ns: string | undefined, podName: string | undefined, param: any) => {
			return request({
				url: '/pod/event/' + ns + '/' + podName,
				method: 'get',
				params: param,
			});
		},
	};
}
