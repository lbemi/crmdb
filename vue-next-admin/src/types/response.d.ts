export interface ResponseType<T = any> {
	code: Number;
	data: {
		data?: any;
		[key: string]: T;
		total: number;
	};
}
