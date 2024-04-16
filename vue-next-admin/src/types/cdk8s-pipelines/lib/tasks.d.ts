/**
 * This file includes the basic objects for creating Tasks and dependencies.
 * While cdk8s supports generating this file from the CRD, the Tekton CRD at
 * the moment does not have the fields specified. See https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml,
 * specifically here:
 *       schema:
 *         openAPIV3Schema:
 *           type: object
 *           # OpenAPIV3 schema allows Kubernetes to perform validation on the schema fields
 *           # and use the schema in tooling such as `kubectl explain`.
 *           # Using "x-kubernetes-preserve-unknown-fields: true"
 *           # at the root of the schema (or within it) allows arbitrary fields.
 *           # We currently perform our own validation separately.
 *           # See https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#specifying-a-structural-schema
 *           # for more info.
 *           x-kubernetes-preserve-unknown-fields: true
 * For the time being, these objects are created manually using the latest
 * Tekton documentation as a reference.
 */
import { ApiObject, ApiObjectMetadata, GroupVersionKind } from '@/types/cdk8s';
import { Construct } from 'constructs';
import { NamedResource, NameKeyPair } from './common';
/**
 * The source for a `env` `valueFrom`.
 */
export interface TaskEnvValueSource {
	readonly secretKeyRef: NameKeyPair;
}
/**
 * Creates an `env` source that can be used in the `env`'s `valueFrom`.
 * @param nk The name and key pair.
 */
export declare function valueFrom(nk: NameKeyPair): TaskEnvValueSource;
/**
 * An `env` for a `Step` on a `Task`.
 */
export interface TaskStepEnv extends NamedResource {
	readonly valueFrom?: TaskEnvValueSource;
}
/**
 * A workspace used by a Task. See https://tekton.dev/docs/pipelines/workspaces/#using-workspaces-in-tasks for more information.
 */
export interface TaskWorkspace extends NamedResource {
	readonly logicalID?: string;
	/**
	 * The description of the workspace.
	 */
	readonly description?: string;
}
/**
 * A `Task` reference. Will be generated as a `taskRef`.
 */
export declare class TaskRef {
	name?: string;
	constructor(name: string);
}
/**
 * A Task parameter value.
 */
export interface TaskParam extends NamedResource {
	/**
	 * The value of the task parameter.
	 */
	readonly value?: string;
}
/**
 * The Task results.
 *
 * @see https://tekton.dev/docs/pipelines/tasks/#emitting-results
 */
export interface TaskSpecResult extends NamedResource {
	/**
	 * The description of the result.
	 */
	description?: string;
	/**
	 * The type of result. This could be an array, but if it is not supplied the
	 * default is a string.
	 */
	type?: string;
}
/**
 * Specifies execution parameters for the Task.
 */
export interface TaskSpecParam extends NamedResource {
	/**
	 * The parameter's type.
	 */
	type?: string;
	/**
	 * The parameter's description.
	 */
	description?: string;
	/**
	 * The default value for a parameter.
	 */
	default?:
		| string
		| {
				[key: string]: string;
		  };
	/**
	 * object param must specify the properties section to define the schema i.e. what keys are available for this object param. See how to define properties section in the following example and the TEP-0075.
	 */
	properties?: {
		[key: string]: {
			type: string;
		};
	};
}
/**
 * The volume mount for the task.
 */
export interface TaskVolumeMount extends NamedResource {
	readonly mountPath?: string;
}
/**
 * The step for a Task. See https://tekton.dev/docs/pipelines/tasks/#defining-steps
 * @see https://tekton.dev/docs/pipelines/container-contract/
 */
export interface TaskStep extends NamedResource {
	/**
	 * The name of the container image to use for the Step.
	 */
	readonly image?: string;
	/**
	 * A script that will be executed on the image. If `script` is specified,
	 * then the `command` value should not be specified.
	 */
	readonly script?: string;
	/**
	 * The command and its arguments (provided in the form of an array) to
	 * execute on the container. If `command` is supplied, you should not supply
	 * `script`.
	 */
	readonly command?: string[];
	/**
	 * Alternatively, you can supply `args` to the `command` value here.
	 */
	readonly args?: string[];
	/**
	 * The volume mounts to use for the task.
	 */
	readonly volumeMounts?: TaskVolumeMount[];
	/**
	 * The name of the working directory for the step
	 */
	readonly workingDir?: string;
	/**
	 * Environment variables for the `Step` on the `Task`.
	 */
	readonly env?: TaskStepEnv[];
}
/**
 * The Task spec.
 *
 * @see https://tekton.dev/docs/pipelines/tasks/#configuring-a-task
 */
export interface TaskSpec {
	/**
	 * The description of the `Task`.
	 * @see https://tekton.dev/docs/pipelines/tasks/#adding-a-description
	 */
	description?: string;
	/**
	 * The `Task`'s parameters.
	 * @see https://tekton.dev/docs/pipelines/tasks/#specifying-parameters
	 */
	params?: TaskSpecParam[];
	results?: TaskSpecResult[];
	workspaces?: TaskWorkspace[];
	/**
	 * The steps that will be executed as part of the Task. The `Step` should
	 * fit the (container contract)[https://tekton.dev/docs/pipelines/container-contract/]
	 * @see https://tekton.dev/docs/pipelines/tasks/#defining-steps
	 */
	steps?: TaskStep[];
}
/**
 * Properties used to create the Task.
 */
export interface TaskProps {
	/**
	 * The object [metadata](https://kubernetes.io/docs/concepts/overview/working-with-objects/#required-fields)
	 * that conforms to standard Kubernetes metadata.
	 */
	metadata?: ApiObjectMetadata;
	/**
	 * The `spec` is the configuration of the `Task` object.
	 */
	spec?: TaskSpec;
}
/**
 * A Tekton Task, which is
 * > a collection of Steps that you define and arrange in
 * > a specific order of execution as part of your continuous integration flow. A
 * > Task executes as a Pod on your Kubernetes cluster. A Task is available within a
 * > specific namespace, while a ClusterTask is available across the entire
 * > cluster.
 *
 * @see https://tekton.dev/docs/pipelines/tasks/
 */
export declare class Task extends ApiObject {
	/**
	 * Returns the apiVersion and kind for "Task"
	 */
	static readonly GVK: GroupVersionKind;
	/**
	 * Renders a Kubernetes manifest for "Task".
	 *
	 * This can be used to inline resource manifests inside other objects (e.g. as templates).
	 *
	 * @param props initialization props
	 */
	static manifest(props?: TaskProps): any;
	private readonly _metadata?;
	private readonly _spec?;
	/**
	 * The configuration for the `Task`.
	 */
	get spec(): TaskSpec | undefined;
	/**
	 * Defines a "Task" API object
	 * @param scope the scope in which to define this object
	 * @param id a scope-local name for the object
	 * @param props initialization props
	 */
	constructor(scope: Construct, id: string, props?: TaskProps);
	/**
	 * Renders the object to Kubernetes JSON.
	 */
	toJson(): any;
}
