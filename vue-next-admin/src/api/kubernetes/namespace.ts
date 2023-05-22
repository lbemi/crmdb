import { Namespace } from 'kubernetes-types/core/v1';

import request from '/@/utils/request';

export function useNamespaceApi() {
	return {
		listNamespace: (param: Object) => {
			return request({
				url: '/namespace',
				method: 'get',
				params: param,
			});
		},
		updateNamespace: (param: Object, namespace: Namespace) => {
			return request({
				url: '/namespace',
				method: 'put',
				params: param,
				data: namespace,
			});
		},
		createNamespace: (param: Object, namespace: Namespace) => {
			return request({
				url: '/namespace',
				method: 'post',
				params: param,
				data: namespace,
			});
		},
		deleteNamespace: (param: Object, namespaceName: string) => {
			return request({
				url: `/namespace/${namespaceName}`,
				method: 'delete',
				params: param,
			});
		},
	};
}
