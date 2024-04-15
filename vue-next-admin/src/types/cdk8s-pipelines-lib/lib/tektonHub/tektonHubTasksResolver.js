"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.TektonHubTask = void 0;
const cdk8s_1 = require("cdk8s");
const cdk8s_pipelines_1 = require("cdk8s-pipelines");
/**
 * This class handles turning a URL that points the YAML for the Tekton Hub Task into a PipelineTask
 */
class TektonHubTask extends cdk8s_pipelines_1.TaskBuilder {
    /**
     * Creates a new Instance of TektonHubTask with a URL that points to the Raw YAML for the task.
     * @link https://hub.tekton.dev/
     * @param scope
     * @param id
     * @param url string Url to the raw yaml for a Tekton Hub Task (i.e https://raw.githubusercontent.com/tektoncd/catalog/main/task/yq/0.4/yq.yaml)
     */
    constructor(scope, id, url) {
        super(scope, id);
        this.url = url;
        this.taskBuild = new cdk8s_pipelines_1.TaskBuilder(scope, id);
    }
    parseYAML() {
        const task = this.readYamlFromUrl();
        this.taskBuild.withName(task.metadata?.name);
        const workspaces = task.spec?.workspaces;
        if (workspaces !== undefined && workspaces?.length !== 0) {
            workspaces.forEach(workspace => {
                this.taskBuild.withWorkspace(new cdk8s_pipelines_1.WorkspaceBuilder(workspace.name)
                    .withName(workspace.name)
                    .withDescription(workspace.description));
            });
        }
        const params = task.spec?.params;
        if (params !== undefined && params.length !== 0) {
            params.forEach(param => {
                this.taskBuild.withStringParam(new cdk8s_pipelines_1.ParameterBuilder(param.name)
                    .withDescription(param.description)
                    .withDefaultValue(param.default)
                    .withPiplineParameter(param.name, param.default));
            });
        }
        const steps = task.spec?.steps;
        if (steps) {
            steps.forEach(step => {
                const sb = new cdk8s_pipelines_1.TaskStepBuilder()
                    .withName(step.name)
                    .fromScriptData(step.script)
                    .withWorkingDir(step.workingDir)
                    .withArgs(step.args)
                    .withImage(step.image);
                // step.env?.forEach(e => {
                //   sb.withEnv(e.name);
                // });
                this.taskBuild.withStep(sb);
            });
        }
        return true;
    }
    readYamlFromUrl() {
        try {
            // Parse the YAML content
            const parsedYaml = cdk8s_1.Yaml.load(this.url);
            return parsedYaml[0];
        }
        catch (error) {
            const errorMessage = error.message; // Type assertion
            throw new Error(`Error reading YAML from URL: ${errorMessage}`);
        }
    }
    /**
     * Returns an instance of PipelineTaskBuilder with the corresponding Tekton Hub Task Link.
     * @returns TaskBuilder
     */
    build() {
        this.parseYAML();
        return this.taskBuild;
    }
}
exports.TektonHubTask = TektonHubTask;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoidGVrdG9uSHViVGFza3NSZXNvbHZlci5qcyIsInNvdXJjZVJvb3QiOiIiLCJzb3VyY2VzIjpbIi4uLy4uL3NyYy90ZWt0b25IdWIvdGVrdG9uSHViVGFza3NSZXNvbHZlci50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiOzs7QUFBQSxpQ0FBNkI7QUFDN0IscURBQXlHO0FBR3pHOztHQUVHO0FBQ0gsTUFBYSxhQUFjLFNBQVEsNkJBQVc7SUFJNUM7Ozs7OztPQU1HO0lBQ0gsWUFBWSxLQUFnQixFQUFFLEVBQVUsRUFBRSxHQUFXO1FBQ25ELEtBQUssQ0FBQyxLQUFLLEVBQUUsRUFBRSxDQUFDLENBQUM7UUFDakIsSUFBSSxDQUFDLEdBQUcsR0FBRyxHQUFHLENBQUM7UUFDZixJQUFJLENBQUMsU0FBUyxHQUFHLElBQUksNkJBQVcsQ0FBQyxLQUFLLEVBQUUsRUFBRSxDQUFDLENBQUM7SUFDOUMsQ0FBQztJQUVPLFNBQVM7UUFDZixNQUFNLElBQUksR0FBRyxJQUFJLENBQUMsZUFBZSxFQUFFLENBQUM7UUFDcEMsSUFBSSxDQUFDLFNBQVMsQ0FBQyxRQUFRLENBQUMsSUFBSSxDQUFDLFFBQVEsRUFBRSxJQUFLLENBQUMsQ0FBQztRQUM5QyxNQUFNLFVBQVUsR0FBRyxJQUFJLENBQUMsSUFBSSxFQUFFLFVBQVUsQ0FBQztRQUN6QyxJQUFJLFVBQVUsS0FBSyxTQUFTLElBQUksVUFBVSxFQUFFLE1BQU0sS0FBSyxDQUFDLEVBQUU7WUFDeEQsVUFBVSxDQUFDLE9BQU8sQ0FBQyxTQUFTLENBQUMsRUFBRTtnQkFDN0IsSUFBSSxDQUFDLFNBQVMsQ0FBQyxhQUFhLENBQUMsSUFBSSxrQ0FBZ0IsQ0FBQyxTQUFTLENBQUMsSUFBSyxDQUFDO3FCQUMvRCxRQUFRLENBQUMsU0FBUyxDQUFDLElBQUssQ0FBQztxQkFDekIsZUFBZSxDQUFDLFNBQVMsQ0FBQyxXQUFZLENBQUMsQ0FBQyxDQUFDO1lBQzlDLENBQUMsQ0FBQyxDQUFDO1NBQ0o7UUFDRCxNQUFNLE1BQU0sR0FBRyxJQUFJLENBQUMsSUFBSSxFQUFFLE1BQU0sQ0FBQztRQUNqQyxJQUFJLE1BQU0sS0FBSyxTQUFTLElBQUksTUFBTSxDQUFDLE1BQU0sS0FBSyxDQUFDLEVBQUU7WUFDL0MsTUFBTSxDQUFDLE9BQU8sQ0FBQyxLQUFLLENBQUMsRUFBRTtnQkFDckIsSUFBSSxDQUFDLFNBQVMsQ0FBQyxlQUFlLENBQUMsSUFBSSxrQ0FBZ0IsQ0FBQyxLQUFLLENBQUMsSUFBSyxDQUFDO3FCQUM3RCxlQUFlLENBQUMsS0FBSyxDQUFDLFdBQVksQ0FBQztxQkFDbkMsZ0JBQWdCLENBQUMsS0FBSyxDQUFDLE9BQVEsQ0FBQztxQkFDaEMsb0JBQW9CLENBQUMsS0FBSyxDQUFDLElBQUssRUFBRSxLQUFLLENBQUMsT0FBUSxDQUFDLENBQUMsQ0FBQztZQUN4RCxDQUFDLENBQUMsQ0FBQztTQUNKO1FBRUQsTUFBTSxLQUFLLEdBQUcsSUFBSSxDQUFDLElBQUksRUFBRSxLQUFLLENBQUM7UUFDL0IsSUFBSSxLQUFLLEVBQUU7WUFDVCxLQUFLLENBQUMsT0FBTyxDQUFDLElBQUksQ0FBQyxFQUFFO2dCQUNuQixNQUFNLEVBQUUsR0FDTixJQUFJLGlDQUFlLEVBQUU7cUJBQ2xCLFFBQVEsQ0FBQyxJQUFJLENBQUMsSUFBSyxDQUFDO3FCQUNwQixjQUFjLENBQUMsSUFBSSxDQUFDLE1BQU8sQ0FBQztxQkFDNUIsY0FBYyxDQUFDLElBQUksQ0FBQyxVQUFXLENBQUM7cUJBQ2hDLFFBQVEsQ0FBQyxJQUFJLENBQUMsSUFBSyxDQUFDO3FCQUNwQixTQUFTLENBQUMsSUFBSSxDQUFDLEtBQU0sQ0FBQyxDQUFDO2dCQUM1QiwyQkFBMkI7Z0JBQzNCLHdCQUF3QjtnQkFDeEIsTUFBTTtnQkFDTixJQUFJLENBQUMsU0FBUyxDQUFDLFFBQVEsQ0FBQyxFQUFFLENBQUMsQ0FBQztZQUM5QixDQUFDLENBQUMsQ0FBQztTQUNKO1FBRUQsT0FBTyxJQUFJLENBQUM7SUFDZCxDQUFDO0lBRU8sZUFBZTtRQUNyQixJQUFJO1lBQ0YseUJBQXlCO1lBQ3pCLE1BQU0sVUFBVSxHQUFHLFlBQUksQ0FBQyxJQUFJLENBQUMsSUFBSSxDQUFDLEdBQUcsQ0FBQyxDQUFDO1lBQ3ZDLE9BQU8sVUFBVSxDQUFDLENBQUMsQ0FBQyxDQUFDO1NBQ3RCO1FBQUMsT0FBTyxLQUFLLEVBQUU7WUFDZCxNQUFNLFlBQVksR0FBSSxLQUFlLENBQUMsT0FBTyxDQUFDLENBQUMsaUJBQWlCO1lBQ2hFLE1BQU0sSUFBSSxLQUFLLENBQUMsZ0NBQWdDLFlBQVksRUFBRSxDQUFDLENBQUM7U0FDakU7SUFDSCxDQUFDO0lBRUQ7OztPQUdHO0lBQ0ksS0FBSztRQUNWLElBQUksQ0FBQyxTQUFTLEVBQUUsQ0FBQztRQUNqQixPQUFPLElBQUksQ0FBQyxTQUFTLENBQUM7SUFDeEIsQ0FBQztDQUNGO0FBN0VELHNDQTZFQyIsInNvdXJjZXNDb250ZW50IjpbImltcG9ydCB7IFlhbWwgfSBmcm9tICdjZGs4cyc7XG5pbXBvcnQgeyBQYXJhbWV0ZXJCdWlsZGVyLCBUYXNrLCBUYXNrQnVpbGRlciwgVGFza1N0ZXBCdWlsZGVyLCBXb3Jrc3BhY2VCdWlsZGVyIH0gZnJvbSAnY2RrOHMtcGlwZWxpbmVzJztcbmltcG9ydCB7IENvbnN0cnVjdCB9IGZyb20gJ2NvbnN0cnVjdHMnO1xuXG4vKipcbiAqIFRoaXMgY2xhc3MgaGFuZGxlcyB0dXJuaW5nIGEgVVJMIHRoYXQgcG9pbnRzIHRoZSBZQU1MIGZvciB0aGUgVGVrdG9uIEh1YiBUYXNrIGludG8gYSBQaXBlbGluZVRhc2tcbiAqL1xuZXhwb3J0IGNsYXNzIFRla3Rvbkh1YlRhc2sgZXh0ZW5kcyBUYXNrQnVpbGRlciB7XG4gIHVybDogc3RyaW5nO1xuICB0YXNrQnVpbGQ6IFRhc2tCdWlsZGVyO1xuXG4gIC8qKlxuICAgKiBDcmVhdGVzIGEgbmV3IEluc3RhbmNlIG9mIFRla3Rvbkh1YlRhc2sgd2l0aCBhIFVSTCB0aGF0IHBvaW50cyB0byB0aGUgUmF3IFlBTUwgZm9yIHRoZSB0YXNrLlxuICAgKiBAbGluayBodHRwczovL2h1Yi50ZWt0b24uZGV2L1xuICAgKiBAcGFyYW0gc2NvcGVcbiAgICogQHBhcmFtIGlkXG4gICAqIEBwYXJhbSB1cmwgc3RyaW5nIFVybCB0byB0aGUgcmF3IHlhbWwgZm9yIGEgVGVrdG9uIEh1YiBUYXNrIChpLmUgaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL3Rla3RvbmNkL2NhdGFsb2cvbWFpbi90YXNrL3lxLzAuNC95cS55YW1sKVxuICAgKi9cbiAgY29uc3RydWN0b3Ioc2NvcGU6IENvbnN0cnVjdCwgaWQ6IHN0cmluZywgdXJsOiBzdHJpbmcpIHtcbiAgICBzdXBlcihzY29wZSwgaWQpO1xuICAgIHRoaXMudXJsID0gdXJsO1xuICAgIHRoaXMudGFza0J1aWxkID0gbmV3IFRhc2tCdWlsZGVyKHNjb3BlLCBpZCk7XG4gIH1cblxuICBwcml2YXRlIHBhcnNlWUFNTCgpOiBCb29sZWFuIHtcbiAgICBjb25zdCB0YXNrID0gdGhpcy5yZWFkWWFtbEZyb21VcmwoKTtcbiAgICB0aGlzLnRhc2tCdWlsZC53aXRoTmFtZSh0YXNrLm1ldGFkYXRhPy5uYW1lISk7XG4gICAgY29uc3Qgd29ya3NwYWNlcyA9IHRhc2suc3BlYz8ud29ya3NwYWNlcztcbiAgICBpZiAod29ya3NwYWNlcyAhPT0gdW5kZWZpbmVkICYmIHdvcmtzcGFjZXM/Lmxlbmd0aCAhPT0gMCkge1xuICAgICAgd29ya3NwYWNlcy5mb3JFYWNoKHdvcmtzcGFjZSA9PiB7XG4gICAgICAgIHRoaXMudGFza0J1aWxkLndpdGhXb3Jrc3BhY2UobmV3IFdvcmtzcGFjZUJ1aWxkZXIod29ya3NwYWNlLm5hbWUhKVxuICAgICAgICAgIC53aXRoTmFtZSh3b3Jrc3BhY2UubmFtZSEpXG4gICAgICAgICAgLndpdGhEZXNjcmlwdGlvbih3b3Jrc3BhY2UuZGVzY3JpcHRpb24hKSk7XG4gICAgICB9KTtcbiAgICB9XG4gICAgY29uc3QgcGFyYW1zID0gdGFzay5zcGVjPy5wYXJhbXM7XG4gICAgaWYgKHBhcmFtcyAhPT0gdW5kZWZpbmVkICYmIHBhcmFtcy5sZW5ndGggIT09IDApIHtcbiAgICAgIHBhcmFtcy5mb3JFYWNoKHBhcmFtID0+IHtcbiAgICAgICAgdGhpcy50YXNrQnVpbGQud2l0aFN0cmluZ1BhcmFtKG5ldyBQYXJhbWV0ZXJCdWlsZGVyKHBhcmFtLm5hbWUhKVxuICAgICAgICAgIC53aXRoRGVzY3JpcHRpb24ocGFyYW0uZGVzY3JpcHRpb24hKVxuICAgICAgICAgIC53aXRoRGVmYXVsdFZhbHVlKHBhcmFtLmRlZmF1bHQhKVxuICAgICAgICAgIC53aXRoUGlwbGluZVBhcmFtZXRlcihwYXJhbS5uYW1lISwgcGFyYW0uZGVmYXVsdCEpKTtcbiAgICAgIH0pO1xuICAgIH1cblxuICAgIGNvbnN0IHN0ZXBzID0gdGFzay5zcGVjPy5zdGVwcztcbiAgICBpZiAoc3RlcHMpIHtcbiAgICAgIHN0ZXBzLmZvckVhY2goc3RlcCA9PiB7XG4gICAgICAgIGNvbnN0IHNiID1cbiAgICAgICAgICBuZXcgVGFza1N0ZXBCdWlsZGVyKClcbiAgICAgICAgICAgIC53aXRoTmFtZShzdGVwLm5hbWUhKVxuICAgICAgICAgICAgLmZyb21TY3JpcHREYXRhKHN0ZXAuc2NyaXB0ISlcbiAgICAgICAgICAgIC53aXRoV29ya2luZ0RpcihzdGVwLndvcmtpbmdEaXIhKVxuICAgICAgICAgICAgLndpdGhBcmdzKHN0ZXAuYXJncyEpXG4gICAgICAgICAgICAud2l0aEltYWdlKHN0ZXAuaW1hZ2UhKTtcbiAgICAgICAgLy8gc3RlcC5lbnY/LmZvckVhY2goZSA9PiB7XG4gICAgICAgIC8vICAgc2Iud2l0aEVudihlLm5hbWUpO1xuICAgICAgICAvLyB9KTtcbiAgICAgICAgdGhpcy50YXNrQnVpbGQud2l0aFN0ZXAoc2IpO1xuICAgICAgfSk7XG4gICAgfVxuXG4gICAgcmV0dXJuIHRydWU7XG4gIH1cblxuICBwcml2YXRlIHJlYWRZYW1sRnJvbVVybCgpOiBUYXNrIHtcbiAgICB0cnkge1xuICAgICAgLy8gUGFyc2UgdGhlIFlBTUwgY29udGVudFxuICAgICAgY29uc3QgcGFyc2VkWWFtbCA9IFlhbWwubG9hZCh0aGlzLnVybCk7XG4gICAgICByZXR1cm4gcGFyc2VkWWFtbFswXTtcbiAgICB9IGNhdGNoIChlcnJvcikge1xuICAgICAgY29uc3QgZXJyb3JNZXNzYWdlID0gKGVycm9yIGFzIEVycm9yKS5tZXNzYWdlOyAvLyBUeXBlIGFzc2VydGlvblxuICAgICAgdGhyb3cgbmV3IEVycm9yKGBFcnJvciByZWFkaW5nIFlBTUwgZnJvbSBVUkw6ICR7ZXJyb3JNZXNzYWdlfWApO1xuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBSZXR1cm5zIGFuIGluc3RhbmNlIG9mIFBpcGVsaW5lVGFza0J1aWxkZXIgd2l0aCB0aGUgY29ycmVzcG9uZGluZyBUZWt0b24gSHViIFRhc2sgTGluay5cbiAgICogQHJldHVybnMgVGFza0J1aWxkZXJcbiAgICovXG4gIHB1YmxpYyBidWlsZCgpOiBUYXNrQnVpbGRlciB7XG4gICAgdGhpcy5wYXJzZVlBTUwoKTtcbiAgICByZXR1cm4gdGhpcy50YXNrQnVpbGQ7XG4gIH1cbn0iXX0=