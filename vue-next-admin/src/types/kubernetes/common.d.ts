import {
	ContainerPort,
	EnvFromSource,
	EnvVar,
	Lifecycle,
	Probe,
	ResourceRequirements,
	SecurityContext,
	VolumeDevice,
	VolumeMount,
} from 'kubernetes-types/core/v1';

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

export interface ContainerType {
	isIntiContainer: boolean;
	/**
	 * Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	 */
	args?: Array<string>;
	/**
	 * Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	 */
	command?: Array<string>;
	/**
	 * List of environment variables to set in the container. Cannot be updated.
	 */
	env?: Array<EnvVar>;
	/**
	 * List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.
	 */
	envFrom?: Array<EnvFromSource>;
	/**
	 * Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.
	 */
	image?: string;
	/**
	 * Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	 *
	 *
	 */
	imagePullPolicy?: string;
	/**
	 * Actions that the management system should take in response to container lifecycle events. Cannot be updated.
	 */
	lifecycle?: Lifecycle;
	/**
	 * Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	 */
	livenessProbe?: Probe;
	/**
	 * Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.
	 */
	name: string;
	/**
	 * List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default "0.0.0.0" address inside a container will be accessible from the network. Modifying this array with strategic merge patch may corrupt the data. For more information See https://github.com/kubernetes/kubernetes/issues/108255. Cannot be updated.
	 */
	ports?: Array<ContainerPort>;
	/**
	 * Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	 */
	readinessProbe?: Probe;
	/**
	 * Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	 */
	resources?: ResourceRequirements;
	/**
	 * SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	 */
	securityContext?: SecurityContext;
	/**
	 * StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	 */
	startupProbe?: Probe;
	/**
	 * Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.
	 */
	stdin?: boolean;
	/**
	 * Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false
	 */
	stdinOnce?: boolean;
	/**
	 * Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.
	 */
	terminationMessagePath?: string;
	/**
	 * Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
	 *
	 *
	 */
	terminationMessagePolicy?: string;
	/**
	 * Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.
	 */
	tty?: boolean;
	/**
	 * volumeDevices is the list of block devices to be used by the container.
	 */
	volumeDevices?: Array<VolumeDevice>;
	/**
	 * Pod volumes to mount into the container's filesystem. Cannot be updated.
	 */
	volumeMounts?: Array<VolumeMount>;
	/**
	 * Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
	 */
	workingDir?: string;
}
