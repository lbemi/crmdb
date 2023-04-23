import {
	V1ObjectMeta,
	V1VolumeMount,
	V1HostPathVolumeSource,
	V1SecretVolumeSource,
	V1ConfigMapVolumeSource,
	V1PersistentVolumeClaimVolumeSource,
	V1EmptyDirVolumeSource,
} from '@kubernetes/client-node';

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

export declare type CreateK8SVolumentData = {
	keySet?: boolean;
	keySetShow?: boolean;
	type: string;
	name: string;
	hostPath?: V1HostPathVolumeSource|undefined;
	secret?: V1SecretVolumeSource|undefined;
	configMap?: V1ConfigMapVolumeSource|undefined;
	persistentVolumeClaim?: V1PersistentVolumeClaimVolumeSource|undefined;
	emptyDir?: V1EmptyDirVolumeSource|undefined;
	volumeMountData?: V1VolumeMount|undefined;
};
