import { number } from '@intlify/core-base';
import { V1ObjectMeta } from '@kubernetes/client-node';

export declare type ResourceType = 'deployment' | 'daemonSet' | 'statefulSet' | 'job' | 'cronJob';

export declare type CreateK8SBindData = {
	metadata?: V1ObjectMeta;
	replicas?: number;
	resourceType?: ResourceType;
};

export interface CreateK8SLabel {
	key: string;
	value: string;
}

export declare type CreateK8SMetaData = {
	loadFromParent: boolean;
	labelData?: Array<CreateK8SLabel>;
	annotationsData?: Array<CreateK8SLabel>;
	replicas?: number;
	resourceType?: ResourceType;
	meta?: V1ObjectMeta;
};
