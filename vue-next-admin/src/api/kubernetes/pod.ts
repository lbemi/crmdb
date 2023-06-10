import request from '/@/utils/request';

export function usePodApi() {
	return {
		listPods: (ns: string, param: any) => {
			return request({
				url: '/pods/' + ns,
				method: 'get',
				params: param,
			});
		},
		getPod: (ns: string,name: string, param: any) => {
			return request({
				url: `/pods/${ns}/${name}`,
				method: 'get',
				params: param,
			});
		},
		searchPods: (ns: string, param: any) => {
			return request({
				url: `/pods/${ns}/search`,
				method: 'get',
				params: param,
			});
		},
		deletePod: (ns: string | undefined, podName: string | undefined, param: any) => {
			return request({
				url: '/pods/' + ns + '/' + podName,
				method: 'delete',
				params: param,
			});
		},
		podLog: (ns: string | undefined, podName: string | undefined, container: string, param: any) => {
			return request({
				url: '/pods/logs/' + ns + '/' + podName + '/' + container,
				method: 'get',
				params: param,
				onDownloadProgress: (e) => {},
			});
		},
		podEvents: (ns: string | undefined, podName: string | undefined, param: any) => {
			return request({
				url: '/pods/events/' + ns + '/' + podName,
				method: 'get',
				params: param,
			});
		},
	};
}
