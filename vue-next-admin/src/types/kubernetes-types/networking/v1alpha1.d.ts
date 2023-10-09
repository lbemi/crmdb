import { ListMeta, ObjectMeta } from "../meta/v1";
import { NodeSelector } from "../core/v1";
/**
 * ClusterCIDR represents a single configuration for per-Node Pod CIDR allocations when the MultiCIDRRangeAllocator is enabled (see the config for kube-controller-manager).  A cluster may have any number of ClusterCIDR resources, all of which will be considered when allocating a CIDR for a Node.  A ClusterCIDR is eligible to be used for a given Node when the node selector matches the node in question and has free CIDRs to allocate.  In case of multiple matching ClusterCIDR resources, the allocator will attempt to break ties using internal heuristics, but any ClusterCIDR whose node selector matches the Node may be used.
 */
export interface ClusterCIDR {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "networking.k8s.io/v1alpha1";
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ClusterCIDR";
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta;
    /**
     * Spec is the desired state of the ClusterCIDR. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    spec?: ClusterCIDRSpec;
}
/**
 * ClusterCIDRList contains a list of ClusterCIDR.
 */
export interface ClusterCIDRList {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: "networking.k8s.io/v1alpha1";
    /**
     * Items is the list of ClusterCIDRs.
     */
    items: Array<ClusterCIDR>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: "ClusterCIDRList";
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ListMeta;
}
/**
 * ClusterCIDRSpec defines the desired state of ClusterCIDR.
 */
export interface ClusterCIDRSpec {
    /**
     * IPv4 defines an IPv4 IP block in CIDR notation(e.g. "10.0.0.0/8"). At least one of IPv4 and IPv6 must be specified. This field is immutable.
     */
    ipv4?: string;
    /**
     * IPv6 defines an IPv6 IP block in CIDR notation(e.g. "2001:db8::/64"). At least one of IPv4 and IPv6 must be specified. This field is immutable.
     */
    ipv6?: string;
    /**
     * NodeSelector defines which nodes the config is applicable to. An empty or nil NodeSelector selects all nodes. This field is immutable.
     */
    nodeSelector?: NodeSelector;
    /**
     * PerNodeHostBits defines the number of host bits to be configured per node. A subnet mask determines how much of the address is used for network bits and host bits. For example an IPv4 address of 192.168.0.0/24, splits the address into 24 bits for the network portion and 8 bits for the host portion. To allocate 256 IPs, set this field to 8 (a /24 mask for IPv4 or a /120 for IPv6). Minimum value is 4 (16 IPs). This field is immutable.
     */
    perNodeHostBits: number;
}
