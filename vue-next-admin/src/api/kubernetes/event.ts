import request from '/@/utils/request';

export function useEventApi() {
    return {
        getEventLog: (namespace :string,param: any) => {
            return request({
                url: '/event/' + namespace,
                method: 'get',
                params: param
            });
        },
    };
}
