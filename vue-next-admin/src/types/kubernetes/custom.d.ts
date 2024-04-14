import { Container, IContainer, IIoK8sApimachineryPkgApisMetaV1ObjectMeta } from 'kubernetes-models/v1';
import {
	VolumeMount,
	HostPathVolumeSource,
	SecretVolumeSource,
	ConfigMapVolumeSource,
	PersistentVolumeClaimVolumeSource,
	EmptyDirVolumeSource,
} from 'kubernetes-models/v1';
export declare type ResourceType = 'deployment' | 'daemonSet' | 'statefulSet' | 'job' | 'cronJob';

export declare type CreateK8SMeta = {
	metadata?: IIoK8sApimachineryPkgApisMetaV1ObjectMeta;
	replicas?: number;
	resourceType?: ResourceType;
};

export declare type CreateK8SMetaData = {
	loadFromParent: boolean;
	replicas?: number;
	resourceType?: ResourceType;
	meta?: IIoK8sApimachineryPkgApisMetaV1ObjectMeta;
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

