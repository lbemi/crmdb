"use strict";
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
exports.ApiObject = void 0;
const JSII_RTTI_SYMBOL_1 = Symbol.for("jsii.rtti");
const constructs_1 = require("constructs");
const _util_1 = require("./_util");
const chart_1 = require("./chart");
const json_patch_1 = require("./json-patch");
const metadata_1 = require("./metadata");
const resolve_1 = require("./resolve");
const API_OBJECT_SYMBOL = Symbol.for('cdk8s.ApiObject');
class ApiObject extends constructs_1.Construct {
    /**
     * Return whether the given object is an `ApiObject`.
     *
     * We do attribute detection since we can't reliably use 'instanceof'.
  
     * @param o The object to check
     */
    static isApiObject(o) {
        return o !== null && typeof o === 'object' && API_OBJECT_SYMBOL in o;
    }
    /**
     * Implements `instanceof ApiObject` using the more reliable `ApiObject.isApiObject` static method
     *
     * @param o The object to check
     * @internal
     */
    static [(_a = JSII_RTTI_SYMBOL_1, Symbol.hasInstance)](o) {
        return ApiObject.isApiObject(o);
    }
    /**
     * Returns the `ApiObject` named `Resource` which is a child of the given
     * construct. If `c` is an `ApiObject`, it is returned directly. Throws an
     * exception if the construct does not have a child named `Default` _or_ if
     * this child is not an `ApiObject`.
     *
     * @param c The higher-level construct
     */
    static of(c) {
        if (c instanceof ApiObject) {
            return c;
        }
        const child = c.node.defaultChild;
        if (!child) {
            throw new Error(`cannot find a (direct or indirect) child of type ApiObject for construct ${c.node.path}`);
        }
        return ApiObject.of(child);
    }
    /**
     * Defines an API object.
     *
     * @param scope the construct scope
     * @param id namespace
     * @param props options
     */
    constructor(scope, id, props) {
        super(scope, id);
        this.props = props;
        this.patches = new Array();
        this.chart = chart_1.Chart.of(this);
        this.kind = props.kind;
        this.apiVersion = props.apiVersion;
        this.apiGroup = parseApiGroup(this.apiVersion);
        this.name = props.metadata?.name ?? this.chart.generateObjectName(this);
        this.metadata = new metadata_1.ApiObjectMetadataDefinition({
            name: this.name,
            // user defined values
            ...props.metadata,
            namespace: props.metadata?.namespace ?? this.chart.namespace,
            labels: {
                ...this.chart.labels,
                ...props.metadata?.labels,
            },
            apiObject: this,
        });
        Object.defineProperty(this, API_OBJECT_SYMBOL, { value: true });
    }
    /**
     * Create a dependency between this ApiObject and other constructs.
     * These can be other ApiObjects, Charts, or custom.
     *
     * @param dependencies the dependencies to add.
     */
    addDependency(...dependencies) {
        this.node.addDependency(...dependencies);
    }
    /**
     * Applies a set of RFC-6902 JSON-Patch operations to the manifest
     * synthesized for this API object.
     *
     * @param ops The JSON-Patch operations to apply.
     *
     * @example
     *
     *   kubePod.addJsonPatch(JsonPatch.replace('/spec/enableServiceLinks', true));
     *
     */
    addJsonPatch(...ops) {
        this.patches.push(...ops);
    }
    /**
     * Renders the object to Kubernetes JSON.
     *
     * To disable sorting of dictionary keys in output object set the
     * `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
     */
    toJson() {
        try {
            const data = {
                ...this.props,
                metadata: this.metadata.toJson(),
            };
            const sortKeys = process.env.CDK8S_DISABLE_SORT ? false : true;
            const json = (0, _util_1.sanitizeValue)((0, resolve_1.resolve)([], data, this), { sortKeys });
            const patched = json_patch_1.JsonPatch.apply(json, ...this.patches);
            // reorder top-level keys so that we first have "apiVersion", "kind" and
            // "metadata" and then all the rest
            const result = {};
            const orderedKeys = ['apiVersion', 'kind', 'metadata', ...Object.keys(patched)];
            for (const k of orderedKeys) {
                if (k in patched) {
                    result[k] = patched[k];
                }
            }
            return result;
        }
        catch (e) {
            throw new Error(`Failed serializing construct at path '${this.node.path}' with name '${this.name}': ${e}`);
        }
    }
}
exports.ApiObject = ApiObject;
ApiObject[_a] = { fqn: "cdk8s.ApiObject", version: "2.68.60" };
function parseApiGroup(apiVersion) {
    const v = apiVersion.split('/');
    // no group means "core"
    // https://kubernetes.io/docs/reference/using-api/api-overview/#api-groups
    if (v.length === 1) {
        return 'core';
    }
    if (v.length === 2) {
        return v[0];
    }
    throw new Error(`invalid apiVersion ${apiVersion}, expecting GROUP/VERSION. See https://kubernetes.io/docs/reference/using-api/api-overview/#api-groups`);
}
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiYXBpLW9iamVjdC5qcyIsInNvdXJjZVJvb3QiOiIiLCJzb3VyY2VzIjpbIi4uL3NyYy9hcGktb2JqZWN0LnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7Ozs7O0FBQUEsMkNBQW1EO0FBQ25ELG1DQUF3QztBQUN4QyxtQ0FBZ0M7QUFDaEMsNkNBQXlDO0FBQ3pDLHlDQUE0RTtBQUM1RSx1Q0FBb0M7QUE0Q3BDLE1BQU0saUJBQWlCLEdBQUcsTUFBTSxDQUFDLEdBQUcsQ0FBQyxpQkFBaUIsQ0FBQyxDQUFDO0FBRXhELE1BQWEsU0FBVSxTQUFRLHNCQUFTO0lBRXRDOzs7Ozs7T0FNRztJQUNILE1BQU0sQ0FBQyxXQUFXLENBQUMsQ0FBTTtRQUN2QixPQUFPLENBQUMsS0FBSyxJQUFJLElBQUksT0FBTyxDQUFDLEtBQUssUUFBUSxJQUFJLGlCQUFpQixJQUFJLENBQUMsQ0FBQztJQUN2RSxDQUFDO0lBRUQ7Ozs7O09BS0c7SUFDSCxNQUFNLENBQUMsMkJBQUMsTUFBTSxDQUFDLFdBQVcsRUFBQyxDQUFDLENBQVU7UUFDcEMsT0FBTyxTQUFTLENBQUMsV0FBVyxDQUFDLENBQUMsQ0FBQyxDQUFDO0lBQ2xDLENBQUM7SUFDRDs7Ozs7OztPQU9HO0lBQ0ksTUFBTSxDQUFDLEVBQUUsQ0FBQyxDQUFhO1FBQzVCLElBQUksQ0FBQyxZQUFZLFNBQVMsRUFBRSxDQUFDO1lBQzNCLE9BQU8sQ0FBQyxDQUFDO1FBQ1gsQ0FBQztRQUVELE1BQU0sS0FBSyxHQUFHLENBQUMsQ0FBQyxJQUFJLENBQUMsWUFBWSxDQUFDO1FBQ2xDLElBQUksQ0FBQyxLQUFLLEVBQUUsQ0FBQztZQUNYLE1BQU0sSUFBSSxLQUFLLENBQUMsNEVBQTRFLENBQUMsQ0FBQyxJQUFJLENBQUMsSUFBSSxFQUFFLENBQUMsQ0FBQztRQUM3RyxDQUFDO1FBRUQsT0FBTyxTQUFTLENBQUMsRUFBRSxDQUFDLEtBQUssQ0FBQyxDQUFDO0lBQzdCLENBQUM7SUEwQ0Q7Ozs7OztPQU1HO0lBQ0gsWUFBWSxLQUFn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QUFDO2dCQUN6QixDQUFDO1lBQ0gsQ0FBQztZQUVELE9BQU8sTUFBTSxDQUFDO1FBQ2hCLENBQUM7UUFBQyxPQUFPLENBQUMsRUFBRSxDQUFDO1lBQ1gsTUFBTSxJQUFJLEtBQUssQ0FBQyx5Q0FBeUMsSUFBSSxDQUFDLElBQUksQ0FBQyxJQUFJLGdCQUFnQixJQUFJLENBQUMsSUFBSSxNQUFNLENBQUMsRUFBRSxDQUFDLENBQUM7UUFDN0csQ0FBQztJQUNILENBQUM7O0FBOUtILDhCQStLQzs7QUFFRCxTQUFTLGFBQWEsQ0FBQyxVQUFrQjtJQUN2QyxNQUFNLENBQUMsR0FBRyxVQUFVLENBQUMsS0FBSyxDQUFDLEdBQUcsQ0FBQyxDQUFDO0lBRWhDLHdCQUF3QjtJQUN4QiwwRUFBMEU7SUFDMUUsSUFBSSxDQUFDLENBQUMsTUFBTSxLQUFLLENBQUMsRUFBRSxDQUFDO1FBQ25CLE9BQU8sTUFBTSxDQUFDO0lBQ2hCLENBQUM7SUFFRCxJQUFJLENBQUMsQ0FBQyxNQUFNLEtBQUssQ0FBQyxFQUFFLENBQUM7UUFDbkIsT0FBTyxDQUFDLENBQUMsQ0FBQyxDQUFDLENBQUM7SUFDZCxDQUFDO0lBRUQsTUFBTSxJQUFJLEtBQUssQ0FBQyxzQkFBc0IsVUFBVSx3R0FBd0csQ0FBQyxDQUFDO0FBQzVKLENBQUMiLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgeyBDb25zdHJ1Y3QsIElDb25zdHJ1Y3QgfSBmcm9tICdjb25zdHJ1Y3RzJztcbmltcG9ydCB7IHNhbml0aXplVmFsdWUgfSBmcm9tICcuL191dGlsJztcbmltcG9ydCB7IENoYXJ0IH0gZnJvbSAnLi9jaGFydCc7XG5pbXBvcnQgeyBKc29uUGF0Y2ggfSBmcm9tICcuL2pzb24tcGF0Y2gnO1xuaW1wb3J0IHsgQXBpT2JqZWN0TWV0YWRhdGEsIEFwaU9iamVjdE1ldGFkYXRhRGVmaW5pdGlvbiB9IGZyb20gJy4vbWV0YWRhdGEnO1xuaW1wb3J0IHsgcmVzb2x2ZSB9IGZyb20gJy4vcmVzb2x2ZSc7XG5cbi8qKlxuICogT3B0aW9ucyBmb3IgZGVmaW5pbmcgQVBJIG9iamVjdHMuXG4gKi9cbmV4cG9ydCBpbnRlcmZhY2UgQXBpT2JqZWN0UHJvcHMge1xuICAvKipcbiAgICogT2JqZWN0IG1ldGFkYXRhLlxuICAgKlxuICAgKiBJZiBgbmFtZWAgaXMgbm90IHNwZWNpZmllZCwgYW4gYXBwLXVuaXF1ZSBuYW1lIHdpbGwgYmUgYWxsb2NhdGVkIGJ5IHRoZVxuICAgKiBmcmFtZXdvcmsgYmFzZWQgb24gdGhlIHBhdGggb2YgdGhlIGNvbnN0cnVjdCB3aXRoaW4gdGhlcyBjb25zdHJ1Y3QgdHJlZS5cbiAgICovXG4gIHJlYWRvbmx5IG1ldGFkYXRhPzogQXBpT2JqZWN0TWV0YWRhdGE7XG5cbiAgLyoqXG4gICAqIEFQSSB2ZXJzaW9uLlxuICAgKi9cbiAgcmVhZG9ubHkgYXBpVmVyc2lvbjogc3RyaW5nO1xuXG4gIC8qKlxuICAgKiBSZXNvdXJjZSBraW5kLlxuICAgKi9cbiAgcmVhZG9ubHkga2luZDogc3RyaW5nO1xuXG4gIC8qKlxuICAgKiBBZGRpdGlvbmFsIGF0dHJpYnV0ZXMgZm9yIHRoaXMgQVBJIG9iamVjdC5cbiAgICogQGpzaWkgaWdub3JlXG4gICAqIEBzZWUgaHR0cHM6Ly9naXRodWIuY29tL2NkazhzLXRlYW0vY2RrOHMtY29yZS9pc3N1ZXMvMTI5N1xuICAgKi9cbiAgcmVhZG9ubHkgW2tleTogc3RyaW5nXTogYW55O1xufVxuXG5leHBvcnQgaW50ZXJmYWNlIEdyb3VwVmVyc2lvbktpbmQge1xuICAvKipcbiAgICogVGhlIG9iamVjdCdzIEFQSSB2ZXJzaW9uIChlLmcuIGBhdXRob3JpemF0aW9uLms4cy5pby92MWApXG4gICAqL1xuICByZWFkb25seSBhcGlWZXJzaW9uOiBzdHJpbmc7XG5cbiAgLyoqXG4gICAqIFRoZSBvYmplY3Qga2luZC5cbiAgICovXG4gIHJlYWRvbmx5IGtpbmQ6IHN0cmluZztcbn1cblxuY29uc3QgQVBJX09CSkVDVF9TWU1CT0wgPSBTeW1ib2wuZm9yKCdjZGs4cy5BcGlPYmplY3QnKTtcblxuZXhwb3J0IGNsYXNzIEFwaU9iamVjdCBleHRlbmRzIENvbnN0cnVjdCB7XG5cbiAgLyoqXG4gICAqIFJldHVybiB3aGV0aGVyIHRoZSBnaXZlbiBvYmplY3QgaXMgYW4gYEFwaU9iamVjdGAuXG4gICAqXG4gICAqIFdlIGRvIGF0dHJpYnV0ZSBkZXRlY3Rpb24gc2luY2Ugd2UgY2FuJ3QgcmVsaWFibHkgdXNlICdpbnN0YW5jZW9mJy5cblxuICAgKiBAcGFyYW0gbyBUaGUgb2JqZWN0IHRvIGNoZWNrXG4gICAqL1xuICBzdGF0aWMgaXNBcGlPYmplY3QobzogYW55KTogbyBpcyBBcGlPYmplY3Qge1xuICAgIHJldHVybiBvICE9PSBudWxsICYmIHR5cGVvZiBvID09PSAnb2JqZWN0JyAmJiBBUElfT0JKRUNUX1NZTUJPTCBpbiBvO1xuICB9XG5cbiAgLyoqXG4gICAqIEltcGxlbWVudHMgYGluc3RhbmNlb2YgQXBpT2JqZWN0YCB1c2luZyB0aGUgbW9yZSByZWxpYWJsZSBgQXBpT2JqZWN0LmlzQXBpT2JqZWN0YCBzdGF0aWMgbWV0aG9kXG4gICAqXG4gICAqIEBwYXJhbSBvIFRoZSBvYmplY3QgdG8gY2hlY2tcbiAgICogQGludGVybmFsXG4gICAqL1xuICBzdGF0aWMgW1N5bWJvbC5oYXNJbnN0YW5jZV0obzogdW5rbm93bikge1xuICAgIHJldHVybiBBcGlPYmplY3QuaXNBcGlPYmplY3Qobyk7XG4gIH1cbiAgLyoqXG4gICAqIFJldHVybnMgdGhlIGBBcGlPYmplY3RgIG5hbWVkIGBSZXNvdXJjZWAgd2hpY2ggaXMgYSBjaGlsZCBvZiB0aGUgZ2l2ZW5cbiAgICogY29uc3RydWN0LiBJZiBgY2AgaXMgYW4gYEFwaU9iamVjdGAsIGl0IGlzIHJldHVybmVkIGRpcmVjdGx5LiBUaHJvd3MgYW5cbiAgICogZXhjZXB0aW9uIGlmIHRoZSBjb25zdHJ1Y3QgZG9lcyBub3QgaGF2ZSBhIGNoaWxkIG5hbWVkIGBEZWZhdWx0YCBfb3JfIGlmXG4gICAqIHRoaXMgY2hpbGQgaXMgbm90IGFuIGBBcGlPYmplY3RgLlxuICAgKlxuICAgKiBAcGFyYW0gYyBUaGUgaGlnaGVyLWxldmVsIGNvbnN0cnVjdFxuICAgKi9cbiAgcHVibGljIHN0YXRpYyBvZihjOiBJQ29uc3RydWN0KTogQXBpT2JqZWN0IHtcbiAgICBpZiAoYyBpbnN0YW5jZW9mIEFwaU9iamVjdCkge1xuICAgICAgcmV0dXJuIGM7XG4gICAgfVxuXG4gICAgY29uc3QgY2hpbGQgPSBjLm5vZGUuZGVmYXVsdENoaWxkO1xuICAgIGlmICghY2hpbGQpIHtcbiAgICAgIHRocm93IG5ldyBFcnJvcihgY2Fubm90IGZpbmQgYSAoZGlyZWN0IG9yIGluZGlyZWN0KSBjaGlsZCBvZiB0eXBlIEFwaU9iamVjdCBmb3IgY29uc3RydWN0ICR7Yy5ub2RlLnBhdGh9YCk7XG4gICAgfVxuXG4gICAgcmV0dXJuIEFwaU9iamVjdC5vZihjaGlsZCk7XG4gIH1cblxuICAvKipcbiAgICogVGhlIG5hbWUgb2YgdGhlIEFQSSBvYmplY3QuXG4gICAqXG4gICAqIElmIGEgbmFtZSBpcyBzcGVjaWZpZWQgaW4gYG1ldGFkYXRhLm5hbWVgIHRoaXMgd2lsbCBiZSB0aGUgbmFtZSByZXR1cm5lZC5cbiAgICogT3RoZXJ3aXNlLCBhIG5hbWUgd2lsbCBiZSBnZW5lcmF0ZWQgYnkgY2FsbGluZ1xuICAgKiBgQ2hhcnQub2YodGhpcykuZ2VuZXJhdGVkT2JqZWN0TmFtZSh0aGlzKWAsIHdoaWNoIGJ5IGRlZmF1bHQgdXNlcyB0aGVcbiAgICogY29uc3RydWN0IHBhdGggdG8gZ2VuZXJhdGUgYSBETlMtY29tcGF0aWJsZSBuYW1lIGZvciB0aGUgcmVzb3VyY2UuXG4gICAqL1xuICBwdWJsaWMgcmVhZG9ubHkgbmFtZTogc3RyaW5nO1xuXG4gIC8qKlxuICAgKiBUaGUgb2JqZWN0J3MgQVBJIHZlcnNpb24gKGUuZy4gYGF1dGhvcml6YXRpb24uazhzLmlvL3YxYClcbiAgICovXG4gIHB1YmxpYyByZWFkb25seSBhcGlWZXJzaW9uOiBzdHJpbmc7XG5cbiAgLyoqXG4gICAqIFRoZSBncm91cCBwb3J0aW9uIG9mIHRoZSBBUEkgdmVyc2lvbiAoZS5nLiBgYXV0aG9yaXphdGlvbi5rOHMuaW9gKVxuICAgKi9cbiAgcHVibGljIHJlYWRvbmx5IGFwaUdyb3VwOiBzdHJpbmc7XG5cbiAgLyoqXG4gICAqIFRoZSBvYmplY3Qga2luZC5cbiAgICovXG4gIHB1YmxpYyByZWFkb25seSBraW5kOiBzdHJpbmc7XG5cbiAgLyoqXG4gICAqIFRoZSBjaGFydCBpbiB3aGljaCB0aGlzIG9iamVjdCBpcyBkZWZpbmVkLlxuICAgKi9cbiAgcHVibGljIHJlYWRvbmx5IGNoYXJ0OiBDaGFydDtcblxuICAvKipcbiAgICogTWV0YWRhdGEgYXNzb2NpYXRlZCB3aXRoIHRoaXMgQVBJIG9iamVjdC5cbiAgICovXG4gIHB1YmxpYyByZWFkb25seSBtZXRhZGF0YTogQXBpT2JqZWN0TWV0YWRhdGFEZWZpbml0aW9uO1xuXG4gIC8qKlxuICAgKiBBIHNldCBvZiBKU09OIHBhdGNoIG9wZXJhdGlvbnMgdG8gYXBwbHkgdG8gdGhlIGRvY3VtZW50IGFmdGVyIHN5bnRoZXNpcy5cbiAgICovXG4gIHByaXZhdGUgcmVhZG9ubHkgcGF0Y2hlczogQXJyYXk8SnNvblBhdGNoPjtcblxuICAvKipcbiAgICogRGVmaW5lcyBhbiBBUEkgb2JqZWN0LlxuICAgKlxuICAgKiBAcGFyYW0gc2NvcGUgdGhlIGNvbnN0cnVjdCBzY29wZVxuICAgKiBAcGFyYW0gaWQgbmFtZXNwYWNlXG4gICAqIEBwYXJhbSBwcm9wcyBvcHRpb25zXG4gICAqL1xuICBjb25zdHJ1Y3RvcihzY29wZTogQ29uc3RydWN0LCBpZDogc3RyaW5nLCBwcml2YXRlIHJlYWRvbmx5IHByb3BzOiBBcGlPYmplY3RQcm9wcykge1xuICAgIHN1cGVyKHNjb3BlLCBpZCk7XG4gICAgdGhpcy5wYXRjaGVzID0gbmV3IEFycmF5PEpzb25QYXRjaD4oKTtcbiAgICB0aGlzLmNoYXJ0ID0gQ2hhcnQub2YodGhpcyk7XG4gICAgdGhpcy5raW5kID0gcHJvcHMua2luZDtcbiAgICB0aGlzLmFwaVZlcnNpb24gPSBwcm9wcy5hcGlWZXJzaW9uO1xuICAgIHRoaXMuYXBpR3JvdXAgPSBwYXJzZUFwaUdyb3VwKHRoaXMuYXBpVmVyc2lvbik7XG5cbiAgICB0aGlzLm5hbWUgPSBwcm9wcy5tZXRhZGF0YT8ubmFtZSA/PyB0aGlzLmNoYXJ0LmdlbmVyYXRlT2JqZWN0TmFtZSh0aGlzKTtcblxuICAgIHRoaXMubWV0YWRhdGEgPSBuZXcgQXBpT2JqZWN0TWV0YWRhdGFEZWZpbml0aW9uKHtcbiAgICAgIG5hbWU6IHRoaXMubmFtZSxcblxuICAgICAgLy8gdXNlciBkZWZpbmVkIHZhbHVlc1xuICAgICAgLi4ucHJvcHMubWV0YWRhdGEsXG5cbiAgICAgIG5hbWVzcGFjZTogcHJvcHMubWV0YWRhdGE/Lm5hbWVzcGFjZSA/PyB0aGlzLmNoYXJ0Lm5hbWVzcGFjZSxcbiAgICAgIGxhYmVsczoge1xuICAgICAgICAuLi50aGlzLmNoYXJ0LmxhYmVscyxcbiAgICAgICAgLi4ucHJvcHMubWV0YWRhdGE/LmxhYmVscyxcbiAgICAgIH0sXG4gICAgICBhcGlPYmplY3Q6IHRoaXMsXG4gICAgfSk7XG5cbiAgICBPYmplY3QuZGVmaW5lUHJvcGVydHkodGhpcywgQVBJX09CSkVDVF9TWU1CT0wsIHsgdmFsdWU6IHRydWUgfSk7XG4gIH1cblxuICAvKipcbiAgICogQ3JlYXRlIGEgZGVwZW5kZW5jeSBiZXR3ZWVuIHRoaXMgQXBpT2JqZWN0IGFuZCBvdGhlciBjb25zdHJ1Y3RzLlxuICAgKiBUaGVzZSBjYW4gYmUgb3RoZXIgQXBpT2JqZWN0cywgQ2hhcnRzLCBvciBjdXN0b20uXG4gICAqXG4gICAqIEBwYXJhbSBkZXBlbmRlbmNpZXMgdGhlIGRlcGVuZGVuY2llcyB0byBhZGQuXG4gICAqL1xuICBwdWJsaWMgYWRkRGVwZW5kZW5jeSguLi5kZXBlbmRlbmNpZXM6IElDb25zdHJ1Y3RbXSkge1xuICAgIHRoaXMubm9kZS5hZGREZXBlbmRlbmN5KC4uLmRlcGVuZGVuY2llcyk7XG4gIH1cblxuICAvKipcbiAgICogQXBwbGllcyBhIHNldCBvZiBSRkMtNjkwMiBKU09OLVBhdGNoIG9wZXJhdGlvbnMgdG8gdGhlIG1hbmlmZXN0XG4gICAqIHN5bnRoZXNpemVkIGZvciB0aGlzIEFQSSBvYmplY3QuXG4gICAqXG4gICAqIEBwYXJhbSBvcHMgVGhlIEpTT04tUGF0Y2ggb3BlcmF0aW9ucyB0byBhcHBseS5cbiAgICpcbiAgICogQGV4YW1wbGVcbiAgICpcbiAgICogICBrdWJlUG9kLmFkZEpzb25QYXRjaChKc29uUGF0Y2gucmVwbGFjZSgnL3NwZWMvZW5hYmxlU2VydmljZUxpbmtzJywgdHJ1ZSkpO1xuICAgKlxuICAgKi9cbiAgcHVibGljIGFkZEpzb25QYXRjaCguLi5vcHM6IEpzb25QYXRjaFtdKSB7XG4gICAgdGhpcy5wYXRjaGVzLnB1c2goLi4ub3BzKTtcbiAgfVxuXG4gIC8qKlxuICAgKiBSZW5kZXJzIHRoZSBvYmplY3QgdG8gS3ViZXJuZXRlcyBKU09OLlxuICAgKlxuICAgKiBUbyBkaXNhYmxlIHNvcnRpbmcgb2YgZGljdGlvbmFyeSBrZXlzIGluIG91dHB1dCBvYmplY3Qgc2V0IHRoZVxuICAgKiBgQ0RLOFNfRElTQUJMRV9TT1JUYCBlbnZpcm9ubWVudCB2YXJpYWJsZSB0byBhbnkgbm9uLWVtcHR5IHZhbHVlLlxuICAgKi9cbiAgcHVibGljIHRvSnNvbigpOiBhbnkge1xuXG4gICAgdHJ5IHtcbiAgICAgIGNvbnN0IGRhdGE6IGFueSA9IHtcbiAgICAgICAgLi4udGhpcy5wcm9wcyxcbiAgICAgICAgbWV0YWRhdGE6IHRoaXMubWV0YWRhdGEudG9Kc29uKCksXG4gICAgICB9O1xuXG4gICAgICBjb25zdCBzb3J0S2V5cyA9IHByb2Nlc3MuZW52LkNESzhTX0RJU0FCTEVfU09SVCA/IGZhbHNlIDogdHJ1ZTtcbiAgICAgIGNvbnN0IGpzb24gPSBzYW5pdGl6ZVZhbHVlKHJlc29sdmUoW10sIGRhdGEsIHRoaXMpLCB7IHNvcnRLZXlzIH0pO1xuICAgICAgY29uc3QgcGF0Y2hlZCA9IEpzb25QYXRjaC5hcHBseShqc29uLCAuLi50aGlzLnBhdGNoZXMpO1xuXG4gICAgICAvLyByZW9yZGVyIHRvcC1sZXZlbCBrZXlzIHNvIHRoYXQgd2UgZmlyc3QgaGF2ZSBcImFwaVZlcnNpb25cIiwgXCJraW5kXCIgYW5kXG4gICAgICAvLyBcIm1ldGFkYXRhXCIgYW5kIHRoZW4gYWxsIHRoZSByZXN0XG4gICAgICBjb25zdCByZXN1bHQ6IGFueSA9IHt9O1xuICAgICAgY29uc3Qgb3JkZXJlZEtleXMgPSBbJ2FwaVZlcnNpb24nLCAna2luZCcsICdtZXRhZGF0YScsIC4uLk9iamVjdC5rZXlzKHBhdGNoZWQpXTtcbiAgICAgIGZvciAoY29uc3QgayBvZiBvcmRlcmVkS2V5cykge1xuICAgICAgICBpZiAoayBpbiBwYXRjaGVkKSB7XG4gICAgICAgICAgcmVzdWx0W2tdID0gcGF0Y2hlZFtrXTtcbiAgICAgICAgfVxuICAgICAgfVxuXG4gICAgICByZXR1cm4gcmVzdWx0O1xuICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgIHRocm93IG5ldyBFcnJvcihgRmFpbGVkIHNlcmlhbGl6aW5nIGNvbnN0cnVjdCBhdCBwYXRoICcke3RoaXMubm9kZS5wYXRofScgd2l0aCBuYW1lICcke3RoaXMubmFtZX0nOiAke2V9YCk7XG4gICAgfVxuICB9XG59XG5cbmZ1bmN0aW9uIHBhcnNlQXBpR3JvdXAoYXBpVmVyc2lvbjogc3RyaW5nKSB7XG4gIGNvbnN0IHYgPSBhcGlWZXJzaW9uLnNwbGl0KCcvJyk7XG5cbiAgLy8gbm8gZ3JvdXAgbWVhbnMgXCJjb3JlXCJcbiAgLy8gaHR0cHM6Ly9rdWJlcm5ldGVzLmlvL2RvY3MvcmVmZXJlbmNlL3VzaW5nLWFwaS9hcGktb3ZlcnZpZXcvI2FwaS1ncm91cHNcbiAgaWYgKHYubGVuZ3RoID09PSAxKSB7XG4gICAgcmV0dXJuICdjb3JlJztcbiAgfVxuXG4gIGlmICh2Lmxlbmd0aCA9PT0gMikge1xuICAgIHJldHVybiB2WzBdO1xuICB9XG5cbiAgdGhyb3cgbmV3IEVycm9yKGBpbnZhbGlkIGFwaVZlcnNpb24gJHthcGlWZXJzaW9ufSwgZXhwZWN0aW5nIEdST1VQL1ZFUlNJT04uIFNlZSBodHRwczovL2t1YmVybmV0ZXMuaW8vZG9jcy9yZWZlcmVuY2UvdXNpbmctYXBpL2FwaS1vdmVydmlldy8jYXBpLWdyb3Vwc2ApO1xufVxuIl19