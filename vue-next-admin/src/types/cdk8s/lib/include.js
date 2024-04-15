"use strict";
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
exports.Include = void 0;
const JSII_RTTI_SYMBOL_1 = Symbol.for("jsii.rtti");
const constructs_1 = require("constructs");
const api_object_1 = require("./api-object");
const yaml_1 = require("./yaml");
/**
 * Reads a YAML manifest from a file or a URL and defines all resources as API
 * objects within the defined scope.
 *
 * The names (`metadata.name`) of imported resources will be preserved as-is
 * from the manifest.
 */
class Include extends constructs_1.Construct {
    constructor(scope, id, props) {
        super(scope, id);
        const objects = yaml_1.Yaml.load(props.url);
        let order = 0;
        for (const obj of objects) {
            const objname = obj.metadata?.name ?? `object${order++}`;
            // render an id: name[-kind][-namespace]
            const objid = [objname, obj.kind?.toLowerCase(), obj.metadata?.namespace].filter(x => x).join('-');
            new api_object_1.ApiObject(this, objid, obj);
        }
    }
    /**
     * Returns all the included API objects.
     */
    get apiObjects() {
        return this.node.children.filter((o) => o instanceof api_object_1.ApiObject);
    }
}
exports.Include = Include;
_a = JSII_RTTI_SYMBOL_1;
Include[_a] = { fqn: "cdk8s.Include", version: "2.68.60" };
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiaW5jbHVkZS5qcyIsInNvdXJjZVJvb3QiOiIiLCJzb3VyY2VzIjpbIi4uL3NyYy9pbmNsdWRlLnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7Ozs7O0FBQUEsMkNBQXVDO0FBQ3ZDLDZDQUF5QztBQUN6QyxpQ0FBOEI7QUFXOUI7Ozs7OztHQU1HO0FBQ0gsTUFBYSxPQUFRLFNBQVEsc0JBQVM7SUFDcEMsWUFBWSxLQUFnQixFQUFFLEVBQVUsRUFBRSxLQUFtQjtRQUMzRCxLQUFLLENBQUMsS0FBSyxFQUFFLEVBQUUsQ0FBQyxDQUFDO1FBRWpCLE1BQU0sT0FBTyxHQUFHLFdBQUksQ0FBQyxJQUFJLENBQUMsS0FBSyxDQUFDLEdBQUcsQ0FBQyxDQUFDO1FBRXJDLElBQUksS0FBSyxHQUFHLENBQUMsQ0FBQztRQUNkLEtBQUssTUFBTSxHQUFHLElBQUksT0FBTyxFQUFFLENBQUM7WUFDMUIsTUFBTSxPQUFPLEdBQUcsR0FBRyxDQUFDLFFBQVEsRUFBRSxJQUFJLElBQUksU0FBUyxLQUFLLEVBQUUsRUFBRSxDQUFDO1lBRXpELHdDQUF3QztZQUN4QyxNQUFNLEtBQUssR0FBRyxDQUFDLE9BQU8sRUFBRSxHQUFHLENBQUMsSUFBSSxFQUFFLFdBQVcsRUFBRSxFQUFFLEdBQUcsQ0FBQyxRQUFRLEVBQUUsU0FBUyxDQUFDLENBQUMsTUFBTSxDQUFDLENBQUMsQ0FBQyxFQUFFLENBQUMsQ0FBQyxDQUFDLENBQUMsSUFBSSxDQUFDLEdBQUcsQ0FBQyxDQUFDO1lBQ25HLElBQUksc0JBQVMsQ0FBQyxJQUFJLEVBQUUsS0FBSyxFQUFFLEdBQUcsQ0FBQyxDQUFDO1FBQ2xDLENBQUM7SUFDSCxDQUFDO0lBRUQ7O09BRUc7SUFDSCxJQUFXLFVBQVU7UUFDbkIsT0FBTyxJQUFJLENBQUMsSUFBSSxDQUFDLFFBQVEsQ0FBQyxNQUFNLENBQUMsQ0FBQyxDQUFDLEVBQWtCLEVBQUUsQ0FBQyxDQUFDLFlBQVksc0JBQVMsQ0FBQyxDQUFDO0lBQ2xGLENBQUM7O0FBckJILDBCQXNCQyIsInNvdXJjZXNDb250ZW50IjpbImltcG9ydCB7IENvbnN0cnVjdCB9IGZyb20gJ2NvbnN0cnVjdHMnO1xuaW1wb3J0IHsgQXBpT2JqZWN0IH0gZnJvbSAnLi9hcGktb2JqZWN0JztcbmltcG9ydCB7IFlhbWwgfSBmcm9tICcuL3lhbWwnO1xuXG5leHBvcnQgaW50ZXJmYWNlIEluY2x1ZGVQcm9wcyB7XG4gIC8qKlxuICAgKiBMb2NhbCBmaWxlIHBhdGggb3IgVVJMIHdoaWNoIGluY2x1ZGVzIGEgS3ViZXJuZXRlcyBZQU1MIG1hbmlmZXN0LlxuICAgKlxuICAgKiBAZXhhbXBsZSBteW1hbmlmZXN0LnlhbWxcbiAgICovXG4gIHJlYWRvbmx5IHVybDogc3RyaW5nO1xufVxuXG4vKipcbiAqIFJlYWRzIGEgWUFNTCBtYW5pZmVzdCBmcm9tIGEgZmlsZSBvciBhIFVSTCBhbmQgZGVmaW5lcyBhbGwgcmVzb3VyY2VzIGFzIEFQSVxuICogb2JqZWN0cyB3aXRoaW4gdGhlIGRlZmluZWQgc2NvcGUuXG4gKlxuICogVGhlIG5hbWVzIChgbWV0YWRhdGEubmFtZWApIG9mIGltcG9ydGVkIHJlc291cmNlcyB3aWxsIGJlIHByZXNlcnZlZCBhcy1pc1xuICogZnJvbSB0aGUgbWFuaWZlc3QuXG4gKi9cbmV4cG9ydCBjbGFzcyBJbmNsdWRlIGV4dGVuZHMgQ29uc3RydWN0IHtcbiAgY29uc3RydWN0b3Ioc2NvcGU6IENvbnN0cnVjdCwgaWQ6IHN0cmluZywgcHJvcHM6IEluY2x1ZGVQcm9wcykge1xuICAgIHN1cGVyKHNjb3BlLCBpZCk7XG5cbiAgICBjb25zdCBvYmplY3RzID0gWWFtbC5sb2FkKHByb3BzLnVybCk7XG5cbiAgICBsZXQgb3JkZXIgPSAwO1xuICAgIGZvciAoY29uc3Qgb2JqIG9mIG9iamVjdHMpIHtcbiAgICAgIGNvbnN0IG9iam5hbWUgPSBvYmoubWV0YWRhdGE/Lm5hbWUgPz8gYG9iamVjdCR7b3JkZXIrK31gO1xuXG4gICAgICAvLyByZW5kZXIgYW4gaWQ6IG5hbWVbLWtpbmRdWy1uYW1lc3BhY2VdXG4gICAgICBjb25zdCBvYmppZCA9IFtvYmpuYW1lLCBvYmoua2luZD8udG9Mb3dlckNhc2UoKSwgb2JqLm1ldGFkYXRhPy5uYW1lc3BhY2VdLmZpbHRlcih4ID0+IHgpLmpvaW4oJy0nKTtcbiAgICAgIG5ldyBBcGlPYmplY3QodGhpcywgb2JqaWQsIG9iaik7XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIFJldHVybnMgYWxsIHRoZSBpbmNsdWRlZCBBUEkgb2JqZWN0cy5cbiAgICovXG4gIHB1YmxpYyBnZXQgYXBpT2JqZWN0cygpOiBBcGlPYmplY3RbXSB7XG4gICAgcmV0dXJuIHRoaXMubm9kZS5jaGlsZHJlbi5maWx0ZXIoKG8pOiBvIGlzIEFwaU9iamVjdCA9PiBvIGluc3RhbmNlb2YgQXBpT2JqZWN0KTtcbiAgfVxufVxuIl19