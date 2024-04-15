"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.EchoMessage = exports.Subscribe = exports.CreateOperatorGroup = exports.RegisterIBMOperatorCatalog = exports.CreateNamespace = exports.ApplyObjectTask = exports.DefaultCatalogSourceNamespace = void 0;
const cdk8s_1 = require("cdk8s");
const cdk8s_pipelines_1 = require("cdk8s-pipelines");
const tektonHubTasks_1 = require("./tektonHub/tektonHubTasks");
exports.DefaultCatalogSourceNamespace = 'openshift-marketplace';
/**
 * Uses the Tekton Hub [openshift_client](https://hub.tekton.dev/tekton/task/openshift-client)
 * and the `oc apply -f` command to apply the YAML representation of `input`.
 *
 * @param scope The parent [Construct](https://cdk8s.io/docs/latest/basics/constructs/).
 * @param id The `id` of the construct. Must be unique for each one in a chart.
 * @param input The input object that will be converted to YAML and applied to the cluster using `oc apply -f`
 * @constructor
 */
function ApplyObjectTask(scope, id, input) {
    return (0, tektonHubTasks_1.openshift_client)(scope, id).withStringParam(new cdk8s_pipelines_1.ParameterBuilder('SCRIPT').withValue(['cat <<EOF | oc apply -f -', cdk8s_1.Yaml.stringify(input), 'EOF', ''].join('\n')));
}
exports.ApplyObjectTask = ApplyObjectTask;
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
function CreateNamespace(scope, id, name) {
    // TODO: validate the namespace and potentially throw an error if it is
    // malformed? Also, consider using objects from the cdk8s-plus for this.
    const createNamespace = {
        apiVersion: 'v1',
        kind: 'Namespace',
        metadata: {
            name: name,
            labels: {
                name: name,
            },
        },
    };
    return ApplyObjectTask(scope, id, createNamespace);
}
exports.CreateNamespace = CreateNamespace;
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
function RegisterIBMOperatorCatalog(scope, id, labels, refreshIntervalMin = 60) {
    const ibmOperatorCatalogSource = {
        apiVersion: 'operators.coreos.com/v1alpha1',
        kind: 'CatalogSource',
        metadata: {
            name: 'ibm-operator-catalog',
            namespace: 'openshift-marketplace',
            labels: labels,
        },
        spec: {
            displayName: 'IBM Operator Catalog',
            publisher: 'IBM',
            sourceType: 'grpc',
            image: 'icr.io/cpopen/ibm-operator-catalog',
            updateStrategy: {
                registryPoll: {
                    interval: `${refreshIntervalMin.toString()}m`,
                },
            },
        },
    };
    return ApplyObjectTask(scope, id, ibmOperatorCatalogSource);
}
exports.RegisterIBMOperatorCatalog = RegisterIBMOperatorCatalog;
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
function CreateOperatorGroup(scope, id, ns, name, targetNs = '') {
    let operatorGroupSource = {};
    if (targetNs) {
        operatorGroupSource = {
            apiVersion: 'operators.coreos.com/v1',
            kind: 'OperatorGroup',
            metadata: {
                name: name,
                namespace: ns,
            },
            spec: {
                targetNamespaces: [targetNs],
            },
        };
    }
    else {
        operatorGroupSource = {
            apiVersion: 'operators.coreos.com/v1',
            kind: 'OperatorGroup',
            metadata: {
                name: name,
                namespace: ns,
            },
        };
    }
    return ApplyObjectTask(scope, id, operatorGroupSource);
}
exports.CreateOperatorGroup = CreateOperatorGroup;
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
function Subscribe(scope, id, ns, name, catalogSource, channel = 'stable') {
    const subscription = {
        apiVersion: 'operators.coreos.com/v1alpha1',
        kind: 'Subscription',
        metadata: {
            name: name,
            namespace: ns,
        },
        spec: {
            channel: channel,
            name: name,
            source: catalogSource,
            sourceNamespace: exports.DefaultCatalogSourceNamespace,
        },
    };
    return ApplyObjectTask(scope, id, subscription);
}
exports.Subscribe = Subscribe;
/**
 * Creates a builder for a pipeline task that uses an busybox image to echo out the
 * given message.
 *
 * @param scope The parent [Construct](https://cdk8s.io/docs/latest/basics/constructs/).
 * @param id The `id` of the construct. Must be unique for each one in a chart.
 * @param message The message for the task to echo to the output.
 * @constructor
 */
function EchoMessage(scope, id, message) {
    // TODO: sanitize input for the message to make sure it has no injection attacks
    return new cdk8s_pipelines_1.TaskBuilder(scope, id)
        .withName('echo-message')
        .withDescription('Echos a message out to the pipeline')
        .withStep(new cdk8s_pipelines_1.TaskStepBuilder()
        .withName('echo-message')
        .withImage('docker.io/library/busybox:latest')
        .withCommand(['echo', `\"${message}\"`]));
}
exports.EchoMessage = EchoMessage;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiY29tbW9udGFza3MuanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyIuLi9zcmMvY29tbW9udGFza3MudHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6Ijs7O0FBQUEsaUNBQTZCO0FBQzdCLHFEQUFpRjtBQUVqRiwrREFBOEQ7QUFFakQsUUFBQSw2QkFBNkIsR0FBVyx1QkFBdUIsQ0FBQztBQUU3RTs7Ozs7Ozs7R0FRRztBQUNILFNBQWdCLGVBQWUsQ0FBQyxLQUFnQixFQUFFLEVBQVUsRUFBRSxLQUFVO0lBQ3RFLE9BQU8sSUFBQSxpQ0FBZ0IsRUFBQyxLQUFLLEVBQUUsRUFBRSxDQUFDLENBQUMsZUFBZSxDQUFDLElBQUksa0NBQWdCLENBQUMsUUFBUSxDQUFDLENBQUMsU0FBUyxDQUN6RixDQUFDLDJCQUEyQixFQUFFLFlBQUksQ0FBQyxTQUFTLENBQUMsS0FBSyxDQUFDLEVBQUUsS0FBSyxFQUFFLEVBQUUsQ0FBQyxDQUFDLElBQUksQ0FBQyxJQUFJLENBQUMsQ0FBQyxDQUM1RSxDQUFDO0FBQ0osQ0FBQztBQUpELDBDQUlDO0FBRUQ7Ozs7Ozs7Ozs7Ozs7Ozs7R0FnQkc7QUFDSCxTQUFnQixlQUFlLENBQUMsS0FBZ0IsRUFBRSxFQUFVLEVBQUUsSUFBWTtJQUN4RSx1RUFBdUU7SUFDdkUsd0VBQXdFO0lBQ3hFLE1BQU0sZUFBZSxHQUFHO1FBQ3RCLFVBQVUsRUFBRSxJQUFJO1FBQ2hCLElBQUksRUFBRSxXQUFXO1FBQ2pCLFFBQVEsRUFBRTtZQUNSLElBQUksRUFBRSxJQUFJO1lBQ1YsTUFBTSxFQUFFO2dCQUNOLElBQUksRUFBRSxJQUFJO2FBQ1g7U0FDRjtLQUNGLENBQUM7SUFFRixPQUFPLGVBQWUsQ0FBQyxLQUFLLEVBQUUsRUFBRSxFQUFFLGVBQWUsQ0FBQyxDQUFDO0FBQ3JELENBQUM7QUFmRCwwQ0FlQztBQUVEOzs7Ozs7Ozs7Ozs7Ozs7O0dBZ0JHO0FBQ0gsU0FBZ0IsMEJBQTBCLENBQUMsS0FBZ0IsRUFBRSxFQUFVLEVBQUUsTUFBVyxFQUFFLHFCQUE2QixFQUFFO0lBQ25ILE1BQU0sd0JBQXdCLEdBQUc7UUFDL0IsVUFBVSxFQUFFLCtCQUErQjtRQUMzQyxJQUFJLEVBQUUsZUFBZTtRQUNyQixRQUFRLEVBQUU7WUFDUixJQUFJLEVBQUUsc0JBQXNCO1lBQzVCLFNBQVMsRUFBRSx1QkFBdUI7WUFDbEMsTUFBTSxFQUFFLE1BQU07U0FDZjtRQUNELElBQUksRUFBRTtZQUNKLFdBQVcsRUFBRSxzQkFBc0I7WUFDbkMsU0FBUyxFQUFFLEtBQUs7WUFDaEIsVUFBVSxFQUFFLE1BQU07WUFDbEIsS0FBSyxFQUFFLG9DQUFvQztZQUMzQyxjQUFjLEVBQUU7Z0JBQ2QsWUFBWSxFQUFFO29CQUNaLFFBQVEsRUFBRSxHQUFHLGtCQUFrQixDQUFDLFFBQVEsRUFBRSxHQUFHO2lCQUM5QzthQUNGO1NBQ0Y7S0FDRixDQUFDO0lBRUYsT0FBTyxlQUFlLENBQUMsS0FBSyxFQUFFLEVBQUUsRUFBRSx3QkFBd0IsQ0FBQyxDQUFDO0FBQzlELENBQUM7QUF2QkQsZ0VBdUJDO0FBRUQ7Ozs7Ozs7Ozs7Ozs7Ozs7O0dBaUJHO0FBQ0gsU0FBZ0IsbUJBQW1CLENBQUMsS0FBZ0IsRUFBRSxFQUFVLEVBQUUsRUFBVSxFQUFFLElBQVksRUFBRSxXQUFtQixFQUFFO0lBQy9HLElBQUksbUJBQW1CLEdBQUcsRUFBRSxDQUFDO0lBQzdCLElBQUksUUFBUSxFQUFFO1FBQ1osbUJBQW1CLEdBQUc7WUFDcEIsVUFBVSxFQUFFLHlCQUF5QjtZQUNyQyxJQUFJLEVBQUUsZUFBZTtZQUNyQixRQUFRLEVBQUU7Z0JBQ1IsSUFBSSxFQUFFLElBQUk7Z0JBQ1YsU0FBUyxFQUFFLEVBQUU7YUFDZDtZQUNELElBQUksRUFBRTtnQkFDSixnQkFBZ0IsRUFBRSxDQUFDLFFBQVEsQ0FBQzthQUM3QjtTQUNGLENBQUM7S0FDSDtTQUFNO1FBQ0wsbUJBQW1CLEdBQUc7WUFDcEIsVUFBVSxFQUFFLHlCQUF5QjtZQUNyQyxJQUFJLEVBQUUsZUFBZTtZQUNyQixRQUFRLEVBQUU7Z0JBQ1IsSUFBSSxFQUFFLElBQUk7Z0JBQ1YsU0FBUyxFQUFFLEVBQUU7YUFDZDtTQUNGLENBQUM7S0FDSDtJQUVELE9BQU8sZUFBZSxDQUFDLEtBQUssRUFBRSxFQUFFLEVBQUUsbUJBQW1CLENBQUMsQ0FBQztBQUN6RCxDQUFDO0FBMUJELGtEQTBCQztBQUVEOzs7Ozs7Ozs7Ozs7Ozs7Ozs7R0FrQkc7QUFDSCxTQUFnQixTQUFTLENBQUMsS0FBZ0IsRUFBRSxFQUFVLEVBQUUsRUFBVSxFQUFFLElBQVksRUFBRSxhQUFxQixFQUFFLFVBQWtCLFFBQVE7SUFDakksTUFBTSxZQUFZLEdBQUc7UUFDbkIsVUFBVSxFQUFFLCtCQUErQjtRQUMzQyxJQUFJLEVBQUUsY0FBYztRQUNwQixRQUFRLEVBQUU7WUFDUixJQUFJLEVBQUUsSUFBSTtZQUNWLFNBQVMsRUFBRSxFQUFFO1NBQ2Q7UUFDRCxJQUFJLEVBQUU7WUFDSixPQUFPLEVBQUUsT0FBTztZQUNoQixJQUFJLEVBQUUsSUFBSTtZQUNWLE1BQU0sRUFBRSxhQUFhO1lBQ3JCLGVBQWUsRUFBRSxxQ0FBNkI7U0FDL0M7S0FDRixDQUFDO0lBQ0YsT0FBTyxlQUFlLENBQUMsS0FBSyxFQUFFLEVBQUUsRUFBRSxZQUFZLENBQUMsQ0FBQztBQUNsRCxDQUFDO0FBaEJELDhCQWdCQztBQUVEOzs7Ozs7OztHQVFHO0FBQ0gsU0FBZ0IsV0FBVyxDQUFDLEtBQWdCLEVBQUUsRUFBVSxFQUFFLE9BQWU7SUFDdkUsZ0ZBQWdGO0lBQ2hGLE9BQU8sSUFBSSw2QkFBVyxDQUFDLEtBQUssRUFBRSxFQUFFLENBQUM7U0FDOUIsUUFBUSxDQUFDLGNBQWMsQ0FBQztTQUN4QixlQUFlLENBQUMscUNBQXFDLENBQUM7U0FDdEQsUUFBUSxDQUFDLElBQUksaUNBQWUsRUFBRTtTQUM1QixRQUFRLENBQUMsY0FBYyxDQUFDO1NBQ3hCLFNBQVMsQ0FBQyxrQ0FBa0MsQ0FBQztTQUM3QyxXQUFXLENBQUMsQ0FBQyxNQUFNLEVBQUUsS0FBSyxPQUFPLElBQUksQ0FBQyxDQUFDLENBQUMsQ0FBQztBQUNoRCxDQUFDO0FBVEQsa0NBU0MiLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgeyBZYW1sIH0gZnJvbSAnY2RrOHMnO1xuaW1wb3J0IHsgUGFyYW1ldGVyQnVpbGRlciwgVGFza0J1aWxkZXIsIFRhc2tTdGVwQnVpbGRlciB9IGZyb20gJ2NkazhzLXBpcGVsaW5lcyc7XG5pbXBvcnQgeyBDb25zdHJ1Y3QgfSBmcm9tICdjb25zdHJ1Y3RzJztcbmltcG9ydCB7IG9wZW5zaGlmdF9jbGllbnQgfSBmcm9tICcuL3Rla3Rvbkh1Yi90ZWt0b25IdWJUYXNrcyc7XG5cbmV4cG9ydCBjb25zdCBEZWZhdWx0Q2F0YWxvZ1NvdXJjZU5hbWVzcGFjZTogc3RyaW5nID0gJ29wZW5zaGlmdC1tYXJrZXRwbGFjZSc7XG5cbi8qKlxuICogVXNlcyB0aGUgVGVrdG9uIEh1YiBbb3BlbnNoaWZ0X2NsaWVudF0oaHR0cHM6Ly9odWIudGVrdG9uLmRldi90ZWt0b24vdGFzay9vcGVuc2hpZnQtY2xpZW50KVxuICogYW5kIHRoZSBgb2MgYXBwbHkgLWZgIGNvbW1hbmQgdG8gYXBwbHkgdGhlIFlBTUwgcmVwcmVzZW50YXRpb24gb2YgYGlucHV0YC5cbiAqXG4gKiBAcGFyYW0gc2NvcGUgVGhlIHBhcmVudCBbQ29uc3RydWN0XShodHRwczovL2NkazhzLmlvL2RvY3MvbGF0ZXN0L2Jhc2ljcy9jb25zdHJ1Y3RzLykuXG4gKiBAcGFyYW0gaWQgVGhlIGBpZGAgb2YgdGhlIGNvbnN0cnVjdC4gTXVzdCBiZSB1bmlxdWUgZm9yIGVhY2ggb25lIGluIGEgY2hhcnQuXG4gKiBAcGFyYW0gaW5wdXQgVGhlIGlucHV0IG9iamVjdCB0aGF0IHdpbGwgYmUgY29udmVydGVkIHRvIFlBTUwgYW5kIGFwcGxpZWQgdG8gdGhlIGNsdXN0ZXIgdXNpbmcgYG9jIGFwcGx5IC1mYFxuICogQGNvbnN0cnVjdG9yXG4gKi9cbmV4cG9ydCBmdW5jdGlvbiBBcHBseU9iamVjdFRhc2soc2NvcGU6IENvbnN0cnVjdCwgaWQ6IHN0cmluZywgaW5wdXQ6IGFueSk6IFRhc2tCdWlsZGVyIHtcbiAgcmV0dXJuIG9wZW5zaGlmdF9jbGllbnQoc2NvcGUsIGlkKS53aXRoU3RyaW5nUGFyYW0obmV3IFBhcmFtZXRlckJ1aWxkZXIoJ1NDUklQVCcpLndpdGhWYWx1ZShcbiAgICBbJ2NhdCA8PEVPRiB8IG9jIGFwcGx5IC1mIC0nLCBZYW1sLnN0cmluZ2lmeShpbnB1dCksICdFT0YnLCAnJ10uam9pbignXFxuJykpLFxuICApO1xufVxuXG4vKipcbiAqIENyZWF0ZXMgYSBOYW1lc3BhY2UgZG9jdW1lbnQgd2l0aCB0aGUgZ2l2ZW4gbmFtZS5cbiAqXG4gKiBUaGlzIGZ1bmN0aW9uIGNyZWF0ZXMgYSBgVGFza0J1aWxkZXJgIHRvIGdlbmVyYXRlIGEgdGFzayB0aGF0IHVzZXMgdGhlXG4gKiBvcGVuc2hpZnRfY2xpZW50IFRla3RvbiBUYXNrIHRvIGFwcGx5IGEgbmFtZXNwYWNlIGRvY3VtZW50IHRvIHRoZVxuICogY2x1c3RlciB0byBjcmVhdGUgYSBuYW1lc3BhY2Ugd2l0aCB0aGUgcHJvdmlkZWQgdmFsdWUuXG4gKlxuICogYGBgdHlwZXNjcmlwdFxuICogY29uc3QgdGFzayA9IENyZWF0ZU5hbWVTcGFjZVRhc2sodGhpcywgJ2NyZWF0ZS1uYW1lc3BhY2UnLCAnbXktZXhhbXBsZS1uYW1lc3BhY2UnKTtcbiAqIGBgYFxuICpcbiAqIEBwYXJhbSBzY29wZSBUaGUgcGFyZW50IFtDb25zdHJ1Y3RdKGh0dHBzOi8vY2RrOHMuaW8vZG9jcy9sYXRlc3QvYmFzaWNzL2NvbnN0cnVjdHMvKS5cbiAqIEBwYXJhbSBpZCBUaGUgYGlkYCBvZiB0aGUgY29uc3RydWN0LiBNdXN0IGJlIHVuaXF1ZSBmb3IgZWFjaCBvbmUgaW4gYSBjaGFydC5cbiAqIEBwYXJhbSBuYW1lIFRoZSBzdHJpbmcgdmFsdWUgbmFtZSBvZiB0aGUgbmFtZXNwYWNlIHRvIGNvbmZpZ3VyZS5cbiAqXG4gKiBAc2VlXG4gKi9cbmV4cG9ydCBmdW5jdGlvbiBDcmVhdGVOYW1lc3BhY2Uoc2NvcGU6IENvbnN0cnVjdCwgaWQ6IHN0cmluZywgbmFtZTogc3RyaW5nKTogVGFza0J1aWxkZXIge1xuICAvLyBUT0RPOiB2YWxpZGF0ZSB0aGUgbmFtZXNwYWNlIGFuZCBwb3RlbnRpYWxseSB0aHJvdyBhbiBlcnJvciBpZiBpdCBpc1xuICAvLyBtYWxmb3JtZWQ/IEFsc28sIGNvbnNpZGVyIHVzaW5nIG9iamVjdHMgZnJvbSB0aGUgY2RrOHMtcGx1cyBmb3IgdGhpcy5cbiAgY29uc3QgY3JlYXRlTmFtZXNwYWNlID0ge1xuICAgIGFwaVZlcnNpb246ICd2MScsXG4gICAga2luZDogJ05hbWVzcGFjZScsXG4gICAgbWV0YWRhdGE6IHtcbiAgICAgIG5hbWU6IG5hbWUsXG4gICAgICBsYWJlbHM6IHtcbiAgICAgICAgbmFtZTogbmFtZSxcbiAgICAgIH0sXG4gICAgfSxcbiAgfTtcblxuICByZXR1cm4gQXBwbHlPYmplY3RUYXNrKHNjb3BlLCBpZCwgY3JlYXRlTmFtZXNwYWNlKTtcbn1cblxuLyoqXG4gKiBDcmVhdGVzIGEgYnVpbGRlciBmb3IgdGhlIHBpcGVsaW5lIHRhc2sgdGhhdCByZWdpc3RlcnMgdGhlIElCTSBPcGVyYXRvclxuICogQ2F0YWxvZyB1c2luZyBgb2MgYXBwbHkgLWZgIHdpdGggYSBjb25maWd1cmF0aW9uIGZpbGUuXG4gKlxuICogVGhpcyBmdW5jdGlvbiBjcmVhdGVzIGEgYFRhc2tCdWlsZGVyYCB0byBnZW5lcmF0ZSBhIHRhc2sgdGhhdCB1c2VzIHRoZVxuICogb3BlbnNoaWZ0X2NsaWVudCBUZWt0b24gVGFzay5cbiAqXG4gKiBgYGB0eXBlc2NyaXB0XG4gKiBjb25zdCB0YXNrID0gUmVnaXN0ZXJJQk1PcGVyYXRvckNhdGFsb2codGhpcywgJ3JlZ2lzdGVyLWlibS1vcGVyYXRvci1jYXRhbG9nJywgJzQ1Jyk7XG4gKiBgYGBcbiAqXG4gKiBAcGFyYW0gc2NvcGUgVGhlIHBhcmVudCBbQ29uc3RydWN0XShodHRwczovL2NkazhzLmlvL2RvY3MvbGF0ZXN0L2Jhc2ljcy9jb25zdHJ1Y3RzLykuXG4gKiBAcGFyYW0gaWQgVGhlIGBpZGAgb2YgdGhlIGNvbnN0cnVjdC4gTXVzdCBiZSB1bmlxdWUgZm9yIGVhY2ggb25lIGluIGEgY2hhcnQuXG4gKiBAcGFyYW0gbGFiZWxzIFRoZSBsYWJlbHMgdG8gZ2l2ZSB0aGUgQ2F0YWxvZ1NvdXJjZS5cbiAqIEBwYXJhbSByZWZyZXNoSW50ZXJ2YWxNaW4gVGhlIGFtb3VudCBvZiB0aW1lLCBpbiBtaW51dGVzLCBiZXR3ZWVuIGNhdGFsb2cgcmVmcmVzaGVzLlxuICogQGNvbnN0cnVjdG9yXG4gKi9cbmV4cG9ydCBmdW5jdGlvbiBSZWdpc3RlcklCTU9wZXJhdG9yQ2F0YWxvZyhzY29wZTogQ29uc3RydWN0LCBpZDogc3RyaW5nLCBsYWJlbHM6IGFueSwgcmVmcmVzaEludGVydmFsTWluOiBudW1iZXIgPSA2MCk6IFRhc2tCdWlsZGVyIHtcbiAgY29uc3QgaWJtT3BlcmF0b3JDYXRhbG9nU291cmNlID0ge1xuICAgIGFwaVZlcnNpb246ICdvcGVyYXRvcnMuY29yZW9zLmNvbS92MWFscGhhMScsXG4gICAga2luZDogJ0NhdGFsb2dTb3VyY2UnLFxuICAgIG1ldGFkYXRhOiB7XG4gICAgICBuYW1lOiAnaWJtLW9wZXJhdG9yLWNhdGFsb2cnLFxuICAgICAgbmFtZXNwYWNlOiAnb3BlbnNoaWZ0LW1hcmtldHBsYWNlJyxcbiAgICAgIGxhYmVsczogbGFiZWxzLFxuICAgIH0sXG4gICAgc3BlYzoge1xuICAgICAgZGlzcGxheU5hbWU6ICdJQk0gT3BlcmF0b3IgQ2F0YWxvZycsXG4gICAgICBwdWJsaXNoZXI6ICdJQk0nLFxuICAgICAgc291cmNlVHlwZTogJ2dycGMnLFxuICAgICAgaW1hZ2U6ICdpY3IuaW8vY3BvcGVuL2libS1vcGVyYXRvci1jYXRhbG9nJyxcbiAgICAgIHVwZGF0ZVN0cmF0ZWd5OiB7XG4gICAgICAgIHJlZ2lzdHJ5UG9sbDoge1xuICAgICAgICAgIGludGVydmFsOiBgJHtyZWZyZXNoSW50ZXJ2YWxNaW4udG9TdHJpbmcoKX1tYCxcbiAgICAgICAgfSxcbiAgICAgIH0sXG4gICAgfSxcbiAgfTtcblxuICByZXR1cm4gQXBwbHlPYmplY3RUYXNrKHNjb3BlLCBpZCwgaWJtT3BlcmF0b3JDYXRhbG9nU291cmNlKTtcbn1cblxuLyoqXG4gKiBDcmVhdGVzIGEgYnVpbGRlciBmb3IgYSBwaXBlbGluZSB0YXNrIHRoYXQgY3JlYXRlcyBhbiBPcGVyYXRvckdyb3VwXG4gKiB1c2luZyBgb2MgYXBwbHkgLWZgIHdpdGggYSBjb25maWd1cmF0aW9uIGZpbGUuXG4gKlxuICogVGhpcyBmdW5jdGlvbiBjcmVhdGVzIGEgYFRhc2tCdWlsZGVyYCB0byBnZW5lcmF0ZSBhIHRhc2sgdGhhdCB1c2VzIHRoZVxuICogb3BlbnNoaWZ0X2NsaWVudCBUZWt0b24gVGFzay5cbiAqXG4gKiBgYGB0eXBlc2NyaXB0XG4gKiBjb25zdCB0YXNrID0gQ3JlYXRlT3BlcmF0b3JHcm91cCh0aGlzLCAnY3JlYXRlLW9wZXJhdG9yLWdyb3VwJywgJ2libS1ldmVudGF1dG9tYXRpb24tb3BlcmF0b3Jncm91cCcsICdpYm0tZXZlbnRzdHJlYW1zJyk7XG4gKiBgYGBcbiAqXG4gKiBAcGFyYW0gc2NvcGUgVGhlIHBhcmVudCBbQ29uc3RydWN0XShodHRwczovL2NkazhzLmlvL2RvY3MvbGF0ZXN0L2Jhc2ljcy9jb25zdHJ1Y3RzLykuXG4gKiBAcGFyYW0gaWQgVGhlIGBpZGAgb2YgdGhlIGNvbnN0cnVjdC4gTXVzdCBiZSB1bmlxdWUgZm9yIGVhY2ggb25lIGluIGEgY2hhcnQuXG4gKiBAcGFyYW0gbnMgVGhlIG5hbWVzcGFjZSBmb3IgdGhlIE9wZXJhdG9yR3JvdXAuXG4gKiBAcGFyYW0gbmFtZSBUaGUgbmFtZSBvZiB0aGUgT3BlcmF0b3JHcm91cC5cbiAqIEBwYXJhbSB0YXJnZXROcyBJZiBwcm92aWRlZCwgaXQgaXMgdGhlIHRhcmdldCBuYW1lc3BhY2UgZm9yIHRoZSBPcGVyYXRvckdyb3VwLlxuICogQGNvbnN0cnVjdG9yXG4gKi9cbmV4cG9ydCBmdW5jdGlvbiBDcmVhdGVPcGVyYXRvckdyb3VwKHNjb3BlOiBDb25zdHJ1Y3QsIGlkOiBzdHJpbmcsIG5zOiBzdHJpbmcsIG5hbWU6IHN0cmluZywgdGFyZ2V0TnM6IHN0cmluZyA9ICcnKTogVGFza0J1aWxkZXIge1xuICBsZXQgb3BlcmF0b3JHcm91cFNvdXJjZSA9IHt9O1xuICBpZiAodGFyZ2V0TnMpIHtcbiAgICBvcGVyYXRvckdyb3VwU291cmNlID0ge1xuICAgICAgYXBpVmVyc2lvbjogJ29wZXJhdG9ycy5jb3Jlb3MuY29tL3YxJyxcbiAgICAgIGtpbmQ6ICdPcGVyYXRvckdyb3VwJyxcbiAgICAgIG1ldGFkYXRhOiB7XG4gICAgICAgIG5hbWU6IG5hbWUsXG4gICAgICAgIG5hbWVzcGFjZTogbnMsXG4gICAgICB9LFxuICAgICAgc3BlYzoge1xuICAgICAgICB0YXJnZXROYW1lc3BhY2VzOiBbdGFyZ2V0TnNdLFxuICAgICAgfSxcbiAgICB9O1xuICB9IGVsc2Uge1xuICAgIG9wZXJhdG9yR3JvdXBTb3VyY2UgPSB7XG4gICAgICBhcGlWZXJzaW9uOiAnb3BlcmF0b3JzLmNvcmVvcy5jb20vdjEnLFxuICAgICAga2luZDogJ09wZXJhdG9yR3JvdXAnLFxuICAgICAgbWV0YWRhdGE6IHtcbiAgICAgICAgbmFtZTogbmFtZSxcbiAgICAgICAgbmFtZXNwYWNlOiBucyxcbiAgICAgIH0sXG4gICAgfTtcbiAgfVxuXG4gIHJldHVybiBBcHBseU9iamVjdFRhc2soc2NvcGUsIGlkLCBvcGVyYXRvckdyb3VwU291cmNlKTtcbn1cblxuLyoqXG4gKiBDcmVhdGVzIGEgYnVpbGRlciBmb3IgYSBwaXBlbGluZSB0YXNrIHRoYXQgY3JlYXRlcyBhIFN1YnNjcmlwdGlvbiB1c2luZ1xuICogYG9jIGFwcGx5IC1mYCB3aXRoIGEgY29uZmlndXJhdGlvbiBmaWxlLlxuICpcbiAqIFRoaXMgZnVuY3Rpb24gY3JlYXRlcyBhIGBUYXNrQnVpbGRlcmAgdG8gZ2VuZXJhdGUgYSB0YXNrIHRoYXQgdXNlcyB0aGVcbiAqIG9wZW5zaGlmdF9jbGllbnQgVGVrdG9uIFRhc2suXG4gKlxuICogYGBgdHlwZXNjcmlwdFxuICogY29uc3QgdGFzayA9ICBTdWJzY3JpYmUodGhpcywgJ2luc3RhbGwtZXZlbnRzdHJlYW1zJywgJ2libS1ldmVudHN0cmVhbXMnLCAnaWJtLWV2ZW50c3RyZWFtcycsICdpYm0tb3BlcmF0b3ItY2F0YWxvZycsICdzdGFibGUnKTtcbiAqIGBgYFxuICpcbiAqIEBwYXJhbSBzY29wZSBUaGUgcGFyZW50IFtDb25zdHJ1Y3RdKGh0dHBzOi8vY2RrOHMuaW8vZG9jcy9sYXRlc3QvYmFzaWNzL2NvbnN0cnVjdHMvKS5cbiAqIEBwYXJhbSBpZCBUaGUgYGlkYCBvZiB0aGUgY29uc3RydWN0LiBNdXN0IGJlIHVuaXF1ZSBmb3IgZWFjaCBvbmUgaW4gYSBjaGFydC5cbiAqIEBwYXJhbSBucyBUaGUgbmFtZXNwYWNlIGZvciB0aGUgU3Vic2NyaXB0aW9uLlxuICogQHBhcmFtIG5hbWUgVGhlIG5hbWUgb2YgdGhlIFN1YnNjcmlwdGlvbi5cbiAqIEBwYXJhbSBjYXRhbG9nU291cmNlIFRoZSBuYW1lIG9mIHRoZSBjYXRhbG9nIHNvdXJjZS4gSWYgdXNpbmcgSUJNIG9wZXJhdG9ycywgdGhhdCBpcyBgaWJtLW9wZXJhdG9yLWNhdGFsb2dgXG4gKiBAcGFyYW0gY2hhbm5lbCBUaGUgbmFtZSBvZiB0aGUgY2hhbm5lbC4gSXQgZGVmYXVsdHMgdG8gYHN0YWJsZWAsIGJ1dCBpdCBjb3VsZCBiZSBzb21ldGhpbmcgbGlrZSBgdjMuM2AuXG4gKiBAY29uc3RydWN0b3JcbiAqL1xuZXhwb3J0IGZ1bmN0aW9uIFN1YnNjcmliZShzY29wZTogQ29uc3RydWN0LCBpZDogc3RyaW5nLCBuczogc3RyaW5nLCBuYW1lOiBzdHJpbmcsIGNhdGFsb2dTb3VyY2U6IHN0cmluZywgY2hhbm5lbDogc3RyaW5nID0gJ3N0YWJsZScpOiBUYXNrQnVpbGRlciB7XG4gIGNvbnN0IHN1YnNjcmlwdGlvbiA9IHtcbiAgICBhcGlWZXJzaW9uOiAnb3BlcmF0b3JzLmNvcmVvcy5jb20vdjFhbHBoYTEnLFxuICAgIGtpbmQ6ICdTdWJzY3JpcHRpb24nLFxuICAgIG1ldGFkYXRhOiB7XG4gICAgICBuYW1lOiBuYW1lLFxuICAgICAgbmFtZXNwYWNlOiBucyxcbiAgICB9LFxuICAgIHNwZWM6IHtcbiAgICAgIGNoYW5uZWw6IGNoYW5uZWwsXG4gICAgICBuYW1lOiBuYW1lLFxuICAgICAgc291cmNlOiBjYXRhbG9nU291cmNlLFxuICAgICAgc291cmNlTmFtZXNwYWNlOiBEZWZhdWx0Q2F0YWxvZ1NvdXJjZU5hbWVzcGFjZSxcbiAgICB9LFxuICB9O1xuICByZXR1cm4gQXBwbHlPYmplY3RUYXNrKHNjb3BlLCBpZCwgc3Vic2NyaXB0aW9uKTtcbn1cblxuLyoqXG4gKiBDcmVhdGVzIGEgYnVpbGRlciBmb3IgYSBwaXBlbGluZSB0YXNrIHRoYXQgdXNlcyBhbiBidXN5Ym94IGltYWdlIHRvIGVjaG8gb3V0IHRoZVxuICogZ2l2ZW4gbWVzc2FnZS5cbiAqXG4gKiBAcGFyYW0gc2NvcGUgVGhlIHBhcmVudCBbQ29uc3RydWN0XShodHRwczovL2NkazhzLmlvL2RvY3MvbGF0ZXN0L2Jhc2ljcy9jb25zdHJ1Y3RzLykuXG4gKiBAcGFyYW0gaWQgVGhlIGBpZGAgb2YgdGhlIGNvbnN0cnVjdC4gTXVzdCBiZSB1bmlxdWUgZm9yIGVhY2ggb25lIGluIGEgY2hhcnQuXG4gKiBAcGFyYW0gbWVzc2FnZSBUaGUgbWVzc2FnZSBmb3IgdGhlIHRhc2sgdG8gZWNobyB0byB0aGUgb3V0cHV0LlxuICogQGNvbnN0cnVjdG9yXG4gKi9cbmV4cG9ydCBmdW5jdGlvbiBFY2hvTWVzc2FnZShzY29wZTogQ29uc3RydWN0LCBpZDogc3RyaW5nLCBtZXNzYWdlOiBzdHJpbmcpOiBUYXNrQnVpbGRlciB7XG4gIC8vIFRPRE86IHNhbml0aXplIGlucHV0IGZvciB0aGUgbWVzc2FnZSB0byBtYWtlIHN1cmUgaXQgaGFzIG5vIGluamVjdGlvbiBhdHRhY2tzXG4gIHJldHVybiBuZXcgVGFza0J1aWxkZXIoc2NvcGUsIGlkKVxuICAgIC53aXRoTmFtZSgnZWNoby1tZXNzYWdlJylcbiAgICAud2l0aERlc2NyaXB0aW9uKCdFY2hvcyBhIG1lc3NhZ2Ugb3V0IHRvIHRoZSBwaXBlbGluZScpXG4gICAgLndpdGhTdGVwKG5ldyBUYXNrU3RlcEJ1aWxkZXIoKVxuICAgICAgLndpdGhOYW1lKCdlY2hvLW1lc3NhZ2UnKVxuICAgICAgLndpdGhJbWFnZSgnZG9ja2VyLmlvL2xpYnJhcnkvYnVzeWJveDpsYXRlc3QnKVxuICAgICAgLndpdGhDb21tYW5kKFsnZWNobycsIGBcXFwiJHttZXNzYWdlfVxcXCJgXSkpO1xufSJdfQ==