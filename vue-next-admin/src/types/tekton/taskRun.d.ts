import { EnvFromSource, EnvVar, ResourceRequirements, SecurityContext, VolumeMount } from 'kubernetes-models/v1';
import { NamedResource } from './common';

export interface Metadata {
	name?: string;
	annotations?: {
		[key: string]: string;
	};
	labels?: {
		[key: string]: string;
	};
	namespace?: string;
	readonly finalizers?: string[];
	readonly ownerReferences?: OwnerReference[];
}

export interface Annotations {
	[key: string]: string;
}

export interface Param {
	name: string;
	type: string;
}

export interface WorkspaceUsage {
	name: string;
	mountPath: string;
}
export interface stdConfig {
	path: string;
}

export interface TaskStep extends NamedResource {
	image?: string;
	script?: string;
	command?: string[];
	args?: string[];
	volumeMounts?: VolumeMount[];
	workingDir?: string;
	env?: EnvVar[];
	envFrom?: EnvFromSource[];
	computeResources: ResourceRequirements;
	volumeDevices: VolumeDevicep[];
	imagePullPolicy?: string;
	securityContext: SecurityContext;
	timeout: Duration;
	workspaces: WorkspaceUsage[];
	onError: 'continue' | 'stopAndFail';
	stdoutConfig: stdConfig;
	stderrConfig: stdConfig;
}

export interface TaskSpecParam extends NamedResource {
	type?: string;
	description?: string;
	default?:
		| string
		| {
				[key: string]: string;
		  };
	properties?: {
		[key: string]: {
			type: string;
		};
	};
}

export interface TaskSpecResult extends NamedResource {
	description?: string;
	type?: string;
	properties?: {
		[key: string]: {
			type: string;
		};
	};
}
export interface TaskWorkspace extends NamedResource {
	description?: string;
	mountPath?: string;
	readOnly?: boolean;
	optional?: boolean;
}
export interface TaskSpec {
	description?: string;
	params?: TaskSpecParam[];
	results?: TaskSpecResult[];
	workspaces?: TaskWorkspace[];
	steps?: TaskStep[];
}

export interface Task {
	kind?: string;
	apiVersion?: string;
	metadata?: Metadata;
	spec?: TaskSpec;
}
