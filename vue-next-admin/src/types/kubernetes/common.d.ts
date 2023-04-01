declare interface QueryType<T = any>{
	query: {
		page: number;
		limit: number;
		[key: string]: T;
	};
	total: number;
	loading: boolean;
}

export interface PageInfo {
	page: number
	limit: number
}
