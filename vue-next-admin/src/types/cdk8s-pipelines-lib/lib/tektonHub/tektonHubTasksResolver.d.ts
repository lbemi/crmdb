import { TaskBuilder } from 'cdk8s-pipelines';
import { Construct } from 'constructs';
/**
 * This class handles turning a URL that points the YAML for the Tekton Hub Task into a PipelineTask
 */
export declare class TektonHubTask extends TaskBuilder {
    url: string;
    taskBuild: TaskBuilder;
    /**
     * Creates a new Instance of TektonHubTask with a URL that points to the Raw YAML for the task.
     * @link https://hub.tekton.dev/
     * @param scope
     * @param id
     * @param url string Url to the raw yaml for a Tekton Hub Task (i.e https://raw.githubusercontent.com/tektoncd/catalog/main/task/yq/0.4/yq.yaml)
     */
    constructor(scope: Construct, id: string, url: string);
    private parseYAML;
    private readYamlFromUrl;
    /**
     * Returns an instance of PipelineTaskBuilder with the corresponding Tekton Hub Task Link.
     * @returns TaskBuilder
     */
    build(): TaskBuilder;
}
