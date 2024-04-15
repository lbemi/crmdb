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
