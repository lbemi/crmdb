"use strict";
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
exports.App = exports.YamlOutputType = void 0;
const JSII_RTTI_SYMBOL_1 = Symbol.for("jsii.rtti");
const fs = require("fs");
const path = require("path");
const constructs_1 = require("constructs");
const api_object_1 = require("./api-object");
const chart_1 = require("./chart");
const dependency_1 = require("./dependency");
const names_1 = require("./names");
const resolve_1 = require("./resolve");
const yaml_1 = require("./yaml");
/** The method to divide YAML output into files */
var YamlOutputType;
(function (YamlOutputType) {
    /** All resources are output into a single YAML file */
    YamlOutputType[YamlOutputType["FILE_PER_APP"] = 0] = "FILE_PER_APP";
    /** Resources are split into seperate files by chart */
    YamlOutputType[YamlOutputType["FILE_PER_CHART"] = 1] = "FILE_PER_CHART";
    /** Each resource is output to its own file */
    YamlOutputType[YamlOutputType["FILE_PER_RESOURCE"] = 2] = "FILE_PER_RESOURCE";
    /** Each chart in its own folder and each resource in its own file */
    YamlOutputType[YamlOutputType["FOLDER_PER_CHART_FILE_PER_RESOURCE"] = 3] = "FOLDER_PER_CHART_FILE_PER_RESOURCE";
})(YamlOutputType || (exports.YamlOutputType = YamlOutputType = {}));
class SynthRequestCache {
    constructor() {
        this.nodeChildrenCache = new Map();
    }
    findAll(node) {
        if (this.nodeChildrenCache.has(node)) {
            return this.nodeChildrenCache.get(node);
        }
        const children = node.findAll();
        this.nodeChildrenCache.set(node, children);
        return children;
    }
}
/**
 * Represents a cdk8s application.
 */
class App extends constructs_1.Construct {
    /**
     * Synthesize a single chart.
     *
     * Each element returned in the resulting array represents a different ApiObject
     * in the scope of the chart.
     *
     * Note that the returned array order is important. It is determined by the various dependencies between
     * the constructs in the chart, where the first element is the one without dependencies, and so on...
     *
     * @returns An array of JSON objects.
     * @param chart the chart to synthesize.
     * @internal
     */
    static _synthChart(chart) {
        const app = App.of(chart);
        const cache = new SynthRequestCache();
        // we must prepare the entire app before synthesizing the chart
        // because the dependency inference happens on the app level.
        resolveDependencies(app, cache);
        // validate the app since we want to call onValidate of the relevant constructs.
        // note this will also call onValidate on constructs from possibly different charts,
        // but thats ok too since we no longer treat constructs as a self-contained synthesis unit.
        validate(app, cache);
        return chartToKube(chart).map(obj => obj.toJson());
    }
    static of(c) {
        const scope = c.node.scope;
        if (!scope) {
            // the app is the only construct without a scope.
            return c;
        }
        return App.of(scope);
    }
    /**
     * Returns all the charts in this app, sorted topologically.
     */
    get charts() {
        const isChart = (x) => x instanceof chart_1.Chart;
        return new dependency_1.DependencyGraph(this.node)
            .topology()
            .filter(isChart);
    }
    /**
     * Defines an app
     * @param props configuration options
     */
    constructor(props = {}) {
        super(undefined, '');
        this.outdir = props.outdir ?? process.env.CDK8S_OUTDIR ?? 'dist';
        this.outputFileExtension = props.outputFileExtension ?? '.k8s.yaml';
        this.yamlOutputType = props.yamlOutputType ?? YamlOutputType.FILE_PER_CHART;
        this.resolvers = [...(props.resolvers ?? []), new resolve_1.LazyResolver(), new resolve_1.ImplicitTokenResolver(), new resolve_1.NumberStringUnionResolver()];
        this.recordConstructMetadata = props.recordConstructMetadata ?? (process.env.CDK8S_RECORD_CONSTRUCT_METADATA === 'true' ? true : false);
    }
    /**
     * Synthesizes all manifests to the output directory
     */
    synth() {
        fs.mkdirSync(this.outdir, { recursive: true });
        const cache = new SynthRequestCache();
        // Since we plan on removing the distributed synth mechanism, we no longer call `Node.synthesize`, but rather simply implement
        // the necessary operations. We do however want to preserve the distributed validation.
        validate(this, cache);
        // this is kind of sucky, eventually I would like the DependencyGraph
        // to be able to answer this question.
        const hasDependantCharts = resolveDependencies(this, cache);
        const charts = this.charts;
        switch (this.yamlOutputType) {
            case YamlOutputType.FILE_PER_APP:
                let apiObjectsList = [];
                for (const chart of charts) {
                    apiObjectsList.push(...Object.values(chart.toJson()));
                }
                if (charts.length > 0) {
                    yaml_1.Yaml.save(path.join(this.outdir, `app${this.outputFileExtension}`), // There is no "app name", so we just hardcode the file name
                    apiObjectsList);
                }
                break;
            case YamlOutputType.FILE_PER_CHART:
                const namer = hasDependantCharts ? new IndexedChartNamer() : new SimpleChartNamer();
                for (const chart of charts) {
                    const chartName = namer.name(chart);
                    const objects = Object.values(chart.toJson());
                    yaml_1.Yaml.save(path.join(this.outdir, chartName + this.outputFileExtension), objects);
                }
                break;
            case YamlOutputType.FILE_PER_RESOURCE:
                for (const chart of charts) {
                    const apiObjects = Object.values(chart.toJson());
                    apiObjects.forEach((apiObject) => {
                        if (!(apiObject === undefined)) {
                            const fileName = `${`${apiObject.kind}.${apiObject.metadata.name}`
                                .replace(/[^0-9a-zA-Z-_.]/g, '')}`;
                            yaml_1.Yaml.save(path.join(this.outdir, fileName + this.outputFileExtension), [apiObject]);
                        }
                    });
                }
                break;
            case YamlOutputType.FOLDER_PER_CHART_FILE_PER_RESOURCE:
                const folderNamer = hasDependantCharts ? new IndexedChartFolderNamer() : new SimpleChartFolderNamer();
                for (const chart of charts) {
                    const chartName = folderNamer.name(chart);
                    const apiObjects = chartToKube(chart);
                    const fullOutDir = path.join(this.outdir, chartName);
                    fs.mkdirSync(fullOutDir, { recursive: true });
                    apiObjects.forEach((apiObject) => {
                        if (!(apiObject === undefined)) {
                            const fileName = `${`${apiObject.kind}.${apiObject.metadata.name}`
                                .replace(/[^0-9a-zA-Z-_.]/g, '')}`;
                            yaml_1.Yaml.save(path.join(fullOutDir, fileName + this.outputFileExtension), [apiObject.toJson()]);
                        }
                    });
                }
                break;
            default:
                break;
        }
        if (this.recordConstructMetadata) {
            const allObjects = this.charts.flatMap(chartToKube);
            this.writeConstructMetadata(allObjects);
        }
    }
    /**
     * Synthesizes the app into a YAML string.
     *
     * @returns A string with all YAML objects across all charts in this app.
     */
    synthYaml() {
        const cache = new SynthRequestCache();
        resolveDependencies(this, cache);
        validate(this, cache);
        const charts = this.charts;
        const docs = [];
        for (const chart of charts) {
            docs.push(...Object.values(chart.toJson()));
        }
        return yaml_1.Yaml.stringify(...docs);
    }
    writeConstructMetadata(apiObjects) {
        const resources = {};
        for (const apiObject of apiObjects) {
            resources[apiObject.name] = { path: apiObject.node.path };
        }
        fs.writeFileSync(path.join(this.outdir, 'construct-metadata.json'), JSON.stringify({
            version: '1.0.0',
            resources: resources,
        }));
    }
}
exports.App = App;
_a = JSII_RTTI_SYMBOL_1;
App[_a] = { fqn: "cdk8s.App", version: "2.68.60" };
function validate(app, cache) {
    const errors = [];
    for (const child of cache.findAll(app.node)) {
        const childErrors = child.node.validate();
        for (const error of childErrors) {
            errors.push(`[${child.node.path}] ${error}`);
        }
    }
    if (errors.length > 0) {
        throw new Error(`Validation failed with the following errors:\n  ${errors.join('\n  ')}`);
    }
}
function buildDependencies(app, cache) {
    const deps = [];
    for (const child of cache.findAll(app.node)) {
        for (const dep of child.node.dependencies) {
            deps.push({ source: child, target: dep });
        }
    }
    return deps;
}
function resolveDependencies(app, cache) {
    let hasDependantCharts = false;
    // create an explicit chart dependency from nested chart relationships
    for (const parentChart of cache.findAll(app.node).filter(x => x instanceof chart_1.Chart)) {
        for (const childChart of parentChart.node.children.filter(x => x instanceof chart_1.Chart)) {
            parentChart.node.addDependency(childChart);
            hasDependantCharts = true;
        }
    }
    // create an explicit chart dependency from implicit construct dependencies
    for (const dep of buildDependencies(app, cache)) {
        const sourceChart = chart_1.Chart.of(dep.source);
        const targetChart = chart_1.Chart.of(dep.target);
        if (sourceChart !== targetChart) {
            sourceChart.node.addDependency(targetChart);
            hasDependantCharts = true;
        }
    }
    // create explicit api object dependencies from implicit construct dependencies
    for (const dep of buildDependencies(app, cache)) {
        const sourceChart = chart_1.Chart.of(dep.source);
        const targetChart = chart_1.Chart.of(dep.target);
        const targetApiObjects = cache.findAll(dep.target.node).filter(c => c instanceof api_object_1.ApiObject).filter(x => chart_1.Chart.of(x) === targetChart);
        const sourceApiObjects = cache.findAll(dep.source.node).filter(c => c instanceof api_object_1.ApiObject).filter(x => chart_1.Chart.of(x) === sourceChart);
        for (const target of targetApiObjects) {
            for (const source of sourceApiObjects) {
                if (target !== source) {
                    source.node.addDependency(target);
                }
            }
        }
    }
    return hasDependantCharts;
}
function chartToKube(chart) {
    return new dependency_1.DependencyGraph(chart.node).topology()
        .filter(x => x instanceof api_object_1.ApiObject)
        .filter(x => chart_1.Chart.of(x) === chart) // include an object only in its closest parent chart
        .map(x => x);
}
class SimpleChartNamer {
    constructor() {
    }
    name(chart) {
        return `${names_1.Names.toDnsLabel(chart)}`;
    }
}
class IndexedChartNamer extends SimpleChartNamer {
    constructor() {
        super();
        this.index = 0;
    }
    name(chart) {
        const name = `${this.index.toString().padStart(4, '0')}-${super.name(chart)}`;
        this.index++;
        return name;
    }
}
class SimpleChartFolderNamer {
    constructor() {
    }
    name(chart) {
        return names_1.Names.toDnsLabel(chart);
    }
}
class IndexedChartFolderNamer extends SimpleChartFolderNamer {
    constructor() {
        super();
        this.index = 0;
    }
    name(chart) {
        const name = `${this.index.toString().padStart(4, '0')}-${super.name(chart)}`;
        this.index++;
        return name;
    }
}
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiYXBwLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsiLi4vc3JjL2FwcC50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiOzs7OztBQUFBLHlCQUF5QjtBQUN6Qiw2QkFBNkI7QUFDN0IsMkNBQXlEO0FBQ3pELDZDQUF5QztBQUN6QyxtQ0FBZ0M7QUFDaEMsNkNBQStDO0FBQy9DLG1DQUFnQztBQUNoQyx1Q0FBc0c7QUFDdEcsaUNBQThCO0FBRTlCLGtEQUFrRDtBQUNsRCxJQUFZLGNBU1g7QUFURCxXQUFZLGNBQWM7SUFDeEIsdURBQXVEO0lBQ3ZELG1FQUFZLENBQUE7SUFDWix1REFBdUQ7SUFDdkQsdUVBQWMsQ0FBQTtJQUNkLDhDQUE4QztJQUM5Qyw2RUFBaUIsQ0FBQTtJQUNqQixxRUFBcUU7SUFDckUsK0dBQWtDLENBQUE7QUFDcEMsQ0FBQyxFQVRXLGNBQWMsOEJBQWQsY0FBYyxRQVN6QjtBQWlERCxNQUFNLGlCQUFpQjtJQUF2QjtRQUNTLHNCQUFpQixHQUE0QixJQUFJLEdBQUcsRUFBc0IsQ0FBQztJQVdwRixDQUFDO0lBVFEsT0FBTyxDQUFDLElBQVU7UUFDdkIsSUFBSSxJQUFJLENBQUMsaUJBQWlCLENBQUMsR0FBRyxDQUFDLElBQUksQ0FBQyxFQUFFLENBQUM7WUFDckMsT0FBTyxJQUFJLENBQUMsaUJBQWlCLENBQUMsR0FBRyxDQUFDLElBQUksQ0FBRSxDQUFDO1FBQzNDLENBQUM7UUFFRCxNQUFNLFFBQVEsR0FBRyxJQUFJLENBQUMsT0FBTyxFQUFFLENBQUM7UUFDaEMsSUFBSSxDQUFDLGlCQUFpQixDQUFDLEdBQUcsQ0FBQyxJQUFJLEVBQUUsUUFBUSxDQUFDLENBQUM7UUFDM0MsT0FBTyxRQUFRLENBQUM7SUFDbEIsQ0FBQztDQUNGO0FBRUQ7O0dBRUc7QUFDSCxNQUFhLEdBQUksU0FBUSxzQkFBUztJQUNoQzs7Ozs7Ozs7Ozs7O09BWUc7SUFDSSxNQUFNLENBQUMsV0FBVyxDQUFDLEtBQVk7UUFFcEMsTUFBTSxHQUFHLEdBQVEs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UUFDOUUsSUFBSSxDQUFDLEtBQUssRUFBRSxDQUFDO1FBQ2IsT0FBTyxJQUFJLENBQUM7SUFDZCxDQUFDO0NBQ0Y7QUFFRCxNQUFNLHNCQUFzQjtJQUMxQjtJQUNBLENBQUM7SUFFTSxJQUFJLENBQUMsS0FBWTtRQUN0QixPQUFPLGFBQUssQ0FBQyxVQUFVLENBQUMsS0FBSyxDQUFDLENBQUM7SUFDakMsQ0FBQztDQUNGO0FBRUQsTUFBTSx1QkFBd0IsU0FBUSxzQkFBc0I7SUFFMUQ7UUFDRSxLQUFLLEVBQUUsQ0FBQztRQUZGLFVBQUssR0FBVyxDQUFDLENBQUM7SUFHMUIsQ0FBQztJQUVNLElBQUksQ0FBQyxLQUFZO1FBQ3RCLE1BQU0sSUFBSSxHQUFHLEdBQUcsSUFBSSxDQUFDLEtBQUssQ0FBQyxRQUFRLEVBQUUsQ0FBQyxRQUFRLENBQUMsQ0FBQyxFQUFFLEdBQUcsQ0FBQyxJQUFJLEtBQUssQ0FBQyxJQUFJLENBQUMsS0FBSyxDQUFDLEVBQUUsQ0FBQztRQUM5RSxJQUFJLENBQUMsS0FBSyxFQUFFLENBQUM7UUFDYixPQUFPLElBQUksQ0FBQztJQUNkLENBQUM7Q0FDRiIsInNvdXJjZXNDb250ZW50IjpbImltcG9ydCAqIGFzIGZzIGZyb20gJ2ZzJztcbmltcG9ydCAqIGFzIHBhdGggZnJvbSAncGF0aCc7XG5pbXBvcnQgeyBDb25zdHJ1Y3QsIElDb25zdHJ1Y3QsIE5vZGUgfSBmcm9tICdjb25zdHJ1Y3RzJztcbmltcG9ydCB7IEFwaU9iamVjdCB9IGZyb20gJy4vYXBpLW9iamVjdCc7XG5pbXBvcnQgeyBDaGFydCB9IGZyb20gJy4vY2hhcnQnO1xuaW1wb3J0IHsgRGVwZW5kZW5jeUdyYXBoIH0gZnJvbSAnLi9kZXBlbmRlbmN5JztcbmltcG9ydCB7IE5hbWVzIH0gZnJvbSAnLi9uYW1lcyc7XG5pbXBvcnQgeyBJUmVzb2x2ZXIsIEltcGxpY2l0VG9rZW5SZXNvbHZlciwgTGF6eVJlc29sdmVyLCBOdW1iZXJTdHJpbmdVbmlvblJlc29sdmVyIH0gZnJvbSAnLi9yZXNvbHZlJztcbmltcG9ydCB7IFlhbWwgfSBmcm9tICcuL3lhbWwnO1xuXG4vKiogVGhlIG1ldGhvZCB0byBkaXZpZGUgWUFNTCBvdXRwdXQgaW50byBmaWxlcyAqL1xuZXhwb3J0IGVudW0gWWFtbE91dHB1dFR5cGUge1xuICAvKiogQWxsIHJlc291cmNlcyBhcmUgb3V0cHV0IGludG8gYSBzaW5nbGUgWUFNTCBmaWxlICovXG4gIEZJTEVfUEVSX0FQUCxcbiAgLyoqIFJlc291cmNlcyBhcmUgc3BsaXQgaW50byBzZXBlcmF0ZSBmaWxlcyBieSBjaGFydCAqL1xuICBGSUxFX1BFUl9DSEFSVCxcbiAgLyoqIEVhY2ggcmVzb3VyY2UgaXMgb3V0cHV0IHRvIGl0cyBvd24gZmlsZSAqL1xuICBGSUxFX1BFUl9SRVNPVVJDRSxcbiAgLyoqIEVhY2ggY2hhcnQgaW4gaXRzIG93biBmb2xkZXIgYW5kIGVhY2ggcmVzb3VyY2UgaW4gaXRzIG93biBmaWxlICovXG4gIEZPTERFUl9QRVJfQ0hBUlRfRklMRV9QRVJfUkVTT1VSQ0UsXG59XG5cbmV4cG9ydCBpbnRlcmZhY2UgQXBwUHJvcHMge1xuICAvKipcbiAgICogVGhlIGRpcmVjdG9yeSB0byBvdXRwdXQgS3ViZXJuZXRlcyBtYW5pZmVzdHMuXG4gICAqXG4gICAqIElmIHlvdSBzeW50aGVzaXplIHlvdXIgYXBwbGljYXRpb24gdXNpbmcgYGNkazhzIHN5bnRoYCwgeW91IG11c3RcbiAgICogYWxzbyBwYXNzIHRoaXMgdmFsdWUgdG8gdGhlIENMSSB1c2luZyB0aGUgYC0tb3V0cHV0YCBvcHRpb24gb3JcbiAgICogdGhlIGBvdXRwdXRgIHByb3BlcnR5IGluIHRoZSBgY2RrOHMueWFtbGAgY29uZmlndXJhdGlvbiBmaWxlLlxuICAgKiBPdGhlcndpc2UsIHRoZSBDTEkgd2lsbCBub3Qga25vdyBhYm91dCB0aGUgb3V0cHV0IGRpcmVjdG9yeSxcbiAgICogYW5kIHN5bnRoZXNpcyB3aWxsIGZhaWwuXG4gICAqXG4gICAqIFRoaXMgcHJvcGVydHkgaXMgaW50ZW5kZWQgZm9yIGludGVybmFsIGFuZCB0ZXN0aW5nIHVzZS5cbiAgICpcbiAgICogQGRlZmF1bHQgLSBDREs4U19PVVRESVIgaWYgZGVmaW5lZCwgb3RoZXJ3aXNlIFwiZGlzdFwiXG4gICAqL1xuICByZWFkb25seSBvdXRkaXI/OiBzdHJpbmc7XG4gIC8qKlxuICAgKiAgVGhlIGZpbGUgZXh0ZW5zaW9uIHRvIHVzZSBmb3IgcmVuZGVyZWQgWUFNTCBmaWxlc1xuICAgKiBAZGVmYXVsdCAuazhzLnlhbWxcbiAgICovXG4gIHJlYWRvbmx5IG91dHB1dEZpbGVFeHRlbnNpb24/OiBzdHJpbmc7XG4gIC8qKlxuICAgKiAgSG93IHRvIGRpdmlkZSB0aGUgWUFNTCBvdXRwdXQgaW50byBmaWxlc1xuICAgKiBAZGVmYXVsdCBZYW1sT3V0cHV0VHlwZS5GSUxFX1BFUl9DSEFSVFxuICAgKi9cbiAgcmVhZG9ubHkgeWFtbE91dHB1dFR5cGU/OiBZYW1sT3V0cHV0VHlwZTtcblxuICAvKipcbiAgICogV2hlbiBzZXQgdG8gdHJ1ZSwgdGhlIG91dHB1dCBkaXJlY3Rvcnkgd2lsbCBjb250YWluIGEgYGNvbnN0cnVjdC1tZXRhZGF0YS5qc29uYCBmaWxlXG4gICAqIHRoYXQgaG9sZHMgY29uc3RydWN0IHJlbGF0ZWQgbWV0YWRhdGEgb24gZXZlcnkgcmVzb3VyY2UgaW4gdGhlIGFwcC5cbiAgICpcbiAgICogQGRlZmF1bHQgZmFsc2VcbiAgICovXG4gIHJlYWRvbmx5IHJlY29yZENvbnN0cnVjdE1ldGFkYXRhPzogYm9vbGVhbjtcblxuICAvKipcbiAgICogQSBsaXN0IG9mIHJlc29sdmVycyB0aGF0IGNhbiBiZSB1c2VkIHRvIHJlcGxhY2UgcHJvcGVydHkgdmFsdWVzIGJlZm9yZVxuICAgKiB0aGV5IGFyZSB3cml0dGVuIHRvIHRoZSBtYW5pZmVzdCBmaWxlLiBXaGVuIG11bHRpcGxlIHJlc29sdmVycyBhcmUgcGFzc2VkLFxuICAgKiB0aGV5IGFyZSBpbnZva2VkIGJ5IG9yZGVyIGluIHRoZSBsaXN0LCBhbmQgb25seSB0aGUgZmlyc3Qgb25lIHRoYXQgYXBwbGllc1xuICAgKiAoZS5nIGNhbGxzIGBjb250ZXh0LnJlcGxhY2VWYWx1ZWApIGlzIGludm9rZWQuXG4gICAqXG4gICAqIEBzZWUgaHR0cHM6Ly9jZGs4cy5pby9kb2NzL2xhdGVzdC9iYXNpY3MvYXBwLyNyZXNvbHZlcnNcbiAgICpcbiAgICogQGRlZmF1bHQgLSBubyByZXNvbHZlcnMuXG4gICAqL1xuICByZWFkb25seSByZXNvbHZlcnM/OiBJUmVzb2x2ZXJbXTtcbn1cblxuY2xhc3MgU3ludGhSZXF1ZXN0Q2FjaGUge1xuICBwdWJsaWMgbm9kZUNoaWxkcmVuQ2FjaGU6IE1hcDxOb2RlLCBJQ29uc3RydWN0W10+ID0gbmV3IE1hcDxOb2RlLCBJQ29uc3RydWN0W10+KCk7XG5cbiAgcHVibGljIGZpbmRBbGwobm9kZTogTm9kZSk6IElDb25zdHJ1Y3RbXSB7XG4gICAgaWYgKHRoaXMubm9kZUNoaWxkcmVuQ2FjaGUuaGFzKG5vZGUpKSB7XG4gICAgICByZXR1cm4gdGhpcy5ub2RlQ2hpbGRyZW5DYWNoZS5nZXQobm9kZSkhO1xuICAgIH1cblxuICAgIGNvbnN0IGNoaWxkcmVuID0gbm9kZS5maW5kQWxsKCk7XG4gICAgdGhpcy5ub2RlQ2hpbGRyZW5DYWNoZS5zZXQobm9kZSwgY2hpbGRyZW4pO1xuICAgIHJldHVybiBjaGlsZHJlbjtcbiAgfVxufVxuXG4vKipcbiAqIFJlcHJlc2VudHMgYSBjZGs4cyBhcHBsaWNhdGlvbi5cbiAqL1xuZXhwb3J0IGNsYXNzIEFwcCBleHRlbmRzIENvbnN0cnVjdCB7XG4gIC8qKlxuICAgKiBTeW50aGVzaXplIGEgc2luZ2xlIGNoYXJ0LlxuICAgKlxuICAgKiBFYWNoIGVsZW1lbnQgcmV0dXJuZWQgaW4gdGhlIHJlc3VsdGluZyBhcnJheSByZXByZXNlbnRzIGEgZGlmZmVyZW50IEFwaU9iamVjdFxuICAgKiBpbiB0aGUgc2NvcGUgb2YgdGhlIGNoYXJ0LlxuICAgKlxuICAgKiBOb3RlIHRoYXQgdGhlIHJldHVybmVkIGFycmF5IG9yZGVyIGlzIGltcG9ydGFudC4gSXQgaXMgZGV0ZXJtaW5lZCBieSB0aGUgdmFyaW91cyBkZXBlbmRlbmNpZXMgYmV0d2VlblxuICAgKiB0aGUgY29uc3RydWN0cyBpbiB0aGUgY2hhcnQsIHdoZXJlIHRoZSBmaXJzdCBlbGVtZW50IGlzIHRoZSBvbmUgd2l0aG91dCBkZXBlbmRlbmNpZXMsIGFuZCBzbyBvbi4uLlxuICAgKlxuICAgKiBAcmV0dXJucyBBbiBhcnJheSBvZiBKU09OIG9iamVjdHMuXG4gICAqIEBwYXJhbSBjaGFydCB0aGUgY2hhcnQgdG8gc3ludGhlc2l6ZS5cbiAgICogQGludGVybmFsXG4gICAqL1xuICBwdWJsaWMgc3RhdGljIF9zeW50aENoYXJ0KGNoYXJ0OiBDaGFydCk6IGFueVtdIHtcblxuICAgIGNvbnN0IGFwcDogQXBwID0gQXBwLm9mKGNoYXJ0KTtcblxuICAgIGNvbnN0IGNhY2hlID0gbmV3IFN5bnRoUmVxdWVzdENhY2hlKCk7XG5cbiAgICAvLyB3ZSBtdXN0IHByZXBhcmUgdGhlIGVudGlyZSBhcHAgYmVmb3JlIHN5bnRoZXNpemluZyB0aGUgY2hhcnRcbiAgICAvLyBiZWNhdXNlIHRoZSBkZXBlbmRlbmN5IGluZmVyZW5jZSBoYXBwZW5zIG9uIHRoZSBhcHAgbGV2ZWwuXG4gICAgcmVzb2x2ZURlcGVuZGVuY2llcyhhcHAsIGNhY2hlKTtcblxuICAgIC8vIHZhbGlkYXRlIHRoZSBhcHAgc2luY2Ugd2Ugd2FudCB0byBjYWxsIG9uVmFsaWRhdGUgb2YgdGhlIHJlbGV2YW50IGNvbnN0cnVjdHMuXG4gICAgLy8gbm90ZSB0aGlzIHdpbGwgYWxzbyBjYWxsIG9uVmFsaWRhdGUgb24gY29uc3RydWN0cyBmcm9tIHBvc3NpYmx5IGRpZmZlcmVudCBjaGFydHMsXG4gICAgLy8gYnV0IHRoYXRzIG9rIHRvbyBzaW5jZSB3ZSBubyBsb25nZXIgdHJlYXQgY29uc3RydWN0cyBhcyBhIHNlbGYtY29udGFpbmVkIHN5bnRoZXNpcyB1bml0LlxuICAgIHZhbGlkYXRlKGFwcCwgY2FjaGUpO1xuXG4gICAgcmV0dXJuIGNoYXJ0VG9LdWJlKGNoYXJ0KS5tYXAob2JqID0+IG9iai50b0pzb24oKSk7XG4gIH1cblxuICBwdWJsaWMgc3RhdGljIG9mKGM6IElDb25zdHJ1Y3QpOiBBcHAge1xuXG4gICAgY29uc3Qgc2NvcGUgPSBjLm5vZGUuc2NvcGU7XG5cbiAgICBpZiAoIXNjb3BlKSB7XG4gICAgICAvLyB0aGUgYXBwIGlzIHRoZSBvbmx5IGNvbnN0cnVjdCB3aXRob3V0IGEgc2NvcGUuXG4gICAgICByZXR1cm4gYyBhcyBBcHA7XG4gICAgfVxuXG4gICAgcmV0dXJuIEFwcC5vZihzY29wZSk7XG4gIH1cblxuICAvKipcbiAgICogVGhlIG91dHB1dCBkaXJlY3RvcnkgaW50byB3aGljaCBtYW5pZmVzdHMgd2lsbCBiZSBzeW50aGVzaXplZC5cbiAgICovXG4gIHB1YmxpYyByZWFkb25seSBvdXRkaXI6IHN0cmluZztcblxuICAvKipcbiAgICogIFRoZSBmaWxlIGV4dGVuc2lvbiB0byB1c2UgZm9yIHJlbmRlcmVkIFlBTUwgZmlsZXNcbiAgICogQGRlZmF1bHQgLms4cy55YW1sXG4gICAqL1xuICBwdWJsaWMgcmVhZG9ubHkgb3V0cHV0RmlsZUV4dGVuc2lvbjogc3RyaW5nO1xuXG4gIC8qKiBIb3cgdG8gZGl2aWRlIHRoZSBZQU1MIG91dHB1dCBpbnRvIGZpbGVzXG4gICAqIEBkZWZhdWx0IFlhbWxPdXRwdXRUeXBlLkZJTEVfUEVSX0NIQVJUXG4gICAqL1xuICBwdWJsaWMgcmVhZG9ubHkgeWFtbE91dHB1dFR5cGU6IFlhbWxPdXRwdXRUeXBlO1xuXG4gIC8qKlxuICAgKiBSZXNvbHZlcnMgdXNlZCBieSB0aGlzIGFwcC4gVGhpcyBpbmNsdWRlcyBib3RoIGN1c3RvbSByZXNvbHZlcnNcbiAgICogcGFzc2VkIGJ5IHRoZSBgcmVzb2x2ZXJzYCBwcm9wZXJ0eSwgYXMgd2VsbCBhcyBidWlsdC1pbiByZXNvbHZlcnMuXG4gICAqL1xuICBwdWJsaWMgcmVhZG9ubHkgcmVzb2x2ZXJzOiBJUmVzb2x2ZXJbXTtcblxuICBwcml2YXRlIHJlYWRvbmx5IHJlY29yZENvbnN0cnVjdE1ldGFkYXRhOiBib29sZWFuO1xuXG4gIC8qKlxuICAgKiBSZXR1cm5zIGFsbCB0aGUgY2hhcnRzIGluIHRoaXMgYXBwLCBzb3J0ZWQgdG9wb2xvZ2ljYWxseS5cbiAgICovXG4gIHB1YmxpYyBnZXQgY2hhcnRzKCk6IENoYXJ0W10ge1xuICAgIGNvbnN0IGlzQ2hhcnQgPSAoeDogSUNvbnN0cnVjdCk6IHggaXMgQ2hhcnQgPT4geCBpbnN0YW5jZW9mIENoYXJ0O1xuICAgIHJldHVybiBuZXcgRGVwZW5kZW5jeUdyYXBoKHRoaXMubm9kZSlcbiAgICAgIC50b3BvbG9neSgpXG4gICAgICAuZmlsdGVyKGlzQ2hhcnQpO1xuICB9XG5cbiAgLyoqXG4gICAqIERlZmluZXMgYW4gYXBwXG4gICAqIEBwYXJhbSBwcm9wcyBjb25maWd1cmF0aW9uIG9wdGlvbnNcbiAgICovXG4gIGNvbnN0cnVjdG9yKHByb3BzOiBBcHBQcm9wcyA9IHt9KSB7XG4gICAgc3VwZXIodW5kZWZpbmVkIGFzIGFueSwgJycpO1xuICAgIHRoaXMub3V0ZGlyID0gcHJvcHMub3V0ZGlyID8/IHByb2Nlc3MuZW52LkNESzhTX09VVERJUiA/PyAnZGlzdCc7XG4gICAgdGhpcy5vdXRwdXRGaWxlRXh0ZW5zaW9uID0gcHJvcHMub3V0cHV0RmlsZUV4dGVuc2lvbiA/PyAnLms4cy55YW1sJztcbiAgICB0aGlzLnlhbWxPdXRwdXRUeXBlID0gcHJvcHMueWFtbE91dHB1dFR5cGUgPz8gWWFtbE91dHB1dFR5cGUuRklMRV9QRVJfQ0hBUlQ7XG4gICAgdGhpcy5yZXNvbHZlcnMgPSBbLi4uKHByb3BzLnJlc29sdmVycyA/PyBbXSksIG5ldyBMYXp5UmVzb2x2ZXIoKSwgbmV3IEltcGxpY2l0VG9rZW5SZXNvbHZlcigpLCBuZXcgTnVtYmVyU3RyaW5nVW5pb25SZXNvbHZlcigpXTtcbiAgICB0aGlzLnJlY29yZENvbnN0cnVjdE1ldGFkYXRhID0gcHJvcHMucmVjb3JkQ29uc3RydWN0TWV0YWRhdGEgPz8gKHByb2Nlc3MuZW52LkNESzhTX1JFQ09SRF9DT05TVFJVQ1RfTUVUQURBVEEgPT09ICd0cnVlJyA/IHRydWUgOiBmYWxzZSk7XG5cbiAgfVxuXG4gIC8qKlxuICAgKiBTeW50aGVzaXplcyBhbGwgbWFuaWZlc3RzIHRvIHRoZSBvdXRwdXQgZGlyZWN0b3J5XG4gICAqL1xuICBwdWJsaWMgc3ludGgoKTogdm9pZCB7XG5cbiAgICBmcy5ta2RpclN5bmModGhpcy5vdXRkaXIsIHsgcmVjdXJzaXZlOiB0cnVlIH0pO1xuXG4gICAgY29uc3QgY2FjaGUgPSBuZXcgU3ludGhSZXF1ZXN0Q2FjaGUoKTtcblxuICAgIC8vIFNpbmNlIHdlIHBsYW4gb24gcmVtb3ZpbmcgdGhlIGRpc3RyaWJ1dGVkIHN5bnRoIG1lY2hhbmlzbSwgd2Ugbm8gbG9uZ2VyIGNhbGwgYE5vZGUuc3ludGhlc2l6ZWAsIGJ1dCByYXRoZXIgc2ltcGx5IGltcGxlbWVudFxuICAgIC8vIHRoZSBuZWNlc3Nhcnkgb3BlcmF0aW9ucy4gV2UgZG8gaG93ZXZlciB3YW50IHRvIHByZXNlcnZlIHRoZSBkaXN0cmlidXRlZCB2YWxpZGF0aW9uLlxuICAgIHZhbGlkYXRlKHRoaXMsIGNhY2hlKTtcblxuICAgIC8vIHRoaXMgaXMga2luZCBvZiBzdWNreSwgZXZlbnR1YWxseSBJIHdvdWxkIGxpa2UgdGhlIERlcGVuZGVuY3lHcmFwaFxuICAgIC8vIHRvIGJlIGFibGUgdG8gYW5zd2VyIHRoaXMgcXVlc3Rpb24uXG4gICAgY29uc3QgaGFzRGVwZW5kYW50Q2hhcnRzID0gcmVzb2x2ZURlcGVuZGVuY2llcyh0aGlzLCBjYWNoZSk7XG4gICAgY29uc3QgY2hhcnRzID0gdGhpcy5jaGFydHM7XG5cbiAgICBzd2l0Y2ggKHRoaXMueWFtbE91dHB1dFR5cGUpIHtcbiAgICAgIGNhc2UgWWFtbE91dHB1dFR5cGUuRklMRV9QRVJfQVBQOlxuICAgICAgICBsZXQgYXBpT2JqZWN0c0xpc3Q6IEFwaU9iamVjdFtdID0gW107XG5cbiAgICAgICAgZm9yIChjb25zdCBjaGFydCBvZiBjaGFydHMpIHtcbiAgICAgICAgICBhcGlPYmplY3RzTGlzdC5wdXNoKC4uLk9iamVjdC52YWx1ZXMoY2hhcnQudG9Kc29uKCkpKTtcbiAgICAgICAgfVxuXG4gICAgICAgIGlmIChjaGFydHMubGVuZ3RoID4gMCkge1xuICAgICAgICAgIFlhbWwuc2F2ZShcbiAgICAgICAgICAgIHBhdGguam9pbih0aGlzLm91dGRpciwgYGFwcCR7dGhpcy5vdXRwdXRGaWxlRXh0ZW5zaW9ufWApLCAvLyBUaGVyZSBpcyBubyBcImFwcCBuYW1lXCIsIHNvIHdlIGp1c3QgaGFyZGNvZGUgdGhlIGZpbGUgbmFtZVxuICAgICAgICAgICAgYXBpT2JqZWN0c0xpc3QpO1xuICAgICAgICB9XG4gICAgICAgIGJyZWFrO1xuXG4gICAgICBjYXNlIFlhbWxPdXRwdXRUeXBlLkZJTEVfUEVSX0NIQVJUOlxuICAgICAgICBjb25zdCBuYW1lcjogQ2hhcnROYW1lciA9IGhhc0RlcGVuZGFudENoYXJ0cyA/IG5ldyBJbmRleGVkQ2hhcnROYW1lcigpIDogbmV3IFNpbXBsZUNoYXJ0TmFtZXIoKTtcbiAgICAgICAgZm9yIChjb25zdCBjaGFydCBvZiBjaGFydHMpIHtcbiAgICAgICAgICBjb25zdCBjaGFydE5hbWUgPSBuYW1lci5uYW1lKGNoYXJ0KTtcbiAgICAgICAgICBjb25zdCBvYmplY3RzID0gT2JqZWN0LnZhbHVlcyhjaGFydC50b0pzb24oKSk7XG4gICAgICAgICAgWWFtbC5zYXZlKHBhdGguam9pbih0aGlzLm91dGRpciwgY2hhcnROYW1lICsgdGhpcy5vdXRwdXRGaWxlRXh0ZW5zaW9uKSwgb2JqZWN0cyk7XG4gICAgICAgIH1cbiAgICAgICAgYnJlYWs7XG5cbiAgICAgIGNhc2UgWWFtbE91dHB1dFR5cGUuRklMRV9QRVJfUkVTT1VSQ0U6XG4gICAgICAgIGZvciAoY29uc3QgY2hhcnQgb2YgY2hhcnRzKSB7XG4gICAgICAgICAgY29uc3QgYXBpT2JqZWN0cyA9IE9iamVjdC52YWx1ZXMoY2hhcnQudG9Kc29uKCkpO1xuXG4gICAgICAgICAgYXBpT2JqZWN0cy5mb3JFYWNoKChhcGlPYmplY3QpID0+IHtcbiAgICAgICAgICAgIGlmICghKGFwaU9iamVjdCA9PT0gdW5kZWZpbmVkKSkge1xuICAgICAgICAgICAgICBjb25zdCBmaWxlTmFtZSA9IGAke2Ake2FwaU9iamVjdC5raW5kfS4ke2FwaU9iamVjdC5tZXRhZGF0YS5uYW1lfWBcbiAgICAgICAgICAgICAgICAucmVwbGFjZSgvW14wLTlhLXpBLVotXy5dL2csICcnKX1gO1xuICAgICAgICAgICAgICBZYW1sLnNhdmUocGF0aC5qb2luKHRoaXMub3V0ZGlyLCBmaWxlTmFtZSArIHRoaXMub3V0cHV0RmlsZUV4dGVuc2lvbiksIFthcGlPYmplY3RdKTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9KTtcbiAgICAgICAgfVxuICAgICAgICBicmVhaztcblxuICAgICAgY2FzZSBZYW1sT3V0cHV0VHlwZS5GT0xERVJfUEVSX0NIQVJUX0ZJTEVfUEVSX1JFU09VUkNFOlxuICAgICAgICBjb25zdCBmb2xkZXJOYW1lcjogQ2hhcnROYW1lciA9IGhhc0RlcGVuZGFudENoYXJ0cyA/IG5ldyBJbmRleGVkQ2hhcnRGb2xkZXJOYW1lcigpIDogbmV3IFNpbXBsZUNoYXJ0Rm9sZGVyTmFtZXIoKTtcbiAgICAgICAgZm9yIChjb25zdCBjaGFydCBvZiBjaGFydHMpIHtcbiAgICAgICAgICBjb25zdCBjaGFydE5hbWUgPSBmb2xkZXJOYW1lci5uYW1lKGNoYXJ0KTtcbiAgICAgICAgICBjb25zdCBhcGlPYmplY3RzID0gY2hhcnRUb0t1YmUoY2hhcnQpO1xuICAgICAgICAgIGNvbnN0IGZ1bGxPdXREaXIgPSBwYXRoLmpvaW4odGhpcy5vdXRkaXIsIGNoYXJ0TmFtZSk7XG4gICAgICAgICAgZnMubWtkaXJTeW5jKGZ1bGxPdXREaXIsIHsgcmVjdXJzaXZlOiB0cnVlIH0pO1xuXG4gICAgICAgICAgYXBpT2JqZWN0cy5mb3JFYWNoKChhcGlPYmplY3QpID0+IHtcbiAgICAgICAgICAgIGlmICghKGFwaU9iamVjdCA9PT0gdW5kZWZpbmVkKSkge1xuICAgICAgICAgICAgICBjb25zdCBmaWxlTmFtZSA9IGAke2Ake2FwaU9iamVjdC5raW5kfS4ke2FwaU9iamVjdC5tZXRhZGF0YS5uYW1lfWBcbiAgICAgICAgICAgICAgICAucmVwbGFjZSgvW14wLTlhLXpBLVotXy5dL2csICcnKX1gO1xuICAgICAgICAgICAgICBZYW1sLnNhdmUocGF0aC5qb2luKGZ1bGxPdXREaXIsIGZpbGVOYW1lICsgdGhpcy5vdXRwdXRGaWxlRXh0ZW5zaW9uKSwgW2FwaU9iamVjdC50b0pzb24oKV0pO1xuICAgICAgICAgICAgfVxuICAgICAgICAgIH0pO1xuICAgICAgICB9XG4gICAgICAgIGJyZWFrO1xuXG4gICAgICBkZWZhdWx0OlxuICAgICAgICBicmVhaztcbiAgICB9XG5cbiAgICBpZiAodGhpcy5yZWNvcmRDb25zdHJ1Y3RNZXRhZGF0YSkge1xuICAgICAgY29uc3QgYWxsT2JqZWN0cyA9IHRoaXMuY2hhcnRzLmZsYXRNYXAoY2hhcnRUb0t1YmUpO1xuICAgICAgdGhpcy53cml0ZUNvbnN0cnVjdE1ldGFkYXRhKGFsbE9iamVjdHMpO1xuICAgIH1cblxuICB9XG5cbiAgLyoqXG4gICAqIFN5bnRoZXNpemVzIHRoZSBhcHAgaW50byBhIFlBTUwgc3RyaW5nLlxuICAgKlxuICAgKiBAcmV0dXJucyBBIHN0cmluZyB3aXRoIGFsbCBZQU1MIG9iamVjdHMgYWNyb3NzIGFsbCBjaGFydHMgaW4gdGhpcyBhcHAuXG4gICAqL1xuICBwdWJsaWMgc3ludGhZYW1sKCk6IHN0cmluZyB7XG4gICAgY29uc3QgY2FjaGUgPSBuZXcgU3ludGhSZXF1ZXN0Q2FjaGUoKTtcblxuICAgIHJlc29sdmVEZXBlbmRlbmNpZXModGhpcywgY2FjaGUpO1xuXG4gICAgdmFsaWRhdGUodGhpcywgY2FjaGUpO1xuXG4gICAgY29uc3QgY2hhcnRzID0gdGhpcy5jaGFydHM7XG4gICAgY29uc3QgZG9jczogYW55W10gPSBbXTtcblxuICAgIGZvciAoY29uc3QgY2hhcnQgb2YgY2hhcnRzKSB7XG4gICAgICBkb2NzLnB1c2goLi4uT2JqZWN0LnZhbHVlcyhjaGFydC50b0pzb24oKSkpO1xuICAgIH1cblxuICAgIHJldHVybiBZYW1sLnN0cmluZ2lmeSguLi5kb2NzKTtcbiAgfVxuXG4gIHByaXZhdGUgd3JpdGVDb25zdHJ1Y3RNZXRhZGF0YShhcGlPYmplY3RzOiBBcGlPYmplY3RbXSkge1xuICAgIGNvbnN0IHJlc291cmNlczogeyBba2V5OiBzdHJpbmddOiBhbnkgfSA9IHt9O1xuICAgIGZvciAoY29uc3QgYXBpT2JqZWN0IG9mIGFwaU9iamVjdHMpIHtcbiAgICAgIHJlc291cmNlc1thcGlPYmplY3QubmFtZV0gPSB7IHBhdGg6IGFwaU9iamVjdC5ub2RlLnBhdGggfTtcbiAgICB9XG4gICAgZnMud3JpdGVGaWxlU3luYyhwYXRoLmpvaW4odGhpcy5vdXRkaXIsICdjb25zdHJ1Y3QtbWV0YWRhdGEuanNvbicpLCBKU09OLnN0cmluZ2lmeSh7XG4gICAgICB2ZXJzaW9uOiAnMS4wLjAnLFxuICAgICAgcmVzb3VyY2VzOiByZXNvdXJjZXMsXG4gICAgfSkpO1xuICB9XG59XG5cbmZ1bmN0aW9uIHZhbGlkYXRlKGFwcDogQXBwLCBjYWNoZTogU3ludGhSZXF1ZXN0Q2FjaGUpIHtcbiAgY29uc3QgZXJyb3JzID0gW107XG4gIGZvciAoY29uc3QgY2hpbGQgb2YgY2FjaGUuZmluZEFsbChhcHAubm9kZSkpIHtcbiAgICBjb25zdCBjaGlsZEVycm9ycyA9IGNoaWxkLm5vZGUudmFsaWRhdGUoKTtcbiAgICBmb3IgKGNvbnN0IGVycm9yIG9mIGNoaWxkRXJyb3JzKSB7XG4gICAgICBlcnJvcnMucHVzaChgWyR7Y2hpbGQubm9kZS5wYXRofV0gJHtlcnJvcn1gKTtcbiAgICB9XG4gIH1cblxuICBpZiAoZXJyb3JzLmxlbmd0aCA+IDApIHtcbiAgICB0aHJvdyBuZXcgRXJyb3IoYFZhbGlkYXRpb24gZmFpbGVkIHdpdGggdGhlIGZvbGxvd2luZyBlcnJvcnM6XFxuICAke2Vycm9ycy5qb2luKCdcXG4gICcpfWApO1xuICB9XG59XG5cbmZ1bmN0aW9uIGJ1aWxkRGVwZW5kZW5jaWVzKGFwcDogQXBwLCBjYWNoZTogU3ludGhSZXF1ZXN0Q2FjaGUpIHtcblxuICBjb25zdCBkZXBzID0gW107XG4gIGZvciAoY29uc3QgY2hpbGQgb2YgY2FjaGUuZmluZEFsbChhcHAubm9kZSkpIHtcbiAgICBmb3IgKGNvbnN0IGRlcCBvZiBjaGlsZC5ub2RlLmRlcGVuZGVuY2llcykge1xuICAgICAgZGVwcy5wdXNoKHsgc291cmNlOiBjaGlsZCwgdGFyZ2V0OiBkZXAgfSk7XG4gICAgfVxuICB9XG5cbiAgcmV0dXJuIGRlcHM7XG5cbn1cblxuZnVuY3Rpb24gcmVzb2x2ZURlcGVuZGVuY2llcyhhcHA6IEFwcCwgY2FjaGU6IFN5bnRoUmVxdWVzdENhY2hlKSB7XG5cbiAgbGV0IGhhc0RlcGVuZGFudENoYXJ0cyA9IGZhbHNlO1xuXG4gIC8vIGNyZWF0ZSBhbiBleHBsaWNpdCBjaGFydCBkZXBlbmRlbmN5IGZyb20gbmVzdGVkIGNoYXJ0IHJlbGF0aW9uc2hpcHNcbiAgZm9yIChjb25zdCBwYXJlbnRDaGFydCBvZiBjYWNoZS5maW5kQWxsKGFwcC5ub2RlKS5maWx0ZXIoeCA9PiB4IGluc3RhbmNlb2YgQ2hhcnQpKSB7XG4gICAgZm9yIChjb25zdCBjaGlsZENoYXJ0IG9mIHBhcmVudENoYXJ0Lm5vZGUuY2hpbGRyZW4uZmlsdGVyKHggPT4geCBpbnN0YW5jZW9mIENoYXJ0KSkge1xuICAgICAgcGFyZW50Q2hhcnQubm9kZS5hZGREZXBlbmRlbmN5KGNoaWxkQ2hhcnQpO1xuICAgICAgaGFzRGVwZW5kYW50Q2hhcnRzID0gdHJ1ZTtcbiAgICB9XG4gIH1cblxuICAvLyBjcmVhdGUgYW4gZXhwbGljaXQgY2hhcnQgZGVwZW5kZW5jeSBmcm9tIGltcGxpY2l0IGNvbnN0cnVjdCBkZXBlbmRlbmNpZXNcbiAgZm9yIChjb25zdCBkZXAgb2YgYnVpbGREZXBlbmRlbmNpZXMoYXBwLCBjYWNoZSkpIHtcblxuICAgIGNvbnN0IHNvdXJjZUNoYXJ0ID0gQ2hhcnQub2YoZGVwLnNvdXJjZSk7XG4gICAgY29uc3QgdGFyZ2V0Q2hhcnQgPSBDaGFydC5vZihkZXAudGFyZ2V0KTtcblxuICAgIGlmIChzb3VyY2VDaGFydCAhPT0gdGFyZ2V0Q2hhcnQpIHtcbiAgICAgIHNvdXJjZUNoYXJ0Lm5vZGUuYWRkRGVwZW5kZW5jeSh0YXJnZXRDaGFydCk7XG4gICAgICBoYXNEZXBlbmRhbnRDaGFydHMgPSB0cnVlO1xuICAgIH1cblxuICB9XG5cbiAgLy8gY3JlYXRlIGV4cGxpY2l0IGFwaSBvYmplY3QgZGVwZW5kZW5jaWVzIGZyb20gaW1wbGljaXQgY29uc3RydWN0IGRlcGVuZGVuY2llc1xuICBmb3IgKGNvbnN0IGRlcCBvZiBidWlsZERlcGVuZGVuY2llcyhhcHAsIGNhY2hlKSkge1xuXG4gICAgY29uc3Qgc291cmNlQ2hhcnQgPSBDaGFydC5vZihkZXAuc291cmNlKTtcbiAgICBjb25zdCB0YXJnZXRDaGFydCA9IENoYXJ0Lm9mKGRlcC50YXJnZXQpO1xuXG4gICAgY29uc3QgdGFyZ2V0QXBpT2JqZWN0cyA9IGNhY2hlLmZpbmRBbGwoZGVwLnRhcmdldC5ub2RlKS5maWx0ZXIoYyA9PiBjIGluc3RhbmNlb2YgQXBpT2JqZWN0KS5maWx0ZXIoeCA9PiBDaGFydC5vZih4KSA9PT0gdGFyZ2V0Q2hhcnQpO1xuICAgIGNvbnN0IHNvdXJjZUFwaU9iamVjdHMgPSBjYWNoZS5maW5kQWxsKGRlcC5zb3VyY2Uubm9kZSkuZmlsdGVyKGMgPT4gYyBpbnN0YW5jZW9mIEFwaU9iamVjdCkuZmlsdGVyKHggPT4gQ2hhcnQub2YoeCkgPT09IHNvdXJjZUNoYXJ0KTtcblxuICAgIGZvciAoY29uc3QgdGFyZ2V0IG9mIHRhcmdldEFwaU9iamVjdHMpIHtcbiAgICAgIGZvciAoY29uc3Qgc291cmNlIG9mIHNvdXJjZUFwaU9iamVjdHMpIHtcbiAgICAgICAgaWYgKHRhcmdldCAhPT0gc291cmNlKSB7XG4gICAgICAgICAgc291cmNlLm5vZGUuYWRkRGVwZW5kZW5jeSh0YXJnZXQpO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgcmV0dXJuIGhhc0RlcGVuZGFudENoYXJ0cztcblxufVxuXG5mdW5jdGlvbiBjaGFydFRvS3ViZShjaGFydDogQ2hhcnQpIHtcbiAgcmV0dXJuIG5ldyBEZXBlbmRlbmN5R3JhcGgoY2hhcnQubm9kZSkudG9wb2xvZ3koKVxuICAgIC5maWx0ZXIoeCA9PiB4IGluc3RhbmNlb2YgQXBpT2JqZWN0KVxuICAgIC5maWx0ZXIoeCA9PiBDaGFydC5vZih4KSA9PT0gY2hhcnQpIC8vIGluY2x1ZGUgYW4gb2JqZWN0IG9ubHkgaW4gaXRzIGNsb3Nlc3QgcGFyZW50IGNoYXJ0XG4gICAgLm1hcCh4ID0+ICh4IGFzIEFwaU9iamVjdCkpO1xufVxuXG5pbnRlcmZhY2UgQ2hhcnROYW1lciB7XG4gIG5hbWUoY2hhcnQ6IENoYXJ0KTogc3RyaW5nO1xufVxuXG5jbGFzcyBTaW1wbGVDaGFydE5hbWVyIGltcGxlbWVudHMgQ2hhcnROYW1lciB7XG4gIGNvbnN0cnVjdG9yKCkge1xuICB9XG5cbiAgcHVibGljIG5hbWUoY2hhcnQ6IENoYXJ0KSB7XG4gICAgcmV0dXJuIGAke05hbWVzLnRvRG5zTGFiZWwoY2hhcnQpfWA7XG4gIH1cbn1cblxuY2xhc3MgSW5kZXhlZENoYXJ0TmFtZXIgZXh0ZW5kcyBTaW1wbGVDaGFydE5hbWVyIGltcGxlbWVudHMgQ2hhcnROYW1lciB7XG4gIHByaXZhdGUgaW5kZXg6IG51bWJlciA9IDA7XG4gIGNvbnN0cnVjdG9yKCkge1xuICAgIHN1cGVyKCk7XG4gIH1cblxuICBwdWJsaWMgbmFtZShjaGFydDogQ2hhcnQpIHtcbiAgICBjb25zdCBuYW1lID0gYCR7dGhpcy5pbmRleC50b1N0cmluZygpLnBhZFN0YXJ0KDQsICcwJyl9LSR7c3VwZXIubmFtZShjaGFydCl9YDtcbiAgICB0aGlzLmluZGV4Kys7XG4gICAgcmV0dXJuIG5hbWU7XG4gIH1cbn1cblxuY2xhc3MgU2ltcGxlQ2hhcnRGb2xkZXJOYW1lciBpbXBsZW1lbnRzIENoYXJ0TmFtZXIge1xuICBjb25zdHJ1Y3RvcigpIHtcbiAgfVxuXG4gIHB1YmxpYyBuYW1lKGNoYXJ0OiBDaGFydCkge1xuICAgIHJldHVybiBOYW1lcy50b0Ruc0xhYmVsKGNoYXJ0KTtcbiAgfVxufVxuXG5jbGFzcyBJbmRleGVkQ2hhcnRGb2xkZXJOYW1lciBleHRlbmRzIFNpbXBsZUNoYXJ0Rm9sZGVyTmFtZXIgaW1wbGVtZW50cyBDaGFydE5hbWVyIHtcbiAgcHJpdmF0ZSBpbmRleDogbnVtYmVyID0gMDtcbiAgY29uc3RydWN0b3IoKSB7XG4gICAgc3VwZXIoKTtcbiAgfVxuXG4gIHB1YmxpYyBuYW1lKGNoYXJ0OiBDaGFydCkge1xuICAgIGNvbnN0IG5hbWUgPSBgJHt0aGlzLmluZGV4LnRvU3RyaW5nKCkucGFkU3RhcnQoNCwgJzAnKX0tJHtzdXBlci5uYW1lKGNoYXJ0KX1gO1xuICAgIHRoaXMuaW5kZXgrKztcbiAgICByZXR1cm4gbmFtZTtcbiAgfVxufVxuIl19