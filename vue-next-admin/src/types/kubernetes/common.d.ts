import { Container, IContainer } from 'kubernetes-models/v1';

// eslint-disable-next-line @typescript-eslint/no-unused-vars
export interface MirrorRepository<T> {
	auths: {
		[T]: {
			username: string;
			password: string;
			email: string;
			auth: string;
		};
	};
}
declare class QueryType<T = any> {
	query: {
		page: number;
		limit: number;
		[key: string]: T;
	};
	data?: any;
	total: number;
	loading: boolean;
}

export interface PageInfo {
	page: number;
	limit: number;
}
