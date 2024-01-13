// export enum ResourceType {
// 	Deployment = 'deployment',
// 	DaemonSet = 'daemonSet',
// 	Job = 'job',
// 	CronJob = 'cronJob',
// 	StatefulSet = 'statefulSet',
// }

import { ObjectMeta } from 'kubernetes-types/meta/v1';
import {
	VolumeMount,
	HostPathVolumeSource,
	SecretVolumeSource,
	ConfigMapVolumeSource,
	PersistentVolumeClaimVolumeSource,
	EmptyDirVolumeSource,
} from 'kubernetes-types/core/v1';
export declare type ResourceType = 'deployment' | 'daemonSet' | 'statefulSet' | 'job' | 'cronJob';

export declare type CreateK8SBindData = {
	metadata?: ObjectMeta;
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
	meta?: ObjectMeta;
};

export declare type CreateK8SVolumeData = {
	type: string;
	name: string;
	hostPath?: HostPathVolumeSource | undefined;
	secret?: SecretVolumeSource | undefined;
	configMap?: ConfigMapVolumeSource | undefined;
	virtualService?: ConfigMapVolumeSource | undefined;
	persistentVolumeClaim?: PersistentVolumeClaimVolumeSource | undefined;
	emptyDir?: EmptyDirVolumeSource | undefined;
	volumeMountData?: VolumeMount | undefined;
};
