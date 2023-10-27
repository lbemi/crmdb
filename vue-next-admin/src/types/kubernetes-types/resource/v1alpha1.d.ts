import { NodeSelector } from "../core/v1";
import { ListMeta, ObjectMeta } from "../meta/v1";
/**
 * AllocationResult contains attributed of an allocated resource.
 */
export interface AllocationResult {
    /**
     * This field will get set by the resource driver after it has allocated the resource driver to inform the scheduler where it can schedule Pods using the ResourceClaim.
     *
     * Setting this field is optional. If null, the resource is available everywhere.
     */
    availableOnNodes?: NodeSelector;
    /**
     * ResourceHandle contains arbitrary data returned by the driver after a successful allocation. This is opaque for Kubernetes. Driver documentation may explain to users how to interpret this data if needed.
     *
     * The maximum size of this field is 16KiB. This may get increased in the future, but not reduced.
     */
    resourceHandle?: string;
    /**
     * Shareable determines whether the resource supports more than one consumer at a time.
     */
    shareable?: boolean;
}
/**
 * PodScheduling objects hold information that is needed to schedule a Pod with ResourceClaims that use "WaitForFirstConsumer" allocation mode.
 *
 * This is an alpha type and requires enabling the DynamicResourceAllocation feature gate.
 */
export interface PodScheduling {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "resource.k8s.io/v1alpha1";
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "PodScheduling";
    /**
     * Standard object metadata
     */
    metadata?: ObjectMeta;
    /**
     * Spec describes where resources for the Pod are needed.
     */
    spec: PodSchedulingSpec;
    /**
     * Status describes where resources for the Pod can be allocated.
     */
    status?: PodSchedulingStatus;
}
/**
 * PodSchedulingList is a collection of Pod scheduling objects.
 */
export interface PodSchedulingList {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "resource.k8s.io/v1alpha1";
    /**
     * Items is the list of PodScheduling objects.
     */
    items: Array<PodScheduling>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "PodSchedulingList";
    /**
     * Standard list metadata
     */
    metadata?: ListMeta;
}
/**
 * PodSchedulingSpec describes where resources for the Pod are needed.
 */
export interface PodSchedulingSpec {
    /**
     * PotentialNodes lists nodes where the Pod might be able to run.
     *
     * The size of this field is limited to 128. This is large enough for many clusters. Larger clusters may need more attempts to find a node that suits all pending resources. This may get increased in the future, but not reduced.
     */
    potentialNodes?: Array<string>;
    /**
     * SelectedNode is the node for which allocation of ResourceClaims that are referenced by the Pod and that use "WaitForFirstConsumer" allocation is to be attempted.
     */
    selectedNode?: string;
}
/**
 * PodSchedulingStatus describes where resources for the Pod can be allocated.
 */
export interface PodSchedulingStatus {
    /**
     * ResourceClaims describes resource availability for each pod.spec.resourceClaim entry where the corresponding ResourceClaim uses "WaitForFirstConsumer" allocation mode.
     */
    resourceClaims?: Array<ResourceClaimSchedulingStatus>;
}
/**
 * ResourceClaim describes which resources are needed by a resource consumer. Its status tracks whether the resource has been allocated and what the resulting attributes are.
 *
 * This is an alpha type and requires enabling the DynamicResourceAllocation feature gate.
 */
export interface ResourceClaim {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "resource.k8s.io/v1alpha1";
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ResourceClaim";
    /**
     * Standard object metadata
     */
    metadata?: ObjectMeta;
    /**
     * Spec describes the desired attributes of a resource that then needs to be allocated. It can only be set once when creating the ResourceClaim.
     */
    spec: ResourceClaimSpec;
    /**
     * Status describes whether the resource is available and with which attributes.
     */
    status?: ResourceClaimStatus;
}
/**
 * ResourceClaimConsumerReference contains enough information to let you locate the consumer of a ResourceClaim. The user must be a resource in the same namespace as the ResourceClaim.
 */
export interface ResourceClaimConsumerReference {
    /**
     * APIGroup is the group for the resource being referenced. It is empty for the core API. This matches the group in the APIVersion that is used when creating the resources.
     */
    apiGroup?: string;
    /**
     * Name is the name of resource being referenced.
     */
    name: string;
    /**
     * Resource is the type of resource being referenced, for example "pods".
     */
    resource: string;
    /**
     * UID identifies exactly one incarnation of the resource.
     */
    uid: string;
}
/**
 * ResourceClaimList is a collection of claims.
 */
export interface ResourceClaimList {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "resource.k8s.io/v1alpha1";
    /**
     * Items is the list of resource claims.
     */
    items: Array<ResourceClaim>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ResourceClaimList";
    /**
     * Standard list metadata
     */
    metadata?: ListMeta;
}
/**
 * ResourceClaimParametersReference contains enough information to let you locate the parameters for a ResourceClaim. The object must be in the same namespace as the ResourceClaim.
 */
export interface ResourceClaimParametersReference {
    /**
     * APIGroup is the group for the resource being referenced. It is empty for the core API. This matches the group in the APIVersion that is used when creating the resources.
     */
    apiGroup?: string;
    /**
     * Kind is the type of resource being referenced. This is the same value as in the parameter object's metadata, for example "ConfigMap".
     */
    kind: string;
    /**
     * Name is the name of resource being referenced.
     */
    name: string;
}
/**
 * ResourceClaimSchedulingStatus contains information about one particular ResourceClaim with "WaitForFirstConsumer" allocation mode.
 */
export interface ResourceClaimSchedulingStatus {
    /**
     * Name matches the pod.spec.resourceClaims[*].Name field.
     */
    name?: string;
    /**
     * UnsuitableNodes lists nodes that the ResourceClaim cannot be allocated for.
     *
     * The size of this field is limited to 128, the same as for PodSchedulingSpec.PotentialNodes. This may get increased in the future, but not reduced.
     */
    unsuitableNodes?: Array<string>;
}
/**
 * ResourceClaimSpec defines how a resource is to be allocated.
 */
export interface ResourceClaimSpec {
    /**
     * Allocation can start immediately or when a Pod wants to use the resource. "WaitForFirstConsumer" is the default.
     */
    allocationMode?: string;
    /**
     * ParametersRef references a separate object with arbitrary parameters that will be used by the driver when allocating a resource for the claim.
     *
     * The object must be in the same namespace as the ResourceClaim.
     */
    parametersRef?: ResourceClaimParametersReference;
    /**
     * ResourceClassName references the driver and additional parameters via the name of a ResourceClass that was created as part of the driver deployment.
     */
    resourceClassName: string;
}
/**
 * ResourceClaimStatus tracks whether the resource has been allocated and what the resulting attributes are.
 */
export interface ResourceClaimStatus {
    /**
     * Allocation is set by the resource driver once a resource has been allocated successfully. If this is not specified, the resource is not yet allocated.
     */
    allocation?: AllocationResult;
    /**
     * DeallocationRequested indicates that a ResourceClaim is to be deallocated.
     *
     * The driver then must deallocate this claim and reset the field together with clearing the Allocation field.
     *
     * While DeallocationRequested is set, no new consumers may be added to ReservedFor.
     */
    deallocationRequested?: boolean;
    /**
     * DriverName is a copy of the driver name from the ResourceClass at the time when allocation started.
     */
    driverName?: string;
    /**
     * ReservedFor indicates which entities are currently allowed to use the claim. A Pod which references a ResourceClaim which is not reserved for that Pod will not be started.
     *
     * There can be at most 32 such reservations. This may get increased in the future, but not reduced.
     */
    reservedFor?: Array<ResourceClaimConsumerReference>;
}
/**
 * ResourceClaimTemplate is used to produce ResourceClaim objects.
 */
export interface ResourceClaimTemplate {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "resource.k8s.io/v1alpha1";
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ResourceClaimTemplate";
    /**
     * Standard object metadata
     */
    metadata?: ObjectMeta;
    /**
     * Describes the ResourceClaim that is to be generated.
     *
     * This field is immutable. A ResourceClaim will get created by the control plane for a Pod when needed and then not get updated anymore.
     */
    spec: ResourceClaimTemplateSpec;
}
/**
 * ResourceClaimTemplateList is a collection of claim templates.
 */
export interface ResourceClaimTemplateList {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "resource.k8s.io/v1alpha1";
    /**
     * Items is the list of resource claim templates.
     */
    items: Array<ResourceClaimTemplate>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ResourceClaimTemplateList";
    /**
     * Standard list metadata
     */
    metadata?: ListMeta;
}
/**
 * ResourceClaimTemplateSpec contains the metadata and fields for a ResourceClaim.
 */
export interface ResourceClaimTemplateSpec {
    /**
     * ObjectMeta may contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.
     */
    metadata?: ObjectMeta;
    /**
     * Spec for the ResourceClaim. The entire content is copied unchanged into the ResourceClaim that gets created from this template. The same fields as in a ResourceClaim are also valid here.
     */
    spec: ResourceClaimSpec;
}
/**
 * ResourceClass is used by administrators to influence how resources are allocated.
 *
 * This is an alpha type and requires enabling the DynamicResourceAllocation feature gate.
 */
export interface ResourceClass {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "resource.k8s.io/v1alpha1";
    /**
     * DriverName defines the name of the dynamic resource driver that is used for allocation of a ResourceClaim that uses this class.
     *
     * Resource drivers have a unique name in forward domain order (acme.example.com).
     */
    driverName: string;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ResourceClass";
    /**
     * Standard object metadata
     */
    metadata?: ObjectMeta;
    /**
     * ParametersRef references an arbitrary separate object that may hold parameters that will be used by the driver when allocating a resource that uses this class. A dynamic resource driver can distinguish between parameters stored here and and those stored in ResourceClaimSpec.
     */
    parametersRef?: ResourceClassParametersReference;
    /**
     * Only nodes matching the selector will be considered by the scheduler when trying to find a Node that fits a Pod when that Pod uses a ResourceClaim that has not been allocated yet.
     *
     * Setting this field is optional. If null, all nodes are candidates.
     */
    suitableNodes?: NodeSelector;
}
/**
 * ResourceClassList is a collection of classes.
 */
export interface ResourceClassList {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "resource.k8s.io/v1alpha1";
    /**
     * Items is the list of resource classes.
     */
    items: Array<ResourceClass>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ResourceClassList";
    /**
     * Standard list metadata
     */
    metadata?: ListMeta;
}
/**
 * ResourceClassParametersReference contains enough information to let you locate the parameters for a ResourceClass.
 */
export interface ResourceClassParametersReference {
    /**
     * APIGroup is the group for the resource being referenced. It is empty for the core API. This matches the group in the APIVersion that is used when creating the resources.
     */
    apiGroup?: string;
    /**
     * Kind is the type of resource being referenced. This is the same value as in the parameter object's metadata.
     */
    kind: string;
    /**
     * Name is the name of resource being referenced.
     */
    name: string;
    /**
     * Namespace that contains the referenced resource. Must be empty for cluster-scoped resources and non-empty for namespaced resources.
     */
    namespace?: string;
}
