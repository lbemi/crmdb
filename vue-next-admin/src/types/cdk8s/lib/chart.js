"use strict";
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
exports.Chart = void 0;
const JSII_RTTI_SYMBOL_1 = Symbol.for("jsii.rtti");
const constructs_1 = require("constructs");
const api_object_1 = require("./api-object");
const app_1 = require("./app");
const names_1 = require("./names");
const CHART_SYMBOL = Symbol.for('cdk8s.Chart');
const CRONJOB = 'CronJob';
class Chart extends constructs_1.Construct {
    /**
     * Return whether the given object is a Chart.
     *
     * We do attribute detection since we can't reliably use 'instanceof'.
     */
    static isChart(x) {
        return x !== null && typeof (x) === 'object' && CHART_SYMBOL in x;
    }
    /**
     * Implements `instanceof Chart` using the more reliable `Chart.isChart` static method
     *
     * @param o The object to check
     * @internal
     */
    static [(_a = JSII_RTTI_SYMBOL_1, Symbol.hasInstance)](o) {
        return Chart.isChart(o);
    }
    /**
     * Finds the chart in which a node is defined.
     * @param c a construct node
     */
    static of(c) {
        if (Chart.isChart(c)) {
            return c;
        }
        const parent = c.node.scope;
        if (!parent) {
            throw new Error('cannot find a parent chart (directly or indirectly)');
        }
        return Chart.of(parent);
    }
    constructor(scope, id, props = {}) {
        super(scope, id);
        this.namespace = props.namespace;
        this._labels = props.labels ?? {};
        this._disableResourceNameHashes = props.disableResourceNameHashes ?? false;
        Object.defineProperty(this, CHART_SYMBOL, { value: true });
    }
    /**
     * Labels applied to all resources in this chart.
     *
     * This is an immutable copy.
     */
    get labels() {
        return { ...this._labels };
    }
    /**
     * Generates a app-unique name for an object given it's construct node path.
     *
     * Different resource types may have different constraints on names
     * (`metadata.name`). The previous version of the name generator was
     * compatible with DNS_SUBDOMAIN but not with DNS_LABEL.
     *
     * For example, `Deployment` names must comply with DNS_SUBDOMAIN while
     * `Service` names must comply with DNS_LABEL.
     *
     * Since there is no formal specification for this, the default name
     * generation scheme for kubernetes objects in cdk8s was changed to DNS_LABEL,
     * since it’s the common denominator for all kubernetes resources
     * (supposedly).
     *
     * You can override this method if you wish to customize object names at the
     * chart level.
     *
     * @param apiObject The API object to generate a name for.
     */
    generateObjectName(apiObject) {
        return names_1.Names.toDnsLabel(apiObject, {
            includeHash: !this._disableResourceNameHashes,
            maxLen: apiObject.kind == CRONJOB ? 52 : undefined,
        });
    }
    /**
     * Create a dependency between this Chart and other constructs.
     * These can be other ApiObjects, Charts, or custom.
     *
     * @param dependencies the dependencies to add.
     */
    addDependency(...dependencies) {
        this.node.addDependency(...dependencies);
    }
    /**
     * Renders this chart to a set of Kubernetes JSON resources.
     * @returns array of resource manifests
     */
    toJson() {
        return app_1.App._synthChart(this);
    }
    /**
     * Returns all the included API objects.
     */
    get apiObjects() {
        return this.node.children.filter((o) => o instanceof api_object_1.ApiObject);
    }
}
exports.Chart = Chart;
Chart[_a] = { fqn: "cdk8s.Chart", version: "2.68.60" };
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiY2hhcnQuanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyIuLi9zcmMvY2hhcnQudHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6Ijs7Ozs7QUFBQSwyQ0FBbUQ7QUFDbkQsNkNBQXlDO0FBQ3pDLCtCQUE0QjtBQUM1QixtQ0FBZ0M7QUFFaEMsTUFBTSxZQUFZLEdBQUcsTUFBTSxDQUFDLEdBQUcsQ0FBQyxhQUFhLENBQUMsQ0FBQztBQUMvQyxNQUFNLE9BQU8sR0FBRyxTQUFTLENBQUM7QUE2QjFCLE1BQWEsS0FBTSxTQUFRLHNCQUFTO0lBQ2xDOzs7O09BSUc7SUFDSSxNQUFNLENBQUMsT0FBTyxDQUFDLENBQU07UUFDMUIsT0FBTyxDQUFDLEtBQUssSUFBSSxJQUFJLE9BQU0sQ0FBQyxDQUFDLENBQUMsS0FBSyxRQUFRLElBQUksWUFBWSxJQUFJLENBQUMsQ0FBQztJQUNuRSxDQUFDO0lBRUQ7Ozs7O09BS0c7SUFDSCxNQUFNLENBQUMsMkJBQUMsTUFBTSxDQUFDLFdBQVcsRUFBQyxDQUFDLENBQVU7UUFDcEMsT0FBTyxLQUFLLENBQUMsT0FBTyxDQUFDLENBQUMsQ0FBQyxDQUFDO0lBQzFCLENBQUM7SUFFRDs7O09BR0c7SUFDSSxNQUFNLENBQUMsRUFBRSxDQUFDLENBQWE7UUFDNUIsSUFBSSxLQUFLLENBQUMsT0FBTyxDQUFDLENBQUMsQ0FBQyxFQUFFLENBQUM7WUFDckIsT0FBTyxDQUFDLENBQUM7UUFDWCxDQUFDO1FBRUQsTUFBTSxNQUFNLEdBQUcsQ0FBQyxDQUFDLElBQUksQ0FBQyxLQUFrQixDQUFDO1FBQ3pDLElBQUksQ0FBQyxNQUFNLEVBQUUsQ0FBQztZQUNaLE1BQU0sSUFBSSxLQUFLLENBQUMscURBQXFELENBQUMsQ0FBQztRQUN6RSxDQUFDO1FBRUQsT0FBTyxLQUFLLENBQUMsRUFBRSxDQUFDLE1BQU0sQ0FBQyxDQUFDO0lBQzFCLENBQUM7SUFpQkQsWUFBWSxLQUFnQixFQUFFLEVBQVUsRUFBRSxRQUFvQixFQUFHO1FBQy9ELEtBQUssQ0FBQyxLQUFLLEVBQUUsRUFBRSxDQUFDLENBQUM7UUFDakIsSUFBSSxDQUFDLFNBQVMsR0FBRyxLQUFLLENBQUMsU0FBUyxDQUFDO1FBQ2pDLElBQUksQ0FBQyxPQUFPLEdBQUcsS0FBSyxDQUFDLE1BQU0sSUFBSSxFQUFFLENBQUM7UUFDbEMsSUFBSSxDQUFDLDBCQUEwQixHQUFHLEtBQUssQ0FBQyx5QkFBeUIsSUFBSSxLQUFLLENBQUM7UUFFM0UsTUFBTSxDQUFDLGNBQWMsQ0FBQyxJQUFJLEVBQUUsWUFBWSxFQUFFLEVBQUUsS0FBSyxFQUFFLElBQUksRUFBRSxDQUFDLENBQUM7SUFDN0QsQ0FBQztJQUVEOzs7O09BSUc7SUFDSCxJQUFXLE1BQU07UUFDZixPQUFPLEVBQUUsR0FBRyxJQUFJLENBQUMsT0FBTyxFQUFFLENBQUM7SUFDN0IsQ0FBQztJQUVEOzs7Ozs7Ozs7Ozs7Ozs7Ozs7O09BbUJHO0lBQ0ksa0JBQWtCLENBQUMsU0FBb0I7UUFDNUMsT0FBTyxhQUFLLENBQUMsVUFBVSxDQUFDLFNBQVMsRUFBRTtZQUNqQyxXQUFXLEVBQUUsQ0FBQyxJQUFJLENBQUMsMEJBQTBCO1lBQzdDLE1BQU0sRUFBRSxTQUFTLENBQUMsSUFBSSxJQUFJLE9BQU8sQ0FBQyxDQUFDLENBQUMsRUFBRSxDQUFDLENBQUMsQ0FBQyxTQUFTO1NBQ25ELENBQUMsQ0FBQztJQUNMLENBQUM7SUFFRDs7Ozs7T0FLRztJQUNJLGFBQWEsQ0FBQyxHQUFHLFlBQTBCO1FBQ2hELElBQUksQ0FBQyxJQUFJLENBQUMsYUFBYSxDQUFDLEdBQUcsWUFBWSxDQUFDLENBQUM7SUFDM0MsQ0FBQztJQUVEOzs7T0FHRztJQUNJLE1BQU07UUFDWCxPQUFPLFNBQUcsQ0FBQyxXQUFXLENBQUMsSUFBSSxDQUFDLENBQUM7SUFDL0IsQ0FBQztJQUVEOztPQUVHO0lBQ0gsSUFBSSxVQUFVO1FBQ1osT0FBTyxJQUFJLENBQUMsSUFBSSxDQUFDLFFBQVEsQ0FBQyxNQUFNLENBQUMsQ0FBQyxDQUFDLEVBQWtCLEVBQUUsQ0FBQyxDQUFDLFlBQVksc0JBQVMsQ0FBQyxDQUFDO0lBQ2xGLENBQUM7O0FBeEhILHNCQXlIQyIsInNvdXJjZXNDb250ZW50IjpbImltcG9ydCB7IENvbnN0cnVjdCwgSUNvbnN0cnVjdCB9IGZyb20gJ2NvbnN0cnVjdHMnO1xuaW1wb3J0IHsgQXBpT2JqZWN0IH0gZnJvbSAnLi9hcGktb2JqZWN0JztcbmltcG9ydCB7IEFwcCB9IGZyb20gJy4vYXBwJztcbmltcG9ydCB7IE5hbWVzIH0gZnJvbSAnLi9uYW1lcyc7XG5cbmNvbnN0IENIQVJUX1NZTUJPTCA9IFN5bWJvbC5mb3IoJ2NkazhzLkNoYXJ0Jyk7XG5jb25zdCBDUk9OSk9CID0gJ0Nyb25Kb2InO1xuXG5leHBvcnQgaW50ZXJmYWNlIENoYXJ0UHJvcHMge1xuICAvKipcbiAgICogVGhlIGRlZmF1bHQgbmFtZXNwYWNlIGZvciBhbGwgb2JqZWN0cyBkZWZpbmVkIGluIHRoaXMgY2hhcnQgKGRpcmVjdGx5IG9yXG4gICAqIGluZGlyZWN0bHkpLiBUaGlzIG5hbWVzcGFjZSB3aWxsIG9ubHkgYXBwbHkgdG8gb2JqZWN0cyB0aGF0IGRvbid0IGhhdmUgYVxuICAgKiBgbmFtZXNwYWNlYCBleHBsaWNpdGx5IGRlZmluZWQgZm9yIHRoZW0uXG4gICAqXG4gICAqIEBkZWZhdWx0IC0gbm8gbmFtZXNwYWNlIGlzIHN5bnRoZXNpemVkICh1c3VhbGx5IHRoaXMgaW1wbGllcyBcImRlZmF1bHRcIilcbiAgICovXG4gIHJlYWRvbmx5IG5hbWVzcGFjZT86IHN0cmluZztcblxuICAvKipcbiAgICogTGFiZWxzIHRvIGFwcGx5IHRvIGFsbCByZXNvdXJjZXMgaW4gdGhpcyBjaGFydC5cbiAgICpcbiAgICogQGRlZmF1bHQgLSBubyBjb21tb24gbGFiZWxzXG4gICAqL1xuICByZWFkb25seSBsYWJlbHM/OiB7IFtuYW1lOiBzdHJpbmddOiBzdHJpbmcgfTtcblxuICAvKipcbiAgICogVGhlIGF1dG9nZW5lcmF0ZWQgcmVzb3VyY2UgbmFtZSBieSBkZWZhdWx0IGlzIHN1ZmZpeGVkIHdpdGggYSBzdGFibGUgaGFzaFxuICAgKiBvZiB0aGUgY29uc3RydWN0IHBhdGguIFNldHRpbmcgdGhpcyBwcm9wZXJ0eSB0byB0cnVlIGRyb3BzIHRoZSBoYXNoIHN1ZmZpeC5cbiAgICpcbiAgICogQGRlZmF1bHQgZmFsc2VcbiAgICovXG4gIHJlYWRvbmx5IGRpc2FibGVSZXNvdXJjZU5hbWVIYXNoZXM/OiBib29sZWFuO1xuXG59XG5cbmV4cG9ydCBjbGFzcyBDaGFydCBleHRlbmRzIENvbnN0cnVjdCB7XG4gIC8qKlxuICAgKiBSZXR1cm4gd2hldGhlciB0aGUgZ2l2ZW4gb2JqZWN0IGlzIGEgQ2hhcnQuXG4gICAqXG4gICAqIFdlIGRvIGF0dHJpYnV0ZSBkZXRlY3Rpb24gc2luY2Ugd2UgY2FuJ3QgcmVsaWFibHkgdXNlICdpbnN0YW5jZW9mJy5cbiAgICovXG4gIHB1YmxpYyBzdGF0aWMgaXNDaGFydCh4OiBhbnkpOiB4IGlzIENoYXJ0IHtcbiAgICByZXR1cm4geCAhPT0gbnVsbCAmJiB0eXBlb2YoeCkgPT09ICdvYmplY3QnICYmIENIQVJUX1NZTUJPTCBpbiB4O1xuICB9XG5cbiAgLyoqXG4gICAqIEltcGxlbWVudHMgYGluc3RhbmNlb2YgQ2hhcnRgIHVzaW5nIHRoZSBtb3JlIHJlbGlhYmxlIGBDaGFydC5pc0NoYXJ0YCBzdGF0aWMgbWV0aG9kXG4gICAqXG4gICAqIEBwYXJhbSBvIFRoZSBvYmplY3QgdG8gY2hlY2tcbiAgICogQGludGVybmFsXG4gICAqL1xuICBzdGF0aWMgW1N5bWJvbC5oYXNJbnN0YW5jZV0obzogdW5rbm93bikge1xuICAgIHJldHVybiBDaGFydC5pc0NoYXJ0KG8pO1xuICB9XG5cbiAgLyoqXG4gICAqIEZpbmRzIHRoZSBjaGFydCBpbiB3aGljaCBhIG5vZGUgaXMgZGVmaW5lZC5cbiAgICogQHBhcmFtIGMgYSBjb25zdHJ1Y3Qgbm9kZVxuICAgKi9cbiAgcHVibGljIHN0YXRpYyBvZihjOiBJQ29uc3RydWN0KTogQ2hhcnQge1xuICAgIGlmIChDaGFydC5pc0NoYXJ0KGMpKSB7XG4gICAgICByZXR1cm4gYztcbiAgICB9XG5cbiAgICBjb25zdCBwYXJlbnQgPSBjLm5vZGUuc2NvcGUgYXMgQ29uc3RydWN0O1xuICAgIGlmICghcGFyZW50KSB7XG4gICAgICB0aHJvdyBuZXcgRXJyb3IoJ2Nhbm5vdCBmaW5kIGEgcGFyZW50IGNoYXJ0IChkaXJlY3RseSBvciBpbmRpcmVjdGx5KScpO1xuICAgIH1cblxuICAgIHJldHVybiBDaGFydC5vZihwYXJlbnQpO1xuICB9XG5cbiAgLyoqXG4gICAqIFRoZSBkZWZhdWx0IG5hbWVzcGFjZSBmb3IgYWxsIG9iamVjdHMgaW4gdGhpcyBjaGFydC5cbiAgICovXG4gIHB1YmxpYyByZWFkb25seSBuYW1lc3BhY2U/OiBzdHJpbmc7XG5cbiAgLyoqXG4gICAqIENoYXJ0LWxldmVsIGxhYmVscy5cbiAgICovXG4gIHByaXZhdGUgcmVhZG9ubHkgX2xhYmVscz86IHsgW25hbWU6IHN0cmluZ106IHN0cmluZyB9O1xuXG4gIC8qKlxuICAgKiBEZXRlcm1pbmVzIGlmIHJlc291cmNlIG5hbWVzIGluIHRoZSBjaGFydCBoYXZlIHRoZSBzdWZmaXhlZCBoYXNoLlxuICAgKi9cbiAgcHJpdmF0ZSByZWFkb25seSBfZGlzYWJsZVJlc291cmNlTmFtZUhhc2hlcz86IGJvb2xlYW47XG5cbiAgY29uc3RydWN0b3Ioc2NvcGU6IENvbnN0cnVjdCwgaWQ6IHN0cmluZywgcHJvcHM6IENoYXJ0UHJvcHMgPSB7IH0pIHtcbiAgICBzdXBlcihzY29wZSwgaWQpO1xuICAgIHRoaXMubmFtZXNwYWNlID0gcHJvcHMubmFtZXNwYWNlO1xuICAgIHRoaXMuX2xhYmVscyA9IHByb3BzLmxhYmVscyA/PyB7fTtcbiAgICB0aGlzLl9kaXNhYmxlUmVzb3VyY2VOYW1lSGFzaGVzID0gcHJvcHMuZGlzYWJsZVJlc291cmNlTmFtZUhhc2hlcyA/PyBmYWxzZTtcblxuICAgIE9iamVjdC5kZWZpbmVQcm9wZXJ0eSh0aGlzLCBDSEFSVF9TWU1CT0wsIHsgdmFsdWU6IHRydWUgfSk7XG4gIH1cblxuICAvKipcbiAgICogTGFiZWxzIGFwcGxpZWQgdG8gYWxsIHJlc291cmNlcyBpbiB0aGlzIGNoYXJ0LlxuICAgKlxuICAgKiBUaGlzIGlzIGFuIGltbXV0YWJsZSBjb3B5LlxuICAgKi9cbiAgcHVibGljIGdldCBsYWJlbHMoKTogeyBbbmFtZTogc3RyaW5nXTogc3RyaW5nIH0ge1xuICAgIHJldHVybiB7IC4uLnRoaXMuX2xhYmVscyB9O1xuICB9XG5cbiAgLyoqXG4gICAqIEdlbmVyYXRlcyBhIGFwcC11bmlxdWUgbmFtZSBmb3IgYW4gb2JqZWN0IGdpdmVuIGl0J3MgY29uc3RydWN0IG5vZGUgcGF0aC5cbiAgICpcbiAgICogRGlmZmVyZW50IHJlc291cmNlIHR5cGVzIG1heSBoYXZlIGRpZmZlcmVudCBjb25zdHJhaW50cyBvbiBuYW1lc1xuICAgKiAoYG1ldGFkYXRhLm5hbWVgKS4gVGhlIHByZXZpb3VzIHZlcnNpb24gb2YgdGhlIG5hbWUgZ2VuZXJhdG9yIHdhc1xuICAgKiBjb21wYXRpYmxlIHdpdGggRE5TX1NVQkRPTUFJTiBidXQgbm90IHdpdGggRE5TX0xBQkVMLlxuICAgKlxuICAgKiBGb3IgZXhhbXBsZSwgYERlcGxveW1lbnRgIG5hbWVzIG11c3QgY29tcGx5IHdpdGggRE5TX1NVQkRPTUFJTiB3aGlsZVxuICAgKiBgU2VydmljZWAgbmFtZXMgbXVzdCBjb21wbHkgd2l0aCBETlNfTEFCRUwuXG4gICAqXG4gICAqIFNpbmNlIHRoZXJlIGlzIG5vIGZvcm1hbCBzcGVjaWZpY2F0aW9uIGZvciB0aGlzLCB0aGUgZGVmYXVsdCBuYW1lXG4gICAqIGdlbmVyYXRpb24gc2NoZW1lIGZvciBrdWJlcm5ldGVzIG9iamVjdHMgaW4gY2RrOHMgd2FzIGNoYW5nZWQgdG8gRE5TX0xBQkVMLFxuICAgKiBzaW5jZSBpdOKAmXMgdGhlIGNvbW1vbiBkZW5vbWluYXRvciBmb3IgYWxsIGt1YmVybmV0ZXMgcmVzb3VyY2VzXG4gICAqIChzdXBwb3NlZGx5KS5cbiAgICpcbiAgICogWW91IGNhbiBvdmVycmlkZSB0aGlzIG1ldGhvZCBpZiB5b3Ugd2lzaCB0byBjdXN0b21pemUgb2JqZWN0IG5hbWVzIGF0IHRoZVxuICAgKiBjaGFydCBsZXZlbC5cbiAgICpcbiAgICogQHBhcmFtIGFwaU9iamVjdCBUaGUgQVBJIG9iamVjdCB0byBnZW5lcmF0ZSBhIG5hbWUgZm9yLlxuICAgKi9cbiAgcHVibGljIGdlbmVyYXRlT2JqZWN0TmFtZShhcGlPYmplY3Q6IEFwaU9iamVjdCkge1xuICAgIHJldHVybiBOYW1lcy50b0Ruc0xhYmVsKGFwaU9iamVjdCwge1xuICAgICAgaW5jbHVkZUhhc2g6ICF0aGlzLl9kaXNhYmxlUmVzb3VyY2VOYW1lSGFzaGVzLFxuICAgICAgbWF4TGVuOiBhcGlPYmplY3Qua2luZCA9PSBDUk9OSk9CID8gNTIgOiB1bmRlZmluZWQsXG4gICAgfSk7XG4gIH1cblxuICAvKipcbiAgICogQ3JlYXRlIGEgZGVwZW5kZW5jeSBiZXR3ZWVuIHRoaXMgQ2hhcnQgYW5kIG90aGVyIGNvbnN0cnVjdHMuXG4gICAqIFRoZXNlIGNhbiBiZSBvdGhlciBBcGlPYmplY3RzLCBDaGFydHMsIG9yIGN1c3RvbS5cbiAgICpcbiAgICogQHBhcmFtIGRlcGVuZGVuY2llcyB0aGUgZGVwZW5kZW5jaWVzIHRvIGFkZC5cbiAgICovXG4gIHB1YmxpYyBhZGREZXBlbmRlbmN5KC4uLmRlcGVuZGVuY2llczogSUNvbnN0cnVjdFtdKSB7XG4gICAgdGhpcy5ub2RlLmFkZERlcGVuZGVuY3koLi4uZGVwZW5kZW5jaWVzKTtcbiAgfVxuXG4gIC8qKlxuICAgKiBSZW5kZXJzIHRoaXMgY2hhcnQgdG8gYSBzZXQgb2YgS3ViZXJuZXRlcyBKU09OIHJlc291cmNlcy5cbiAgICogQHJldHVybnMgYXJyYXkgb2YgcmVzb3VyY2UgbWFuaWZlc3RzXG4gICAqL1xuICBwdWJsaWMgdG9Kc29uKCk6IGFueVtdIHtcbiAgICByZXR1cm4gQXBwLl9zeW50aENoYXJ0KHRoaXMpO1xuICB9XG5cbiAgLyoqXG4gICAqIFJldHVybnMgYWxsIHRoZSBpbmNsdWRlZCBBUEkgb2JqZWN0cy5cbiAgICovXG4gIGdldCBhcGlPYmplY3RzKCk6IEFwaU9iamVjdFtdIHtcbiAgICByZXR1cm4gdGhpcy5ub2RlLmNoaWxkcmVuLmZpbHRlcigobyk6IG8gaXMgQXBpT2JqZWN0ID0+IG8gaW5zdGFuY2VvZiBBcGlPYmplY3QpO1xuICB9XG59XG4iXX0=