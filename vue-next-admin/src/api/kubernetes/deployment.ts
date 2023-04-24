import request from '/@/utils/request';

export function useDeploymentApi() {
	return {
		listDeployment: (ns: string, param: any) => {
			return request({
				url: '/deployment/' + ns,
				method: 'get',
				params: param,
			});
		},
		createDeployment: (param: any, data: any) => {
			return request({
				url: '/deployment',
				method: 'post',
				params: param,
				data: data,
			});
		},
		updateLabel: (param: any, data: any) => {
			return request({
				url: '/node',
				method: 'patch',
				params: param,
				data: data,
			});
		},
		detailDeployment: (ns: string, depName: string, param: any) => {
			return request({
				url: '/deployment/' + ns + '/' + depName + '/pod',
				method: 'get',
				params: param,
			});
		},
		deleteDeployment: (ns: string, name: string, param: any) => {
			return request({
				url: `/deployment/${ns}/${name}`,
				method: 'delete',
				params: param,
			});
		},
		scaleDeployment: (ns: string, name: string, n: number, param: any) => {
			return request({
				url: `/deployment/${ns}/${name}/${n}`,
				method: 'put',
				params: param,
			});
		},
	};
}
