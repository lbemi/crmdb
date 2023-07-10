//返回结果
declare type Result = {
	code: number;
	data?: {
		data: any;
		total: number;
	};
	message: string;
};

declare type ParamType<T = any> = {
	page: number;
	limit: number;
	[key: string]: T;
};
