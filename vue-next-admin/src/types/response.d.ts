//返回结果
export type ResponseType = {
	code: number;
	data?: {
		data: any;
		total: number;
	};
	message: string;
};
