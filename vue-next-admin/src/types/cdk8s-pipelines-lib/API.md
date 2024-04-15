[![build](https://github.com/cloud-native-toolkit/cdk8s-pipelines-lib/actions/workflows/build.yml/badge.svg)](https://github.com/cloud-native-toolkit/cdk8s-pipelines-lib/actions/workflows/build.yml)

[![View on Construct Hub](https://constructs.dev/badge?package=cdk8s-pipelines-lib)](https://constructs.dev/packages/cdk8s-pipelines-lib)

# Pipeline Library of cdk8s Constructs

This is a library of several "pattern" pipelines that are intended to be
basic and therefore easily reusable.

Additionally, using the `TaskBuilder`, each `Task`
(see [Tasks](https://tekton.dev/docs/getting-started/tasks/))
from [Tekton Hub](https://hub.tekton.dev/) can be found in this library as a construct.

## Using tasks from Tekton Hub

The following is an example chart that uses a Tekton Hub Task for
an [OpenShift client](https://hub.tekton.dev/tekton/task/openshift-client).

```ts
import { App, Chart, ChartProps } from 'cdk8s';
import { ParameterBuilder, PipelineBuilder } from 'cdk8s-pipelines';
import { openshift_client } from 'cdk8s-pipelines-lib';
import { Construct } from 'constructs';

export class MyChart extends Chart {
  constructor(scope: Construct, id: string, props: ChartProps = {}) {
    super(scope, id, props);

    const projectName = 'my-project';

    const createProject = openshift_client(this, 'create-project')
      .withStringParam(new ParameterBuilder('SCRIPT')
        .withValue(`oc create ${projectName}`));

    new PipelineBuilder(this, 'create-some-namespace')
      .withDescription('Creates a namespace and then does some other stuff')
      .withTask(createProject)
      // ... more tasks go here
      .buildPipeline({ includeDependencies: true });
  }
}
const app = new App();
new MyChart(app, 'hello');
app.synth();
```

The result of this code will include the dependent tasks. The output will look like this:

```yaml
apiVersion: tekton.dev/v1
kind: Task
metadata:
  name: openshift-client
spec:
  description: null
  workspaces:
    - name: manifest-dir
      description: The workspace which contains kubernetes manifests which we want to apply on the cluster.
    - name: kubeconfig-dir
      description: The workspace which contains the the kubeconfig file if in case we want to run the oc command on another cluster.
  params:
    - name: SCRIPT
      description: ""
      default: null
    - name: VERSION
      description: The OpenShift Version to use
      default: "4.7"
  steps:
    - name: oc
      image: quay.io/openshift/origin-cli:$(params.VERSION)
      script: |
        #!/usr/bin/env bash

        [[ "$(workspaces.manifest-dir.bound)" == "true" ]] && \
        cd $(workspaces.manifest-dir.path)

        [[ "$(workspaces.kubeconfig-dir.bound)" == "true" ]] && \
        [[ -f $(workspaces.kubeconfig-dir.path)/kubeconfig ]] && \
        export KUBECONFIG=$(workspaces.kubeconfig-dir.path)/kubeconfig

        $(params.SCRIPT)
      workingDir: null
      env: null
---
apiVersion: tekton.dev/v1
kind: Pipeline
metadata:
  name: create-some-namespace
spec:
  description: Creates a namespace and then does some other stuff
  params:
    - name: VERSION
      type: string
  workspaces:
    - name: manifest-dir
      description: The workspace which contains kubernetes manifests which we want to apply on the cluster.
    - name: kubeconfig-dir
      description: The workspace which contains the the kubeconfig file if in case we want to run the oc command on another cluster.
  tasks:
    - name: create-project
      taskRef:
        name: openshift-client
      params:
        - name: SCRIPT
          value: oc create my-project
        - name: VERSION
          value: $(params.VERSION)
      workspaces:
        - name: manifest-dir
          workspace: manifest-dir
        - name: kubeconfig-dir
          workspace: kubeconfig-dir
```

## Using in a build

The goal of using cdk8s-pipeline and cdk8s-pipeline-lib should be to produce YAML artifacts in a build process that are
included in the release of a project. As an example, see [this example AWS CDK project](https://github.ibm.com/Nathan-Good/example-cdk-aws-ec2-vm),
which demonstrates how to include the output created by the CDK `synth()` in the output of the build as a versioned
release.

# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### AWSCDKPipelineChart <a name="AWSCDKPipelineChart" id="cdk8s-pipelines-lib.AWSCDKPipelineChart"></a>

The chart for creating a Tekton Pipeline that will use an AWS CDK project to create resources in AWS for re-usable artifacts.

#### Initializers <a name="Initializers" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.Initializer"></a>

```typescript
import { AWSCDKPipelineChart } from 'cdk8s-pipelines-lib'

new AWSCDKPipelineChart(scope: Construct, id: string, props?: AWSCDKPipelineChartProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChart.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChart.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChart.Initializer.parameter.props">props</a></code> | <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChartProps">AWSCDKPipelineChartProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Optional</sup> <a name="props" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.Initializer.parameter.props"></a>

- *Type:* <a href="#cdk8s-pipelines-lib.AWSCDKPipelineChartProps">AWSCDKPipelineChartProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChart.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChart.addDependency">addDependency</a></code> | Create a dependency between this Chart and other constructs. |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChart.generateObjectName">generateObjectName</a></code> | Generates a app-unique name for an object given it's construct node path. |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChart.toJson">toJson</a></code> | Renders this chart to a set of Kubernetes JSON resources. |

---

##### `toString` <a name="toString" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `addDependency` <a name="addDependency" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.addDependency"></a>

```typescript
public addDependency(dependencies: IConstruct): void
```

Create a dependency between this Chart and other constructs.

These can be other ApiObjects, Charts, or custom.

###### `dependencies`<sup>Required</sup> <a name="dependencies" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.addDependency.parameter.dependencies"></a>

- *Type:* constructs.IConstruct

the dependencies to add.

---

##### `generateObjectName` <a name="generateObjectName" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.generateObjectName"></a>

```typescript
public generateObjectName(apiObject: ApiObject): string
```

Generates a app-unique name for an object given it's construct node path.

Different resource types may have different constraints on names
(`metadata.name`). The previous version of the name generator was
compatible with DNS_SUBDOMAIN but not with DNS_LABEL.

For example, `Deployment` names must comply with DNS_SUBDOMAIN while
`Service` names must comply with DNS_LABEL.

Since there is no formal specification for this, the default name
generation scheme for kubernetes objects in cdk8s was changed to DNS_LABEL,
since it’s the common denominator for all kubernetes resources
(supposedly).

You can override this method if you wish to customize object names at the
chart level.

###### `apiObject`<sup>Required</sup> <a name="apiObject" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.generateObjectName.parameter.apiObject"></a>

- *Type:* cdk8s.ApiObject

The API object to generate a name for.

---

##### `toJson` <a name="toJson" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.toJson"></a>

```typescript
public toJson(): any[]
```

Renders this chart to a set of Kubernetes JSON resources.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChart.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChart.isChart">isChart</a></code> | Return whether the given object is a Chart. |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChart.of">of</a></code> | Finds the chart in which a node is defined. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.isConstruct"></a>

```typescript
import { AWSCDKPipelineChart } from 'cdk8s-pipelines-lib'

AWSCDKPipelineChart.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isChart` <a name="isChart" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.isChart"></a>

```typescript
import { AWSCDKPipelineChart } from 'cdk8s-pipelines-lib'

AWSCDKPipelineChart.isChart(x: any)
```

Return whether the given object is a Chart.

We do attribute detection since we can't reliably use 'instanceof'.

###### `x`<sup>Required</sup> <a name="x" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.isChart.parameter.x"></a>

- *Type:* any

---

##### `of` <a name="of" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.of"></a>

```typescript
import { AWSCDKPipelineChart } from 'cdk8s-pipelines-lib'

AWSCDKPipelineChart.of(c: IConstruct)
```

Finds the chart in which a node is defined.

###### `c`<sup>Required</sup> <a name="c" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.of.parameter.c"></a>

- *Type:* constructs.IConstruct

a construct node.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChart.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChart.property.apiObjects">apiObjects</a></code> | <code>cdk8s.ApiObject[]</code> | Returns all the included API objects. |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChart.property.labels">labels</a></code> | <code>{[ key: string ]: string}</code> | Labels applied to all resources in this chart. |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChart.property.namespace">namespace</a></code> | <code>string</code> | The default namespace for all objects in this chart. |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `apiObjects`<sup>Required</sup> <a name="apiObjects" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.property.apiObjects"></a>

```typescript
public readonly apiObjects: ApiObject[];
```

- *Type:* cdk8s.ApiObject[]

Returns all the included API objects.

---

##### `labels`<sup>Required</sup> <a name="labels" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.property.labels"></a>

```typescript
public readonly labels: {[ key: string ]: string};
```

- *Type:* {[ key: string ]: string}

Labels applied to all resources in this chart.

This is an immutable copy.

---

##### `namespace`<sup>Optional</sup> <a name="namespace" id="cdk8s-pipelines-lib.AWSCDKPipelineChart.property.namespace"></a>

```typescript
public readonly namespace: string;
```

- *Type:* string

The default namespace for all objects in this chart.

---


## Structs <a name="Structs" id="Structs"></a>

### AWSCDKPipelineChartProps <a name="AWSCDKPipelineChartProps" id="cdk8s-pipelines-lib.AWSCDKPipelineChartProps"></a>

Initialization properties for the AWSCDKPipelineChart.

#### Initializer <a name="Initializer" id="cdk8s-pipelines-lib.AWSCDKPipelineChartProps.Initializer"></a>

```typescript
import { AWSCDKPipelineChartProps } from 'cdk8s-pipelines-lib'

const aWSCDKPipelineChartProps: AWSCDKPipelineChartProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChartProps.property.disableResourceNameHashes">disableResourceNameHashes</a></code> | <code>boolean</code> | The autogenerated resource name by default is suffixed with a stable hash of the construct path. |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChartProps.property.labels">labels</a></code> | <code>{[ key: string ]: string}</code> | Labels to apply to all resources in this chart. |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChartProps.property.namespace">namespace</a></code> | <code>string</code> | The default namespace for all objects defined in this chart (directly or indirectly). |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipelineChartProps.property.params">params</a></code> | <code>string[]</code> | *No description.* |

---

##### `disableResourceNameHashes`<sup>Optional</sup> <a name="disableResourceNameHashes" id="cdk8s-pipelines-lib.AWSCDKPipelineChartProps.property.disableResourceNameHashes"></a>

```typescript
public readonly disableResourceNameHashes: boolean;
```

- *Type:* boolean
- *Default:* false

The autogenerated resource name by default is suffixed with a stable hash of the construct path.

Setting this property to true drops the hash suffix.

---

##### `labels`<sup>Optional</sup> <a name="labels" id="cdk8s-pipelines-lib.AWSCDKPipelineChartProps.property.labels"></a>

```typescript
public readonly labels: {[ key: string ]: string};
```

- *Type:* {[ key: string ]: string}
- *Default:* no common labels

Labels to apply to all resources in this chart.

---

##### `namespace`<sup>Optional</sup> <a name="namespace" id="cdk8s-pipelines-lib.AWSCDKPipelineChartProps.property.namespace"></a>

```typescript
public readonly namespace: string;
```

- *Type:* string
- *Default:* no namespace is synthesized (usually this implies "default")

The default namespace for all objects defined in this chart (directly or indirectly).

This namespace will only apply to objects that don't have a
`namespace` explicitly defined for them.

---

##### `params`<sup>Optional</sup> <a name="params" id="cdk8s-pipelines-lib.AWSCDKPipelineChartProps.property.params"></a>

```typescript
public readonly params: string[];
```

- *Type:* string[]

---

### GitRepoConfig <a name="GitRepoConfig" id="cdk8s-pipelines-lib.GitRepoConfig"></a>

Contains the information for the GitHub repo and the stack so we can go get it and generate the AWS CDK pipeline.

#### Initializer <a name="Initializer" id="cdk8s-pipelines-lib.GitRepoConfig.Initializer"></a>

```typescript
import { GitRepoConfig } from 'cdk8s-pipelines-lib'

const gitRepoConfig: GitRepoConfig = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk8s-pipelines-lib.GitRepoConfig.property.ghUrl">ghUrl</a></code> | <code>string</code> | The URL for the GitHub or GHE API. |
| <code><a href="#cdk8s-pipelines-lib.GitRepoConfig.property.owner">owner</a></code> | <code>string</code> | The owner of the GitHub repository. |
| <code><a href="#cdk8s-pipelines-lib.GitRepoConfig.property.release">release</a></code> | <code>string</code> | The release tag for the release in which the AWS CDK template should be found. |
| <code><a href="#cdk8s-pipelines-lib.GitRepoConfig.property.repo">repo</a></code> | <code>string</code> | The name of the repository. |
| <code><a href="#cdk8s-pipelines-lib.GitRepoConfig.property.stackName">stackName</a></code> | <code>string</code> | The name of the AWS CDK stack. |
| <code><a href="#cdk8s-pipelines-lib.GitRepoConfig.property.token">token</a></code> | <code>string</code> | The personal access token (PAT) for accessing the library in GitHub. |

---

##### `ghUrl`<sup>Optional</sup> <a name="ghUrl" id="cdk8s-pipelines-lib.GitRepoConfig.property.ghUrl"></a>

```typescript
public readonly ghUrl: string;
```

- *Type:* string

The URL for the GitHub or GHE API.

The value should look like https://api.github.com or
https://github.mycompany.com/api/v3.

---

##### `owner`<sup>Optional</sup> <a name="owner" id="cdk8s-pipelines-lib.GitRepoConfig.property.owner"></a>

```typescript
public readonly owner: string;
```

- *Type:* string

The owner of the GitHub repository.

---

##### `release`<sup>Optional</sup> <a name="release" id="cdk8s-pipelines-lib.GitRepoConfig.property.release"></a>

```typescript
public readonly release: string;
```

- *Type:* string

The release tag for the release in which the AWS CDK template should be found.

---

##### `repo`<sup>Optional</sup> <a name="repo" id="cdk8s-pipelines-lib.GitRepoConfig.property.repo"></a>

```typescript
public readonly repo: string;
```

- *Type:* string

The name of the repository.

---

##### `stackName`<sup>Optional</sup> <a name="stackName" id="cdk8s-pipelines-lib.GitRepoConfig.property.stackName"></a>

```typescript
public readonly stackName: string;
```

- *Type:* string

The name of the AWS CDK stack.

This should be a generated template that is included
in the release.

---

##### `token`<sup>Optional</sup> <a name="token" id="cdk8s-pipelines-lib.GitRepoConfig.property.token"></a>

```typescript
public readonly token: string;
```

- *Type:* string

The personal access token (PAT) for accessing the library in GitHub.

---

## Classes <a name="Classes" id="Classes"></a>

### AWSCDKPipeline <a name="AWSCDKPipeline" id="cdk8s-pipelines-lib.AWSCDKPipeline"></a>

Creator for the AWSCDKPipelineChart.

#### Initializers <a name="Initializers" id="cdk8s-pipelines-lib.AWSCDKPipeline.Initializer"></a>

```typescript
import { AWSCDKPipeline } from 'cdk8s-pipelines-lib'

new AWSCDKPipeline()
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |

---


#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk8s-pipelines-lib.AWSCDKPipeline.createFrom">createFrom</a></code> | Generates the AWS CDK Pipeline (AWSCDKPipelineChart) based on the actual project located in GitHub and specified by the configuration. |

---

##### `createFrom` <a name="createFrom" id="cdk8s-pipelines-lib.AWSCDKPipeline.createFrom"></a>

```typescript
import { AWSCDKPipeline } from 'cdk8s-pipelines-lib'

AWSCDKPipeline.createFrom(config: GitRepoConfig)
```

Generates the AWS CDK Pipeline (AWSCDKPipelineChart) based on the actual project located in GitHub and specified by the configuration.

###### `config`<sup>Required</sup> <a name="config" id="cdk8s-pipelines-lib.AWSCDKPipeline.createFrom.parameter.config"></a>

- *Type:* <a href="#cdk8s-pipelines-lib.GitRepoConfig">GitRepoConfig</a>

---



### InstallFromIBMOperatorPipeline <a name="InstallFromIBMOperatorPipeline" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline"></a>

A basic pipeline that starts with a subscription to the IBM operator catalog.

The following steps are included in this pipeline, so you do not need to add
them. The pipeline:

1. Creates the specified namespace.
1. Registers the IBM operator.
1. Creates an OperatorGroup.
1. Subscribes to the given `name` and `channel`

#### Initializers <a name="Initializers" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.Initializer"></a>

```typescript
import { InstallFromIBMOperatorPipeline } from 'cdk8s-pipelines-lib'

new InstallFromIBMOperatorPipeline(scope: Construct, id: string, ns: string, subscription: string, channel: string)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | The parent [Construct](https://cdk8s.io/docs/latest/basics/constructs/). |
| <code><a href="#cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.Initializer.parameter.id">id</a></code> | <code>string</code> | The `id` of the construct. |
| <code><a href="#cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.Initializer.parameter.ns">ns</a></code> | <code>string</code> | The namespace to create and to use for subscribing to the product and channel. |
| <code><a href="#cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.Initializer.parameter.subscription">subscription</a></code> | <code>string</code> | The name of the subscription. |
| <code><a href="#cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.Initializer.parameter.channel">channel</a></code> | <code>string</code> | The name of the channel (e.g., `v3.3`, `stable`). |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

The parent [Construct](https://cdk8s.io/docs/latest/basics/constructs/).

---

##### `id`<sup>Required</sup> <a name="id" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.Initializer.parameter.id"></a>

- *Type:* string

The `id` of the construct.

Must be unique for each one in a chart.

---

##### `ns`<sup>Required</sup> <a name="ns" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.Initializer.parameter.ns"></a>

- *Type:* string

The namespace to create and to use for subscribing to the product and channel.

---

##### `subscription`<sup>Required</sup> <a name="subscription" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.Initializer.parameter.subscription"></a>

- *Type:* string

The name of the subscription.

For example, for IBM Event Streams is it `ibm-eventstreams`. For Red Hat Serverless, it is `serverless-operator`.

---

##### `channel`<sup>Required</sup> <a name="channel" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.Initializer.parameter.channel"></a>

- *Type:* string

The name of the channel (e.g., `v3.3`, `stable`).

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.buildPipeline">buildPipeline</a></code> | Builds the actual [Pipeline](https://tekton.dev/docs/getting-started/pipelines/) from the settings configured using the fluid syntax. |
| <code><a href="#cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.withDescription">withDescription</a></code> | Provides the name for the pipeline task and will be rendered as the `name` property. |
| <code><a href="#cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.withName">withName</a></code> | Provides the name for the pipeline task and will be rendered as the `name` property. |
| <code><a href="#cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.withTask">withTask</a></code> | *No description.* |

---

##### `buildPipeline` <a name="buildPipeline" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.buildPipeline"></a>

```typescript
public buildPipeline(opts?: BuilderOptions): void
```

Builds the actual [Pipeline](https://tekton.dev/docs/getting-started/pipelines/) from the settings configured using the fluid syntax.

###### `opts`<sup>Optional</sup> <a name="opts" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.buildPipeline.parameter.opts"></a>

- *Type:* cdk8s-pipelines.BuilderOptions

---

##### `withDescription` <a name="withDescription" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.withDescription"></a>

```typescript
public withDescription(description: string): PipelineBuilder
```

Provides the name for the pipeline task and will be rendered as the `name` property.

###### `description`<sup>Required</sup> <a name="description" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.withDescription.parameter.description"></a>

- *Type:* string

---

##### `withName` <a name="withName" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.withName"></a>

```typescript
public withName(name: string): PipelineBuilder
```

Provides the name for the pipeline task and will be rendered as the `name` property.

###### `name`<sup>Required</sup> <a name="name" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.withName.parameter.name"></a>

- *Type:* string

---

##### `withTask` <a name="withTask" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.withTask"></a>

```typescript
public withTask(taskB: TaskBuilder): PipelineBuilder
```

###### `taskB`<sup>Required</sup> <a name="taskB" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.withTask.parameter.taskB"></a>

- *Type:* cdk8s-pipelines.TaskBuilder

---


#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.property.name">name</a></code> | <code>string</code> | Gets the name of the pipeline. |
| <code><a href="#cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.property.params">params</a></code> | <code>cdk8s-pipelines.PipelineParam[]</code> | Returns the array of `PipelineParam` objects that represent the parameters configured for the `Pipeline`. |
| <code><a href="#cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.property.workspaces">workspaces</a></code> | <code>cdk8s-pipelines.PipelineWorkspace[]</code> | Returns the array of `PipelineWorkspace` objects that represent the workspaces configured for the `Pipeline`. |

---

##### `name`<sup>Required</sup> <a name="name" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

Gets the name of the pipeline.

---

##### `params`<sup>Required</sup> <a name="params" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.property.params"></a>

```typescript
public readonly params: PipelineParam[];
```

- *Type:* cdk8s-pipelines.PipelineParam[]

Returns the array of `PipelineParam` objects that represent the parameters configured for the `Pipeline`.

Note this is an "expensive" get because it loops through the tasks in the
pipeline and checks for duplicates in the pipeline parameters for each task
parameter found. You should avoid calling this in a loop--instead, declare
a local variable before the loop and reference that instead.

---

##### `workspaces`<sup>Required</sup> <a name="workspaces" id="cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline.property.workspaces"></a>

```typescript
public readonly workspaces: PipelineWorkspace[];
```

- *Type:* cdk8s-pipelines.PipelineWorkspace[]

Returns the array of `PipelineWorkspace` objects that represent the workspaces configured for the `Pipeline`.

This is an "expensive" get because it loops through the workspaces in the
pipeline and checks for duplicates in the pipeline workspaces for each task
workspace found. You should avoid calling this in a loop--instead, declare
a local variable before the loop and reference that instead.

---



