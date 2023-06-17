import request from '@/utils/request';


export function useRoleApi() {
    return {
        listRole: (query: any) => {
            return request({
                url: '/roles',
                method: 'get',
                params: query
            });
        },
        updateRole: ( id: number , data :any) => {
            return request({
                url: '/roles/' + id ,
                method: 'put',
                data: data
            });
        },
        addRole: (data: any) => {
            return request({
                url: '/roles',
                method: 'post',
                data: data
            });
        },
        getRoleMenu: (id: number, params?: any) => {
            return request({
                url: '/roles/'+id +'/menus',
                method: 'get',
                params: params
            });
        },
        setRoleAuth: (id:number,data: any) => {
            return request({
                url: '/roles/'+ id +'/menus',
                method: 'post',
                data: data
            });
        },
        deleteRole: (id: number) => {
            return request({
                url: '/roles/' + id,
                method: 'delete',
            });
        },

        updateRoleStatus: (id: number, status: number) => {
            return request({
                url: '/roles/' + id + '/status/' + status,
                method: 'put',
            });
        },

    };
}
