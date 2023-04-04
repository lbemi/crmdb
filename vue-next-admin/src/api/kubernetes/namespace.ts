import request from '/@/utils/request';

export function useNamespaceApi() {
    return {
        listNamespace: (param: string) => {
            return request({
                url: '/namespace',
                method: 'get',
                params: param
            });
        },
    };
}
