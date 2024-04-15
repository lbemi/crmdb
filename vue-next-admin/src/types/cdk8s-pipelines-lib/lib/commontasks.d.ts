import { TaskBuilder } from 'cdk8s-pipelines';
import { Construct } from 'constructs';
export declare const DefaultCatalogSourceNamespace: string;
/**
 * Uses the Tekton Hub [openshift_client](https://hub.tekton.dev/tekton/task/openshift-client)
 * and the `oc apply -f` command to apply the YAML representation of `input`.
 *
 * @param scope The parent [Construct](https://cdk8s.io/docs/latest/basics/constructs/).
 * @param id The `id` of the construct. Must be unique for each one in a chart.
 * @param input The input object that will be converted to YAML and applied to the cluster using `oc apply -f`
 * @constructor
 */
export declare function ApplyObjectTask(scope: Construct, id: string, input: any): TaskBuilder;
/**
 * Creates a Namespace document with the given name.
 *
 * This function creates a `TaskBuilder` to generate a task that uses the
 * openshift_client Tekton Task to apply a namespace document to the
 * cluster to create a namespace with the provided value.
 *
 * ```typescript
 * const task = CreateNameSpaceTask(this, 'create-namespace', 'my-example-namespace');
 * ```
 *
 * @param scope The parent [Construct](https://cdk8s.io/docs/latest/basics/constructs/).
 * @param id The `id` of the construct. Must be unique for each one in a chart.
 * @param name The string value name of the namespace to configure.
 *
 * @see
 */
export declare function CreateNamespace(scope: Construct, id: string, name: string): TaskBuilder;
/**
 * Creates a builder for the pipeline task that registers the IBM Operator
 * Catalog using `oc apply -f` with a configuration file.
 *
 * This function creates a `TaskBuilder` to generate a task that uses the
 * openshift_client Tekton Task.
 *
 * ```typescript
 * const task = RegisterIBMOperatorCatalog(this, 'register-ibm-operator-catalog', '45');
 * ```
 *
 * @param scope The parent [Construct](https://cdk8s.io/docs/latest/basics/constructs/).
 * @param id The `id` of the construct. Must be unique for each one in a chart.
 * @param labels The labels to give the CatalogSource.
 * @param refreshIntervalMin The amount of time, in minutes, between catalog refreshes.
 * @constructor
 */
export declare function RegisterIBMOperatorCatalog(scope: Construct, id: string, labels: any, refreshIntervalMin?: number): TaskBuilder;
/**
 * Creates a builder for a pipeline task that creates an OperatorGroup
 * using `oc apply -f` with a configuration file.
 *
 * This function creates a `TaskBuilder` to generate a task that uses the
 * openshift_client Tekton Task.
 *
 * ```typescript
 * const task = CreateOperatorGroup(this, 'create-operator-group', 'ibm-eventautomation-operatorgroup', 'ibm-eventstreams');
 * ```
 *
 * @param scope The parent [Construct](https://cdk8s.io/docs/latest/basics/constructs/).
 * @param id The `id` of the construct. Must be unique for each one in a chart.
 * @param ns The namespace for the OperatorGroup.
 * @param name The name of the OperatorGroup.
 * @param targetNs If provided, it is the target namespace for the OperatorGroup.
 * @constructor
 */
export declare function CreateOperatorGroup(scope: Construct, id: string, ns: string, name: string, targetNs?: string): TaskBuilder;
/**
 * Creates a builder for a pipeline task that creates a Subscription using
 * `oc apply -f` with a configuration file.
 *
 * This function creates a `TaskBuilder` to generate a task that uses the
 * openshift_client Tekton Task.
 *
 * ```typescript
 * const task =  Subscribe(this, 'install-eventstreams', 'ibm-eventstreams', 'ibm-eventstreams', 'ibm-operator-catalog', 'stable');
 * ```
 *
 * @param scope The parent [Construct](https://cdk8s.io/docs/latest/basics/constructs/).
 * @param id The `id` of the construct. Must be unique for each one in a chart.
 * @param ns The namespace for the Subscription.
 * @param name The name of the Subscription.
 * @param catalogSource The name of the catalog source. If using IBM operators, that is `ibm-operator-catalog`
 * @param channel The name of the channel. It defaults to `stable`, but it could be something like `v3.3`.
 * @constructor
 */
export declare function Subscribe(scope: Construct, id: string, ns: string, name: string, catalogSource: string, channel?: string): TaskBuilder;
/**
 * Creates a builder for a pipeline task that uses an busybox image to echo out the
 * given message.
 *
 * @param scope The parent [Construct](https://cdk8s.io/docs/latest/basics/constructs/).
 * @param id The `id` of the construct. Must be unique for each one in a chart.
 * @param message The message for the task to echo to the output.
 * @constructor
 */
export declare function EchoMessage(scope: Construct, id: string, message: string): TaskBuilder;
