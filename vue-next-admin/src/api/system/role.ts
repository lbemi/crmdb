import request from '/@/utils/request';


export function useRoleApi() {
    return {
        listRole: (query: any) => {
            return request({
                url: '/role',
                method: 'get',
                params: query
            });
        },
        updateRole: ( id: number , data :any) => {
            return request({
                url: '/role/' + id ,
                method: 'put',
                data: data
            });
        },
        addRole: (data: any) => {
            return request({
                url: '/role',
                method: 'post',
                data: data
            });
        },
        setRoleAuth: (id:number,data: any) => {
            return request({
                url: '/role/'+ id,
                method: 'post',
                data: data
            });
        },
        deleteRole: (id: number) => {
            return request({
                url: '/role/' + id,
                method: 'delete',
            });
        },

        updateRoleStatus: (id: number, status: number) => {
            return request({
                url: '/role/' + id + '/status/' + status,
                method: 'put',
            });
        },

    };
}
