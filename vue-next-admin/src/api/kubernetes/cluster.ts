import request from '/@/utils/request';

export function useClusterApi() {
	return {
		listCluster: () => {
			return request({
				url: '/clusters',
				method: 'get',
			});
		},
		deleteCluster: (id: number) => {
			return request({
				url: '/clusters/' + id,
				method: 'delete',
			});
		},
		createCluster: (data: any, header: any) => {
			return request({
				url: '/clusters',
				method: 'post',
				data: data,
				headers: header,
			});
		},
		getCluster: (clusterName: string) => {
			return request({
				url: '/clusters/' + clusterName,
				method: 'get',
			});
		},
	};
}
