import { Host } from '@/types/asset/hosts';
import request from '@/utils/request';
import { Group } from 'jsplumb';

export function useHostApi() {
	return {
		/**
		 * Retrieves a list of hosts based on the given query parameters.
		 *
		 * @param {any} query - The query parameters for retrieving hosts.
		 * @return {Promise<any>} - A promise that resolves to the list of hosts.
		 */
		lisHost: (query: any): Promise<any> => {
			return request({
				url: '/hosts',
				method: 'get',
				params: query,
			});
		},
		/**
		 * Fetches the list of hosts by group.
		 *
		 * @param {any} query - The query parameters.
		 * @param {Group[]} data - The data parameters.
		 * @return {Promise<any>} A promise that resolves with the list of hosts.
		 */
		lisHostByGroup: (query: any, data: Group[]): Promise<any> => {
			return request({
				url: '/hosts/groups',
				method: 'get',
				params: query,
				data: data,
			});
		},
		/**
		 * Update the host with the given data.
		 *
		 * @param {Host} data - The data to update the host with.
		 * @return {Promise<any>} - A promise that resolves with the updated host.
		 */
		updateHost: (data: Host): Promise<any> => {
			return request({
				url: '/hosts',
				method: 'put',
				data: data,
			});
		},
		/**
		 * Adds a host to the system.
		 *
		 * @param {Host} data - The data for the host.
		 * @return {any} The response from the server.
		 */
		addHost: (data: Host): any => {
			return request({
				url: '/hosts',
				method: 'post',
				data: data,
			});
		},
		/**
		 * Deletes a host with the given ID.
		 *
		 * @param {number} id - The ID of the host to be deleted.
		 * @return {Promise<void>} A promise that resolves when the host is successfully deleted.
		 */
		deleteHost: (id: number): Promise<void> => {
			return request({
				url: '/hosts/' + id,
				method: 'delete',
			});
		},
	};
}
