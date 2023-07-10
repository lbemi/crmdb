import { HostParams } from '@/types/asset/hosts';
import { ParamsState } from '@/types/views';
import request from '@/utils/request';

export function useAccountApi() {
	return {
	
		/**
		 * Retrieves a list of accounts.
		 *
		 * @param {HostParams} [query] - Optional query parameters.
		 * @return {Promise<Account[]>} A promise that resolves to an array of Account objects.
		 */
		listAccount: (query?: HostParams): Promise<Account[]> => {
			return request({
				url: '/accounts',
				method: 'get',
				params: query,
			});
		},
		
		/**
		 * Updates an account.
		 *
		 * @param {Account} data - The data of the account to be updated.
		 * @return {Promise<Result>} A promise that resolves to the result of the update operation.
		 */
		updateAccount: (data: Account): Promise<Result> => {
			return request({
				url: '/accounts',
				method: 'put',
				data: data,
			});
		},

		
		/**
		 * Adds an account.
		 *
		 * @param {Account} data - The account data to be added.
		 * @return {Promise<Result>} A promise that resolves to the result of the operation.
		 */
		addAccount: (data: Account): Promise<Result> => {
			return request({
				url: '/accounts',
				method: 'post',
				data: data,
			});
		},
	
		/**
		 * Deletes an account.
		 *
		 * @param {number} id - The ID of the account to be deleted.
		 * @return {Promise<Result>} A promise that resolves to the result of the deletion.
		 */
		deleteAccount: (id: number): Promise<Result> => {
			return request({
				url: '/accounts/' + id,
				method: 'delete',
			});
		},
	};
}
