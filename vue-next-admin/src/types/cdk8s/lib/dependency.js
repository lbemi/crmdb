"use strict";
var _a, _b;
Object.defineProperty(exports, "__esModule", { value: true });
exports.DependencyVertex = exports.DependencyGraph = void 0;
const JSII_RTTI_SYMBOL_1 = Symbol.for("jsii.rtti");
/**
 * Represents the dependency graph for a given Node.
 *
 * This graph includes the dependency relationships between all nodes in the
 * node (construct) sub-tree who's root is this Node.
 *
 * Note that this means that lonely nodes (no dependencies and no dependants) are also included in this graph as
 * childless children of the root node of the graph.
 *
 * The graph does not include cross-scope dependencies. That is, if a child on the current scope depends on a node
 * from a different scope, that relationship is not represented in this graph.
 *
 */
class DependencyGraph {
    constructor(node) {
        this._fosterParent = new DependencyVertex();
        const nodes = {};
        function putVertex(construct) {
            nodes[construct.node.path] = new DependencyVertex(construct);
        }
        function getVertex(construct) {
            return nodes[construct.node.path];
        }
        // create all vertices of the graph.
        for (const n of node.findAll()) {
            putVertex(n);
        }
        const deps = [];
        for (const child of node.findAll()) {
            for (const dep of child.node.dependencies) {
                deps.push({ source: child, target: dep });
            }
        }
        // create all the edges of the graph.
        for (const dep of deps) {
            if (!getVertex(dep.target)) {
                // dont cross scope boundaries.
                // since charts only renders its own children, this is ok and
                // has the benefit of simplifying the graph. we should reconsider this behavior when moving
                // to a more general purpose use-case.
                continue;
            }
            const sourceDepNode = getVertex(dep.source);
            const targetDepNode = getVertex(dep.target);
            sourceDepNode.addChild(targetDepNode);
        }
        // create the root.
        for (const n of Object.values(nodes)) {
            if (n.inbound.length === 0) {
                // orphans are dependency roots. lets adopt them!
                this._fosterParent.addChild(n);
            }
        }
    }
    /**
     * Returns the root of the graph.
     *
     * Note that this vertex will always have `null` as its `.value` since it is an artifical root
     * that binds all the connected spaces of the graph.
     */
    get root() {
        return this._fosterParent;
    }
    /**
     * @see Vertex.topology()
     */
    topology() {
        return this._fosterParent.topology();
    }
}
exports.DependencyGraph = DependencyGraph;
_a = JSII_RTTI_SYMBOL_1;
DependencyGraph[_a] = { fqn: "cdk8s.DependencyGraph", version: "2.68.60" };
/**
 * Represents a vertex in the graph.
 *
 * The value of each vertex is an `IConstruct` that is accessible via the `.value` getter.
 */
class DependencyVertex {
    constructor(value = undefined) {
        this._children = new Set();
        this._parents = new Set();
        this._value = value;
    }
    /**
     * Returns the IConstruct this graph vertex represents.
     *
     * `null` in case this is the root of the graph.
     */
    get value() {
        return this._value;
    }
    /**
     * Returns the children of the vertex (i.e dependencies)
     */
    get outbound() {
        return Array.from(this._children);
    }
    /**
     * Returns the parents of the vertex (i.e dependants)
     */
    get inbound() {
        return Array.from(this._parents);
    }
    /**
     * Returns a topologically sorted array of the constructs in the sub-graph.
     */
    topology() {
        const found = new Set();
        const topology = [];
        function visit(n) {
            for (const c of n.outbound) {
                visit(c);
            }
            if (!found.has(n)) {
                topology.push(n);
                found.add(n);
            }
        }
        visit(this);
        return topology.filter(d => d.value).map(d => d.value);
    }
    /**
     * Adds a vertex as a dependency of the current node.
     * Also updates the parents of `dep`, so that it contains this node as a parent.
     *
     * This operation will fail in case it creates a cycle in the graph.
     *
     * @param dep The dependency
     */
    addChild(dep) {
        const cycle = dep.findRoute(this);
        if (cycle.length !== 0) {
            cycle.push(dep);
            throw new Error(`Dependency cycle detected: ${cycle.filter(d => d.value).map(d => d.value.node.path).join(' => ')}`);
        }
        this._children.add(dep);
        dep.addParent(this);
    }
    addParent(dep) {
        this._parents.add(dep);
    }
    findRoute(dst) {
        const route = [];
        visit(this);
        return route;
        function visit(n) {
            route.push(n);
            let found = false;
            for (const c of n.outbound) {
                if (c === dst) {
                    route.push(c);
                    return true;
                }
                found = visit(c);
            }
            if (!found) {
                route.pop();
            }
            return found;
        }
    }
}
exports.DependencyVertex = DependencyVertex;
_b = JSII_RTTI_SYMBOL_1;
DependencyVertex[_b] = { fqn: "cdk8s.DependencyVertex", version: "2.68.60" };
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiZGVwZW5kZW5jeS5qcyIsInNvdXJjZVJvb3QiOiIiLCJzb3VyY2VzIjpbIi4uL3NyYy9kZXBlbmRlbmN5LnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7Ozs7O0FBR0E7Ozs7Ozs7Ozs7OztHQVlHO0FBQ0gsTUFBYSxlQUFlO0lBSTFCLFlBQVksSUFBVTtRQUVwQixJQUFJLENBQUMsYUFBYSxHQUFHLElBQUksZ0JBQWdCLEVBQUUsQ0FBQztRQUU1QyxNQUFNLEtBQUssR0FBcUMsRUFBRSxDQUFDO1FBRW5ELFNBQVMsU0FBUyxDQUFDLFNBQXFCO1lBQ3RDLEtBQUssQ0FBQyxTQUFTLENBQUMsSUFBSSxDQUFDLElBQUksQ0FBQyxHQUFHLElBQUksZ0JBQWdCLENBQUMsU0FBUyxDQUFDLENBQUM7UUFDL0QsQ0FBQztRQUVELFNBQVMsU0FBUyxDQUFDLFNBQXFCO1lBQ3RDLE9BQU8sS0FBSyxDQUFDLFNBQVMsQ0FBQyxJQUFJLENBQUMsSUFBSSxDQUFDLENBQUM7UUFDcEMsQ0FBQztRQUVELG9DQUFvQztRQUNwQyxLQUFLLE1BQU0sQ0FBQyxJQUFJLElBQUksQ0FBQyxPQUFPLEVBQUUsRUFBRSxDQUFDO1lBQy9CLFNBQVMsQ0FBQyxDQUFDLENBQUMsQ0FBQztRQUNmLENBQUM7UUFFRCxNQUFNLElBQUksR0FBRyxFQUFFLENBQUM7UUFDaEIsS0FBSyxNQUFNLEtBQUssSUFBSSxJQUFJLENBQUMsT0FBTyxFQUFFLEVBQUUsQ0FBQztZQUNuQyxLQUFLLE1BQU0sR0FBRyxJQUFJLEtBQUssQ0FBQyxJQUFJLENBQUMsWUFBWSxFQUFFLENBQUM7Z0JBQzFDLElBQUksQ0FBQyxJQUFJLENBQUMsRUFBRSxNQUFNLEVBQUUsS0FBSyxFQUFFLE1BQU0sRUFBRSxHQUFHLEVBQUUsQ0FBQyxDQUFDO1lBQzVDLENBQUM7UUFDSCxDQUFDO1FBRUQscUNBQXFDO1FBQ3JDLEtBQUssTUFBTSxHQUFHLElBQUksSUFBSSxFQUFFLENBQUM7WUFFdkIsSUFBSSxDQUFDLFNBQVMsQ0FBQyxH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Q2QsQ0FBQztZQUNELE9BQU8sS0FBSyxDQUFDO1FBRWYsQ0FBQztJQUVILENBQUM7O0FBeEdILDRDQXlHQyIsInNvdXJjZXNDb250ZW50IjpbImltcG9ydCB7IE5vZGUsIElDb25zdHJ1Y3QgfSBmcm9tICdjb25zdHJ1Y3RzJztcblxuXG4vKipcbiAqIFJlcHJlc2VudHMgdGhlIGRlcGVuZGVuY3kgZ3JhcGggZm9yIGEgZ2l2ZW4gTm9kZS5cbiAqXG4gKiBUaGlzIGdyYXBoIGluY2x1ZGVzIHRoZSBkZXBlbmRlbmN5IHJlbGF0aW9uc2hpcHMgYmV0d2VlbiBhbGwgbm9kZXMgaW4gdGhlXG4gKiBub2RlIChjb25zdHJ1Y3QpIHN1Yi10cmVlIHdobydzIHJvb3QgaXMgdGhpcyBOb2RlLlxuICpcbiAqIE5vdGUgdGhhdCB0aGlzIG1lYW5zIHRoYXQgbG9uZWx5IG5vZGVzIChubyBkZXBlbmRlbmNpZXMgYW5kIG5vIGRlcGVuZGFudHMpIGFyZSBhbHNvIGluY2x1ZGVkIGluIHRoaXMgZ3JhcGggYXNcbiAqIGNoaWxkbGVzcyBjaGlsZHJlbiBvZiB0aGUgcm9vdCBub2RlIG9mIHRoZSBncmFwaC5cbiAqXG4gKiBUaGUgZ3JhcGggZG9lcyBub3QgaW5jbHVkZSBjcm9zcy1zY29wZSBkZXBlbmRlbmNpZXMuIFRoYXQgaXMsIGlmIGEgY2hpbGQgb24gdGhlIGN1cnJlbnQgc2NvcGUgZGVwZW5kcyBvbiBhIG5vZGVcbiAqIGZyb20gYSBkaWZmZXJlbnQgc2NvcGUsIHRoYXQgcmVsYXRpb25zaGlwIGlzIG5vdCByZXByZXNlbnRlZCBpbiB0aGlzIGdyYXBoLlxuICpcbiAqL1xuZXhwb3J0IGNsYXNzIERlcGVuZGVuY3lHcmFwaCB7XG5cbiAgcHJpdmF0ZSByZWFkb25seSBfZm9zdGVyUGFyZW50OiBEZXBlbmRlbmN5VmVydGV4O1xuXG4gIGNvbnN0cnVjdG9yKG5vZGU6IE5vZGUpIHtcblxuICAgIHRoaXMuX2Zvc3RlclBhcmVudCA9IG5ldyBEZXBlbmRlbmN5VmVydGV4KCk7XG5cbiAgICBjb25zdCBub2RlczogUmVjb3JkPHN0cmluZywgRGVwZW5kZW5jeVZlcnRleD4gPSB7fTtcblxuICAgIGZ1bmN0aW9uIHB1dFZlcnRleChjb25zdHJ1Y3Q6IElDb25zdHJ1Y3QpIHtcbiAgICAgIG5vZGVzW2NvbnN0cnVjdC5ub2RlLnBhdGhdID0gbmV3IERlcGVuZGVuY3lWZXJ0ZXgoY29uc3RydWN0KTtcbiAgICB9XG5cbiAgICBmdW5jdGlvbiBnZXRWZXJ0ZXgoY29uc3RydWN0OiBJQ29uc3RydWN0KTogRGVwZW5kZW5jeVZlcnRleCB7XG4gICAgICByZXR1cm4gbm9kZXNbY29uc3RydWN0Lm5vZGUucGF0aF07XG4gICAgfVxuXG4gICAgLy8gY3JlYXRlIGFsbCB2ZXJ0aWNlcyBvZiB0aGUgZ3JhcGguXG4gICAgZm9yIChjb25zdCBuIG9mIG5vZGUuZmluZEFsbCgpKSB7XG4gICAgICBwdXRWZXJ0ZXgobik7XG4gICAgfVxuXG4gICAgY29uc3QgZGVwcyA9IFtdO1xuICAgIGZvciAoY29uc3QgY2hpbGQgb2Ygbm9kZS5maW5kQWxsKCkpIHtcbiAgICAgIGZvciAoY29uc3QgZGVwIG9mIGNoaWxkLm5vZGUuZGVwZW5kZW5jaWVzKSB7XG4gICAgICAgIGRlcHMucHVzaCh7IHNvdXJjZTogY2hpbGQsIHRhcmdldDogZGVwIH0pO1xuICAgICAgfVxuICAgIH1cblxuICAgIC8vIGNyZWF0ZSBhbGwgdGhlIGVkZ2VzIG9mIHRoZSBncmFwaC5cbiAgICBmb3IgKGNvbnN0IGRlcCBvZiBkZXBzKSB7XG5cbiAgICAgIGlmICghZ2V0VmVydGV4KGRlcC50YXJnZXQpKSB7XG4gICAgICAgIC8vIGRvbnQgY3Jvc3Mgc2NvcGUgYm91bmRhcmllcy5cbiAgICAgICAgLy8gc2luY2UgY2hhcnRzIG9ubHkgcmVuZGVycyBpdHMgb3duIGNoaWxkcmVuLCB0aGlzIGlzIG9rIGFuZFxuICAgICAgICAvLyBoYXMgdGhlIGJlbmVmaXQgb2Ygc2ltcGxpZnlpbmcgdGhlIGdyYXBoLiB3ZSBzaG91bGQgcmVjb25zaWRlciB0aGlzIGJlaGF2aW9yIHdoZW4gbW92aW5nXG4gICAgICAgIC8vIHRvIGEgbW9yZSBnZW5lcmFsIHB1cnBvc2UgdXNlLWNhc2UuXG4gICAgICAgIGNvbnRpbnVlO1xuICAgICAgfVxuXG4gICAgICBjb25zdCBzb3VyY2VEZXBOb2RlID0gZ2V0VmVydGV4KGRlcC5zb3VyY2UpO1xuICAgICAgY29uc3QgdGFyZ2V0RGVwTm9kZSA9IGdldFZlcnRleChkZXAudGFyZ2V0KTtcblxuICAgICAgc291cmNlRGVwTm9kZS5hZGRDaGlsZCh0YXJnZXREZXBOb2RlKTtcblxuICAgIH1cblxuICAgIC8vIGNyZWF0ZSB0aGUgcm9vdC5cbiAgICBmb3IgKGNvbnN0IG4gb2YgT2JqZWN0LnZhbHVlcyhub2RlcykpIHtcbiAgICAgIGlmIChuLmluYm91bmQubGVuZ3RoID09PSAwKSB7XG4gICAgICAgIC8vIG9ycGhhbnMgYXJlIGRlcGVuZGVuY3kgcm9vdHMuIGxldHMgYWRvcHQgdGhlbSFcbiAgICAgICAgdGhpcy5fZm9zdGVyUGFyZW50LmFkZENoaWxkKG4pO1xuICAgICAgfVxuICAgIH1cblxuICB9XG5cbiAgLyoqXG4gICAqIFJldHVybnMgdGhlIHJvb3Qgb2YgdGhlIGdyYXBoLlxuICAgKlxuICAgKiBOb3RlIHRoYXQgdGhpcyB2ZXJ0ZXggd2lsbCBhbHdheXMgaGF2ZSBgbnVsbGAgYXMgaXRzIGAudmFsdWVgIHNpbmNlIGl0IGlzIGFuIGFydGlmaWNhbCByb290XG4gICAqIHRoYXQgYmluZHMgYWxsIHRoZSBjb25uZWN0ZWQgc3BhY2VzIG9mIHRoZSBncmFwaC5cbiAgICovXG4gIHB1YmxpYyBnZXQgcm9vdCgpOiBEZXBlbmRlbmN5VmVydGV4IHtcbiAgICByZXR1cm4gdGhpcy5fZm9zdGVyUGFyZW50O1xuICB9XG5cbiAgLyoqXG4gICAqIEBzZWUgVmVydGV4LnRvcG9sb2d5KClcbiAgICovXG4gIHB1YmxpYyB0b3BvbG9neSgpOiBJQ29uc3RydWN0W10ge1xuICAgIHJldHVybiB0aGlzLl9mb3N0ZXJQYXJlbnQudG9wb2xvZ3koKTtcbiAgfVxufVxuXG4vKipcbiAqIFJlcHJlc2VudHMgYSB2ZXJ0ZXggaW4gdGhlIGdyYXBoLlxuICpcbiAqIFRoZSB2YWx1ZSBvZiBlYWNoIHZlcnRleCBpcyBhbiBgSUNvbnN0cnVjdGAgdGhhdCBpcyBhY2Nlc3NpYmxlIHZpYSB0aGUgYC52YWx1ZWAgZ2V0dGVyLlxuICovXG5leHBvcnQgY2xhc3MgRGVwZW5kZW5jeVZlcnRleCB7XG5cbiAgcHJpdmF0ZSByZWFkb25seSBfdmFsdWU6IElDb25zdHJ1Y3QgfCB1bmRlZmluZWQ7XG4gIHByaXZhdGUgcmVhZG9ubHkgX2NoaWxkcmVuOiBTZXQ8RGVwZW5kZW5jeVZlcnRleD4gPSBuZXcgU2V0PERlcGVuZGVuY3lWZXJ0ZXg+KCk7XG4gIHByaXZhdGUgcmVhZG9ubHkgX3BhcmVudHM6IFNldDxEZXBlbmRlbmN5VmVydGV4PiA9IG5ldyBTZXQ8RGVwZW5kZW5jeVZlcnRleD4oKTtcblxuICBjb25zdHJ1Y3Rvcih2YWx1ZTogSUNvbnN0cnVjdCB8IHVuZGVmaW5lZCA9IHVuZGVmaW5lZCkge1xuICAgIHRoaXMuX3ZhbHVlID0gdmFsdWU7XG4gIH1cblxuICAvKipcbiAgICogUmV0dXJucyB0aGUgSUNvbnN0cnVjdCB0aGlzIGdyYXBoIHZlcnRleCByZXByZXNlbnRzLlxuICAgKlxuICAgKiBgbnVsbGAgaW4gY2FzZSB0aGlzIGlzIHRoZSByb290IG9mIHRoZSBncmFwaC5cbiAgICovXG4gIHB1YmxpYyBnZXQgdmFsdWUoKTogSUNvbnN0cnVjdCB8IHVuZGVmaW5lZCB7XG4gICAgcmV0dXJuIHRoaXMuX3ZhbHVlO1xuICB9XG5cbiAgLyoqXG4gICAqIFJldHVybnMgdGhlIGNoaWxkcmVuIG9mIHRoZSB2ZXJ0ZXggKGkuZSBkZXBlbmRlbmNpZXMpXG4gICAqL1xuICBwdWJsaWMgZ2V0IG91dGJvdW5kKCk6IEFycmF5PERlcGVuZGVuY3lWZXJ0ZXg+IHtcbiAgICByZXR1cm4gQXJyYXkuZnJvbSh0aGlzLl9jaGlsZHJlbik7XG4gIH1cblxuICAvKipcbiAgICogUmV0dXJucyB0aGUgcGFyZW50cyBvZiB0aGUgdmVydGV4IChpLmUgZGVwZW5kYW50cylcbiAgICovXG4gIHB1YmxpYyBnZXQgaW5ib3VuZCgpOiBBcnJheTxEZXBlbmRlbmN5VmVydGV4PiB7XG4gICAgcmV0dXJuIEFycmF5LmZyb20odGhpcy5fcGFyZW50cyk7XG4gIH1cblxuICAvKipcbiAgICogUmV0dXJucyBhIHRvcG9sb2dpY2FsbHkgc29ydGVkIGFycmF5IG9mIHRoZSBjb25zdHJ1Y3RzIGluIHRoZSBzdWItZ3JhcGguXG4gICAqL1xuICBwdWJsaWMgdG9wb2xvZ3koKTogSUNvbnN0cnVjdFtdIHtcblxuICAgIGNvbnN0IGZvdW5kID0gbmV3IFNldDxEZXBlbmRlbmN5VmVydGV4PigpO1xuICAgIGNvbnN0IHRvcG9sb2d5OiBEZXBlbmRlbmN5VmVydGV4W10gPSBbXTtcblxuICAgIGZ1bmN0aW9uIHZpc2l0KG46IERlcGVuZGVuY3lWZXJ0ZXgpIHtcbiAgICAgIGZvciAoY29uc3QgYyBvZiBuLm91dGJvdW5kKSB7XG4gICAgICAgIHZpc2l0KGMpO1xuICAgICAgfVxuICAgICAgaWYgKCFmb3VuZC5oYXMobikpIHtcbiAgICAgICAgdG9wb2xvZ3kucHVzaChuKTtcbiAgICAgICAgZm91bmQuYWRkKG4pO1xuICAgICAgfVxuICAgIH1cblxuICAgIHZpc2l0KHRoaXMpO1xuXG4gICAgcmV0dXJuIHRvcG9sb2d5LmZpbHRlcihkID0+IGQudmFsdWUpLm1hcChkID0+IGQudmFsdWUhKTtcblxuICB9XG5cbiAgLyoqXG4gICAqIEFkZHMgYSB2ZXJ0ZXggYXMgYSBkZXBlbmRlbmN5IG9mIHRoZSBjdXJyZW50IG5vZGUuXG4gICAqIEFsc28gdXBkYXRlcyB0aGUgcGFyZW50cyBvZiBgZGVwYCwgc28gdGhhdCBpdCBjb250YWlucyB0aGlzIG5vZGUgYXMgYSBwYXJlbnQuXG4gICAqXG4gICAqIFRoaXMgb3BlcmF0aW9uIHdpbGwgZmFpbCBpbiBjYXNlIGl0IGNyZWF0ZXMgYSBjeWNsZSBpbiB0aGUgZ3JhcGguXG4gICAqXG4gICAqIEBwYXJhbSBkZXAgVGhlIGRlcGVuZGVuY3lcbiAgICovXG4gIHB1YmxpYyBhZGRDaGlsZChkZXA6IERlcGVuZGVuY3lWZXJ0ZXgpIHtcblxuICAgIGNvbnN0IGN5Y2xlOiBEZXBlbmRlbmN5VmVydGV4W10gPSBkZXAuZmluZFJvdXRlKHRoaXMpO1xuICAgIGlmIChjeWNsZS5sZW5ndGggIT09IDApIHtcbiAgICAgIGN5Y2xlLnB1c2goZGVwKTtcbiAgICAgIHRocm93IG5ldyBFcnJvcihgRGVwZW5kZW5jeSBjeWNsZSBkZXRlY3RlZDogJHtjeWNsZS5maWx0ZXIoZCA9PiBkLnZhbHVlKS5tYXAoZCA9PiBkLnZhbHVlIS5ub2RlLnBhdGgpLmpvaW4oJyA9PiAnKX1gKTtcbiAgICB9XG5cbiAgICB0aGlzLl9jaGlsZHJlbi5hZGQoZGVwKTtcbiAgICBkZXAuYWRkUGFyZW50KHRoaXMpO1xuICB9XG5cbiAgcHJpdmF0ZSBhZGRQYXJlbnQoZGVwOiBEZXBlbmRlbmN5VmVydGV4KSB7XG4gICAgdGhpcy5fcGFyZW50cy5hZGQoZGVwKTtcbiAgfVxuXG4gIHByaXZhdGUgZmluZFJvdXRlKGRzdDogRGVwZW5kZW5jeVZlcnRleCk6IERlcGVuZGVuY3lWZXJ0ZXhbXSB7XG5cbiAgICBjb25zdCByb3V0ZTogRGVwZW5kZW5jeVZlcnRleFtdID0gW107XG4gICAgdmlzaXQodGhpcyk7XG4gICAgcmV0dXJuIHJvdXRlO1xuXG4gICAgZnVuY3Rpb24gdmlzaXQobjogRGVwZW5kZW5jeVZlcnRleCk6IGJvb2xlYW4ge1xuICAgICAgcm91dGUucHVzaChuKTtcbiAgICAgIGxldCBmb3VuZCA9IGZhbHNlO1xuICAgICAgZm9yIChjb25zdCBjIG9mIG4ub3V0Ym91bmQpIHtcbiAgICAgICAgaWYgKGMgPT09IGRzdCkge1xuICAgICAgICAgIHJvdXRlLnB1c2goYyk7XG4gICAgICAgICAgcmV0dXJuIHRydWU7XG4gICAgICAgIH1cbiAgICAgICAgZm91bmQgPSB2aXNpdChjKTtcbiAgICAgIH1cbiAgICAgIGlmICghZm91bmQpIHtcbiAgICAgICAgcm91dGUucG9wKCk7XG4gICAgICB9XG4gICAgICByZXR1cm4gZm91bmQ7XG5cbiAgICB9XG5cbiAgfVxufVxuIl19