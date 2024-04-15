import { Chart, ChartProps } from 'cdk8s';
import { Construct } from 'constructs';
/**
 * Initialization properties for the AWSCDKPipelineChart
 */
export interface AWSCDKPipelineChartProps extends ChartProps {
    readonly params?: string[];
}
/**
 * The chart for creating a Tekton Pipeline that will use an AWS CDK project
 * to create resources in AWS for re-usable artifacts.
 */
export declare class AWSCDKPipelineChart extends Chart {
    /**
     * Initializes an instance of the AWSCDKPipelineChart.
     *
     * @param scope
     * @param id
     * @param props
     */
    constructor(scope: Construct, id: string, props?: AWSCDKPipelineChartProps);
}
/**
 * Contains the information for the GitHub repo and the stack so we can go get
 * it and generate the AWS CDK pipeline.
 */
export interface GitRepoConfig {
    /**
     * The URL for the GitHub or GHE API. The value should look like https://api.github.com or
     * https://github.mycompany.com/api/v3.
     */
    readonly ghUrl?: string;
    /**
     * The owner of the GitHub repository.
     */
    readonly owner?: string;
    /**
     * The release tag for the release in which the AWS CDK template should be found.
     */
    readonly release?: string;
    /**
     * The name of the repository.
     */
    readonly repo?: string;
    /**
     * The name of the AWS CDK stack. This should be a generated template that is included
     * in the release.
     */
    readonly stackName?: string;
    /**
     * The personal access token (PAT) for accessing the library in GitHub.
     */
    readonly token?: string;
}
/**
 * Creator for the AWSCDKPipelineChart
 */
export declare class AWSCDKPipeline {
    /**
     * Generates the AWS CDK Pipeline (AWSCDKPipelineChart) based on the actual project
     * located in GitHub and specified by the configuration.
     * @param config
     */
    static createFrom(config: GitRepoConfig): void;
}
