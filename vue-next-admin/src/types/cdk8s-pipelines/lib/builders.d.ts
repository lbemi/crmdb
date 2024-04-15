/**
 * This file has builders in it for the various pipeline constructs.
 */
import { ApiObjectProps } from 'cdk8s';
import { Construct } from 'constructs';
import { PipelineParam, PipelineWorkspace } from './pipelines';
import { TaskEnvValueSource, TaskStep } from './tasks';
/**
 * The options for builders for the `buildXX()` methods.
 */
export interface BuilderOptions {
    /**
     * If true, all the dependent objects are generated with the build. This is
     * designed to run on as minimal cluster as possible, with as few pre steps
     * as possible.
     */
    readonly includeDependencies?: boolean;
}
/**
 * The default options for the builders.
 */
export declare const DefaultBuilderOptions: BuilderOptions;
/**
 * Builds the Workspaces for use by Tasks and Pipelines.
 */
export declare class WorkspaceBuilder {
    private readonly _logicalID;
    private _name?;
    private _description?;
    /**
     * Creates the `WorkspaceBuilder`, using the given `id` as the logical ID for
     * the workspace.
     * @param id
     */
    constructor(id: string);
    /**
     * Gets the logical ID of the `Workspace`.
     */
    get logicalID(): string | undefined;
    /**
     * Gets the name of the workspace.
     */
    get name(): string | undefined;
    /**
     * Gets the description of the workspace.
     */
    get description(): string;
    withName(name: string): WorkspaceBuilder;
    withDescription(desc: string): WorkspaceBuilder;
}
/**
 * Builds the parameters for use by Tasks and Pipelines.
 */
export declare class ParameterBuilder {
    private readonly _logicalID;
    private _name?;
    private _description?;
    private _type?;
    private _value?;
    private _defaultValue?;
    private _requiresPipelineParam;
    constructor(id: string);
    /**
     * Gets the logicalID for the `ParameterBuilder`, which is used by the underlying
     * construct.
     */
    get logicalID(): string | undefined;
    get name(): string | undefined;
    get description(): string;
    /**
     * Sets the name of the parameter.
     * @param name
     */
    withName(name: string): ParameterBuilder;
    /**
     * Sets the description of the parameter.
     * @param desc
     */
    withDescription(desc: string): ParameterBuilder;
    /**
     * Sets the type of the parameter
     * @param type
     */
    ofType(type: string): ParameterBuilder;
    /**
     * Gets the type of the parameter
     */
    get type(): string | undefined;
    /**
     * Sets the value for the parameter
     * @param val
     */
    withValue(val: string): ParameterBuilder;
    /**
     * Gets the value of the parameter
     */
    get value(): string | undefined;
    /**
     * Sets the default value for the parameter.
     * @param val
     */
    withDefaultValue(val: string): ParameterBuilder;
    get defaultValue(): string | undefined;
    /**
     * Sets the default value for the parameter.
     * @param pipelineParamName
     * @param defaultValue
     */
    withPiplineParameter(pipelineParamName: string, defaultValue?: string): ParameterBuilder;
    /**
     * Returns true if this parameter expects input at the pipeline level.
     */
    get requiresPipelineParameter(): boolean;
}
/**
 * Creates a `Step` in a `Task`.
 */
export declare class TaskStepBuilder {
    private _name?;
    private _dir?;
    private _image?;
    private _cmd?;
    private _args?;
    private _env?;
    private _script?;
    /**
     *
     */
    constructor();
    /**
     * The name of the `Step` of the `Task`.
     */
    get name(): string | undefined;
    /**
     * The name of the container `image` used to execute the `Step` of the
     * `Task`.
     */
    get image(): string | undefined;
    get scriptData(): string | undefined;
    /**
     * Gets the command-line arguments that will be supplied to the `command`.
     */
    get args(): string[] | undefined;
    /**
     * Gets the command used for the `Step` on the `Task`.
     */
    get command(): string[] | undefined;
    get workingDir(): string | undefined;
    withName(name: string): TaskStepBuilder;
    /**
     * The name of the image to use when executing the `Step` on the `Task`
     * @param img
     */
    withImage(img: string): TaskStepBuilder;
    /**
     * The name of the command to use when running the `Step` of the `Task`. If
     * `command` is specified, do not specify `script`.
     * @param cmd
     */
    withCommand(cmd: string[]): TaskStepBuilder;
    /**
     * The args to use with the `command`.
     * @param args
     */
    withArgs(args: string[]): TaskStepBuilder;
    /**
     * If supplied, uses the content found at the given URL for the
     * `script` value of the step. Use this as an alternative to "heredoc", which
     * is embedding hard-coded shell or other scripts in the step.
     *
     * If you supply this, do not supply a value for `fromScriptObject`.
     * @param url
     */
    fromScriptUrl(url: string): TaskStepBuilder;
    /**
     * If supplied, uses the cdk8s `ApiObject` supplied as the body of the
     * `script` for the `Task`. This is most useful when used with `oc apply` or
     * other tasks in which you want to apply the object during the step in the
     * pipeline.
     *
     * If you supply this, do not supply a value for `fromScriptUrl`.
     * @param obj
     */
    fromScriptObject(obj: any): TaskStepBuilder;
    /**
     * If supplied, uses the provided script data as-is for the script value.
     *
     * Use this when you have the script data from a source other than a file or
     * an object. Use the other methods, such as `fromScriptUrl` (when the script
     * is in a file) or `scriptFromObject` (when the script is a CDK8s object)
     * rather than resolving those yourself.
     *
     * @param data
     */
    fromScriptData(data: string): TaskStepBuilder;
    /**
     * The `workingDir` of the `Task`.
     * @param dir
     */
    withWorkingDir(dir: string): TaskStepBuilder;
    withEnv(name: string, valueFrom: TaskEnvValueSource): TaskStepBuilder;
    buildTaskStep(): TaskStep | undefined;
}
/**
 * Builds Tekton `Task` objects that are independent of a `Pipeline`.
 *
 * To use a builder for tasks that will be used in a Pipeline, use the
 * `PipelineBuilder` instead.
 */
export declare class TaskBuilder {
    private readonly _scope;
    private readonly _id;
    private _steps?;
    private _name?;
    private _description?;
    private _workspaces;
    private _params;
    private _results;
    private _annotations?;
    private _labels?;
    /**
     * Creates a new instance of the `TaskBuilder` using the given `scope` and
     * `id`.
     * @param scope
     * @param id
     */
    constructor(scope: Construct, id: string);
    get logicalID(): string;
    /**
     * Adds a label to the `Task` with the provided label key and value.
     *
     * @see https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
     *
     * @param key
     * @param value
     */
    withLabel(key: string, value: string): TaskBuilder;
    /**
     * Adds an annotation to the `Task` `metadata` with the provided key and value.
     *
     * @see https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
     *
     * @param key The annotation's key.
     * @param value The annotation's value.
     */
    withAnnotation(key: string, value: string): TaskBuilder;
    /**
     * Gets the name of the `Task` built by the `TaskBuilder`.
     */
    get name(): string | undefined;
    /**
     * Sets the name of the `Task` being built.
     * @param name
     */
    withName(name: string): TaskBuilder;
    /**
     * Sets the `description` of the `Task` being built.
     * @param description
     */
    withDescription(description: string): TaskBuilder;
    /**
     * Gets the `description` of the `Task`.
     */
    get description(): string | undefined;
    /**
     * Adds the specified workspace to the `Task`.
     * @param workspace
     */
    withWorkspace(workspace: WorkspaceBuilder): TaskBuilder;
    /**
     * Gets the workspaces for the `Task`.
     */
    get workspaces(): WorkspaceBuilder[] | undefined;
    /**
     * Adds a parameter of type string to the `Task`.
     *
     * @param param
     */
    withStringParam(param: ParameterBuilder): TaskBuilder;
    get parameters(): ParameterBuilder[] | undefined;
    /**
     * Allows you to add an result to the Task.
     *
     * @see https://tekton.dev/docs/pipelines/tasks/#emitting-results
     *
     * @param name The name of the result.
     * @param description The result's description.
     */
    withResult(name: string, description: string): TaskBuilder;
    /**
     * Adds the given `step` (`TaskStepBuilder`) to the `Task`.
     * @param step
     */
    withStep(step: TaskStepBuilder): TaskBuilder;
    /**
     * Builds the `Task`.
     */
    buildTask(): void;
}
/**
 *
 */
export declare class PipelineBuilder {
    private readonly _scope;
    private readonly _id;
    private _name?;
    private _description?;
    private _tasks?;
    constructor(scope: Construct, id: string);
    /**
     * Provides the name for the pipeline task and will be
     * rendered as the `name` property.
     * @param name
     */
    withName(name: string): PipelineBuilder;
    /**
     * Gets the name of the pipeline
     */
    get name(): string;
    /**
     * Provides the name for the pipeline task and will be
     * rendered as the `name` property.
     * @param description
     */
    withDescription(description: string): PipelineBuilder;
    withTask(taskB: TaskBuilder): PipelineBuilder;
    /**
     * Returns the array of `PipelineParam` objects that represent the parameters
     * configured for the `Pipeline`.
     *
     * Note this is an "expensive" get because it loops through the tasks in the
     * pipeline and checks for duplicates in the pipeline parameters for each task
     * parameter found. You should avoid calling this in a loop--instead, declare
     * a local variable before the loop and reference that instead.
     *
     * @returns PipelineParam[] An array of the pipeline parameters.
     */
    get params(): PipelineParam[];
    /**
     * Returns the array of `PipelineWorkspace` objects that represent the workspaces
     * configured for the `Pipeline`.
     *
     * This is an "expensive" get because it loops through the workspaces in the
     * pipeline and checks for duplicates in the pipeline workspaces for each task
     * workspace found. You should avoid calling this in a loop--instead, declare
     * a local variable before the loop and reference that instead.
     *
     * @returns PipelineWorkspace[] An array of the pipeline workspaces.
     */
    get workspaces(): PipelineWorkspace[];
    /**
     * Builds the actual [Pipeline](https://tekton.dev/docs/getting-started/pipelines/)
     * from the settings configured using the fluid syntax.
     */
    buildPipeline(opts?: BuilderOptions): void;
}
/**
 * Builds a `PipelineRun` using the supplied configuration.
 *
 * @see https://tekton.dev/docs/pipelines/pipelineruns/
 */
export declare class PipelineRunBuilder {
    private readonly _scope;
    private readonly _id;
    private readonly _pipeline;
    private readonly _runParams;
    private readonly _runWorkspaces;
    private _sa;
    private _crbProps;
    /**
     * Creates a new instance of the `PipelineRunBuilder` for the specified
     * `Pipeline` that is built by the `PipelineBuilder` supplied here.
     *
     * A pipeline run is configured only for a specific pipeline, so it did not
     * make any sense here to allow the run to be created without the pipeline
     * specified.
     *
     * @param scope The `Construct` in which to create the `PipelineRun`.
     * @param id The logical ID of the `PipelineRun` construct.
     * @param pipeline The `Pipeline` for which to create this run, using the `PipelineBuilder`.
     */
    constructor(scope: Construct, id: string, pipeline: PipelineBuilder);
    /**
     * Adds a run parameter to the `PipelineRun`. It will throw an error if you try
     * to add a parameter that does not exist on the pipeline.
     *
     * @param name The name of the parameter added to the pipeline run.
     * @param value The value of the parameter added to the pipeline run.
     */
    withRunParam(name: string, value: string): PipelineRunBuilder;
    /**
     * Allows you to specify the name of a `PersistentVolumeClaim` but does not
     * do any compile-time validation on the volume claim's name or existence.
     *
     * @see https://kubernetes.io/docs/tasks/configure-pod-container/configure-persistent-volume-storage/#create-a-persistentvolumeclaim
     *
     * @param name The name of the workspace in the `PipelineRun` that will be used by the `Pipeline`.
     * @param claimName The name of the `PersistentVolumeClaim` to use for the `workspace`.
     * @param subPath The sub path on the `persistentVolumeClaim` to use for the `workspace`.
     */
    withWorkspace(name: string, claimName: string, subPath: string): PipelineRunBuilder;
    withClusterRoleBindingProps(props: ApiObjectProps): PipelineRunBuilder;
    /**
     * Uses the provided role name for the `serviceAccountName` on the
     * `PipelineRun`. If this method is not called prior to `buildPipelineRun()`,
     * then the default service account will be used, which is _default:pipeline_.
     *
     * @param sa The name of the service account (`serviceAccountName`) to use.
     */
    withServiceAccount(sa: string): PipelineRunBuilder;
    /**
     * Builds the `PipelineRun` for the configured `Pipeline` used in the constructor.
     * @param opts
     */
    buildPipelineRun(opts?: BuilderOptions): void;
}
