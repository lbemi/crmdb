"use strict";
var _a, _b, _c, _d, _e, _f;
Object.defineProperty(exports, "__esModule", { value: true });
exports.PipelineRunBuilder = exports.PipelineBuilder = exports.TaskBuilder = exports.TaskStepBuilder = exports.ParameterBuilder = exports.WorkspaceBuilder = exports.DefaultBuilderOptions = void 0;
const JSII_RTTI_SYMBOL_1 = Symbol.for("jsii.rtti");
/**
 * This file has builders in it for the various pipeline constructs.
 */
const fs = require("fs");
const cdk8s_1 = require("cdk8s");
const common_1 = require("./common");
const pipelines_1 = require("./pipelines");
const tasks_1 = require("./tasks");
const DefaultPipelineServiceAccountName = 'default:pipeline';
/**
 * Creates the properties for a `ClusterRoleBinding`
 * @param bindingName
 * @param bindingNs
 * @param rolename
 * @param sa
 * @param saNamespace
 */
function createRoleBindingProps(bindingName, bindingNs, rolename, sa, saNamespace) {
    return {
        apiVersion: 'rbac.authorization.k8s.io/v1',
        kind: 'ClusterRoleBinding',
        metadata: {
            name: bindingName,
            namespace: bindingNs,
        },
        roleRef: {
            kind: 'ClusterRole',
            name: rolename,
        },
        subjects: [
            {
                kind: 'ServiceAccount',
                name: sa,
                namespace: saNamespace,
            },
        ],
    };
}
const DefaultClusterRoleBindingProps = createRoleBindingProps('pipeline-admin-default-crb', 'default', 'cluster-admin', 'pipeline', 'default');
/**
 * The default options for the builders.
 */
exports.DefaultBuilderOptions = {
    includeDependencies: false,
};
/**
 * Builds the Workspaces for use by Tasks and Pipelines.
 */
class WorkspaceBuilder {
    /**
     * Creates the `WorkspaceBuilder`, using the given `id` as the logical ID for
     * the workspace.
     * @param id
     */
    constructor(id) {
        this._logicalID = id;
    }
    /**
     * Gets the logical ID of the `Workspace`.
     */
    get logicalID() {
        return this._logicalID;
    }
    /**
     * Gets the name of the workspace.
     */
    get name() {
        return this._name;
    }
    /**
     * Gets the description of the workspace.
     */
    get description() {
        return this._description || '';
    }
    withName(name) {
        this._name = name;
        return this;
    }
    withDescription(desc) {
        this._description = desc;
        return this;
    }
}
_a = JSII_RTTI_SYMBOL_1;
WorkspaceBuilder[_a] = { fqn: "cdk8s-pipelines.WorkspaceBuilder", version: "0.0.16" };
exports.WorkspaceBuilder = WorkspaceBuilder;
/**
 * Builds the parameters for use by Tasks and Pipelines.
 */
class ParameterBuilder {
    constructor(id) {
        this._logicalID = id;
        this._requiresPipelineParam = false;
    }
    /**
     * Gets the logicalID for the `ParameterBuilder`, which is used by the underlying
     * construct.
     */
    get logicalID() {
        return this._logicalID;
    }
    get name() {
        return this._name;
    }
    get description() {
        return this._description || '';
    }
    /**
     * Sets the name of the parameter.
     * @param name
     */
    withName(name) {
        this._name = name;
        return this;
    }
    /**
     * Sets the description of the parameter.
     * @param desc
     */
    withDescription(desc) {
        this._description = desc;
        return this;
    }
    /**
     * Sets the type of the parameter
     * @param type
     */
    ofType(type) {
        this._type = type;
        return this;
    }
    /**
     * Gets the type of the parameter
     */
    get type() {
        return this._type;
    }
    /**
     * Sets the value for the parameter
     * @param val
     */
    withValue(val) {
        // If you are giving it a value here, then you do not
        // need the Pipeline parameter for this parameter.
        this._requiresPipelineParam = false;
        this._value = val;
        return this;
    }
    /**
     * Gets the value of the parameter
     */
    get value() {
        return this._value;
    }
    /**
     * Sets the default value for the parameter.
     * @param val
     */
    withDefaultValue(val) {
        this._defaultValue = val;
        return this;
    }
    get defaultValue() {
        return this._defaultValue;
    }
    /**
     * Sets the default value for the parameter.
     * @param pipelineParamName
     * @param defaultValue
     */
    withPiplineParameter(pipelineParamName, defaultValue = '') {
        this._requiresPipelineParam = true;
        this._name = pipelineParamName;
        this._defaultValue = defaultValue;
        this._value = (0, common_1.buildParam)(pipelineParamName);
        return this;
    }
    /**
     * Returns true if this parameter expects input at the pipeline level.
     */
    get requiresPipelineParameter() {
        return this._requiresPipelineParam;
    }
}
_b = JSII_RTTI_SYMBOL_1;
ParameterBuilder[_b] = { fqn: "cdk8s-pipelines.ParameterBuilder", version: "0.0.16" };
exports.ParameterBuilder = ParameterBuilder;
/**
 * Resolves the provided object into a YAML string.
 */
class ObjScriptResolver {
    /**
     * Creates an instance of the `ObjScriptResolver`.
     * @param obj The object to serialize to YAML for the script.
     */
    constructor(obj) {
        this._obj = obj;
    }
    /**
     * Gets the body of the script as a YAML representation of the object.
     */
    scriptData() {
        return cdk8s_1.Yaml.stringify(this._obj);
    }
}
/**
 * Gets the content from the provided URL and returns it as the script data.
 */
class UrlScriptResolver {
    /**
     * Creates an instance of the `UrlScriptResolver` with the provided URL.
     * @param url
     */
    constructor(url) {
        this._url = url;
    }
    /**
     * Gets the body of the script from the provided URL.
     * @return string Script data.
     */
    scriptData() {
        const data = fs.readFileSync(this._url, {
            encoding: 'utf8',
            flag: 'r',
        });
        return data.replace(/\n/g, '\\n');
    }
}
/**
 * Gets the content from the static value provided.
 */
class StaticScriptResolver {
    /**
     * Creates an instance of the `StaticScriptResolver`.
     * @param data
     */
    constructor(data) {
        this._script = data;
    }
    /**
     * Returns the static value provided.
     */
    scriptData() {
        return this._script;
    }
}
/**
 * Creates a `Step` in a `Task`.
 */
class TaskStepBuilder {
    /**
     *
     */
    constructor() {
    }
    /**
     * The name of the `Step` of the `Task`.
     */
    get name() {
        return this._name;
    }
    /**
     * The name of the container `image` used to execute the `Step` of the
     * `Task`.
     */
    get image() {
        return this._image;
    }
    get scriptData() {
        return this._script?.scriptData();
    }
    /**
     * Gets the command-line arguments that will be supplied to the `command`.
     */
    get args() {
        return this._args;
    }
    /**
     * Gets the command used for the `Step` on the `Task`.
     */
    get command() {
        return this._cmd;
    }
    get workingDir() {
        return this._dir;
    }
    withName(name) {
        this._name = name;
        return this;
    }
    /**
     * The name of the image to use when executing the `Step` on the `Task`
     * @param img
     */
    withImage(img) {
        this._image = img;
        return this;
    }
    /**
     * The name of the command to use when running the `Step` of the `Task`. If
     * `command` is specified, do not specify `script`.
     * @param cmd
     */
    withCommand(cmd) {
        this._cmd = cmd;
        return this;
    }
    /**
     * The args to use with the `command`.
     * @param args
     */
    withArgs(args) {
        this._args = args;
        return this;
    }
    /**
     * If supplied, uses the content found at the given URL for the
     * `script` value of the step. Use this as an alternative to "heredoc", which
     * is embedding hard-coded shell or other scripts in the step.
     *
     * If you supply this, do not supply a value for `fromScriptObject`.
     * @param url
     */
    fromScriptUrl(url) {
        this._script = new UrlScriptResolver(url);
        return this;
    }
    /**
     * If supplied, uses the cdk8s `ApiObject` supplied as the body of the
     * `script` for the `Task`. This is most useful when used with `oc apply` or
     * other tasks in which you want to apply the object during the step in the
     * pipeline.
     *
     * If you supply this, do not supply a value for `fromScriptUrl`.
     * @param obj
     */
    fromScriptObject(obj) {
        this._script = new ObjScriptResolver(obj);
        return this;
    }
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
    fromScriptData(data) {
        this._script = new StaticScriptResolver(data);
        return this;
    }
    /**
     * The `workingDir` of the `Task`.
     * @param dir
     */
    withWorkingDir(dir) {
        this._dir = dir;
        return this;
    }
    withEnv(name, valueFrom) {
        if (!this._env) {
            this._env = new Array();
        }
        this._env.push({
            name: name,
            valueFrom: valueFrom,
        });
        return this;
    }
    buildTaskStep() {
        if (this._script) {
            return {
                name: this.name,
                image: this.image,
                script: this.scriptData,
                workingDir: this.workingDir,
                env: this._env,
            };
        }
        else {
            return {
                name: this.name,
                image: this.image,
                command: this.command,
                args: this.args,
                workingDir: this.workingDir,
                env: this._env,
            };
        }
    }
}
_c = JSII_RTTI_SYMBOL_1;
TaskStepBuilder[_c] = { fqn: "cdk8s-pipelines.TaskStepBuilder", version: "0.0.16" };
exports.TaskStepBuilder = TaskStepBuilder;
/**
 * Builds Tekton `Task` objects that are independent of a `Pipeline`.
 *
 * To use a builder for tasks that will be used in a Pipeline, use the
 * `PipelineBuilder` instead.
 */
class TaskBuilder {
    /**
     * Creates a new instance of the `TaskBuilder` using the given `scope` and
     * `id`.
     * @param scope
     * @param id
     */
    constructor(scope, id) {
        // These were initially arrays, but converted them to maps so that if
        // multiple values are added that the last one will win.
        this._workspaces = new Map;
        this._params = new Map;
        this._results = new Array;
        this._scope = scope;
        this._id = id;
        // These are required, and it's better to just create it rather than
        // check each time.
        this._steps = new Array();
    }
    get logicalID() {
        return this._id;
    }
    /**
     * Adds a label to the `Task` with the provided label key and value.
     *
     * @see https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
     *
     * @param key
     * @param value
     */
    withLabel(key, value) {
        if (!this._labels) {
            this._labels = {};
        }
        this._labels[key] = value;
        return this;
    }
    /**
     * Adds an annotation to the `Task` `metadata` with the provided key and value.
     *
     * @see https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
     *
     * @param key The annotation's key.
     * @param value The annotation's value.
     */
    withAnnotation(key, value) {
        if (!this._annotations) {
            this._annotations = {};
        }
        this._annotations[key] = value;
        return this;
    }
    /**
     * Gets the name of the `Task` built by the `TaskBuilder`.
     */
    get name() {
        return this._name;
    }
    /**
     * Sets the name of the `Task` being built.
     * @param name
     */
    withName(name) {
        this._name = name;
        return this;
    }
    /**
     * Sets the `description` of the `Task` being built.
     * @param description
     */
    withDescription(description) {
        this._description = description;
        return this;
    }
    /**
     * Gets the `description` of the `Task`.
     */
    get description() {
        return this._description;
    }
    /**
     * Adds the specified workspace to the `Task`.
     * @param workspace
     */
    withWorkspace(workspace) {
        this._workspaces.set(workspace.logicalID, workspace);
        return this;
    }
    /**
     * Gets the workspaces for the `Task`.
     */
    get workspaces() {
        return Array.from(this._workspaces?.values());
    }
    /**
     * Adds a parameter of type string to the `Task`.
     *
     * @param param
     */
    withStringParam(param) {
        this._params.set(param.logicalID, param.ofType('string'));
        return this;
    }
    get parameters() {
        return Array.from(this._params?.values());
    }
    /**
     * Allows you to add an result to the Task.
     *
     * @see https://tekton.dev/docs/pipelines/tasks/#emitting-results
     *
     * @param name The name of the result.
     * @param description The result's description.
     */
    withResult(name, description) {
        // First, check to see if there is already a result with this name
        const existing = this._results.find((obj) => obj.name === name);
        if (existing) {
            throw new Error(`Cannot add result ${name}, as it already exists.`);
        }
        this._results.push({
            name: name,
            description: description,
        });
        return this;
    }
    /**
     * Adds the given `step` (`TaskStepBuilder`) to the `Task`.
     * @param step
     */
    withStep(step) {
        this._steps.push(step);
        return this;
    }
    /**
     * Builds the `Task`.
     */
    buildTask() {
        const taskSteps = new Array();
        this._steps?.forEach((s) => {
            const step = s.buildTaskStep();
            if (step) {
                taskSteps.push(step);
            }
        });
        const taskParams = new Array();
        this._params?.forEach((p) => {
            taskParams.push({
                name: p.logicalID,
                description: p.description,
                default: p.defaultValue,
            });
        });
        const taskWorkspaces = new Array();
        this._workspaces?.forEach((ws) => {
            taskWorkspaces.push({
                name: ws.logicalID,
                description: ws.description,
            });
        });
        const props = {
            metadata: {
                name: this.name,
                labels: this._labels,
                annotations: this._annotations,
            },
            spec: {
                description: this.description,
                workspaces: taskWorkspaces,
                params: taskParams,
                steps: taskSteps,
                results: this._results,
            },
        };
        new tasks_1.Task(this._scope, this._id, props);
    }
}
_d = JSII_RTTI_SYMBOL_1;
TaskBuilder[_d] = { fqn: "cdk8s-pipelines.TaskBuilder", version: "0.0.16" };
exports.TaskBuilder = TaskBuilder;
/**
 *
 */
class PipelineBuilder {
    constructor(scope, id) {
        this._scope = scope;
        this._id = id;
    }
    /**
     * Provides the name for the pipeline task and will be
     * rendered as the `name` property.
     * @param name
     */
    withName(name) {
        this._name = name;
        return this;
    }
    /**
     * Gets the name of the pipeline
     */
    get name() {
        return this._name || this._id;
    }
    /**
     * Provides the name for the pipeline task and will be
     * rendered as the `name` property.
     * @param description
     */
    withDescription(description) {
        this._description = description;
        return this;
    }
    // Adds the task to the pipeline.
    withTask(taskB) {
        // Add the task to the list of tasks...
        if (!this._tasks) {
            this._tasks = new Array();
        }
        this._tasks.push(taskB);
        return this;
    }
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
    get params() {
        // Not trying to prematurely optimize here, but this could be an expensive
        // operation, so we only need to do it if the state of the object has not
        // changed.
        const pipelineParams = new Map();
        this._tasks?.forEach((t) => {
            t.parameters?.forEach(p => {
                const pp = pipelineParams.get(p.name);
                if (!pp) {
                    // Do not add it to the pipeline if there is no need to add it...
                    if (p.requiresPipelineParameter) {
                        pipelineParams.set(p.name, {
                            name: p.name,
                            type: p.type,
                        });
                    }
                }
            });
        });
        return Array.from(pipelineParams.values());
    }
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
    get workspaces() {
        const pipelineWorkspaces = new Map();
        this._tasks?.forEach((t) => {
            t.workspaces?.forEach((w) => {
                // Only add the workspace on the pipeline level if it is not already
                // there...
                const ws = pipelineWorkspaces.get(w.name);
                if (!ws) {
                    pipelineWorkspaces.set(w.name, {
                        name: w.name,
                        description: w.description,
                    });
                }
            });
        });
        return Array.from(pipelineWorkspaces.values());
    }
    /**
     * Builds the actual [Pipeline](https://tekton.dev/docs/getting-started/pipelines/)
     * from the settings configured using the fluid syntax.
     */
    buildPipeline(opts = exports.DefaultBuilderOptions) {
        // TODO: validate the object
        const pipelineTasks = new Array();
        // For making a list to make sure that tasks aren't duplicated when doing
        // the build. Not that it really hurts anything, but it makes the multidoc
        // YAML file bigger and more complex than it needs to be.
        const taskList = new Array();
        this._tasks?.forEach((t, i) => {
            const taskParams = new Array();
            const taskWorkspaces = new Array();
            t.parameters?.forEach(p => {
                taskParams.push({
                    name: p.logicalID,
                    value: p.value,
                });
            });
            t.workspaces?.forEach((w) => {
                taskWorkspaces.push({
                    name: w.logicalID,
                    workspace: w.name,
                });
            });
            const pt = createOrderedPipelineTask(t, ((i > 0) ? this._tasks[i - 1].logicalID : ''), taskParams, taskWorkspaces);
            pipelineTasks.push(pt);
            if (opts.includeDependencies) {
                // Build the task if the user has asked for the dependencies to be
                // built along with the pipeline, but only if we haven't already
                // built the task yet.
                if (!taskList.find(it => {
                    return it == t.name;
                })) {
                    t.buildTask();
                }
                taskList.push(t.name);
            }
        });
        new pipelines_1.Pipeline(this._scope, this._id, {
            metadata: {
                name: this.name,
            },
            spec: {
                description: this._description,
                params: this.params,
                workspaces: this.workspaces,
                tasks: pipelineTasks,
            },
        });
    }
}
_e = JSII_RTTI_SYMBOL_1;
PipelineBuilder[_e] = { fqn: "cdk8s-pipelines.PipelineBuilder", version: "0.0.16" };
exports.PipelineBuilder = PipelineBuilder;
function createOrderedPipelineTask(t, after, params, ws) {
    if (after) {
        return {
            name: t.logicalID,
            taskRef: {
                name: t.name,
            },
            runAfter: [after],
            params: params,
            workspaces: ws,
        };
    }
    return {
        name: t.logicalID,
        taskRef: {
            name: t.name,
        },
        params: params,
        workspaces: ws,
    };
}
/**
 * Builds a `PipelineRun` using the supplied configuration.
 *
 * @see https://tekton.dev/docs/pipelines/pipelineruns/
 */
class PipelineRunBuilder {
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
    constructor(scope, id, pipeline) {
        this._scope = scope;
        this._id = id;
        this._pipeline = pipeline;
        this._sa = DefaultPipelineServiceAccountName;
        this._crbProps = DefaultClusterRoleBindingProps;
        this._runParams = new Array();
        this._runWorkspaces = new Array();
    }
    /**
     * Adds a run parameter to the `PipelineRun`. It will throw an error if you try
     * to add a parameter that does not exist on the pipeline.
     *
     * @param name The name of the parameter added to the pipeline run.
     * @param value The value of the parameter added to the pipeline run.
     */
    withRunParam(name, value) {
        const params = this._pipeline.params;
        const p = params.find((obj) => obj.name === name);
        if (p) {
            this._runParams.push({
                name: name,
                value: value,
            });
        }
        else {
            throw new Error(`PipelineRun parameter '${name}' does not exist in pipeline '${this._pipeline.name}'`);
        }
        return this;
    }
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
    withWorkspace(name, claimName, subPath) {
        this._runWorkspaces.push({
            name: name,
            persistentVolumeClaim: {
                claimName: claimName,
            },
            subPath: subPath,
        });
        return this;
    }
    withClusterRoleBindingProps(props) {
        this._crbProps = props;
        return this;
    }
    /**
     * Uses the provided role name for the `serviceAccountName` on the
     * `PipelineRun`. If this method is not called prior to `buildPipelineRun()`,
     * then the default service account will be used, which is _default:pipeline_.
     *
     * @param sa The name of the service account (`serviceAccountName`) to use.
     */
    withServiceAccount(sa) {
        this._sa = sa;
        return this;
    }
    /**
     * Builds the `PipelineRun` for the configured `Pipeline` used in the constructor.
     * @param opts
     */
    buildPipelineRun(opts = exports.DefaultBuilderOptions) {
        if (opts && opts.includeDependencies) {
            // Generate the ClusterRoleBinding document, if configured to do so.
            new cdk8s_1.ApiObject(this._scope, this._crbProps.metadata?.name, this._crbProps);
        }
        // Throw an error here if the parameters are not defined that are required
        // by the Pipeline, because there is really no point in going any further.
        const params = this._pipeline.params;
        params.forEach((p) => {
            const prp = this._runParams.find((obj) => obj.name == p.name);
            if (!prp) {
                throw new Error(`Pipeline parameter '${p.name}' is not defined in PipelineRun '${this._id}'`);
            }
        });
        // Do the same thing for workspaces. Check to make sure that the workspaces
        // expected by the Pipeline are defined in the PipelineRun.
        const workspaces = this._pipeline.workspaces;
        workspaces.forEach((ws) => {
            const pws = this._runWorkspaces.find((obj) => obj.name == ws.name);
            if (!pws) {
                throw new Error(`Pipeline workspace '${ws.name}' is not defined in PipelineRun '${this._id}'`);
            }
        });
        new pipelines_1.PipelineRun(this._scope, this._id, {
            metadata: {
                name: this._id,
            },
            serviceAccountName: this._sa,
            spec: {
                pipelineRef: {
                    name: this._pipeline.name,
                },
                params: this._runParams,
                workspaces: this._runWorkspaces,
            },
        });
    }
}
_f = JSII_RTTI_SYMBOL_1;
PipelineRunBuilder[_f] = { fqn: "cdk8s-pipelines.PipelineRunBuilder", version: "0.0.16" };
exports.PipelineRunBuilder = PipelineRunBuilder;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiYnVpbGRlcnMuanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyIuLi9zcmMvYnVpbGRlcnMudHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6Ijs7Ozs7QUFBQTs7R0FFRztBQUdILHlCQUF5QjtBQUN6QixpQ0FBd0Q7QUFFeEQscUNBQXNDO0FBQ3RDLDJDQVNxQjtBQUNyQixtQ0FBOEk7QUFFOUksTUFBTSxpQ0FBaUMsR0FBRyxrQkFBa0IsQ0FBQztBQUU3RDs7Ozs7OztHQU9HO0FBQ0gsU0FBUyxzQkFBc0IsQ0FBQyxXQUFtQixFQUFFLFNBQWlCLEVBQUUsUUFBZ0IsRUFBRSxFQUFVLEVBQUUsV0FBbUI7SUFDdkgsT0FBTztRQUNMLFVBQVUsRUFBRSw4QkFBOEI7UUFDMUMsSUFBSSxFQUFFLG9CQUFvQjtRQUMxQixRQUFRLEVBQUU7WUFDUixJQUFJLEVBQUUsV0FBVztZQUNqQixTQUFTLEVBQUUsU0FBUztTQUNyQjtRQUNELE9BQU8sRUFBRTtZQUNQLElBQUksRUFBRSxhQUFhO1lBQ25CLElBQUksRUFBRSxRQUFRO1NBQ2Y7UUFDRCxRQUFRLEVBQUU7WUFDUjtnQkFDRSxJQUFJLEVBQUUsZ0JBQWdCO2dCQUN0QixJQUFJLEVBQUUsRUFBRTtnQkFDUixTQUFTLEVBQUUsV0FBVzthQUN2QjtTQUNGO0tBQ0YsQ0FBQztBQUNKLENBQUM7QUFFRCxNQUFNLDhCQUE4QixHQUFHLHNCQUFzQixDQUMzRCw0QkFBNEIsRUFDNUIsU0FBUyxFQUNULGVBQWUsRUFDZixVQUFVLEVBQ1YsU0FBUyxDQUFDLENBQUM7QUFjYjs7R0FFRztBQUNVLFFBQUEscUJBQXFCLEdBQW1CO0lBQ25ELG1CQUFtQixFQUFFLEtBQUs7Q0FDM0IsQ0FBQztBQUVGOztHQUVHO0FBQ0gsTUFBYSxnQkFBZ0I7SUFLM0I7Ozs7T0FJRztJQUNILFlBQVksRUFBVTtRQUNwQixJQUFJLENBQUMsVUFBVSxHQUFHLEVBQUUsQ0FBQztJQUN2QixDQUFDO0lBRUQ7O09BRUc7SUFDSCxJQUFXLFNBQVM7UUFD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LENBQUMsQ0FBQyxHQUFHLEVBQUUsRUFBRSxDQUFDLEdBQUcsQ0FBQyxJQUFJLElBQUksRUFBRSxDQUFDLElBQUksQ0FBQyxDQUFDO1lBQ25FLElBQUksQ0FBRSxHQUFHLEVBQUU7Z0JBQ1QsTUFBTSxJQUFJLEtBQUssQ0FBQyx1QkFBdUIsRUFBRSxDQUFDLElBQUksb0NBQW9DLElBQUksQ0FBQyxHQUFHLEdBQUcsQ0FBQyxDQUFDO2FBQ2hHO1FBQ0gsQ0FBQyxDQUFDLENBQUM7UUFFSCxJQUFJLHVCQUFXLENBQUMsSUFBSSxDQUFDLE1BQU0sRUFBRSxJQUFJLENBQUMsR0FBRyxFQUFFO1lBQ3JDLFFBQVEsRUFBRTtnQkFDUixJQUFJLEVBQUUsSUFBSSxDQUFDLEdBQUc7YUFDZjtZQUNELGtCQUFrQixFQUFFLElBQUksQ0FBQyxHQUFHO1lBQzVCLElBQUksRUFBRTtnQkFDSixXQUFXLEVBQUU7b0JBQ1gsSUFBSSxFQUFFLElBQUksQ0FBQyxTQUFTLENBQUMsSUFBSTtpQkFDMUI7Z0JBQ0QsTUFBTSxFQUFFLElBQUksQ0FBQyxVQUFVO2dCQUN2QixVQUFVLEVBQUUsSUFBSSxDQUFDLGNBQWM7YUFDaEM7U0FDRixDQUFDLENBQUM7SUFDTCxDQUFDOzs7O0FBcklVLGdEQUFrQiIsInNvdXJjZXNDb250ZW50IjpbIi8qKlxuICogVGhpcyBmaWxlIGhhcyBidWlsZGVycyBpbiBpdCBmb3IgdGhlIHZhcmlvdXMgcGlwZWxpbmUgY29uc3RydWN0cy5cbiAqL1xuXG5cbmltcG9ydCAqIGFzIGZzIGZyb20gJ2ZzJztcbmltcG9ydCB7IEFwaU9iamVjdCwgQXBpT2JqZWN0UHJvcHMsIFlhbWwgfSBmcm9tICdjZGs4cyc7XG5pbXBvcnQgeyBDb25zdHJ1Y3QgfSBmcm9tICdjb25zdHJ1Y3RzJztcbmltcG9ydCB7IGJ1aWxkUGFyYW0gfSBmcm9tICcuL2NvbW1vbic7XG5pbXBvcnQge1xuICBQaXBlbGluZSxcbiAgUGlwZWxpbmVQYXJhbSxcbiAgUGlwZWxpbmVSdW4sXG4gIFBpcGVsaW5lUnVuUGFyYW0sXG4gIFBpcGVsaW5lUnVuV29ya3NwYWNlLFxuICBQaXBlbGluZVRhc2ssXG4gIFBpcGVsaW5lVGFza1dvcmtzcGFjZSxcbiAgUGlwZWxpbmVXb3Jrc3BhY2UsXG59IGZyb20gJy4vcGlwZWxpbmVzJztcbmltcG9ydCB7IFRhc2ssIFRhc2tFbnZWYWx1ZVNvdXJjZSwgVGFza1BhcmFtLCBUYXNrUHJvcHMsIFRhc2tTcGVjUGFyYW0sIFRhc2tTcGVjUmVzdWx0LCBUYXNrU3RlcCwgVGFza1N0ZXBFbnYsIFRhc2tXb3Jrc3BhY2UgfSBmcm9tICcuL3Rhc2tzJztcblxuY29uc3QgRGVmYXVsdFBpcGVsaW5lU2VydmljZUFjY291bnROYW1lID0gJ2RlZmF1bHQ6cGlwZWxpbmUnO1xuXG4vKipcbiAqIENyZWF0ZXMgdGhlIHByb3BlcnRpZXMgZm9yIGEgYENsdXN0ZXJSb2xlQmluZGluZ2BcbiAqIEBwYXJhbSBiaW5kaW5nTmFtZVxuICogQHBhcmFtIGJpbmRpbmdOc1xuICogQHBhcmFtIHJvbGVuYW1lXG4gKiBAcGFyYW0gc2FcbiAqIEBwYXJhbSBzYU5hbWVzcGFjZVxuICovXG5mdW5jdGlvbiBjcmVhdGVSb2xlQmluZGluZ1Byb3BzKGJpbmRpbmdOYW1lOiBzdHJpbmcsIGJpbmRpbmdOczogc3RyaW5nLCByb2xlbmFtZTogc3RyaW5nLCBzYTogc3RyaW5nLCBzYU5hbWVzcGFjZTogc3RyaW5nKTogQXBpT2JqZWN0UHJvcHMge1xuICByZXR1cm4ge1xuICAgIGFwaVZlcnNpb246ICdyYmFjLmF1dGhvcml6YXRpb24uazhzLmlvL3YxJyxcbiAgICBraW5kOiAnQ2x1c3RlclJvbGVCaW5kaW5nJyxcbiAgICBtZXRhZGF0YToge1xuICAgICAgbmFtZTogYmluZGluZ05hbWUsXG4gICAgICBuYW1lc3BhY2U6IGJpbmRpbmdOcyxcbiAgICB9LFxuICAgIHJvbGVSZWY6IHtcbiAgICAgIGtpbmQ6ICdDbHVzdGVyUm9sZScsXG4gICAgICBuYW1lOiByb2xlbmFtZSxcbiAgICB9LFxuICAgIHN1YmplY3RzOiBbXG4gICAgICB7XG4gICAgICAgIGtpbmQ6ICdTZXJ2aWNlQWNjb3VudCcsXG4gICAgICAgIG5hbWU6IHNhLFxuICAgICAgICBuYW1lc3BhY2U6IHNhTmFtZXNwYWNlLFxuICAgICAgfSxcbiAgICBdLFxuICB9O1xufVxuXG5jb25zdCBEZWZhdWx0Q2x1c3RlclJvbGVCaW5kaW5nUHJvcHMgPSBjcmVhdGVSb2xlQmluZGluZ1Byb3BzKFxuICAncGlwZWxpbmUtYWRtaW4tZGVmYXVsdC1jcmInLFxuICAnZGVmYXVsdCcsXG4gICdjbHVzdGVyLWFkbWluJyxcbiAgJ3BpcGVsaW5lJyxcbiAgJ2RlZmF1bHQnKTtcblxuLyoqXG4gKiBUaGUgb3B0aW9ucyBmb3IgYnVpbGRlcnMgZm9yIHRoZSBgYnVpbGRYWCgpYCBtZXRob2RzLlxuICovXG5leHBvcnQgaW50ZXJmYWNlIEJ1aWxkZXJPcHRpb25zIHtcbiAgLyoqXG4gICAqIElmIHRydWUsIGFsbCB0aGUgZGVwZW5kZW50IG9iamVjdHMgYXJlIGdlbmVyYXRlZCB3aXRoIHRoZSBidWlsZC4gVGhpcyBpc1xuICAgKiBkZXNpZ25lZCB0byBydW4gb24gYXMgbWluaW1hbCBjbHVzdGVyIGFzIHBvc3NpYmxlLCB3aXRoIGFzIGZldyBwcmUgc3RlcHNcbiAgICogYXMgcG9zc2libGUuXG4gICAqL1xuICByZWFkb25seSBpbmNsdWRlRGVwZW5kZW5jaWVzPzogYm9vbGVhbjtcbn1cblxuLyoqXG4gKiBUaGUgZGVmYXVsdCBvcHRpb25zIGZvciB0aGUgYnVpbGRlcnMuXG4gKi9cbmV4cG9ydCBjb25zdCBEZWZhdWx0QnVpbGRlck9wdGlvbnM6IEJ1aWxkZXJPcHRpb25zID0ge1xuICBpbmNsdWRlRGVwZW5kZW5jaWVzOiBmYWxzZSxcbn07XG5cbi8qKlxuICogQnVpbGRzIHRoZSBXb3Jrc3BhY2VzIGZvciB1c2UgYnkgVGFza3MgYW5kIFBpcGVsaW5lcy5cbiAqL1xuZXhwb3J0IGNsYXNzIFdvcmtzcGFjZUJ1aWxkZXIge1xuICBwcml2YXRlIHJlYWRvbmx5IF9sb2dpY2FsSUQ6IHN0cmluZztcbiAgcHJpdmF0ZSBfbmFtZT86IHN0cmluZztcbiAgcHJpdmF0ZSBfZGVzY3JpcHRpb24/OiBzdHJpbmc7XG5cbiAgLyoqXG4gICAqIENyZWF0ZXMgdGhlIGBXb3Jrc3BhY2VCdWlsZGVyYCwgdXNpbmcgdGhlIGdpdmVuIGBpZGAgYXMgdGhlIGxvZ2ljYWwgSUQgZm9yXG4gICAqIHRoZSB3b3Jrc3BhY2UuXG4gICAqIEBwYXJhbSBpZFxuICAgKi9cbiAgY29uc3RydWN0b3IoaWQ6IHN0cmluZykge1xuICAgIHRoaXMuX2xvZ2ljYWxJRCA9IGlkO1xuICB9XG5cbiAgLyoqXG4gICAqIEdldHMgdGhlIGxvZ2ljYWwgSUQgb2YgdGhlIGBXb3Jrc3BhY2VgLlxuICAgKi9cbiAgcHVibGljIGdldCBsb2dpY2FsSUQoKTogc3RyaW5nIHwgdW5kZWZpbmVkIHtcbiAgICByZXR1cm4gdGhpcy5fbG9naWNhbElEO1xuICB9XG5cbiAgLyoqXG4gICAqIEdldHMgdGhlIG5hbWUgb2YgdGhlIHdvcmtzcGFjZS5cbiAgICovXG4gIHB1YmxpYyBnZXQgbmFtZSgpOiBzdHJpbmcgfCB1bmRlZmluZWQge1xuICAgIHJldHVybiB0aGlzLl9uYW1lO1xuICB9XG5cbiAgLyoqXG4gICAqIEdldHMgdGhlIGRlc2NyaXB0aW9uIG9mIHRoZSB3b3Jrc3BhY2UuXG4gICAqL1xuICBwdWJsaWMgZ2V0IGRlc2NyaXB0aW9uKCk6IHN0cmluZyB7XG4gICAgcmV0dXJuIHRoaXMuX2Rlc2NyaXB0aW9uIHx8ICcnO1xuICB9XG5cbiAgcHVibGljIHdpdGhOYW1lKG5hbWU6IHN0cmluZyk6IFdvcmtzcGFjZUJ1aWxkZXIge1xuICAgIHRoaXMuX25hbWUgPSBuYW1lO1xuICAgIHJldHVybiB0aGlzO1xuICB9XG5cbiAgcHVibGljIHdpdGhEZXNjcmlwdGlvbihkZXNjOiBzdHJpbmcpOiBXb3Jrc3BhY2VCdWlsZGVyIHtcbiAgICB0aGlzLl9kZXNjcmlwdGlvbiA9IGRlc2M7XG4gICAgcmV0dXJuIHRoaXM7XG4gIH1cblxufVxuXG4vKipcbiAqIEJ1aWxkcyB0aGUgcGFyYW1ldGVycyBmb3IgdXNlIGJ5IFRhc2tzIGFuZCBQaXBlbGluZXMuXG4gKi9cbmV4cG9ydCBjbGFzcyBQYXJhbWV0ZXJCdWlsZGVyIHtcbiAgcHJpdmF0ZSByZWFkb25seSBfbG9naWNhbElEOiBzdHJpbmc7XG4gIHByaXZhdGUgX25hbWU/OiBzdHJpbmc7XG4gIHByaXZhdGUgX2Rlc2NyaXB0aW9uPzogc3RyaW5nO1xuICBwcml2YXRlIF90eXBlPzogc3RyaW5nO1xuICBwcml2YXRlIF92YWx1ZT86IHN0cmluZztcbiAgcHJpdmF0ZSBfZGVmYXVsdFZhbHVlPzogc3RyaW5nO1xuICBwcml2YXRlIF9yZXF1aXJlc1BpcGVsaW5lUGFyYW06IGJvb2xlYW47XG5cbiAgY29uc3RydWN0b3IoaWQ6IHN0cmluZykge1xuICAgIHRoaXMuX2xvZ2ljYWxJRCA9IGlkO1xuICAgIHRoaXMuX3JlcXVpcmVzUGlwZWxpbmVQYXJhbSA9IGZhbHNlO1xuICB9XG5cbiAgLyoqXG4gICAqIEdldHMgdGhlIGxvZ2ljYWxJRCBmb3IgdGhlIGBQYXJhbWV0ZXJCdWlsZGVyYCwgd2hpY2ggaXMgdXNlZCBieSB0aGUgdW5kZXJseWluZ1xuICAgKiBjb25zdHJ1Y3QuXG4gICAqL1xuICBwdWJsaWMgZ2V0IGxvZ2ljYWxJRCgpOiBzdHJpbmcgfCB1bmRlZmluZWQge1xuICAgIHJldHVybiB0aGlzLl9sb2dpY2FsSUQ7XG4gIH1cblxuICBwdWJsaWMgZ2V0IG5hbWUoKTogc3RyaW5nIHwgdW5kZWZpbmVkIHtcbiAgICByZXR1cm4gdGhpcy5fbmFtZTtcbiAgfVxuXG4gIHB1YmxpYyBnZXQgZGVzY3JpcHRpb24oKTogc3RyaW5nIHtcbiAgICByZXR1cm4gdGhpcy5fZGVzY3JpcHRpb24gfHwgJyc7XG4gIH1cblxuICAvKipcbiAgICogU2V0cyB0aGUgbmFtZSBvZiB0aGUgcGFyYW1ldGVyLlxuICAgKiBAcGFyYW0gbmFtZVxuICAgKi9cbiAgcHVibGljIHdpdGhOYW1lKG5hbWU6IHN0cmluZyk6IFBhcmFtZXRlckJ1aWxkZXIge1xuICAgIHRoaXMuX25hbWUgPSBuYW1lO1xuICAgIHJldHVybiB0aGlzO1xuICB9XG5cbiAgLyoqXG4gICAqIFNldHMgdGhlIGRlc2NyaXB0aW9uIG9mIHRoZSBwYXJhbWV0ZXIuXG4gICAqIEBwYXJhbSBkZXNjXG4gICAqL1xuICBwdWJsaWMgd2l0aERlc2NyaXB0aW9uKGRlc2M6IHN0cmluZyk6IFBhcmFtZXRlckJ1aWxkZXIge1xuICAgIHRoaXMuX2Rlc2NyaXB0aW9uID0gZGVzYztcbiAgICByZXR1cm4gdGhpcztcbiAgfVxuXG4gIC8qKlxuICAgKiBTZXRzIHRoZSB0eXBlIG9mIHRoZSBwYXJhbWV0ZXJcbiAgICogQHBhcmFtIHR5cGVcbiAgICovXG4gIHB1YmxpYyBvZlR5cGUodHlwZTogc3RyaW5nKTogUGFyYW1ldGVyQnVpbGRlciB7XG4gICAgdGhpcy5fdHlwZSA9IHR5cGU7XG4gICAgcmV0dXJuIHRoaXM7XG4gIH1cblxuICAvKipcbiAgICogR2V0cyB0aGUgdHlwZSBvZiB0aGUgcGFyYW1ldGVyXG4gICAqL1xuICBwdWJsaWMgZ2V0IHR5cGUoKTogc3RyaW5nIHwgdW5kZWZpbmVkIHtcbiAgICByZXR1cm4gdGhpcy5fdHlwZTtcbiAgfVxuXG4gIC8qKlxuICAgKiBTZXRzIHRoZSB2YWx1ZSBmb3IgdGhlIHBhcmFtZXRlclxuICAgKiBAcGFyYW0gdmFsXG4gICAqL1xuICBwdWJsaWMgd2l0aFZhbHVlKHZhbDogc3RyaW5nKTogUGFyYW1ldGVyQnVpbGRlciB7XG4gICAgLy8gSWYgeW91IGFyZSBnaXZpbmcgaXQgYSB2YWx1ZSBoZXJlLCB0aGVuIHlvdSBkbyBub3RcbiAgICAvLyBuZWVkIHRoZSBQaXBlbGluZSBwYXJhbWV0ZXIgZm9yIHRoaXMgcGFyYW1ldGVyLlxuICAgIHRoaXMuX3JlcXVpcmVzUGlwZWxpbmVQYXJhbSA9IGZhbHNlO1xuICAgIHRoaXMuX3ZhbHVlID0gdmFsO1xuICAgIHJldHVybiB0aGlzO1xuICB9XG5cbiAgLyoqXG4gICAqIEdldHMgdGhlIHZhbHVlIG9mIHRoZSBwYXJhbWV0ZXJcbiAgICovXG4gIHB1YmxpYyBnZXQgdmFsdWUoKTogc3RyaW5nIHwgdW5kZWZpbmVkIHtcbiAgICByZXR1cm4gdGhpcy5fdmFsdWU7XG4gIH1cblxuICAvKipcbiAgICogU2V0cyB0aGUgZGVmYXVsdCB2YWx1ZSBmb3IgdGhlIHBhcmFtZXRlci5cbiAgICogQHBhcmFtIHZhbFxuICAgKi9cbiAgcHVibGljIHdpdGhEZWZhdWx0VmFsdWUodmFsOiBzdHJpbmcpOiBQYXJhbWV0ZXJCdWlsZGVyIHtcbiAgICB0aGlzLl9kZWZhdWx0VmFsdWUgPSB2YWw7XG4gICAgcmV0dXJuIHRoaXM7XG4gIH1cblxuICBwdWJsaWMgZ2V0IGRlZmF1bHRWYWx1ZSgpOiBzdHJpbmcgfCB1bmRlZmluZWQge1xuICAgIHJldHVybiB0aGlzLl9kZWZhdWx0VmFsdWU7XG4gIH1cblxuICAvKipcbiAgICogU2V0cyB0aGUgZGVmYXVsdCB2YWx1ZSBmb3IgdGhlIHBhcmFtZXRlci5cbiAgICogQHBhcmFtIHBpcGVsaW5lUGFyYW1OYW1lXG4gICAqIEBwYXJhbSBkZWZhdWx0VmFsdWVcbiAgICovXG4gIHB1YmxpYyB3aXRoUGlwbGluZVBhcmFtZXRlcihwaXBlbGluZVBhcmFtTmFtZTogc3RyaW5nLCBkZWZhdWx0VmFsdWU6IHN0cmluZyA9ICcnKTogUGFyYW1ldGVyQnVpbGRlciB7XG4gICAgdGhpcy5fcmVxdWlyZXNQaXBlbGluZVBhcmFtID0gdHJ1ZTtcbiAgICB0aGlzLl9uYW1lID0gcGlwZWxpbmVQYXJhbU5hbWU7XG4gICAgdGhpcy5fZGVmYXVsdFZhbHVlID0gZGVmYXVsdFZhbHVlO1xuICAgIHRoaXMuX3ZhbHVlID0gYnVpbGRQYXJhbShwaXBlbGluZVBhcmFtTmFtZSk7XG4gICAgcmV0dXJuIHRoaXM7XG4gIH1cblxuICAvKipcbiAgICogUmV0dXJucyB0cnVlIGlmIHRoaXMgcGFyYW1ldGVyIGV4cGVjdHMgaW5wdXQgYXQgdGhlIHBpcGVsaW5lIGxldmVsLlxuICAgKi9cbiAgcHVibGljIGdldCByZXF1aXJlc1BpcGVsaW5lUGFyYW1ldGVyKCk6IGJvb2xlYW4ge1xuICAgIHJldHVybiB0aGlzLl9yZXF1aXJlc1BpcGVsaW5lUGFyYW07XG4gIH1cbn1cblxuLyoqXG4gKiBSZXNvbHZlcyB0aGUgYHNjcmlwdGAgdGhyb3VnaCBkaWZmZXJlbnQgbWVhbnMuXG4gKi9cbmludGVyZmFjZSBTY3JpcHRSZXNvbHZlciB7XG4gIC8qKlxuICAgKiBHZXRzIHRoZSBib2R5IG9mIHRoZSBzY3JpcHQuXG4gICAqIEByZXR1cm5zIHN0cmluZyBUaGUgc2NyaXB0LlxuICAgKi9cbiAgc2NyaXB0RGF0YSgpOiBzdHJpbmc7XG59XG5cbi8qKlxuICogUmVzb2x2ZXMgdGhlIHByb3ZpZGVkIG9iamVjdCBpbnRvIGEgWUFNTCBzdHJpbmcuXG4gKi9cbmNsYXNzIE9ialNjcmlwdFJlc29sdmVyIGltcGxlbWVudHMgU2NyaXB0UmVzb2x2ZXIge1xuICByZWFkb25seSBfb2JqOiBhbnk7XG5cbiAgLyoqXG4gICAqIENyZWF0ZXMgYW4gaW5zdGFuY2Ugb2YgdGhlIGBPYmpTY3JpcHRSZXNvbHZlcmAuXG4gICAqIEBwYXJhbSBvYmogVGhlIG9iamVjdCB0byBzZXJpYWxpemUgdG8gWUFNTCBmb3IgdGhlIHNjcmlwdC5cbiAgICovXG4gIGNvbnN0cnVjdG9yKG9iajogYW55KSB7XG4gICAgdGhpcy5fb2JqID0gb2JqO1xuICB9XG5cbiAgLyoqXG4gICAqIEdldHMgdGhlIGJvZHkgb2YgdGhlIHNjcmlwdCBhcyBhIFlBTUwgcmVwcmVzZW50YXRpb24gb2YgdGhlIG9iamVjdC5cbiAgICovXG4gIHB1YmxpYyBzY3JpcHREYXRhKCk6IHN0cmluZyB7XG4gICAgcmV0dXJuIFlhbWwuc3RyaW5naWZ5KHRoaXMuX29iaik7XG4gIH1cbn1cblxuLyoqXG4gKiBHZXRzIHRoZSBjb250ZW50IGZyb20gdGhlIHByb3ZpZGVkIFVSTCBhbmQgcmV0dXJucyBpdCBhcyB0aGUgc2NyaXB0IGRhdGEuXG4gKi9cbmNsYXNzIFVybFNjcmlwdFJlc29sdmVyIGltcGxlbWVudHMgU2NyaXB0UmVzb2x2ZXIge1xuICByZWFkb25seSBfdXJsOiBzdHJpbmc7XG5cbiAgLyoqXG4gICAqIENyZWF0ZXMgYW4gaW5zdGFuY2Ugb2YgdGhlIGBVcmxTY3JpcHRSZXNvbHZlcmAgd2l0aCB0aGUgcHJvdmlkZWQgVVJMLlxuICAgKiBAcGFyYW0gdXJsXG4gICAqL1xuICBjb25zdHJ1Y3Rvcih1cmw6IHN0cmluZykge1xuICAgIHRoaXMuX3VybCA9IHVybDtcbiAgfVxuXG4gIC8qKlxuICAgKiBHZXRzIHRoZSBib2R5IG9mIHRoZSBzY3JpcHQgZnJvbSB0aGUgcHJvdmlkZWQgVVJMLlxuICAgKiBAcmV0dXJuIHN0cmluZyBTY3JpcHQgZGF0YS5cbiAgICovXG4gIHB1YmxpYyBzY3JpcHREYXRhKCk6IHN0cmluZyB7XG4gICAgY29uc3QgZGF0YSA9IGZzLnJlYWRGaWxlU3luYyh0aGlzLl91cmwsIHtcbiAgICAgIGVuY29kaW5nOiAndXRmOCcsXG4gICAgICBmbGFnOiAncicsXG4gICAgfSk7XG5cbiAgICByZXR1cm4gZGF0YS5yZXBsYWNlKC9cXG4vZywgJ1xcXFxuJyk7XG4gIH1cbn1cblxuLyoqXG4gKiBHZXRzIHRoZSBjb250ZW50IGZyb20gdGhlIHN0YXRpYyB2YWx1ZSBwcm92aWRlZC5cbiAqL1xuY2xhc3MgU3RhdGljU2NyaXB0UmVzb2x2ZXIgaW1wbGVtZW50cyBTY3JpcHRSZXNvbHZlciB7XG4gIHJlYWRvbmx5IF9zY3JpcHQ6IHN0cmluZztcblxuICAvKipcbiAgICogQ3JlYXRlcyBhbiBpbnN0YW5jZSBvZiB0aGUgYFN0YXRpY1NjcmlwdFJlc29sdmVyYC5cbiAgICogQHBhcmFtIGRhdGFcbiAgICovXG4gIGNvbnN0cnVjdG9yKGRhdGE6IHN0cmluZykge1xuICAgIHRoaXMuX3NjcmlwdCA9IGRhdGE7XG4gIH1cblxuICAvKipcbiAgICogUmV0dXJucyB0aGUgc3RhdGljIHZhbHVlIHByb3ZpZGVkLlxuICAgKi9cbiAgcHVibGljIHNjcmlwdERhdGEoKTogc3RyaW5nIHtcbiAgICByZXR1cm4gdGhpcy5fc2NyaXB0O1xuICB9XG59XG5cbi8qKlxuICogQ3JlYXRlcyBhIGBTdGVwYCBpbiBhIGBUYXNrYC5cbiAqL1xuZXhwb3J0IGNsYXNzIFRhc2tTdGVwQnVpbGRlciB7XG4gIHByaXZhdGUgX25hbWU/OiBzdHJpbmc7XG4gIHByaXZhdGUgX2Rpcj86IHN0cmluZztcbiAgcHJpdmF0ZSBfaW1hZ2U/OiBzdHJpbmc7XG4gIHByaXZhdGUgX2NtZD86IHN0cmluZ1tdO1xuICBwcml2YXRlIF9hcmdzPzogc3RyaW5nW107XG4gIHByaXZhdGUgX2Vudj86IFRhc2tTdGVwRW52W107XG4gIHByaXZhdGUgX3NjcmlwdD86IFNjcmlwdFJlc29sdmVyO1xuXG4gIC8qKlxuICAgKlxuICAgKi9cbiAgcHVibGljIGNvbnN0cnVjdG9yKCkge1xuXG4gIH1cblxuICAvKipcbiAgICogVGhlIG5hbWUgb2YgdGhlIGBTdGVwYCBvZiB0aGUgYFRhc2tgLlxuICAgKi9cbiAgcHVibGljIGdldCBuYW1lKCk6IHN0cmluZyB8IHVuZGVmaW5lZCB7XG4gICAgcmV0dXJuIHRoaXMuX25hbWU7XG4gIH1cblxuICAvKipcbiAgICogVGhlIG5hbWUgb2YgdGhlIGNvbnRhaW5lciBgaW1hZ2VgIHVzZWQgdG8gZXhlY3V0ZSB0aGUgYFN0ZXBgIG9mIHRoZVxuICAgKiBgVGFza2AuXG4gICAqL1xuICBwdWJsaWMgZ2V0IGltYWdlKCk6IHN0cmluZyB8IHVuZGVmaW5lZCB7XG4gICAgcmV0dXJuIHRoaXMuX2ltYWdlO1xuICB9XG5cbiAgcHVibGljIGdldCBzY3JpcHREYXRhKCk6IHN0cmluZyB8IHVuZGVmaW5lZCB7XG4gICAgcmV0dXJuIHRoaXMuX3NjcmlwdD8uc2NyaXB0RGF0YSgpO1xuICB9XG5cbiAgLyoqXG4gICAqIEdldHMgdGhlIGNvbW1hbmQtbGluZSBhcmd1bWVudHMgdGhhdCB3aWxsIGJlIHN1cHBsaWVkIHRvIHRoZSBgY29tbWFuZGAuXG4gICAqL1xuICBwdWJsaWMgZ2V0IGFyZ3MoKTogc3RyaW5nW10gfCB1bmRlZmluZWQge1xuICAgIHJldHVybiB0aGlzLl9hcmdzO1xuICB9XG5cbiAgLyoqXG4gICAqIEdldHMgdGhlIGNvbW1hbmQgdXNlZCBmb3IgdGhlIGBTdGVwYCBvbiB0aGUgYFRhc2tgLlxuICAgKi9cbiAgcHVibGljIGdldCBjb21tYW5kKCk6IHN0cmluZ1tdIHwgdW5kZWZpbmVkIHtcbiAgICByZXR1cm4gdGhpcy5fY21kO1xuICB9XG5cbiAgcHVibGljIGdldCB3b3JraW5nRGlyKCk6IHN0cmluZyB8IHVuZGVmaW5lZCB7XG4gICAgcmV0dXJuIHRoaXMuX2RpcjtcbiAgfVxuXG4gIHB1YmxpYyB3aXRoTmFtZShuYW1lOiBzdHJpbmcpOiBUYXNrU3RlcEJ1aWxkZXIge1xuICAgIHRoaXMuX25hbWUgPSBuYW1lO1xuICAgIHJldHVybiB0aGlzO1xuICB9XG5cbiAgLyoqXG4gICAqIFRoZSBuYW1lIG9mIHRoZSBpbWFnZSB0byB1c2Ugd2hlbiBleGVjdXRpbmcgdGhlIGBTdGVwYCBvbiB0aGUgYFRhc2tgXG4gICAqIEBwYXJhbSBpbWdcbiAgICovXG4gIHB1YmxpYyB3aXRoSW1hZ2UoaW1nOiBzdHJpbmcpOiBUYXNrU3RlcEJ1aWxkZXIge1xuICAgIHRoaXMuX2ltYWdlID0gaW1nO1xuICAgIHJldHVybiB0aGlzO1xuICB9XG5cbiAgLyoqXG4gICAqIFRoZSBuYW1lIG9mIHRoZSBjb21tYW5kIHRvIHVzZSB3aGVuIHJ1bm5pbmcgdGhlIGBTdGVwYCBvZiB0aGUgYFRhc2tgLiBJZlxuICAgKiBgY29tbWFuZGAgaXMgc3BlY2lmaWVkLCBkbyBub3Qgc3BlY2lmeSBgc2NyaXB0YC5cbiAgICogQHBhcmFtIGNtZFxuICAgKi9cbiAgcHVibGljIHdpdGhDb21tYW5kKGNtZDogc3RyaW5nW10pOiBUYXNrU3RlcEJ1aWxkZXIge1xuICAgIHRoaXMuX2NtZCA9IGNtZDtcbiAgICByZXR1cm4gdGhpcztcbiAgfVxuXG4gIC8qKlxuICAgKiBUaGUgYXJncyB0byB1c2Ugd2l0aCB0aGUgYGNvbW1hbmRgLlxuICAgKiBAcGFyYW0gYXJnc1xuICAgKi9cbiAgcHVibGljIHdpdGhBcmdzKGFyZ3M6IHN0cmluZ1tdKTogVGFza1N0ZXBCdWlsZGVyIHtcbiAgICB0aGlzLl9hcmdzID0gYXJncztcbiAgICByZXR1cm4gdGhpcztcbiAgfVxuXG4gIC8qKlxuICAgKiBJZiBzdXBwbGllZCwgdXNlcyB0aGUgY29udGVudCBmb3VuZCBhdCB0aGUgZ2l2ZW4gVVJMIGZvciB0aGVcbiAgICogYHNjcmlwdGAgdmFsdWUgb2YgdGhlIHN0ZXAuIFVzZSB0aGlzIGFzIGFuIGFsdGVybmF0aXZlIHRvIFwiaGVyZWRvY1wiLCB3aGljaFxuICAgKiBpcyBlbWJlZGRpbmcgaGFyZC1jb2RlZCBzaGVsbCBvciBvdGhlciBzY3JpcHRzIGluIHRoZSBzdGVwLlxuICAgKlxuICAgKiBJZiB5b3Ugc3VwcGx5IHRoaXMsIGRvIG5vdCBzdXBwbHkgYSB2YWx1ZSBmb3IgYGZyb21TY3JpcHRPYmplY3RgLlxuICAgKiBAcGFyYW0gdXJsXG4gICAqL1xuICBwdWJsaWMgZnJvbVNjcmlwdFVybCh1cmw6IHN0cmluZyk6IFRhc2tTdGVwQnVpbGRlciB7XG4gICAgdGhpcy5fc2NyaXB0ID0gbmV3IFVybFNjcmlwdFJlc29sdmVyKHVybCk7XG4gICAgcmV0dXJuIHRoaXM7XG4gIH1cblxuICAvKipcbiAgICogSWYgc3VwcGxpZWQsIHVzZXMgdGhlIGNkazhzIGBBcGlPYmplY3RgIHN1cHBsaWVkIGFzIHRoZSBib2R5IG9mIHRoZVxuICAgKiBgc2NyaXB0YCBmb3IgdGhlIGBUYXNrYC4gVGhpcyBpcyBtb3N0IHVzZWZ1bCB3aGVuIHVzZWQgd2l0aCBgb2MgYXBwbHlgIG9yXG4gICAqIG90aGVyIHRhc2tzIGluIHdoaWNoIHlvdSB3YW50IHRvIGFwcGx5IHRoZSBvYmplY3QgZHVyaW5nIHRoZSBzdGVwIGluIHRoZVxuICAgKiBwaXBlbGluZS5cbiAgICpcbiAgICogSWYgeW91IHN1cHBseSB0aGlzLCBkbyBub3Qgc3VwcGx5IGEgdmFsdWUgZm9yIGBmcm9tU2NyaXB0VXJsYC5cbiAgICogQHBhcmFtIG9ialxuICAgKi9cbiAgcHVibGljIGZyb21TY3JpcHRPYmplY3Qob2JqOiBhbnkpOiBUYXNrU3RlcEJ1aWxkZXIge1xuICAgIHRoaXMuX3NjcmlwdCA9IG5ldyBPYmpTY3JpcHRSZXNvbHZlcihvYmopO1xuICAgIHJldHVybiB0aGlzO1xuICB9XG5cbiAgLyoqXG4gICAqIElmIHN1cHBsaWVkLCB1c2VzIHRoZSBwcm92aWRlZCBzY3JpcHQgZGF0YSBhcy1pcyBmb3IgdGhlIHNjcmlwdCB2YWx1ZS5cbiAgICpcbiAgICogVXNlIHRoaXMgd2hlbiB5b3UgaGF2ZSB0aGUgc2NyaXB0IGRhdGEgZnJvbSBhIHNvdXJjZSBvdGhlciB0aGFuIGEgZmlsZSBvclxuICAgKiBhbiBvYmplY3QuIFVzZSB0aGUgb3RoZXIgbWV0aG9kcywgc3VjaCBhcyBgZnJvbVNjcmlwdFVybGAgKHdoZW4gdGhlIHNjcmlwdFxuICAgKiBpcyBpbiBhIGZpbGUpIG9yIGBzY3JpcHRGcm9tT2JqZWN0YCAod2hlbiB0aGUgc2NyaXB0IGlzIGEgQ0RLOHMgb2JqZWN0KVxuICAgKiByYXRoZXIgdGhhbiByZXNvbHZpbmcgdGhvc2UgeW91cnNlbGYuXG4gICAqXG4gICAqIEBwYXJhbSBkYXRhXG4gICAqL1xuICBwdWJsaWMgZnJvbVNjcmlwdERhdGEoZGF0YTogc3RyaW5nKTogVGFza1N0ZXBCdWlsZGVyIHtcbiAgICB0aGlzLl9zY3JpcHQgPSBuZXcgU3RhdGljU2NyaXB0UmVzb2x2ZXIoZGF0YSk7XG4gICAgcmV0dXJuIHRoaXM7XG4gIH1cblxuICAvKipcbiAgICogVGhlIGB3b3JraW5nRGlyYCBvZiB0aGUgYFRhc2tgLlxuICAgKiBAcGFyYW0gZGlyXG4gICAqL1xuICBwdWJsaWMgd2l0aFdvcmtpbmdEaXIoZGlyOiBzdHJpbmcpOiBUYXNrU3RlcEJ1aWxkZXIge1xuICAgIHRoaXMuX2RpciA9IGRpcjtcbiAgICByZXR1cm4gdGhpcztcbiAgfVxuXG4gIHdpdGhFbnYobmFtZTogc3RyaW5nLCB2YWx1ZUZyb206IFRhc2tFbnZWYWx1ZVNvdXJjZSk6IFRhc2tTdGVwQnVpbGRlciB7XG4gICAgaWYgKCF0aGlzLl9lbnYpIHtcbiAgICAgIHRoaXMuX2VudiA9IG5ldyBBcnJheTxUYXNrU3RlcEVudj4oKTtcbiAgICB9XG4gICAgdGhpcy5fZW52LnB1c2goe1xuICAgICAgbmFtZTogbmFtZSxcbiAgICAgIHZhbHVlRnJvbTogdmFsdWVGcm9tLFxuICAgIH0pO1xuICAgIHJldHVybiB0aGlzO1xuICB9XG5cbiAgcHVibGljIGJ1aWxkVGFza1N0ZXAoKTogVGFza1N0ZXAgfCB1bmRlZmluZWQge1xuICAgIGlmICh0aGlzLl9zY3JpcHQpIHtcbiAgICAgIHJldHVybiB7XG4gICAgICAgIG5hbWU6IHRoaXMubmFtZSxcbiAgICAgICAgaW1hZ2U6IHRoaXMuaW1hZ2UsXG4gICAgICAgIHNjcmlwdDogdGhpcy5zY3JpcHREYXRhLFxuICAgICAgICB3b3JraW5nRGlyOiB0aGlzLndvcmtpbmdEaXIsXG4gICAgICAgIGVudjogdGhpcy5fZW52LFxuICAgICAgfTtcbiAgICB9IGVsc2Uge1xuICAgICAgcmV0dXJuIHtcbiAgICAgICAgbmFtZTogdGhpcy5uYW1lLFxuICAgICAgICBpbWFnZTogdGhpcy5pbWFnZSxcbiAgICAgICAgY29tbWFuZDogdGhpcy5jb21tYW5kLFxuICAgICAgICBhcmdzOiB0aGlzLmFyZ3MsXG4gICAgICAgIHdvcmtpbmdEaXI6IHRoaXMud29ya2luZ0RpcixcbiAgICAgICAgZW52OiB0aGlzLl9lbnYsXG4gICAgICB9O1xuICAgIH1cbiAgfVxufVxuXG4vKipcbiAqIEJ1aWxkcyBUZWt0b24gYFRhc2tgIG9iamVjdHMgdGhhdCBhcmUgaW5kZXBlbmRlbnQgb2YgYSBgUGlwZWxpbmVgLlxuICpcbiAqIFRvIHVzZSBhIGJ1aWxkZXIgZm9yIHRhc2tzIHRoYXQgd2lsbCBiZSB1c2VkIGluIGEgUGlwZWxpbmUsIHVzZSB0aGVcbiAqIGBQaXBlbGluZUJ1aWxkZXJgIGluc3RlYWQuXG4gKi9cbmV4cG9ydCBjbGFzcyBUYXNrQnVpbGRlciB7XG5cbiAgcHJpdmF0ZSByZWFkb25seSBfc2NvcGU6IENvbnN0cnVjdDtcbiAgcHJpdmF0ZSByZWFkb25seSBfaWQ6IHN0cmluZztcbiAgcHJpdmF0ZSBfc3RlcHM/OiBUYXNrU3RlcEJ1aWxkZXJbXTtcbiAgcHJpdmF0ZSBfbmFtZT86IHN0cmluZztcbiAgcHJpdmF0ZSBfZGVzY3JpcHRpb24/OiBzdHJpbmc7XG4gIC8vIFRoZXNlIHdlcmUgaW5pdGlhbGx5IGFycmF5cywgYnV0IGNvbnZlcnRlZCB0aGVtIHRvIG1hcHMgc28gdGhhdCBpZlxuICAvLyBtdWx0aXBsZSB2YWx1ZXMgYXJlIGFkZGVkIHRoYXQgdGhlIGxhc3Qgb25lIHdpbGwgd2luLlxuICBwcml2YXRlIF93b3Jrc3BhY2VzID0gbmV3IE1hcDxzdHJpbmcsIFdvcmtzcGFjZUJ1aWxkZXI+O1xuICBwcml2YXRlIF9wYXJhbXMgPSBuZXcgTWFwPHN0cmluZywgUGFyYW1ldGVyQnVpbGRlcj47XG4gIHByaXZhdGUgX3Jlc3VsdHM6IFRhc2tTcGVjUmVzdWx0W10gPSBuZXcgQXJyYXk8VGFza1NwZWNSZXN1bHQ+O1xuICBwcml2YXRlIF9hbm5vdGF0aW9ucz86IHtcbiAgICBba2V5OiBzdHJpbmddOiBzdHJpbmc7XG4gIH07XG4gIHByaXZhdGUgX2xhYmVscz86IHtcbiAgICBba2V5OiBzdHJpbmddOiBzdHJpbmc7XG4gIH07XG5cbiAgLyoqXG4gICAqIENyZWF0ZXMgYSBuZXcgaW5zdGFuY2Ugb2YgdGhlIGBUYXNrQnVpbGRlcmAgdXNpbmcgdGhlIGdpdmVuIGBzY29wZWAgYW5kXG4gICAqIGBpZGAuXG4gICAqIEBwYXJhbSBzY29wZVxuICAgKiBAcGFyYW0gaWRcbiAgICovXG4gIHB1YmxpYyBjb25zdHJ1Y3RvcihzY29wZTogQ29uc3RydWN0LCBpZDogc3RyaW5nKSB7XG4gICAgdGhpcy5fc2NvcGUgPSBzY29wZTtcbiAgICB0aGlzLl9pZCA9IGlkO1xuICAgIC8vIFRoZXNlIGFyZSByZXF1aXJlZCwgYW5kIGl0J3MgYmV0dGVyIHRvIGp1c3QgY3JlYXRlIGl0IHJhdGhlciB0aGFuXG4gICAgLy8gY2hlY2sgZWFjaCB0aW1lLlxuICAgIHRoaXMuX3N0ZXBzID0gbmV3IEFycmF5PFRhc2tTdGVwQnVpbGRlcj4oKTtcbiAgfVxuXG4gIHB1YmxpYyBnZXQgbG9naWNhbElEKCk6IHN0cmluZyB7XG4gICAgcmV0dXJuIHRoaXMuX2lkO1xuICB9XG5cbiAgLyoqXG4gICAqIEFkZHMgYSBsYWJlbCB0byB0aGUgYFRhc2tgIHdpdGggdGhlIHByb3ZpZGVkIGxhYmVsIGtleSBhbmQgdmFsdWUuXG4gICAqXG4gICAqIEBzZWUgaHR0cHM6Ly9rdWJlcm5ldGVzLmlvL2RvY3MvY29uY2VwdHMvb3ZlcnZpZXcvd29ya2luZy13aXRoLW9iamVjdHMvbGFiZWxzL1xuICAgKlxuICAgKiBAcGFyYW0ga2V5XG4gICAqIEBwYXJhbSB2YWx1ZVxuICAgKi9cbiAgcHVibGljIHdpdGhMYWJlbChrZXk6IHN0cmluZywgdmFsdWU6IHN0cmluZyk6IFRhc2tCdWlsZGVyIHtcbiAgICBpZiAoISB0aGlzLl9sYWJlbHMpIHtcbiAgICAgIHRoaXMuX2xhYmVscyA9IHt9O1xuICAgIH1cbiAgICB0aGlzLl9sYWJlbHNba2V5XSA9IHZhbHVlO1xuICAgIHJldHVybiB0aGlzO1xuICB9XG5cbiAgLyoqXG4gICAqIEFkZHMgYW4gYW5ub3RhdGlvbiB0byB0aGUgYFRhc2tgIGBtZXRhZGF0YWAgd2l0aCB0aGUgcHJvdmlkZWQga2V5IGFuZCB2YWx1ZS5cbiAgICpcbiAgICogQHNlZSBodHRwczovL2t1YmVybmV0ZXMuaW8vZG9jcy9jb25jZXB0cy9vdmVydmlldy93b3JraW5nLXdpdGgtb2JqZWN0cy9hbm5vdGF0aW9ucy9cbiAgICpcbiAgICogQHBhcmFtIGtleSBUaGUgYW5ub3RhdGlvbidzIGtleS5cbiAgICogQHBhcmFtIHZhbHVlIFRoZSBhbm5vdGF0aW9uJ3MgdmFsdWUuXG4gICAqL1xuICBwdWJsaWMgd2l0aEFubm90YXRpb24oa2V5OiBzdHJpbmcsIHZhbHVlOiBzdHJpbmcpOiBUYXNrQnVpbGRlciB7XG4gICAgaWYgKCEgdGhpcy5fYW5ub3RhdGlvbnMpIHtcbiAgICAgIHRoaXMuX2Fubm90YXRpb25zID0ge307XG4gICAgfVxuICAgIHRoaXMuX2Fubm90YXRpb25zW2tleV0gPSB2YWx1ZTtcbiAgICByZXR1cm4gdGhpcztcbiAgfVxuXG4gIC8qKlxuICAgKiBHZXRzIHRoZSBuYW1lIG9mIHRoZSBgVGFza2AgYnVpbHQgYnkgdGhlIGBUYXNrQnVpbGRlcmAuXG4gICAqL1xuICBwdWJsaWMgZ2V0IG5hbWUoKTogc3RyaW5nIHwgdW5kZWZpbmVkIHtcbiAgICByZXR1cm4gdGhpcy5fbmFtZTtcbiAgfVxuXG4gIC8qKlxuICAgKiBTZXRzIHRoZSBuYW1lIG9mIHRoZSBgVGFza2AgYmVpbmcgYnVpbHQuXG4gICAqIEBwYXJhbSBuYW1lXG4gICAqL1xuICBwdWJsaWMgd2l0aE5hbWUobmFtZTogc3RyaW5nKTogVGFza0J1aWxkZXIge1xuICAgIHRoaXMuX25hbWUgPSBuYW1lO1xuICAgIHJldHVybiB0aGlzO1xuICB9XG5cbiAgLyoqXG4gICAqIFNldHMgdGhlIGBkZXNjcmlwdGlvbmAgb2YgdGhlIGBUYXNrYCBiZWluZyBidWlsdC5cbiAgICogQHBhcmFtIGRlc2NyaXB0aW9uXG4gICAqL1xuICBwdWJsaWMgd2l0aERlc2NyaXB0aW9uKGRlc2NyaXB0aW9uOiBzdHJpbmcpOiBUYXNrQnVpbGRlciB7XG4gICAgdGhpcy5fZGVzY3JpcHRpb24gPSBkZXNjcmlwdGlvbjtcbiAgICByZXR1cm4gdGhpcztcbiAgfVxuXG4gIC8qKlxuICAgKiBHZXRzIHRoZSBgZGVzY3JpcHRpb25gIG9mIHRoZSBgVGFza2AuXG4gICAqL1xuICBwdWJsaWMgZ2V0IGRlc2NyaXB0aW9uKCk6IHN0cmluZyB8IHVuZGVmaW5lZCB7XG4gICAgcmV0dXJuIHRoaXMuX2Rlc2NyaXB0aW9uO1xuICB9XG5cbiAgLyoqXG4gICAqIEFkZHMgdGhlIHNwZWNpZmllZCB3b3Jrc3BhY2UgdG8gdGhlIGBUYXNrYC5cbiAgICogQHBhcmFtIHdvcmtzcGFjZVxuICAgKi9cbiAgcHVibGljIHdpdGhXb3Jrc3BhY2Uod29ya3NwYWNlOiBXb3Jrc3BhY2VCdWlsZGVyKTogVGFza0J1aWxkZXIge1xuICAgIHRoaXMuX3dvcmtzcGFjZXMuc2V0KHdvcmtzcGFjZS5sb2dpY2FsSUQhLCB3b3Jrc3BhY2UpO1xuICAgIHJldHVybiB0aGlzO1xuICB9XG5cbiAgLyoqXG4gICAqIEdldHMgdGhlIHdvcmtzcGFjZXMgZm9yIHRoZSBgVGFza2AuXG4gICAqL1xuICBwdWJsaWMgZ2V0IHdvcmtzcGFjZXMoKTogV29ya3NwYWNlQnVpbGRlcltdIHwgdW5kZWZpbmVkIHtcbiAgICByZXR1cm4gQXJyYXkuZnJvbSh0aGlzLl93b3Jrc3BhY2VzPy52YWx1ZXMoKSk7XG4gIH1cblxuICAvKipcbiAgICogQWRkcyBhIHBhcmFtZXRlciBvZiB0eXBlIHN0cmluZyB0byB0aGUgYFRhc2tgLlxuICAgKlxuICAgKiBAcGFyYW0gcGFyYW1cbiAgICovXG4gIHB1YmxpYyB3aXRoU3RyaW5nUGFyYW0ocGFyYW06IFBhcmFtZXRlckJ1aWxkZXIpOiBUYXNrQnVpbGRlciB7XG4gICAgdGhpcy5fcGFyYW1zLnNldChwYXJhbS5sb2dpY2FsSUQhLCBwYXJhbS5vZlR5cGUoJ3N0cmluZycpKTtcbiAgICByZXR1cm4gdGhpcztcbiAgfVxuXG4gIHB1YmxpYyBnZXQgcGFyYW1ldGVycygpOiBQYXJhbWV0ZXJCdWlsZGVyW10gfCB1bmRlZmluZWQge1xuICAgIHJldHVybiBBcnJheS5mcm9tKHRoaXMuX3BhcmFtcz8udmFsdWVzKCkpO1xuICB9XG5cbiAgLyoqXG4gICAqIEFsbG93cyB5b3UgdG8gYWRkIGFuIHJlc3VsdCB0byB0aGUgVGFzay5cbiAgICpcbiAgICogQHNlZSBodHRwczovL3Rla3Rvbi5kZXYvZG9jcy9waXBlbGluZXMvdGFza3MvI2VtaXR0aW5nLXJlc3VsdHNcbiAgICpcbiAgICogQHBhcmFtIG5hbWUgVGhlIG5hbWUgb2YgdGhlIHJlc3VsdC5cbiAgICogQHBhcmFtIGRlc2NyaXB0aW9uIFRoZSByZXN1bHQncyBkZXNjcmlwdGlvbi5cbiAgICovXG4gIHB1YmxpYyB3aXRoUmVzdWx0KG5hbWU6IHN0cmluZywgZGVzY3JpcHRpb246IHN0cmluZyk6IFRhc2tCdWlsZGVyIHtcbiAgICAvLyBGaXJzdCwgY2hlY2sgdG8gc2VlIGlmIHRoZXJlIGlzIGFscmVhZHkgYSByZXN1bHQgd2l0aCB0aGlzIG5hbWVcbiAgICBjb25zdCBleGlzdGluZyA9IHRoaXMuX3Jlc3VsdHMuZmluZCgob2JqKSA9PiBvYmoubmFtZSA9PT0gbmFtZSk7XG4gICAgaWYgKGV4aXN0aW5nKSB7XG4gICAgICB0aHJvdyBuZXcgRXJyb3IoYENhbm5vdCBhZGQgcmVzdWx0ICR7bmFtZX0sIGFzIGl0IGFscmVhZHkgZXhpc3RzLmApO1xuICAgIH1cbiAgICB0aGlzLl9yZXN1bHRzLnB1c2goe1xuICAgICAgbmFtZTogbmFtZSxcbiAgICAgIGRlc2NyaXB0aW9uOiBkZXNjcmlwdGlvbixcbiAgICB9KTtcbiAgICByZXR1cm4gdGhpcztcbiAgfVxuXG4gIC8qKlxuICAgKiBBZGRzIHRoZSBnaXZlbiBgc3RlcGAgKGBUYXNrU3RlcEJ1aWxkZXJgKSB0byB0aGUgYFRhc2tgLlxuICAgKiBAcGFyYW0gc3RlcFxuICAgKi9cbiAgcHVibGljIHdpdGhTdGVwKHN0ZXA6IFRhc2tTdGVwQnVpbGRlcik6IFRhc2tCdWlsZGVyIHtcbiAgICB0aGlzLl9zdGVwcyEucHVzaChzdGVwKTtcbiAgICByZXR1cm4gdGhpcztcbiAgfVxuXG4gIC8qKlxuICAgKiBCdWlsZHMgdGhlIGBUYXNrYC5cbiAgICovXG4gIHB1YmxpYyBidWlsZFRhc2soKTogdm9pZCB7XG5cbiAgICBjb25zdCB0YXNrU3RlcHMgPSBuZXcgQXJyYXk8VGFza1N0ZXA+KCk7XG5cbiAgICB0aGlzLl9zdGVwcz8uZm9yRWFjaCgocykgPT4ge1xuICAgICAgY29uc3Qgc3RlcCA9IHMuYnVpbGRUYXNrU3RlcCgpO1xuICAgICAgaWYgKHN0ZXApIHtcbiAgICAgICAgdGFza1N0ZXBzLnB1c2goc3RlcCk7XG4gICAgICB9XG4gICAgfSk7XG5cbiAgICBjb25zdCB0YXNrUGFyYW1zID0gbmV3IEFycmF5PFRhc2tTcGVjUGFyYW0+KCk7XG4gICAgdGhpcy5fcGFyYW1zPy5mb3JFYWNoKChwKSA9PiB7XG4gICAgICB0YXNrUGFyYW1zLnB1c2goe1xuICAgICAgICBuYW1lOiBwLmxvZ2ljYWxJRCxcbiAgICAgICAgZGVzY3JpcHRpb246IHAuZGVzY3JpcHRpb24sXG4gICAgICAgIGRlZmF1bHQ6IHAuZGVmYXVsdFZhbHVlLFxuICAgICAgfSk7XG4gICAgfSk7XG5cbiAgICBjb25zdCB0YXNrV29ya3NwYWNlcyA9IG5ldyBBcnJheTxUYXNrV29ya3NwYWNlPigpO1xuICAgIHRoaXMuX3dvcmtzcGFjZXM/LmZvckVhY2goKHdzKSA9PiB7XG4gICAgICB0YXNrV29ya3NwYWNlcy5wdXNoKHtcbiAgICAgICAgbmFtZTogd3MubG9naWNhbElELFxuICAgICAgICBkZXNjcmlwdGlvbjogd3MuZGVzY3JpcHRpb24sXG4gICAgICB9KTtcbiAgICB9KTtcblxuICAgIGNvbnN0IHByb3BzOiBUYXNrUHJvcHMgPSB7XG4gICAgICBtZXRhZGF0YToge1xuICAgICAgICBuYW1lOiB0aGlzLm5hbWUsXG4gICAgICAgIGxhYmVsczogdGhpcy5fbGFiZWxzLFxuICAgICAgICBhbm5vdGF0aW9uczogdGhpcy5fYW5ub3RhdGlvbnMsXG4gICAgICB9LFxuICAgICAgc3BlYzoge1xuICAgICAgICBkZXNjcmlwdGlvbjogdGhpcy5kZXNjcmlwdGlvbixcbiAgICAgICAgd29ya3NwYWNlczogdGFza1dvcmtzcGFjZXMsXG4gICAgICAgIHBhcmFtczogdGFza1BhcmFtcyxcbiAgICAgICAgc3RlcHM6IHRhc2tTdGVwcyxcbiAgICAgICAgcmVzdWx0czogdGhpcy5fcmVzdWx0cyxcbiAgICAgIH0sXG4gICAgfTtcblxuICAgIG5ldyBUYXNrKHRoaXMuX3Njb3BlISwgdGhpcy5faWQhLCBwcm9wcyk7XG5cbiAgfVxufVxuXG4vKipcbiAqXG4gKi9cbmV4cG9ydCBjbGFzcyBQaXBlbGluZUJ1aWxkZXIge1xuICBwcml2YXRlIHJlYWRvbmx5IF9zY29wZTogQ29uc3RydWN0O1xuICBwcml2YXRlIHJlYWRvbmx5IF9pZDogc3RyaW5nO1xuICBwcml2YXRlIF9uYW1lPzogc3RyaW5nO1xuICBwcml2YXRlIF9kZXNjcmlwdGlvbj86IHN0cmluZztcbiAgcHJpdmF0ZSBfdGFza3M/OiBUYXNrQnVpbGRlcltdO1xuXG4gIHB1YmxpYyBjb25zdHJ1Y3RvcihzY29wZTogQ29uc3RydWN0LCBpZDogc3RyaW5nKSB7XG4gICAgdGhpcy5fc2NvcGUgPSBzY29wZTtcbiAgICB0aGlzLl9pZCA9IGlkO1xuICB9XG5cbiAgLyoqXG4gICAqIFByb3ZpZGVzIHRoZSBuYW1lIGZvciB0aGUgcGlwZWxpbmUgdGFzayBhbmQgd2lsbCBiZVxuICAgKiByZW5kZXJlZCBhcyB0aGUgYG5hbWVgIHByb3BlcnR5LlxuICAgKiBAcGFyYW0gbmFtZVxuICAgKi9cbiAgcHVibGljIHdpdGhOYW1lKG5hbWU6IHN0cmluZyk6IFBpcGVsaW5lQnVpbGRlciB7XG4gICAgdGhpcy5fbmFtZSA9IG5hbWU7XG4gICAgcmV0dXJuIHRoaXM7XG4gIH1cblxuICAvKipcbiAgICogR2V0cyB0aGUgbmFtZSBvZiB0aGUgcGlwZWxpbmVcbiAgICovXG4gIHB1YmxpYyBnZXQgbmFtZSgpOiBzdHJpbmcge1xuICAgIHJldHVybiB0aGlzLl9uYW1lIHx8IHRoaXMuX2lkO1xuICB9XG5cbiAgLyoqXG4gICAqIFByb3ZpZGVzIHRoZSBuYW1lIGZvciB0aGUgcGlwZWxpbmUgdGFzayBhbmQgd2lsbCBiZVxuICAgKiByZW5kZXJlZCBhcyB0aGUgYG5hbWVgIHByb3BlcnR5LlxuICAgKiBAcGFyYW0gZGVzY3JpcHRpb25cbiAgICovXG4gIHB1YmxpYyB3aXRoRGVzY3JpcHRpb24oZGVzY3JpcHRpb246IHN0cmluZyk6IFBpcGVsaW5lQnVpbGRlciB7XG4gICAgdGhpcy5fZGVzY3JpcHRpb24gPSBkZXNjcmlwdGlvbjtcbiAgICByZXR1cm4gdGhpcztcbiAgfVxuXG4gIC8vIEFkZHMgdGhlIHRhc2sgdG8gdGhlIHBpcGVsaW5lLlxuICBwdWJsaWMgd2l0aFRhc2sodGFza0I6IFRhc2tCdWlsZGVyKTogUGlwZWxpbmVCdWlsZGVyIHtcbiAgICAvLyBBZGQgdGhlIHRhc2sgdG8gdGhlIGxpc3Qgb2YgdGFza3MuLi5cbiAgICBpZiAoIXRoaXMuX3Rhc2tzKSB7XG4gICAgICB0aGlzLl90YXNrcyA9IG5ldyBBcnJheTxUYXNrQnVpbGRlcj4oKTtcbiAgICB9XG4gICAgdGhpcy5fdGFza3MucHVzaCh0YXNrQik7XG4gICAgcmV0dXJuIHRoaXM7XG4gIH1cblxuICAvKipcbiAgICogUmV0dXJucyB0aGUgYXJyYXkgb2YgYFBpcGVsaW5lUGFyYW1gIG9iamVjdHMgdGhhdCByZXByZXNlbnQgdGhlIHBhcmFtZXRlcnNcbiAgICogY29uZmlndXJlZCBmb3IgdGhlIGBQaXBlbGluZWAuXG4gICAqXG4gICAqIE5vdGUgdGhpcyBpcyBhbiBcImV4cGVuc2l2ZVwiIGdldCBiZWNhdXNlIGl0IGxvb3BzIHRocm91Z2ggdGhlIHRhc2tzIGluIHRoZVxuICAgKiBwaXBlbGluZSBhbmQgY2hlY2tzIGZvciBkdXBsaWNhdGVzIGluIHRoZSBwaXBlbGluZSBwYXJhbWV0ZXJzIGZvciBlYWNoIHRhc2tcbiAgICogcGFyYW1ldGVyIGZvdW5kLiBZb3Ugc2hvdWxkIGF2b2lkIGNhbGxpbmcgdGhpcyBpbiBhIGxvb3AtLWluc3RlYWQsIGRlY2xhcmVcbiAgICogYSBsb2NhbCB2YXJpYWJsZSBiZWZvcmUgdGhlIGxvb3AgYW5kIHJlZmVyZW5jZSB0aGF0IGluc3RlYWQuXG4gICAqXG4gICAqIEByZXR1cm5zIFBpcGVsaW5lUGFyYW1bXSBBbiBhcnJheSBvZiB0aGUgcGlwZWxpbmUgcGFyYW1ldGVycy5cbiAgICovXG4gIHB1YmxpYyBnZXQgcGFyYW1zKCk6IFBpcGVsaW5lUGFyYW1bXSB7XG4gICAgLy8gTm90IHRyeWluZyB0byBwcmVtYXR1cmVseSBvcHRpbWl6ZSBoZXJlLCBidXQgdGhpcyBjb3VsZCBiZSBhbiBleHBlbnNpdmVcbiAgICAvLyBvcGVyYXRpb24sIHNvIHdlIG9ubHkgbmVlZCB0byBkbyBpdCBpZiB0aGUgc3RhdGUgb2YgdGhlIG9iamVjdCBoYXMgbm90XG4gICAgLy8gY2hhbmdlZC5cbiAgICBjb25zdCBwaXBlbGluZVBhcmFtcyA9IG5ldyBNYXA8c3RyaW5nLCBQaXBlbGluZVBhcmFtPigpO1xuICAgIHRoaXMuX3Rhc2tzPy5mb3JFYWNoKCh0KSA9PiB7XG4gICAgICB0LnBhcmFtZXRlcnM/LmZvckVhY2gocCA9PiB7XG4gICAgICAgIGNvbnN0IHBwID0gcGlwZWxpbmVQYXJhbXMuZ2V0KHAubmFtZSEpO1xuICAgICAgICBpZiAoIXBwKSB7XG4gICAgICAgICAgLy8gRG8gbm90IGFkZCBpdCB0byB0aGUgcGlwZWxpbmUgaWYgdGhlcmUgaXMgbm8gbmVlZCB0byBhZGQgaXQuLi5cbiAgICAgICAgICBpZiAocC5yZXF1aXJlc1BpcGVsaW5lUGFyYW1ldGVyKSB7XG4gICAgICAgICAgICBwaXBlbGluZVBhcmFtcy5zZXQocC5uYW1lISwge1xuICAgICAgICAgICAgICBuYW1lOiBwLm5hbWUsXG4gICAgICAgICAgICAgIHR5cGU6IHAudHlwZSxcbiAgICAgICAgICAgIH0pO1xuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgfSk7XG4gICAgfSk7XG4gICAgcmV0dXJuIEFycmF5LmZyb20ocGlwZWxpbmVQYXJhbXMudmFsdWVzKCkpO1xuICB9XG5cbiAgLyoqXG4gICAqIFJldHVybnMgdGhlIGFycmF5IG9mIGBQaXBlbGluZVdvcmtzcGFjZWAgb2JqZWN0cyB0aGF0IHJlcHJlc2VudCB0aGUgd29ya3NwYWNlc1xuICAgKiBjb25maWd1cmVkIGZvciB0aGUgYFBpcGVsaW5lYC5cbiAgICpcbiAgICogVGhpcyBpcyBhbiBcImV4cGVuc2l2ZVwiIGdldCBiZWNhdXNlIGl0IGxvb3BzIHRocm91Z2ggdGhlIHdvcmtzcGFjZXMgaW4gdGhlXG4gICAqIHBpcGVsaW5lIGFuZCBjaGVja3MgZm9yIGR1cGxpY2F0ZXMgaW4gdGhlIHBpcGVsaW5lIHdvcmtzcGFjZXMgZm9yIGVhY2ggdGFza1xuICAgKiB3b3Jrc3BhY2UgZm91bmQuIFlvdSBzaG91bGQgYXZvaWQgY2FsbGluZyB0aGlzIGluIGEgbG9vcC0taW5zdGVhZCwgZGVjbGFyZVxuICAgKiBhIGxvY2FsIHZhcmlhYmxlIGJlZm9yZSB0aGUgbG9vcCBhbmQgcmVmZXJlbmNlIHRoYXQgaW5zdGVhZC5cbiAgICpcbiAgICogQHJldHVybnMgUGlwZWxpbmVXb3Jrc3BhY2VbXSBBbiBhcnJheSBvZiB0aGUgcGlwZWxpbmUgd29ya3NwYWNlcy5cbiAgICovXG4gIHB1YmxpYyBnZXQgd29ya3NwYWNlcygpOiBQaXBlbGluZVdvcmtzcGFjZVtdIHtcbiAgICBjb25zdCBwaXBlbGluZVdvcmtzcGFjZXMgPSBuZXcgTWFwPHN0cmluZywgUGlwZWxpbmVXb3Jrc3BhY2U+KCk7XG4gICAgdGhpcy5fdGFza3M/LmZvckVhY2goKHQpID0+IHtcbiAgICAgIHQud29ya3NwYWNlcz8uZm9yRWFjaCgodykgPT4ge1xuICAgICAgICAvLyBPbmx5IGFkZCB0aGUgd29ya3NwYWNlIG9uIHRoZSBwaXBlbGluZSBsZXZlbCBpZiBpdCBpcyBub3QgYWxyZWFkeVxuICAgICAgICAvLyB0aGVyZS4uLlxuICAgICAgICBjb25zdCB3cyA9IHBpcGVsaW5lV29ya3NwYWNlcy5nZXQody5uYW1lISk7XG4gICAgICAgIGlmICghd3MpIHtcbiAgICAgICAgICBwaXBlbGluZVdvcmtzcGFjZXMuc2V0KHcubmFtZSEsIHtcbiAgICAgICAgICAgIG5hbWU6IHcubmFtZSxcbiAgICAgICAgICAgIGRlc2NyaXB0aW9uOiB3LmRlc2NyaXB0aW9uLFxuICAgICAgICAgIH0pO1xuICAgICAgICB9XG4gICAgICB9KTtcbiAgICB9KTtcbiAgICByZXR1cm4gQXJyYXkuZnJvbShwaXBlbGluZVdvcmtzcGFjZXMudmFsdWVzKCkpO1xuICB9XG5cbiAgLyoqXG4gICAqIEJ1aWxkcyB0aGUgYWN0dWFsIFtQaXBlbGluZV0oaHR0cHM6Ly90ZWt0b24uZGV2L2RvY3MvZ2V0dGluZy1zdGFydGVkL3BpcGVsaW5lcy8pXG4gICAqIGZyb20gdGhlIHNldHRpbmdzIGNvbmZpZ3VyZWQgdXNpbmcgdGhlIGZsdWlkIHN5bnRheC5cbiAgICovXG4gIHB1YmxpYyBidWlsZFBpcGVsaW5lKG9wdHM6IEJ1aWxkZXJPcHRpb25zID0gRGVmYXVsdEJ1aWxkZXJPcHRpb25zKTogdm9pZCB7XG4gICAgLy8gVE9ETzogdmFsaWRhdGUgdGhlIG9iamVjdFxuXG4gICAgY29uc3QgcGlwZWxpbmVUYXNrczogUGlwZWxpbmVUYXNrW10gPSBuZXcgQXJyYXk8UGlwZWxpbmVUYXNrPigpO1xuICAgIC8vIEZvciBtYWtpbmcgYSBsaXN0IHRvIG1ha2Ugc3VyZSB0aGF0IHRhc2tzIGFyZW4ndCBkdXBsaWNhdGVkIHdoZW4gZG9pbmdcbiAgICAvLyB0aGUgYnVpbGQuIE5vdCB0aGF0IGl0IHJlYWxseSBodXJ0cyBhbnl0aGluZywgYnV0IGl0IG1ha2VzIHRoZSBtdWx0aWRvY1xuICAgIC8vIFlBTUwgZmlsZSBiaWdnZXIgYW5kIG1vcmUgY29tcGxleCB0aGFuIGl0IG5lZWRzIHRvIGJlLlxuICAgIGNvbnN0IHRhc2tMaXN0OiBzdHJpbmdbXSA9IG5ldyBBcnJheTxzdHJpbmc+KCk7XG5cbiAgICB0aGlzLl90YXNrcz8uZm9yRWFjaCgodCwgaSkgPT4ge1xuXG4gICAgICBjb25zdCB0YXNrUGFyYW1zOiBUYXNrUGFyYW1bXSA9IG5ldyBBcnJheTxUYXNrUGFyYW0+KCk7XG4gICAgICBjb25zdCB0YXNrV29ya3NwYWNlczogUGlwZWxpbmVUYXNrV29ya3NwYWNlW10gPSBuZXcgQXJyYXk8VGFza1dvcmtzcGFjZT4oKTtcblxuICAgICAgdC5wYXJhbWV0ZXJzPy5mb3JFYWNoKHAgPT4ge1xuICAgICAgICB0YXNrUGFyYW1zLnB1c2goe1xuICAgICAgICAgIG5hbWU6IHAubG9naWNhbElELFxuICAgICAgICAgIHZhbHVlOiBwLnZhbHVlLFxuICAgICAgICB9KTtcbiAgICAgIH0pO1xuXG4gICAgICB0LndvcmtzcGFjZXM/LmZvckVhY2goKHcpID0+IHtcbiAgICAgICAgdGFza1dvcmtzcGFjZXMucHVzaCh7XG4gICAgICAgICAgbmFtZTogdy5sb2dpY2FsSUQsXG4gICAgICAgICAgd29ya3NwYWNlOiB3Lm5hbWUsXG4gICAgICAgIH0pO1xuICAgICAgfSk7XG5cbiAgICAgIGNvbnN0IHB0ID0gY3JlYXRlT3JkZXJlZFBpcGVsaW5lVGFzayh0LCAoKGkgPiAwKSA/IHRoaXMuX3Rhc2tzIVtpIC0gMV0ubG9naWNhbElEIDogJycpLCB0YXNrUGFyYW1zLCB0YXNrV29ya3NwYWNlcyk7XG5cbiAgICAgIHBpcGVsaW5lVGFza3MucHVzaChwdCk7XG5cbiAgICAgIGlmIChvcHRzLmluY2x1ZGVEZXBlbmRlbmNpZXMpIHtcbiAgICAgICAgLy8gQnVpbGQgdGhlIHRhc2sgaWYgdGhlIHVzZXIgaGFzIGFza2VkIGZvciB0aGUgZGVwZW5kZW5jaWVzIHRvIGJlXG4gICAgICAgIC8vIGJ1aWx0IGFsb25nIHdpdGggdGhlIHBpcGVsaW5lLCBidXQgb25seSBpZiB3ZSBoYXZlbid0IGFscmVhZHlcbiAgICAgICAgLy8gYnVpbHQgdGhlIHRhc2sgeWV0LlxuICAgICAgICBpZiAoIXRhc2tMaXN0LmZpbmQoaXQgPT4ge1xuICAgICAgICAgIHJldHVybiBpdCA9PSB0Lm5hbWU7XG4gICAgICAgIH0pKSB7XG4gICAgICAgICAgdC5idWlsZFRhc2soKTtcbiAgICAgICAgfVxuICAgICAgICB0YXNrTGlzdC5wdXNoKHQubmFtZSEpO1xuICAgICAgfVxuICAgIH0pO1xuXG4gICAgbmV3IFBpcGVsaW5lKHRoaXMuX3Njb3BlISwgdGhpcy5faWQhLCB7XG4gICAgICBtZXRhZGF0YTpcbiAgICAgICAge1xuICAgICAgICAgIG5hbWU6IHRoaXMubmFtZSxcbiAgICAgICAgfSxcbiAgICAgIHNwZWM6IHtcbiAgICAgICAgZGVzY3JpcHRpb246IHRoaXMuX2Rlc2NyaXB0aW9uLFxuICAgICAgICBwYXJhbXM6IHRoaXMucGFyYW1zLFxuICAgICAgICB3b3Jrc3BhY2VzOiB0aGlzLndvcmtzcGFjZXMsXG4gICAgICAgIHRhc2tzOiBwaXBlbGluZVRhc2tzLFxuICAgICAgfSxcbiAgICB9KTtcbiAgfVxufVxuXG5mdW5jdGlvbiBjcmVhdGVPcmRlcmVkUGlwZWxpbmVUYXNrKHQ6IFRhc2tCdWlsZGVyLCBhZnRlcjogc3RyaW5nLCBwYXJhbXM6IFRhc2tQYXJhbVtdLCB3czogVGFza1dvcmtzcGFjZVtdKTogUGlwZWxpbmVUYXNrIHtcbiAgaWYgKGFmdGVyKSB7XG4gICAgcmV0dXJuIHtcbiAgICAgIG5hbWU6IHQubG9naWNhbElELFxuICAgICAgdGFza1JlZjoge1xuICAgICAgICBuYW1lOiB0Lm5hbWUsXG4gICAgICB9LFxuICAgICAgcnVuQWZ0ZXI6IFthZnRlcl0sXG4gICAgICBwYXJhbXM6IHBhcmFtcyxcbiAgICAgIHdvcmtzcGFjZXM6IHdzLFxuICAgIH07XG4gIH1cbiAgcmV0dXJuIHtcbiAgICBuYW1lOiB0LmxvZ2ljYWxJRCxcbiAgICB0YXNrUmVmOiB7XG4gICAgICBuYW1lOiB0Lm5hbWUsXG4gICAgfSxcbiAgICBwYXJhbXM6IHBhcmFtcyxcbiAgICB3b3Jrc3BhY2VzOiB3cyxcbiAgfTtcbn1cblxuLyoqXG4gKiBCdWlsZHMgYSBgUGlwZWxpbmVSdW5gIHVzaW5nIHRoZSBzdXBwbGllZCBjb25maWd1cmF0aW9uLlxuICpcbiAqIEBzZWUgaHR0cHM6Ly90ZWt0b24uZGV2L2RvY3MvcGlwZWxpbmVzL3BpcGVsaW5lcnVucy9cbiAqL1xuZXhwb3J0IGNsYXNzIFBpcGVsaW5lUnVuQnVpbGRlciB7XG4gIHByaXZhdGUgcmVhZG9ubHkgX3Njb3BlOiBDb25zdHJ1Y3Q7XG4gIHByaXZhdGUgcmVhZG9ubHkgX2lkOiBzdHJpbmc7XG4gIHByaXZhdGUgcmVhZG9ubHkgX3BpcGVsaW5lOiBQaXBlbGluZUJ1aWxkZXI7XG4gIHByaXZhdGUgcmVhZG9ubHkgX3J1blBhcmFtczogUGlwZWxpbmVSdW5QYXJhbVtdO1xuICBwcml2YXRlIHJlYWRvbmx5IF9ydW5Xb3Jrc3BhY2VzOiBQaXBlbGluZVJ1bldvcmtzcGFjZVtdO1xuICBwcml2YXRlIF9zYTogc3RyaW5nO1xuICBwcml2YXRlIF9jcmJQcm9wczogQXBpT2JqZWN0UHJvcHM7XG5cbiAgLyoqXG4gICAqIENyZWF0ZXMgYSBuZXcgaW5zdGFuY2Ugb2YgdGhlIGBQaXBlbGluZVJ1bkJ1aWxkZXJgIGZvciB0aGUgc3BlY2lmaWVkXG4gICAqIGBQaXBlbGluZWAgdGhhdCBpcyBidWlsdCBieSB0aGUgYFBpcGVsaW5lQnVpbGRlcmAgc3VwcGxpZWQgaGVyZS5cbiAgICpcbiAgICogQSBwaXBlbGluZSBydW4gaXMgY29uZmlndXJlZCBvbmx5IGZvciBhIHNwZWNpZmljIHBpcGVsaW5lLCBzbyBpdCBkaWQgbm90XG4gICAqIG1ha2UgYW55IHNlbnNlIGhlcmUgdG8gYWxsb3cgdGhlIHJ1biB0byBiZSBjcmVhdGVkIHdpdGhvdXQgdGhlIHBpcGVsaW5lXG4gICAqIHNwZWNpZmllZC5cbiAgICpcbiAgICogQHBhcmFtIHNjb3BlIFRoZSBgQ29uc3RydWN0YCBpbiB3aGljaCB0byBjcmVhdGUgdGhlIGBQaXBlbGluZVJ1bmAuXG4gICAqIEBwYXJhbSBpZCBUaGUgbG9naWNhbCBJRCBvZiB0aGUgYFBpcGVsaW5lUnVuYCBjb25zdHJ1Y3QuXG4gICAqIEBwYXJhbSBwaXBlbGluZSBUaGUgYFBpcGVsaW5lYCBmb3Igd2hpY2ggdG8gY3JlYXRlIHRoaXMgcnVuLCB1c2luZyB0aGUgYFBpcGVsaW5lQnVpbGRlcmAuXG4gICAqL1xuICBwdWJsaWMgY29uc3RydWN0b3Ioc2NvcGU6IENvbnN0cnVjdCwgaWQ6IHN0cmluZywgcGlwZWxpbmU6IFBpcGVsaW5lQnVpbGRlcikge1xuICAgIHRoaXMuX3Njb3BlID0gc2NvcGU7XG4gICAgdGhpcy5faWQgPSBpZDtcbiAgICB0aGlzLl9waXBlbGluZSA9IHBpcGVsaW5lO1xuICAgIHRoaXMuX3NhID0gRGVmYXVsdFBpcGVsaW5lU2VydmljZUFjY291bnROYW1lO1xuICAgIHRoaXMuX2NyYlByb3BzID0gRGVmYXVsdENsdXN0ZXJSb2xlQmluZGluZ1Byb3BzO1xuICAgIHRoaXMuX3J1blBhcmFtcyA9IG5ldyBBcnJheTxQaXBlbGluZVJ1blBhcmFtPigpO1xuICAgIHRoaXMuX3J1bldvcmtzcGFjZXMgPSBuZXcgQXJyYXk8UGlwZWxpbmVSdW5Xb3Jrc3BhY2U+KCk7XG4gIH1cblxuICAvKipcbiAgICogQWRkcyBhIHJ1biBwYXJhbWV0ZXIgdG8gdGhlIGBQaXBlbGluZVJ1bmAuIEl0IHdpbGwgdGhyb3cgYW4gZXJyb3IgaWYgeW91IHRyeVxuICAgKiB0byBhZGQgYSBwYXJhbWV0ZXIgdGhhdCBkb2VzIG5vdCBleGlzdCBvbiB0aGUgcGlwZWxpbmUuXG4gICAqXG4gICAqIEBwYXJhbSBuYW1lIFRoZSBuYW1lIG9mIHRoZSBwYXJhbWV0ZXIgYWRkZWQgdG8gdGhlIHBpcGVsaW5lIHJ1bi5cbiAgICogQHBhcmFtIHZhbHVlIFRoZSB2YWx1ZSBvZiB0aGUgcGFyYW1ldGVyIGFkZGVkIHRvIHRoZSBwaXBlbGluZSBydW4uXG4gICAqL1xuICBwdWJsaWMgd2l0aFJ1blBhcmFtKG5hbWU6IHN0cmluZywgdmFsdWU6IHN0cmluZyk6IFBpcGVsaW5lUnVuQnVpbGRlciB7XG4gICAgY29uc3QgcGFyYW1zID0gdGhpcy5fcGlwZWxpbmUucGFyYW1zO1xuICAgIGNvbnN0IHAgPSBwYXJhbXMuZmluZCgob2JqKSA9PiBvYmoubmFtZSA9PT0gbmFtZSk7XG4gICAgaWYgKHApIHtcbiAgICAgIHRoaXMuX3J1blBhcmFtcy5wdXNoKHtcbiAgICAgICAgbmFtZTogbmFtZSxcbiAgICAgICAgdmFsdWU6IHZhbHVlLFxuICAgICAgfSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHRocm93IG5ldyBFcnJvcihgUGlwZWxpbmVSdW4gcGFyYW1ldGVyICcke25hbWV9JyBkb2VzIG5vdCBleGlzdCBpbiBwaXBlbGluZSAnJHt0aGlzLl9waXBlbGluZS5uYW1lfSdgKTtcbiAgICB9XG4gICAgcmV0dXJuIHRoaXM7XG4gIH1cblxuICAvKipcbiAgICogQWxsb3dzIHlvdSB0byBzcGVjaWZ5IHRoZSBuYW1lIG9mIGEgYFBlcnNpc3RlbnRWb2x1bWVDbGFpbWAgYnV0IGRvZXMgbm90XG4gICAqIGRvIGFueSBjb21waWxlLXRpbWUgdmFsaWRhdGlvbiBvbiB0aGUgdm9sdW1lIGNsYWltJ3MgbmFtZSBvciBleGlzdGVuY2UuXG4gICAqXG4gICAqIEBzZWUgaHR0cHM6Ly9rdWJlcm5ldGVzLmlvL2RvY3MvdGFza3MvY29uZmlndXJlLXBvZC1jb250YWluZXIvY29uZmlndXJlLXBlcnNpc3RlbnQtdm9sdW1lLXN0b3JhZ2UvI2NyZWF0ZS1hLXBlcnNpc3RlbnR2b2x1bWVjbGFpbVxuICAgKlxuICAgKiBAcGFyYW0gbmFtZSBUaGUgbmFtZSBvZiB0aGUgd29ya3NwYWNlIGluIHRoZSBgUGlwZWxpbmVSdW5gIHRoYXQgd2lsbCBiZSB1c2VkIGJ5IHRoZSBgUGlwZWxpbmVgLlxuICAgKiBAcGFyYW0gY2xhaW1OYW1lIFRoZSBuYW1lIG9mIHRoZSBgUGVyc2lzdGVudFZvbHVtZUNsYWltYCB0byB1c2UgZm9yIHRoZSBgd29ya3NwYWNlYC5cbiAgICogQHBhcmFtIHN1YlBhdGggVGhlIHN1YiBwYXRoIG9uIHRoZSBgcGVyc2lzdGVudFZvbHVtZUNsYWltYCB0byB1c2UgZm9yIHRoZSBgd29ya3NwYWNlYC5cbiAgICovXG4gIHB1YmxpYyB3aXRoV29ya3NwYWNlKG5hbWU6IHN0cmluZywgY2xhaW1OYW1lOiBzdHJpbmcsIHN1YlBhdGg6IHN0cmluZyk6IFBpcGVsaW5lUnVuQnVpbGRlciB7XG4gICAgdGhpcy5fcnVuV29ya3NwYWNlcy5wdXNoKHtcbiAgICAgIG5hbWU6IG5hbWUsXG4gICAgICBwZXJzaXN0ZW50Vm9sdW1lQ2xhaW06IHtcbiAgICAgICAgY2xhaW1OYW1lOiBjbGFpbU5hbWUsXG4gICAgICB9LFxuICAgICAgc3ViUGF0aDogc3ViUGF0aCxcbiAgICB9KTtcbiAgICByZXR1cm4gdGhpcztcbiAgfVxuXG4gIHB1YmxpYyB3aXRoQ2x1c3RlclJvbGVCaW5kaW5nUHJvcHMocHJvcHM6IEFwaU9iamVjdFByb3BzKTogUGlwZWxpbmVSdW5CdWlsZGVyIHtcbiAgICB0aGlzLl9jcmJQcm9wcyA9IHByb3BzO1xuICAgIHJldHVybiB0aGlzO1xuICB9XG5cbiAgLyoqXG4gICAqIFVzZXMgdGhlIHByb3ZpZGVkIHJvbGUgbmFtZSBmb3IgdGhlIGBzZXJ2aWNlQWNjb3VudE5hbWVgIG9uIHRoZVxuICAgKiBgUGlwZWxpbmVSdW5gLiBJZiB0aGlzIG1ldGhvZCBpcyBub3QgY2FsbGVkIHByaW9yIHRvIGBidWlsZFBpcGVsaW5lUnVuKClgLFxuICAgKiB0aGVuIHRoZSBkZWZhdWx0IHNlcnZpY2UgYWNjb3VudCB3aWxsIGJlIHVzZWQsIHdoaWNoIGlzIF9kZWZhdWx0OnBpcGVsaW5lXy5cbiAgICpcbiAgICogQHBhcmFtIHNhIFRoZSBuYW1lIG9mIHRoZSBzZXJ2aWNlIGFjY291bnQgKGBzZXJ2aWNlQWNjb3VudE5hbWVgKSB0byB1c2UuXG4gICAqL1xuICBwdWJsaWMgd2l0aFNlcnZpY2VBY2NvdW50KHNhOiBzdHJpbmcpOiBQaXBlbGluZVJ1bkJ1aWxkZXIge1xuICAgIHRoaXMuX3NhID0gc2E7XG4gICAgcmV0dXJuIHRoaXM7XG4gIH1cblxuICAvKipcbiAgICogQnVpbGRzIHRoZSBgUGlwZWxpbmVSdW5gIGZvciB0aGUgY29uZmlndXJlZCBgUGlwZWxpbmVgIHVzZWQgaW4gdGhlIGNvbnN0cnVjdG9yLlxuICAgKiBAcGFyYW0gb3B0c1xuICAgKi9cbiAgcHVibGljIGJ1aWxkUGlwZWxpbmVSdW4ob3B0czogQnVpbGRlck9wdGlvbnMgPSBEZWZhdWx0QnVpbGRlck9wdGlvbnMpOiB2b2lkIHtcbiAgICBpZiAob3B0cyAmJiBvcHRzLmluY2x1ZGVEZXBlbmRlbmNpZXMpIHtcbiAgICAgIC8vIEdlbmVyYXRlIHRoZSBDbHVzdGVyUm9sZUJpbmRpbmcgZG9jdW1lbnQsIGlmIGNvbmZpZ3VyZWQgdG8gZG8gc28uXG4gICAgICBuZXcgQXBpT2JqZWN0KHRoaXMuX3Njb3BlLCB0aGlzLl9jcmJQcm9wcy5tZXRhZGF0YT8ubmFtZSEsIHRoaXMuX2NyYlByb3BzKTtcbiAgICB9XG5cbiAgICAvLyBUaHJvdyBhbiBlcnJvciBoZXJlIGlmIHRoZSBwYXJhbWV0ZXJzIGFyZSBub3QgZGVmaW5lZCB0aGF0IGFyZSByZXF1aXJlZFxuICAgIC8vIGJ5IHRoZSBQaXBlbGluZSwgYmVjYXVzZSB0aGVyZSBpcyByZWFsbHkgbm8gcG9pbnQgaW4gZ29pbmcgYW55IGZ1cnRoZXIuXG4gICAgY29uc3QgcGFyYW1zID0gdGhpcy5fcGlwZWxpbmUucGFyYW1zO1xuICAgIHBhcmFtcy5mb3JFYWNoKChwKSA9PiB7XG4gICAgICBjb25zdCBwcnAgPSB0aGlzLl9ydW5QYXJhbXMuZmluZCgob2JqKSA9PiBvYmoubmFtZSA9PSBwLm5hbWUpO1xuICAgICAgaWYgKCFwcnApIHtcbiAgICAgICAgdGhyb3cgbmV3IEVycm9yKGBQaXBlbGluZSBwYXJhbWV0ZXIgJyR7cC5uYW1lfScgaXMgbm90IGRlZmluZWQgaW4gUGlwZWxpbmVSdW4gJyR7dGhpcy5faWR9J2ApO1xuICAgICAgfVxuICAgIH0pO1xuXG4gICAgLy8gRG8gdGhlIHNhbWUgdGhpbmcgZm9yIHdvcmtzcGFjZXMuIENoZWNrIHRvIG1ha2Ugc3VyZSB0aGF0IHRoZSB3b3Jrc3BhY2VzXG4gICAgLy8gZXhwZWN0ZWQgYnkgdGhlIFBpcGVsaW5lIGFyZSBkZWZpbmVkIGluIHRoZSBQaXBlbGluZVJ1bi5cbiAgICBjb25zdCB3b3Jrc3BhY2VzOiBQaXBlbGluZVdvcmtzcGFjZVtdID0gdGhpcy5fcGlwZWxpbmUud29ya3NwYWNlcztcbiAgICB3b3Jrc3BhY2VzLmZvckVhY2goKHdzKSA9PiB7XG4gICAgICBjb25zdCBwd3MgPSB0aGlzLl9ydW5Xb3Jrc3BhY2VzLmZpbmQoKG9iaikgPT4gb2JqLm5hbWUgPT0gd3MubmFtZSk7XG4gICAgICBpZiAoISBwd3MpIHtcbiAgICAgICAgdGhyb3cgbmV3IEVycm9yKGBQaXBlbGluZSB3b3Jrc3BhY2UgJyR7d3MubmFtZX0nIGlzIG5vdCBkZWZpbmVkIGluIFBpcGVsaW5lUnVuICcke3RoaXMuX2lkfSdgKTtcbiAgICAgIH1cbiAgICB9KTtcblxuICAgIG5ldyBQaXBlbGluZVJ1bih0aGlzLl9zY29wZSwgdGhpcy5faWQsIHtcbiAgICAgIG1ldGFkYXRhOiB7XG4gICAgICAgIG5hbWU6IHRoaXMuX2lkLFxuICAgICAgfSxcbiAgICAgIHNlcnZpY2VBY2NvdW50TmFtZTogdGhpcy5fc2EsXG4gICAgICBzcGVjOiB7XG4gICAgICAgIHBpcGVsaW5lUmVmOiB7XG4gICAgICAgICAgbmFtZTogdGhpcy5fcGlwZWxpbmUubmFtZSxcbiAgICAgICAgfSxcbiAgICAgICAgcGFyYW1zOiB0aGlzLl9ydW5QYXJhbXMsXG4gICAgICAgIHdvcmtzcGFjZXM6IHRoaXMuX3J1bldvcmtzcGFjZXMsXG4gICAgICB9LFxuICAgIH0pO1xuICB9XG59Il19