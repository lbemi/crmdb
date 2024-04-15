"use strict";
var _a, _b;
Object.defineProperty(exports, "__esModule", { value: true });
exports.AWSCDKPipeline = exports.AWSCDKPipelineChart = void 0;
const JSII_RTTI_SYMBOL_1 = Symbol.for("jsii.rtti");
const node_util_1 = require("node:util");
const cdk8s_1 = require("cdk8s");
const cdk8s_pipelines_1 = require("cdk8s-pipelines");
const octokit_1 = require("octokit");
/**
 * The chart for creating a Tekton Pipeline that will use an AWS CDK project
 * to create resources in AWS for re-usable artifacts.
 */
class AWSCDKPipelineChart extends cdk8s_1.Chart {
    /**
     * Initializes an instance of the AWSCDKPipelineChart.
     *
     * @param scope
     * @param id
     * @param props
     */
    constructor(scope, id, props = {}) {
        super(scope, id, props);
        // Create the pipeline that will run in the OpenShift or K8s cluster.
        // const  = new Pipeline(this, 'aws-cdk-pipeline', {
        //   name: 'aws-cdk-pipeline',
        // });
        // props.params?.forEach(s => {
        //   pipeline.addStringParam(s);
        // });
        const pipeline = new cdk8s_pipelines_1.PipelineBuilder(this, 'aws-cdk-pipeline')
            .withName('deploy-cdk-project')
            .withDescription('A pipeline for deploying a AWS CDK project from a GitHub repository to a cluster.')
            .withTask(new cdk8s_pipelines_1.TaskBuilder(this, 'git-clone')
            .withName('fetch-project')
            // .withWorkspace('ssh-creds', 'ssh-credentials', 'The location of the SSH keys and credentials'),
            .withWorkspace(new cdk8s_pipelines_1.WorkspaceBuilder('output')
            .withName('shared-data')
            .withDescription('The AWS CDK project files.'))
            .withWorkspace(new cdk8s_pipelines_1.WorkspaceBuilder('ssh-creds')
            .withName('ssh-credentials')
            .withDescription('The location of the SSH keys and credentials')));
        // Now build out the
        const awsCdkTask = new cdk8s_pipelines_1.TaskBuilder(this, 'aws-cdk-synth')
            .withName('synth-cdk-pipeline')
            .withWorkspace(new cdk8s_pipelines_1.WorkspaceBuilder('projectdata')
            .withName('shared-data')
            .withDescription('The AWS CDK project files'));
        props.params?.forEach((s) => {
            awsCdkTask.withStringParam(new cdk8s_pipelines_1.ParameterBuilder(s)
                .withPiplineParameter(s));
        });
        pipeline.withTask(awsCdkTask);
        pipeline.buildPipeline();
    }
}
exports.AWSCDKPipelineChart = AWSCDKPipelineChart;
_a = JSII_RTTI_SYMBOL_1;
AWSCDKPipelineChart[_a] = { fqn: "cdk8s-pipelines-lib.AWSCDKPipelineChart", version: "0.0.12" };
/**
 * Creator for the AWSCDKPipelineChart
 */
class AWSCDKPipeline {
    /**
     * Generates the AWS CDK Pipeline (AWSCDKPipelineChart) based on the actual project
     * located in GitHub and specified by the configuration.
     * @param config
     */
    static createFrom(config) {
        const octokit = new octokit_1.Octokit({
            auth: config.token,
            baseUrl: config.ghUrl,
        });
        octokit.rest.repos.getReleaseByTag({
            owner: config.owner,
            repo: config.repo,
            tag: config.release,
        }).then(function (releaseResponse) {
            const releaseId = releaseResponse.data.id;
            octokit.rest.repos.listReleaseAssets({
                owner: config.owner,
                repo: config.repo,
                release_id: releaseId,
            }).then(function (templateResponse) {
                const asset = templateResponse.data.find(a => a.name == `${config.stackName}.template.json`);
                const assetId = asset?.id;
                // Now that I have my asset ID, I can download the asset...
                octokit.rest.repos.getReleaseAsset({
                    owner: config.owner,
                    repo: config.repo,
                    asset_id: Number(assetId),
                    headers: {
                        accept: 'application/octet-stream',
                    },
                }).then(function (assetResponse) {
                    const template = JSON.parse(new node_util_1.TextDecoder().decode(assetResponse.data));
                    const app = new cdk8s_1.App();
                    const templateParams = new Array();
                    // Now, we are going to automagically add the AWS-required parameters before adding the
                    // template parameters.
                    templateParams.push('AwsAccountId');
                    templateParams.push('AwsAccessKeyId');
                    templateParams.push('AwsSecretKeyId');
                    templateParams.push('AwsRegion');
                    Object.keys(template.Parameters).forEach(key => {
                        templateParams.push(key);
                    });
                    new AWSCDKPipelineChart(app, 'example-cdk8s-pipeline', {
                        params: templateParams,
                    });
                    app.synth();
                }).catch(function (assetErr) {
                    console.error(assetErr);
                });
            }).catch(function (templateErr) {
                console.error(templateErr);
            });
        }).catch(function (releaseErr) {
            console.error(releaseErr);
        });
    }
}
exports.AWSCDKPipeline = AWSCDKPipeline;
_b = JSII_RTTI_SYMBOL_1;
AWSCDKPipeline[_b] = { fqn: "cdk8s-pipelines-lib.AWSCDKPipeline", version: "0.0.12" };
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiYXdzY2RrcGlwZWxpbmUuanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyIuLi9zcmMvYXdzY2RrcGlwZWxpbmUudHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6Ijs7Ozs7QUFBQSx5Q0FBd0M7QUFDeEMsaUNBQStDO0FBQy9DLHFEQUFtRztBQUVuRyxxQ0FBa0M7QUFTbEM7OztHQUdHO0FBQ0gsTUFBYSxtQkFBb0IsU0FBUSxhQUFLO0lBQzVDOzs7Ozs7T0FNRztJQUNILFlBQVksS0FBZ0IsRUFBRSxFQUFVLEVBQUUsUUFBa0MsRUFBRTtRQUM1RSxLQUFLLENBQUMsS0FBSyxFQUFFLEVBQUUsRUFBRSxLQUFLLENBQUMsQ0FBQztRQUN4QixxRUFBcUU7UUFDckUsb0RBQW9EO1FBQ3BELDhCQUE4QjtRQUM5QixNQUFNO1FBQ04sK0JBQStCO1FBQy9CLGdDQUFnQztRQUNoQyxNQUFNO1FBQ04sTUFBTSxRQUFRLEdBQUcsSUFBSSxpQ0FBZSxDQUFDLElBQUksRUFBRSxrQkFBa0IsQ0FBQzthQUMzRCxRQUFRLENBQUMsb0JBQW9CLENBQUM7YUFDOUIsZUFBZSxDQUFDLG1GQUFtRixDQUFDO2FBQ3BHLFFBQVEsQ0FBQyxJQUFJLDZCQUFXLENBQUMsSUFBSSxFQUFFLFdBQVcsQ0FBQzthQUN6QyxRQUFRLENBQUMsZUFBZSxDQUFDO1lBQzFCLGtHQUFrRzthQUNqRyxhQUFhLENBQUMsSUFBSSxrQ0FBZ0IsQ0FBQyxRQUFRLENBQUM7YUFDMUMsUUFBUSxDQUFDLGFBQWEsQ0FBQzthQUN2QixlQUFlLENBQUMsNEJBQTRCLENBQUMsQ0FBQzthQUNoRCxhQUFhLENBQUMsSUFBSSxrQ0FBZ0IsQ0FBQyxXQUFXLENBQUM7YUFDN0MsUUFBUSxDQUFDLGlCQUFpQixDQUFDO2FBQzNCLGVBQWUsQ0FBQyw4Q0FBOEMsQ0FBQyxDQUFDLENBQUMsQ0FBQztRQUV6RSxvQkFBb0I7UUFDcEIsTUFBTSxVQUFVLEdBQUcsSUFBSSw2QkFBVyxDQUFDLElBQUksRUFBRSxlQUFlLENBQUM7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QUFNLENBQUMsSUFBSSxDQUFDLFFBQVEsQ0FBQyxVQUFVLENBQUMsQ0FBQyxPQUFPLENBQUMsR0FBRyxDQUFDLEVBQUU7d0JBQzdDLGNBQWMsQ0FBQyxJQUFJLENBQUMsR0FBRyxDQUFDLENBQUM7b0JBQzNCLENBQUMsQ0FBQyxDQUFDO29CQUNILElBQUksbUJBQW1CLENBQUMsR0FBRyxFQUFFLHdCQUF3QixFQUFFO3dCQUNyRCxNQUFNLEVBQUUsY0FBYztxQkFDdkIsQ0FBQyxDQUFDO29CQUNILEdBQUcsQ0FBQyxLQUFLLEVBQUUsQ0FBQztnQkFDZCxDQUFDLENBQUMsQ0FBQyxLQUFLLENBQUMsVUFBVSxRQUFRO29CQUN6QixPQUFPLENBQUMsS0FBSyxDQUFDLFFBQVEsQ0FBQyxDQUFDO2dCQUMxQixDQUFDLENBQUMsQ0FBQztZQUNMLENBQUMsQ0FBQyxDQUFDLEtBQUssQ0FBQyxVQUFVLFdBQVc7Z0JBQzVCLE9BQU8sQ0FBQyxLQUFLLENBQUMsV0FBVyxDQUFDLENBQUM7WUFDN0IsQ0FBQyxDQUFDLENBQUM7UUFDTCxDQUFDLENBQUMsQ0FBQyxLQUFLLENBQUMsVUFBVSxVQUFVO1lBQzNCLE9BQU8sQ0FBQyxLQUFLLENBQUMsVUFBVSxDQUFDLENBQUM7UUFDNUIsQ0FBQyxDQUFDLENBQUM7SUFDTCxDQUFDOztBQTVESCx3Q0E2REMiLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgeyBUZXh0RGVjb2RlciB9IGZyb20gJ25vZGU6dXRpbCc7XG5pbXBvcnQgeyBBcHAsIENoYXJ0LCBDaGFydFByb3BzIH0gZnJvbSAnY2RrOHMnO1xuaW1wb3J0IHsgUGFyYW1ldGVyQnVpbGRlciwgUGlwZWxpbmVCdWlsZGVyLCBUYXNrQnVpbGRlciwgV29ya3NwYWNlQnVpbGRlciB9IGZyb20gJ2NkazhzLXBpcGVsaW5lcyc7XG5pbXBvcnQgeyBDb25zdHJ1Y3QgfSBmcm9tICdjb25zdHJ1Y3RzJztcbmltcG9ydCB7IE9jdG9raXQgfSBmcm9tICdvY3Rva2l0JztcblxuLyoqXG4gKiBJbml0aWFsaXphdGlvbiBwcm9wZXJ0aWVzIGZvciB0aGUgQVdTQ0RLUGlwZWxpbmVDaGFydFxuICovXG5leHBvcnQgaW50ZXJmYWNlIEFXU0NES1BpcGVsaW5lQ2hhcnRQcm9wcyBleHRlbmRzIENoYXJ0UHJvcHMge1xuICByZWFkb25seSBwYXJhbXM/OiBzdHJpbmdbXTtcbn1cblxuLyoqXG4gKiBUaGUgY2hhcnQgZm9yIGNyZWF0aW5nIGEgVGVrdG9uIFBpcGVsaW5lIHRoYXQgd2lsbCB1c2UgYW4gQVdTIENESyBwcm9qZWN0XG4gKiB0byBjcmVhdGUgcmVzb3VyY2VzIGluIEFXUyBmb3IgcmUtdXNhYmxlIGFydGlmYWN0cy5cbiAqL1xuZXhwb3J0IGNsYXNzIEFXU0NES1BpcGVsaW5lQ2hhcnQgZXh0ZW5kcyBDaGFydCB7XG4gIC8qKlxuICAgKiBJbml0aWFsaXplcyBhbiBpbnN0YW5jZSBvZiB0aGUgQVdTQ0RLUGlwZWxpbmVDaGFydC5cbiAgICpcbiAgICogQHBhcmFtIHNjb3BlXG4gICAqIEBwYXJhbSBpZFxuICAgKiBAcGFyYW0gcHJvcHNcbiAgICovXG4gIGNvbnN0cnVjdG9yKHNjb3BlOiBDb25zdHJ1Y3QsIGlkOiBzdHJpbmcsIHByb3BzOiBBV1NDREtQaXBlbGluZUNoYXJ0UHJvcHMgPSB7fSkge1xuICAgIHN1cGVyKHNjb3BlLCBpZCwgcHJvcHMpO1xuICAgIC8vIENyZWF0ZSB0aGUgcGlwZWxpbmUgdGhhdCB3aWxsIHJ1biBpbiB0aGUgT3BlblNoaWZ0IG9yIEs4cyBjbHVzdGVyLlxuICAgIC8vIGNvbnN0ICA9IG5ldyBQaXBlbGluZSh0aGlzLCAnYXdzLWNkay1waXBlbGluZScsIHtcbiAgICAvLyAgIG5hbWU6ICdhd3MtY2RrLXBpcGVsaW5lJyxcbiAgICAvLyB9KTtcbiAgICAvLyBwcm9wcy5wYXJhbXM/LmZvckVhY2gocyA9PiB7XG4gICAgLy8gICBwaXBlbGluZS5hZGRTdHJpbmdQYXJhbShzKTtcbiAgICAvLyB9KTtcbiAgICBjb25zdCBwaXBlbGluZSA9IG5ldyBQaXBlbGluZUJ1aWxkZXIodGhpcywgJ2F3cy1jZGstcGlwZWxpbmUnKVxuICAgICAgLndpdGhOYW1lKCdkZXBsb3ktY2RrLXByb2plY3QnKVxuICAgICAgLndpdGhEZXNjcmlwdGlvbignQSBwaXBlbGluZSBmb3IgZGVwbG95aW5nIGEgQVdTIENESyBwcm9qZWN0IGZyb20gYSBHaXRIdWIgcmVwb3NpdG9yeSB0byBhIGNsdXN0ZXIuJylcbiAgICAgIC53aXRoVGFzayhuZXcgVGFza0J1aWxkZXIodGhpcywgJ2dpdC1jbG9uZScpXG4gICAgICAgIC53aXRoTmFtZSgnZmV0Y2gtcHJvamVjdCcpXG4gICAgICAgIC8vIC53aXRoV29ya3NwYWNlKCdzc2gtY3JlZHMnLCAnc3NoLWNyZWRlbnRpYWxzJywgJ1RoZSBsb2NhdGlvbiBvZiB0aGUgU1NIIGtleXMgYW5kIGNyZWRlbnRpYWxzJyksXG4gICAgICAgIC53aXRoV29ya3NwYWNlKG5ldyBXb3Jrc3BhY2VCdWlsZGVyKCdvdXRwdXQnKVxuICAgICAgICAgIC53aXRoTmFtZSgnc2hhcmVkLWRhdGEnKVxuICAgICAgICAgIC53aXRoRGVzY3JpcHRpb24oJ1RoZSBBV1MgQ0RLIHByb2plY3QgZmlsZXMuJykpXG4gICAgICAgIC53aXRoV29ya3NwYWNlKG5ldyBXb3Jrc3BhY2VCdWlsZGVyKCdzc2gtY3JlZHMnKVxuICAgICAgICAgIC53aXRoTmFtZSgnc3NoLWNyZWRlbnRpYWxzJylcbiAgICAgICAgICAud2l0aERlc2NyaXB0aW9uKCdUaGUgbG9jYXRpb24gb2YgdGhlIFNTSCBrZXlzIGFuZCBjcmVkZW50aWFscycpKSk7XG5cbiAgICAvLyBOb3cgYnVpbGQgb3V0IHRoZVxuICAgIGNvbnN0IGF3c0Nka1Rhc2sgPSBuZXcgVGFza0J1aWxkZXIodGhpcywgJ2F3cy1jZGstc3ludGgnKVxuICAgICAgLndpdGhOYW1lKCdzeW50aC1jZGstcGlwZWxpbmUnKVxuICAgICAgLndpdGhXb3Jrc3BhY2UobmV3IFdvcmtzcGFjZUJ1aWxkZXIoJ3Byb2plY3RkYXRhJylcbiAgICAgICAgLndpdGhOYW1lKCdzaGFyZWQtZGF0YScpXG4gICAgICAgIC53aXRoRGVzY3JpcHRpb24oJ1RoZSBBV1MgQ0RLIHByb2plY3QgZmlsZXMnKSk7XG5cbiAgICBwcm9wcy5wYXJhbXM/LmZvckVhY2goKHMpID0+IHtcbiAgICAgIGF3c0Nka1Rhc2sud2l0aFN0cmluZ1BhcmFtKG5ldyBQYXJhbWV0ZXJCdWlsZGVyKHMpXG4gICAgICAgIC53aXRoUGlwbGluZVBhcmFtZXRlcihzKSk7XG4gICAgfSk7XG4gICAgcGlwZWxpbmUud2l0aFRhc2soYXdzQ2RrVGFzayk7XG4gICAgcGlwZWxpbmUuYnVpbGRQaXBlbGluZSgpO1xuICB9XG59XG5cbi8qKlxuICogQ29udGFpbnMgdGhlIGluZm9ybWF0aW9uIGZvciB0aGUgR2l0SHViIHJlcG8gYW5kIHRoZSBzdGFjayBzbyB3ZSBjYW4gZ28gZ2V0XG4gKiBpdCBhbmQgZ2VuZXJhdGUgdGhlIEFXUyBDREsgcGlwZWxpbmUuXG4gKi9cbmV4cG9ydCBpbnRlcmZhY2UgR2l0UmVwb0NvbmZpZyB7XG4gIC8qKlxuICAgKiBUaGUgVVJMIGZvciB0aGUgR2l0SHViIG9yIEdIRSBBUEkuIFRoZSB2YWx1ZSBzaG91bGQgbG9vayBsaWtlIGh0dHBzOi8vYXBpLmdpdGh1Yi5jb20gb3JcbiAgICogaHR0cHM6Ly9naXRodWIubXljb21wYW55LmNvbS9hcGkvdjMuXG4gICAqL1xuICByZWFkb25seSBnaFVybD86IHN0cmluZztcbiAgLyoqXG4gICAqIFRoZSBvd25lciBvZiB0aGUgR2l0SHViIHJlcG9zaXRvcnkuXG4gICAqL1xuICByZWFkb25seSBvd25lcj86IHN0cmluZztcbiAgLyoqXG4gICAqIFRoZSByZWxlYXNlIHRhZyBmb3IgdGhlIHJlbGVhc2UgaW4gd2hpY2ggdGhlIEFXUyBDREsgdGVtcGxhdGUgc2hvdWxkIGJlIGZvdW5kLlxuICAgKi9cbiAgcmVhZG9ubHkgcmVsZWFzZT86IHN0cmluZztcbiAgLyoqXG4gICAqIFRoZSBuYW1lIG9mIHRoZSByZXBvc2l0b3J5LlxuICAgKi9cbiAgcmVhZG9ubHkgcmVwbz86IHN0cmluZztcbiAgLyoqXG4gICAqIFRoZSBuYW1lIG9mIHRoZSBBV1MgQ0RLIHN0YWNrLiBUaGlzIHNob3VsZCBiZSBhIGdlbmVyYXRlZCB0ZW1wbGF0ZSB0aGF0IGlzIGluY2x1ZGVkXG4gICAqIGluIHRoZSByZWxlYXNlLlxuICAgKi9cbiAgcmVhZG9ubHkgc3RhY2tOYW1lPzogc3RyaW5nO1xuICAvKipcbiAgICogVGhlIHBlcnNvbmFsIGFjY2VzcyB0b2tlbiAoUEFUKSBmb3IgYWNjZXNzaW5nIHRoZSBsaWJyYXJ5IGluIEdpdEh1Yi5cbiAgICovXG4gIHJlYWRvbmx5IHRva2VuPzogc3RyaW5nO1xufVxuXG4vKipcbiAqIENyZWF0b3IgZm9yIHRoZSBBV1NDREtQaXBlbGluZUNoYXJ0XG4gKi9cbmV4cG9ydCBjbGFzcyBBV1NDREtQaXBlbGluZSB7XG5cbiAgLyoqXG4gICAqIEdlbmVyYXRlcyB0aGUgQVdTIENESyBQaXBlbGluZSAoQVdTQ0RLUGlwZWxpbmVDaGFydCkgYmFzZWQgb24gdGhlIGFjdHVhbCBwcm9qZWN0XG4gICAqIGxvY2F0ZWQgaW4gR2l0SHViIGFuZCBzcGVjaWZpZWQgYnkgdGhlIGNvbmZpZ3VyYXRpb24uXG4gICAqIEBwYXJhbSBjb25maWdcbiAgICovXG4gIHB1YmxpYyBzdGF0aWMgY3JlYXRlRnJvbShjb25maWc6IEdpdFJlcG9Db25maWcpOiB2b2lkIHtcbiAgICBjb25zdCBvY3Rva2l0ID0gbmV3IE9jdG9raXQoe1xuICAgICAgYXV0aDogY29uZmlnLnRva2VuLFxuICAgICAgYmFzZVVybDogY29uZmlnLmdoVXJsLFxuICAgIH0pO1xuXG4gICAgb2N0b2tpdC5yZXN0LnJlcG9zLmdldFJlbGVhc2VCeVRhZyh7XG4gICAgICBvd25lcjogY29uZmlnLm93bmVyISxcbiAgICAgIHJlcG86IGNvbmZpZy5yZXBvISxcbiAgICAgIHRhZzogY29uZmlnLnJlbGVhc2UhLFxuICAgIH0pLnRoZW4oZnVuY3Rpb24gKHJlbGVhc2VSZXNwb25zZSkge1xuICAgICAgY29uc3QgcmVsZWFzZUlkID0gcmVsZWFzZVJlc3BvbnNlLmRhdGEuaWQ7XG4gICAgICBvY3Rva2l0LnJlc3QucmVwb3MubGlzdFJlbGVhc2VBc3NldHMoe1xuICAgICAgICBvd25lcjogY29uZmlnLm93bmVyISxcbiAgICAgICAgcmVwbzogY29uZmlnLnJlcG8hLFxuICAgICAgICByZWxlYXNlX2lkOiByZWxlYXNlSWQsXG4gICAgICB9KS50aGVuKGZ1bmN0aW9uICh0ZW1wbGF0ZVJlc3BvbnNlKSB7XG4gICAgICAgIGNvbnN0IGFzc2V0ID0gdGVtcGxhdGVSZXNwb25zZS5kYXRhLmZpbmQoYSA9PiBhLm5hbWUgPT0gYCR7Y29uZmlnLnN0YWNrTmFtZX0udGVtcGxhdGUuanNvbmApO1xuICAgICAgICBjb25zdCBhc3NldElkID0gYXNzZXQ/LmlkO1xuICAgICAgICAvLyBOb3cgdGhhdCBJIGhhdmUgbXkgYXNzZXQgSUQsIEkgY2FuIGRvd25sb2FkIHRoZSBhc3NldC4uLlxuICAgICAgICBvY3Rva2l0LnJlc3QucmVwb3MuZ2V0UmVsZWFzZUFzc2V0KHtcbiAgICAgICAgICBvd25lcjogY29uZmlnLm93bmVyISxcbiAgICAgICAgICByZXBvOiBjb25maWcucmVwbyEsXG4gICAgICAgICAgYXNzZXRfaWQ6IE51bWJlcihhc3NldElkKSxcbiAgICAgICAgICBoZWFkZXJzOiB7XG4gICAgICAgICAgICBhY2NlcHQ6ICdhcHBsaWNhdGlvbi9vY3RldC1zdHJlYW0nLFxuICAgICAgICAgIH0sXG4gICAgICAgIH0pLnRoZW4oZnVuY3Rpb24gKGFzc2V0UmVzcG9uc2UpIHtcbiAgICAgICAgICBjb25zdCB0ZW1wbGF0ZSA9IEpTT04ucGFyc2UobmV3IFRleHREZWNvZGVyKCkuZGVjb2RlKGFzc2V0UmVzcG9uc2UuZGF0YSBhcyB1bmtub3duIGFzIEFycmF5QnVmZmVyKSk7XG4gICAgICAgICAgY29uc3QgYXBwID0gbmV3IEFwcCgpO1xuICAgICAgICAgIGNvbnN0IHRlbXBsYXRlUGFyYW1zOiBzdHJpbmdbXSA9IG5ldyBBcnJheTxzdHJpbmc+KCk7XG4gICAgICAgICAgLy8gTm93LCB3ZSBhcmUgZ29pbmcgdG8gYXV0b21hZ2ljYWxseSBhZGQgdGhlIEFXUy1yZXF1aXJlZCBwYXJhbWV0ZXJzIGJlZm9yZSBhZGRpbmcgdGhlXG4gICAgICAgICAgLy8gdGVtcGxhdGUgcGFyYW1ldGVycy5cbiAgICAgICAgICB0ZW1wbGF0ZVBhcmFtcy5wdXNoKCdBd3NBY2NvdW50SWQnKTtcbiAgICAgICAgICB0ZW1wbGF0ZVBhcmFtcy5wdXNoKCdBd3NBY2Nlc3NLZXlJZCcpO1xuICAgICAgICAgIHRlbXBsYXRlUGFyYW1zLnB1c2goJ0F3c1NlY3JldEtleUlkJyk7XG4gICAgICAgICAgdGVtcGxhdGVQYXJhbXMucHVzaCgnQXdzUmVnaW9uJyk7XG4gICAgICAgICAgT2JqZWN0LmtleXModGVtcGxhdGUuUGFyYW1ldGVycykuZm9yRWFjaChrZXkgPT4ge1xuICAgICAgICAgICAgdGVtcGxhdGVQYXJhbXMucHVzaChrZXkpO1xuICAgICAgICAgIH0pO1xuICAgICAgICAgIG5ldyBBV1NDREtQaXBlbGluZUNoYXJ0KGFwcCwgJ2V4YW1wbGUtY2RrOHMtcGlwZWxpbmUnLCB7XG4gICAgICAgICAgICBwYXJhbXM6IHRlbXBsYXRlUGFyYW1zLFxuICAgICAgICAgIH0pO1xuICAgICAgICAgIGFwcC5zeW50aCgpO1xuICAgICAgICB9KS5jYXRjaChmdW5jdGlvbiAoYXNzZXRFcnIpIHtcbiAgICAgICAgICBjb25zb2xlLmVycm9yKGFzc2V0RXJyKTtcbiAgICAgICAgfSk7XG4gICAgICB9KS5jYXRjaChmdW5jdGlvbiAodGVtcGxhdGVFcnIpIHtcbiAgICAgICAgY29uc29sZS5lcnJvcih0ZW1wbGF0ZUVycik7XG4gICAgICB9KTtcbiAgICB9KS5jYXRjaChmdW5jdGlvbiAocmVsZWFzZUVycikge1xuICAgICAgY29uc29sZS5lcnJvcihyZWxlYXNlRXJyKTtcbiAgICB9KTtcbiAgfVxufVxuIl19