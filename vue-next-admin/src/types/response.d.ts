//返回结果
export interface ResponseType {
	code: Number;
	data: {
		data: any;
		total: number;
	};
	message: string;
}
