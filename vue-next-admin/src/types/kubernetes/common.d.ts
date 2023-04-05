declare class QueryType<T = any>{
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
	page: number
	limit: number
}
