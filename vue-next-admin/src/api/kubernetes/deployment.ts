import request from '/@/utils/request';

export function useDeploymentApi() {
    return {
        listDeployment: (ns: string,param: any) => {
            return request({
                url: '/deployment/'+ns,
                method: 'get',
                params: param,
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
    };
}
