import { Container, IContainer } from 'kubernetes-models/v1';

export interface IInitContainer extends IContainer {
	isInitContainer: boolean;
}
// 继承 IContainer 接口并添加 isInitContainer 属性
export  class CustomizeContainer extends Container implements IInitContainer {
	isInitContainer: boolean;
	constructor(options: IContainer & { isInitContainer: boolean }) {
		super(options);
		this.isInitContainer = options.isInitContainer;
	}
}
