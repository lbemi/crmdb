import request from '@/utils/request';

export function useHBAApi() {
	return {
		/**
		 * Retrieves a list of host bing accounts.
		 *
		 * @param {ParamType} query - Optional query parameters.
		 * @return {Promise<HostBingAccounts[]>} A promise that resolves to an array of host bing accounts.
		 */
		lisHba: (query?: ParamType): Promise<HostBingAccounts[]> => {
			return request({
				url: '/hbas',
				method: 'get',
				params: query,
			});
		},
		/**
		 * Updates the host bing accounts.
		 *
		 * @param {HostBingAccounts} data - The data to update.
		 * @return {Promise<Result>} A promise that resolves with the result.
		 */
		updateHba: (data: HostBingAccounts): Promise<Result> => {
			return request({
				url: '/hbas',
				method: 'put',
				data: data,
			});
		},

		/**
		 * Adds a host bing account.
		 *
		 * @param {HostBingAccounts} data - The host bing account data.
		 * @return {Promise<Result>} The result of the operation.
		 */
		addHba: (data: HostBingAccounts): Promise<Result> => {
			return request({
				url: '/hbas',
				method: 'post',
				data: data,
			});
		},
		/**
		 * Deletes an HBA by its ID.
		 *
		 * @param {number} id - The ID of the HBA to delete.
		 * @return {Promise<Result>} A promise that resolves with the result of the deletion.
		 */
		deleteHba: (id: number): Promise<Result> => {
			return request({
				url: '/hbas/' + id,
				method: 'delete',
			});
		},
	};
}
