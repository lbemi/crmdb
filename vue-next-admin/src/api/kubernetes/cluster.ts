import request from '/@/utils/request';

export function useClusterApi() {
	return {
		listCluster: () => {
			return request({
				url: '/cluster',
				method: 'get',
			});
		},
		deleteCluster: (id: number) => {
			return request({
				url: '/cluster/' + id,
                method: 'delete'
			});
		},
        createCluster: (data:any, header:any) => {
            return request({
                url: '/cluster',
                method:  "post",
                data: data,
                headers: header
            })
        },
        getCluster: (clusterName: string)=>{
            return request({
                url: '/cluster/' + clusterName,
                method: 'get'
            })
        }
	};
}