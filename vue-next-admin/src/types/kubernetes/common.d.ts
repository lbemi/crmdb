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

export interface IInitContainer extends IContainer {
	isInitContainer: boolean;
}
// 继承 IContainer 接口并添加 isInitContainer 属性
export declare class CustomizeContainer extends Container implements IInitContainer {
	isInitContainer: boolean;
	constructor(options: IContainer & { isInitContainer: boolean }) {
		super(options);
		this.isInitContainer = options.isInitContainer;
	}
}
