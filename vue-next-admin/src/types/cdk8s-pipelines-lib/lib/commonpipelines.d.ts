import { BuilderOptions, PipelineBuilder } from 'cdk8s-pipelines';
import { Construct } from 'constructs';
/**
 * A basic pipeline that starts with a subscription to the IBM operator catalog.
 *
 * The following steps are included in this pipeline, so you do not need to add
 * them. The pipeline:
 *
 * 1. Creates the specified namespace.
 * 1. Registers the IBM operator.
 * 1. Creates an OperatorGroup.
 * 1. Subscribes to the given `name` and `channel`
 */
export declare class InstallFromIBMOperatorPipeline extends PipelineBuilder {
    /**
     *
     * @param scope The parent [Construct](https://cdk8s.io/docs/latest/basics/constructs/).
     * @param id The `id` of the construct. Must be unique for each one in a chart.
     * @param ns The namespace to create and to use for subscribing to the product and channel.
     * @param subscription The name of the subscription. For example, for IBM Event Streams is it `ibm-eventstreams`. For Red Hat Serverless, it is `serverless-operator`.
     * @param channel The name of the channel (e.g., `v3.3`, `stable`).
     */
    constructor(scope: Construct, id: string, ns: string, subscription: string, channel: string);
    /**
     *
     * @param opts
     */
    buildPipeline(opts?: BuilderOptions): void;
}
