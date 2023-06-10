import request from '/@/utils/request';

export function useEventApi() {
    return {
        getEventLog: (namespace :string,param: any) => {
            return request({
                url: '/events/namespaces/' + namespace,
                method: 'get',
                params: param
            });
        },
    };
}
