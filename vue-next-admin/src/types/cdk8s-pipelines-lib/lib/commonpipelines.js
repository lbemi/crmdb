"use strict";
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
exports.InstallFromIBMOperatorPipeline = void 0;
const JSII_RTTI_SYMBOL_1 = Symbol.for("jsii.rtti");
const cdk8s_pipelines_1 = require("cdk8s-pipelines");
const commontasks_1 = require("./commontasks");
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
class InstallFromIBMOperatorPipeline extends cdk8s_pipelines_1.PipelineBuilder {
    /**
     *
     * @param scope The parent [Construct](https://cdk8s.io/docs/latest/basics/constructs/).
     * @param id The `id` of the construct. Must be unique for each one in a chart.
     * @param ns The namespace to create and to use for subscribing to the product and channel.
     * @param subscription The name of the subscription. For example, for IBM Event Streams is it `ibm-eventstreams`. For Red Hat Serverless, it is `serverless-operator`.
     * @param channel The name of the channel (e.g., `v3.3`, `stable`).
     */
    constructor(scope, id, ns, subscription, channel) {
        super(scope, id);
        const labels = {};
        super.withTask((0, commontasks_1.CreateNamespace)(scope, 'create-namespace', ns));
        super.withTask((0, commontasks_1.RegisterIBMOperatorCatalog)(scope, 'register-ibm-operators', labels, 45));
        super.withTask((0, commontasks_1.CreateOperatorGroup)(scope, 'create-operator-group', ns, `${subscription}-operator-group`));
        super.withTask((0, commontasks_1.Subscribe)(scope, 'subscribe', ns, subscription, 'ibm-operator-catalog', channel));
    }
    /**
     *
     * @param opts
     */
    buildPipeline(opts = cdk8s_pipelines_1.DefaultBuilderOptions) {
        // Add the
        super.buildPipeline(opts);
    }
}
exports.InstallFromIBMOperatorPipeline = InstallFromIBMOperatorPipeline;
_a = JSII_RTTI_SYMBOL_1;
InstallFromIBMOperatorPipeline[_a] = { fqn: "cdk8s-pipelines-lib.InstallFromIBMOperatorPipeline", version: "0.0.12" };
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiY29tbW9ucGlwZWxpbmVzLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsiLi4vc3JjL2NvbW1vbnBpcGVsaW5lcy50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiOzs7OztBQUFBLHFEQUF5RjtBQUV6RiwrQ0FBNEc7QUFFNUc7Ozs7Ozs7Ozs7R0FVRztBQUNILE1BQWEsOEJBQStCLFNBQVEsaUNBQWU7SUFFakU7Ozs7Ozs7T0FPRztJQUNILFlBQW1CLEtBQWdCLEVBQUUsRUFBVSxFQUFFLEVBQVUsRUFBRSxZQUFvQixFQUFFLE9BQWU7UUFDaEcsS0FBSyxDQUFDLEtBQUssRUFBRSxFQUFFLENBQUMsQ0FBQztRQUNqQixNQUFNLE1BQU0sR0FBRyxFQUFFLENBQUM7UUFDbEIsS0FBSyxDQUFDLFFBQVEsQ0FBQyxJQUFBLDZCQUFlLEVBQUMsS0FBSyxFQUFFLGtCQUFrQixFQUFFLEVBQUUsQ0FBQyxDQUFDLENBQUM7UUFDL0QsS0FBSyxDQUFDLFFBQVEsQ0FBQyxJQUFBLHdDQUEwQixFQUFDLEtBQUssRUFBRSx3QkFBd0IsRUFBRSxNQUFNLEVBQUUsRUFBRSxDQUFDLENBQUMsQ0FBQztRQUN4RixLQUFLLENBQUMsUUFBUSxDQUFDLElBQUEsaUNBQW1CLEVBQUMsS0FBSyxFQUFFLHVCQUF1QixFQUFFLEVBQUUsRUFBRSxHQUFHLFlBQVksaUJBQWlCLENBQUMsQ0FBQyxDQUFDO1FBQzFHLEtBQUssQ0FBQyxRQUFRLENBQUMsSUFBQSx1QkFBUyxFQUFDLEtBQUssRUFBRSxXQUFXLEVBQUUsRUFBRSxFQUFFLFlBQVksRUFBRSxzQkFBc0IsRUFBRSxPQUFPLENBQUMsQ0FBQyxDQUFDO0lBQ25HLENBQUM7SUFFRDs7O09BR0c7SUFDSSxhQUFhLENBQUMsT0FBdUIsdUNBQXFCO1FBQy9ELFVBQVU7UUFDVixLQUFLLENBQUMsYUFBYSxDQUFDLElBQUksQ0FBQyxDQUFDO0lBQzVCLENBQUM7O0FBMUJILHdFQTJCQyIsInNvdXJjZXNDb250ZW50IjpbImltcG9ydCB7IEJ1aWxkZXJPcHRpb25zLCBEZWZhdWx0QnVpbGRlck9wdGlvbnMsIFBpcGVsaW5lQnVpbGRlciB9IGZyb20gJ2NkazhzLXBpcGVsaW5lcyc7XG5pbXBvcnQgeyBDb25zdHJ1Y3QgfSBmcm9tICdjb25zdHJ1Y3RzJztcbmltcG9ydCB7IENyZWF0ZU5hbWVzcGFjZSwgQ3JlYXRlT3BlcmF0b3JHcm91cCwgUmVnaXN0ZXJJQk1PcGVyYXRvckNhdGFsb2csIFN1YnNjcmliZSB9IGZyb20gJy4vY29tbW9udGFza3MnO1xuXG4vKipcbiAqIEEgYmFzaWMgcGlwZWxpbmUgdGhhdCBzdGFydHMgd2l0aCBhIHN1YnNjcmlwdGlvbiB0byB0aGUgSUJNIG9wZXJhdG9yIGNhdGFsb2cuXG4gKlxuICogVGhlIGZvbGxvd2luZyBzdGVwcyBhcmUgaW5jbHVkZWQgaW4gdGhpcyBwaXBlbGluZSwgc28geW91IGRvIG5vdCBuZWVkIHRvIGFkZFxuICogdGhlbS4gVGhlIHBpcGVsaW5lOlxuICpcbiAqIDEuIENyZWF0ZXMgdGhlIHNwZWNpZmllZCBuYW1lc3BhY2UuXG4gKiAxLiBSZWdpc3RlcnMgdGhlIElCTSBvcGVyYXRvci5cbiAqIDEuIENyZWF0ZXMgYW4gT3BlcmF0b3JHcm91cC5cbiAqIDEuIFN1YnNjcmliZXMgdG8gdGhlIGdpdmVuIGBuYW1lYCBhbmQgYGNoYW5uZWxgXG4gKi9cbmV4cG9ydCBjbGFzcyBJbnN0YWxsRnJvbUlCTU9wZXJhdG9yUGlwZWxpbmUgZXh0ZW5kcyBQaXBlbGluZUJ1aWxkZXIge1xuXG4gIC8qKlxuICAgKlxuICAgKiBAcGFyYW0gc2NvcGUgVGhlIHBhcmVudCBbQ29uc3RydWN0XShodHRwczovL2NkazhzLmlvL2RvY3MvbGF0ZXN0L2Jhc2ljcy9jb25zdHJ1Y3RzLykuXG4gICAqIEBwYXJhbSBpZCBUaGUgYGlkYCBvZiB0aGUgY29uc3RydWN0LiBNdXN0IGJlIHVuaXF1ZSBmb3IgZWFjaCBvbmUgaW4gYSBjaGFydC5cbiAgICogQHBhcmFtIG5zIFRoZSBuYW1lc3BhY2UgdG8gY3JlYXRlIGFuZCB0byB1c2UgZm9yIHN1YnNjcmliaW5nIHRvIHRoZSBwcm9kdWN0IGFuZCBjaGFubmVsLlxuICAgKiBAcGFyYW0gc3Vic2NyaXB0aW9uIFRoZSBuYW1lIG9mIHRoZSBzdWJzY3JpcHRpb24uIEZvciBleGFtcGxlLCBmb3IgSUJNIEV2ZW50IFN0cmVhbXMgaXMgaXQgYGlibS1ldmVudHN0cmVhbXNgLiBGb3IgUmVkIEhhdCBTZXJ2ZXJsZXNzLCBpdCBpcyBgc2VydmVybGVzcy1vcGVyYXRvcmAuXG4gICAqIEBwYXJhbSBjaGFubmVsIFRoZSBuYW1lIG9mIHRoZSBjaGFubmVsIChlLmcuLCBgdjMuM2AsIGBzdGFibGVgKS5cbiAgICovXG4gIHB1YmxpYyBjb25zdHJ1Y3RvcihzY29wZTogQ29uc3RydWN0LCBpZDogc3RyaW5nLCBuczogc3RyaW5nLCBzdWJzY3JpcHRpb246IHN0cmluZywgY2hhbm5lbDogc3RyaW5nKSB7XG4gICAgc3VwZXIoc2NvcGUsIGlkKTtcbiAgICBjb25zdCBsYWJlbHMgPSB7fTtcbiAgICBzdXBlci53aXRoVGFzayhDcmVhdGVOYW1lc3BhY2Uoc2NvcGUsICdjcmVhdGUtbmFtZXNwYWNlJywgbnMpKTtcbiAgICBzdXBlci53aXRoVGFzayhSZWdpc3RlcklCTU9wZXJhdG9yQ2F0YWxvZyhzY29wZSwgJ3JlZ2lzdGVyLWlibS1vcGVyYXRvcnMnLCBsYWJlbHMsIDQ1KSk7XG4gICAgc3VwZXIud2l0aFRhc2soQ3JlYXRlT3BlcmF0b3JHcm91cChzY29wZSwgJ2NyZWF0ZS1vcGVyYXRvci1ncm91cCcsIG5zLCBgJHtzdWJzY3JpcHRpb259LW9wZXJhdG9yLWdyb3VwYCkpO1xuICAgIHN1cGVyLndpdGhUYXNrKFN1YnNjcmliZShzY29wZSwgJ3N1YnNjcmliZScsIG5zLCBzdWJzY3JpcHRpb24sICdpYm0tb3BlcmF0b3ItY2F0YWxvZycsIGNoYW5uZWwpKTtcbiAgfVxuXG4gIC8qKlxuICAgKlxuICAgKiBAcGFyYW0gb3B0c1xuICAgKi9cbiAgcHVibGljIGJ1aWxkUGlwZWxpbmUob3B0czogQnVpbGRlck9wdGlvbnMgPSBEZWZhdWx0QnVpbGRlck9wdGlvbnMpOiB2b2lkIHtcbiAgICAvLyBBZGQgdGhlXG4gICAgc3VwZXIuYnVpbGRQaXBlbGluZShvcHRzKTtcbiAgfVxufSJdfQ==