import { EnvFromSource, EnvVar, ResourceRequirements, SecurityContext, VolumeDevice, VolumeMount } from 'kubernetes-models/v1';
import { NamedResource } from '../common';
export interface OwnerReference {
	/**
	 * API version of the referent.
	 */
	readonly apiVersion: string;
	/**
	 * If true, AND if the owner has the "foregroundDeletion" finalizer, then the
	 * owner cannot be deleted from the key-value store until this reference is
	 * removed. Defaults to false. To set this field, a user needs "delete"
	 * permission of the owner, otherwise 422 (Unprocessable Entity) will be
	 * returned.
	 *
	 * @default false. To set this field, a user needs "delete" permission of the
	 * owner, otherwise 422 (Unprocessable Entity) will be returned.
	 */
	readonly blockOwnerDeletion?: boolean;
	/**
	 * If true, this reference points to the managing controller.
	 */
	readonly controller?: boolean;
	/**
	 * Kind of the referent.
	 *
	 * @see https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	 */
	readonly kind: string;
	/**
	 * Name of the referent.
	 *
	 * @see http://kubernetes.io/docs/user-guide/identifiers#names
	 */
	readonly name: string;
	/**
	 * UID of the referent.
	 *
	 * @see http://kubernetes.io/docs/user-guide/identifiers#uids
	 */
	readonly uid: string;
}

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
	volumeDevices: VolumeDevice[];
	imagePullPolicy?: string;
	securityContext: SecurityContext;
	timeout: string;
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
				// eslint-disable-next-line no-mixed-spaces-and-tabs
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
