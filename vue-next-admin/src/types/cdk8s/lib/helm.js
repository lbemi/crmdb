"use strict";
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
exports.Helm = void 0;
const JSII_RTTI_SYMBOL_1 = Symbol.for("jsii.rtti");
const fs = require("fs");
const os = require("os");
const path = require("path");
const _child_process_1 = require("./_child_process");
const include_1 = require("./include");
const names_1 = require("./names");
const yaml_1 = require("./yaml");
const MAX_HELM_BUFFER = 10 * 1024 * 1024;
/**
 * Represents a Helm deployment.
 *
 * Use this construct to import an existing Helm chart and incorporate it into your constructs.
 */
class Helm extends include_1.Include {
    constructor(scope, id, props) {
        const workdir = fs.mkdtempSync(path.join(os.tmpdir(), 'cdk8s-helm-'));
        const args = new Array();
        args.push('template');
        // values
        if (props.values && Object.keys(props.values).length > 0) {
            const valuesPath = path.join(workdir, 'overrides.yaml');
            fs.writeFileSync(valuesPath, yaml_1.Yaml.stringify(props.values));
            args.push('-f', valuesPath);
        }
        if (props.repo) {
            args.push('--repo', props.repo);
        }
        if (props.version) {
            args.push('--version', props.version);
        }
        if (props.namespace) {
            args.push('--namespace', props.namespace);
        }
        // custom flags
        if (props.helmFlags) {
            args.push(...props.helmFlags);
        }
        // release name
        // constraints: https://github.com/helm/helm/issues/6006
        const releaseName = props.releaseName ?? names_1.Names.toDnsLabel(scope, { maxLen: 53, extra: [id] });
        args.push(releaseName);
        // chart
        args.push(props.chart);
        const prog = props.helmExecutable ?? 'helm';
        const outputFile = renderTemplate(workdir, prog, args);
        super(scope, id, { url: outputFile });
        this.releaseName = releaseName;
    }
}
exports.Helm = Helm;
_a = JSII_RTTI_SYMBOL_1;
Helm[_a] = { fqn: "cdk8s.Helm", version: "2.68.60" };
function renderTemplate(workdir, prog, args) {
    const helm = _child_process_1._child_process.spawnSync(prog, args, {
        maxBuffer: MAX_HELM_BUFFER,
    });
    if (helm.error) {
        const err = helm.error.message;
        if (err.includes('ENOENT')) {
            throw new Error(`unable to execute '${prog}' to render Helm chart. Is it installed on your system?`);
        }
        throw new Error(`error while rendering a helm chart: ${err}`);
    }
    if (helm.status !== 0) {
        throw new Error(helm.stderr.toString());
    }
    const outputFile = path.join(workdir, 'chart.yaml');
    const stdout = helm.stdout;
    fs.writeFileSync(outputFile, stdout);
    return outputFile;
}
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiaGVsbS5qcyIsInNvdXJjZVJvb3QiOiIiLCJzb3VyY2VzIjpbIi4uL3NyYy9oZWxtLnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7Ozs7O0FBQUEseUJBQXlCO0FBQ3pCLHlCQUF5QjtBQUN6Qiw2QkFBNkI7QUFFN0IscURBQWtEO0FBQ2xELHVDQUFvQztBQUNwQyxtQ0FBZ0M7QUFDaEMsaUNBQThCO0FBRTlCLE1BQU0sZUFBZSxHQUFHLEVBQUUsR0FBRyxJQUFJLEdBQUcsSUFBSSxDQUFDO0FBcUV6Qzs7OztHQUlHO0FBQ0gsTUFBYSxJQUFLLFNBQVEsaUJBQU87SUFNL0IsWUFBWSxLQUFnQixFQUFFLEVBQVUsRUFBRSxLQUFnQjtRQUN4RCxNQUFNLE9BQU8sR0FBRyxFQUFFLENBQUMsV0FBVyxDQUFDLElBQUksQ0FBQyxJQUFJLENBQUMsRUFBRSxDQUFDLE1BQU0sRUFBRSxFQUFFLGFBQWEsQ0FBQyxDQUFDLENBQUM7UUFFdEUsTUFBTSxJQUFJLEdBQUcsSUFBSSxLQUFLLEVBQVUsQ0FBQztRQUNqQyxJQUFJLENBQUMsSUFBSSxDQUFDLFVBQVUsQ0FBQyxDQUFDO1FBRXRCLFNBQVM7UUFDVCxJQUFJLEtBQUssQ0FBQyxNQUFNLElBQUksTUFBTSxDQUFDLElBQUksQ0FBQyxLQUFLLENBQUMsTUFBTSxDQUFDLENBQUMsTUFBTSxHQUFHLENBQUMsRUFBRSxDQUFDO1lBQ3pELE1BQU0sVUFBVSxHQUFHLElBQUksQ0FBQyxJQUFJLENBQUMsT0FBTyxFQUFFLGdCQUFnQixDQUFDLENBQUM7WUFDeEQsRUFBRSxDQUFDLGFBQWEsQ0FBQyxVQUFVLEVBQUUsV0FBSSxDQUFDLFNBQVMsQ0FBQyxLQUFLLENBQUMsTUFBTSxDQUFDLENBQUMsQ0FBQztZQUMzRCxJQUFJLENBQUMsSUFBSSxDQUFDLElBQUksRUFBRSxVQUFVLENBQUMsQ0FBQztRQUM5QixDQUFDO1FBRUQsSUFBSSxLQUFLLENBQUMsSUFBSSxFQUFFLENBQUM7WUFDZixJQUFJLENBQUMsSUFBSSxDQUFDLFFBQVEsRUFBRSxLQUFLLENBQUMsSUFBSSxDQUFDLENBQUM7UUFDbEMsQ0FBQztRQUVELElBQUksS0FBSyxDQUFDLE9BQU8sRUFBRSxDQUFDO1lBQ2xCLElBQUksQ0FBQyxJQUFJLENBQUMsV0FBVyxFQUFFLEtBQUssQ0FBQyxPQUFPLENBQUMsQ0FBQztRQUN4QyxDQUFDO1FBRUQsSUFBSSxLQUFLLENBQUMsU0FBUyxFQUFFLENBQUM7WUFDcEIsSUFBSSxDQUFDLElBQUksQ0FBQyxhQUFhLEVBQUUsS0FBSyxDQUFDLFNBQVMsQ0FBQyxDQUFDO1FBQzVDLENBQUM7UUFFRCxlQUFlO1FBQ2YsSUFBSSxLQUFLLENBQUMsU0FBUyxFQUFFLENBQUM7WUFDcEIsSUFBSSxDQUFDLElBQUksQ0FBQyxHQUFHLEtBQUssQ0FBQyxTQUFTLENBQUMsQ0FBQztRQUNoQyxDQUFDO1FBRUQsZUFBZTtRQUNmLHdEQUF3RDtRQUN4RCxNQUFNLFdBQVcsR0FBRyxLQUFLLENBQUMsV0FBVyxJQUFJLGFBQUssQ0FBQyxVQUFVLENBQUMsS0FBSyxFQUFFLEVBQUUsTUFBTSxFQUFFLEVBQUUsRUFBRSxLQUFLLEVBQUUsQ0FBQyxFQUFFLENBQUMsRUFBRSxDQUFDLENBQUM7UUFDOUYsSUFBSSxDQUFDLElBQUksQ0FBQyxXQUFXLENBQUMsQ0FBQztRQUV2QixRQUFRO1FBQ1IsSUFBSSxDQUFDLElBQUksQ0FBQyxLQUFLLENBQUMsS0FBSyxDQUFDLENBQUM7UUFFdkIsTUFBTSxJQUFJLEdBQUcsS0FBSyxDQUFDLGNBQWMsSUFBSSxNQUFNLENBQUM7UUFDNUMsTUFBTSxVQUFVLEdBQUcsY0FBYyxDQUFDLE9BQU8sRUFBRSxJQUFJLEVBQUUsSUFBSSxDQUFDLENBQUM7UUFFdkQsS0FBSyxDQUFDLEtBQUssRUFBRSxFQUFFLEVBQUUsRUFBRSxHQUFHLEVBQUUsVUFBVSxFQUFFLENBQUMsQ0FBQztRQUV0QyxJQUFJLENBQUMsV0FBVyxHQUFHLFdBQVcsQ0FBQztJQUNqQyxDQUFDOztBQWxESCxvQkFtREM7OztBQUVELFNBQVMsY0FBYyxDQUFDLE9BQWUsRUFBRSxJQUFZLEVBQUUsSUFBYztJQUNuRSxNQUFNLElBQUksR0FBRywrQkFBYyxDQUFDLFNBQVMsQ0FBQyxJQUFJLEVBQUUsSUFBSSxFQUFFO1FBQ2hELFNBQVMsRUFBRSxlQUFlO0tBQzNCLENBQUMsQ0FBQztJQUVILElBQUksSUFBSSxDQUFDLEtBQUssRUFBRSxDQUFDO1FBQ2YsTUFBTSxHQUFHLEdBQUcsSUFBSSxDQUFDLEtBQUssQ0FBQyxPQUFPLENBQUM7UUFDL0IsSUFBSSxHQUFHLENBQUMsUUFBUSxDQUFDLFFBQVEsQ0FBQyxFQUFFLENBQUM7WUFDM0IsTUFBTSxJQUFJLEtBQUssQ0FBQyxzQkFBc0IsSUFBSSx5REFBeUQsQ0FBQyxDQUFDO1FBQ3ZHLENBQUM7UUFFRCxNQUFNLElBQUksS0FBSyxDQUFDLHVDQUF1QyxHQUFHLEVBQUUsQ0FBQyxDQUFDO0lBQ2hFLENBQUM7SUFFRCxJQUFJLElBQUksQ0FBQyxNQUFNLEtBQUssQ0FBQyxFQUFFLENBQUM7UUFDdEIsTUFBTSxJQUFJLEtBQUssQ0FBQyxJQUFJLENBQUMsTUFBTSxDQUFDLFFBQVEsRUFBRSxDQUFDLENBQUM7SUFDMUMsQ0FBQztJQUVELE1BQU0sVUFBVSxHQUFHLElBQUksQ0FBQyxJQUFJLENBQUMsT0FBTyxFQUFFLFlBQVksQ0FBQyxDQUFDO0lBQ3BELE1BQU0sTUFBTSxHQUFHLElBQUksQ0FBQyxNQUFNLENBQUM7SUFDM0IsRUFBRSxDQUFDLGFBQWEsQ0FBQyxVQUFVLEVBQUUsTUFBTSxDQUFDLENBQUM7SUFDckMsT0FBTyxVQUFVLENBQUM7QUFDcEIsQ0FBQyIsInNvdXJjZXNDb250ZW50IjpbImltcG9ydCAqIGFzIGZzIGZyb20gJ2ZzJztcbmltcG9ydCAqIGFzIG9zIGZyb20gJ29zJztcbmltcG9ydCAqIGFzIHBhdGggZnJvbSAncGF0aCc7XG5pbXBvcnQgeyBDb25zdHJ1Y3QgfSBmcm9tICdjb25zdHJ1Y3RzJztcbmltcG9ydCB7IF9jaGlsZF9wcm9jZXNzIH0gZnJvbSAnLi9fY2hpbGRfcHJvY2Vzcyc7XG5pbXBvcnQgeyBJbmNsdWRlIH0gZnJvbSAnLi9pbmNsdWRlJztcbmltcG9ydCB7IE5hbWVzIH0gZnJvbSAnLi9uYW1lcyc7XG5pbXBvcnQgeyBZYW1sIH0gZnJvbSAnLi95YW1sJztcblxuY29uc3QgTUFYX0hFTE1fQlVGRkVSID0gMTAgKiAxMDI0ICogMTAyNDtcblxuLyoqXG4gKiBPcHRpb25zIGZvciBgSGVsbWAuXG4gKi9cbmV4cG9ydCBpbnRlcmZhY2UgSGVsbVByb3BzIHtcbiAgLyoqXG4gICAqIFRoZSBjaGFydCBuYW1lIHRvIHVzZS4gSXQgY2FuIGJlIGEgY2hhcnQgZnJvbSBhIGhlbG0gcmVwb3NpdG9yeSBvciBhIGxvY2FsIGRpcmVjdG9yeS5cbiAgICpcbiAgICogVGhpcyBuYW1lIGlzIHBhc3NlZCB0byBgaGVsbSB0ZW1wbGF0ZWAgYW5kIGhhcyBhbGwgdGhlIHJlbGV2YW50IHNlbWFudGljcy5cbiAgICpcbiAgICogQGV4YW1wbGUgXCIuL215c3FsXCJcbiAgICogQGV4YW1wbGUgXCJiaXRuYW1pL3JlZGlzXCJcbiAgICovXG4gIHJlYWRvbmx5IGNoYXJ0OiBzdHJpbmc7XG5cbiAgLyoqXG4gICAqIENoYXJ0IHJlcG9zaXRvcnkgdXJsIHdoZXJlIHRvIGxvY2F0ZSB0aGUgcmVxdWVzdGVkIGNoYXJ0XG4gICAqL1xuICByZWFkb25seSByZXBvPzogc3RyaW5nO1xuXG4gIC8qKlxuICAgKiBWZXJzaW9uIGNvbnN0cmFpbnQgZm9yIHRoZSBjaGFydCB2ZXJzaW9uIHRvIHVzZS5cbiAgICogVGhpcyBjb25zdHJhaW50IGNhbiBiZSBhIHNwZWNpZmljIHRhZyAoZS5nLiAxLjEuMSlcbiAgICogb3IgaXQgbWF5IHJlZmVyZW5jZSBhIHZhbGlkIHJhbmdlIChlLmcuIF4yLjAuMCkuXG4gICAqIElmIHRoaXMgaXMgbm90IHNwZWNpZmllZCwgdGhlIGxhdGVzdCB2ZXJzaW9uIGlzIHVzZWRcbiAgICpcbiAgICogVGhpcyBuYW1lIGlzIHBhc3NlZCB0byBgaGVsbSB0ZW1wbGF0ZSAtLXZlcnNpb25gIGFuZCBoYXMgYWxsIHRoZSByZWxldmFudCBzZW1hbnRpY3MuXG4gICAqXG4gICAqIEBleGFtcGxlIFwiMS4xLjFcIlxuICAgKiBAZXhhbXBsZSBcIl4yLjAuMFwiXG4gICAqL1xuICByZWFkb25seSB2ZXJzaW9uPzogc3RyaW5nO1xuXG4gIC8qKlxuICAgKiBTY29wZSBhbGwgcmVzb3VyY2VzIGluIHRvIGEgZ2l2ZW4gbmFtZXNwYWNlLlxuICAgKi9cbiAgcmVhZG9ubHkgbmFtZXNwYWNlPzogc3RyaW5nO1xuXG4gIC8qKlxuICAgKiBUaGUgcmVsZWFzZSBuYW1lLlxuICAgKlxuICAgKiBAc2VlIGh0dHBzOi8vaGVsbS5zaC9kb2NzL2ludHJvL3VzaW5nX2hlbG0vI3RocmVlLWJpZy1jb25jZXB0c1xuICAgKiBAZGVmYXVsdCAtIGlmIHVuc3BlY2lmaWVkLCBhIG5hbWUgd2lsbCBiZSBhbGxvY2F0ZWQgYmFzZWQgb24gdGhlIGNvbnN0cnVjdCBwYXRoXG4gICAqL1xuICByZWFkb25seSByZWxlYXNlTmFtZT86IHN0cmluZztcblxuICAvKipcbiAgICogVmFsdWVzIHRvIHBhc3MgdG8gdGhlIGNoYXJ0LlxuICAgKlxuICAgKiBAZGVmYXVsdCAtIElmIG5vIHZhbHVlcyBhcmUgc3BlY2lmaWVkLCBjaGFydCB3aWxsIHVzZSB0aGUgZGVmYXVsdHMuXG4gICAqL1xuICByZWFkb25seSB2YWx1ZXM/OiB7IFtrZXk6IHN0cmluZ106IGFueSB9O1xuXG4gIC8qKlxuICAgKiBUaGUgbG9jYWwgaGVsbSBleGVjdXRhYmxlIHRvIHVzZSBpbiBvcmRlciB0byBjcmVhdGUgdGhlIG1hbmlmZXN0IHRoZSBjaGFydC5cbiAgICpcbiAgICogQGRlZmF1bHQgXCJoZWxtXCJcbiAgICovXG4gIHJlYWRvbmx5IGhlbG1FeGVjdXRhYmxlPzogc3RyaW5nO1xuXG4gIC8qKlxuICAgKiBBZGRpdGlvbmFsIGZsYWdzIHRvIGFkZCB0byB0aGUgYGhlbG1gIGV4ZWN1dGlvbi5cbiAgICpcbiAgICogQGRlZmF1bHQgW11cbiAgICovXG4gIHJlYWRvbmx5IGhlbG1GbGFncz86IHN0cmluZ1tdO1xufVxuXG4vKipcbiAqIFJlcHJlc2VudHMgYSBIZWxtIGRlcGxveW1lbnQuXG4gKlxuICogVXNlIHRoaXMgY29uc3RydWN0IHRvIGltcG9ydCBhbiBleGlzdGluZyBIZWxtIGNoYXJ0IGFuZCBpbmNvcnBvcmF0ZSBpdCBpbnRvIHlvdXIgY29uc3RydWN0cy5cbiAqL1xuZXhwb3J0IGNsYXNzIEhlbG0gZXh0ZW5kcyBJbmNsdWRlIHtcbiAgLyoqXG4gICAqIFRoZSBoZWxtIHJlbGVhc2UgbmFtZS5cbiAgICovXG4gIHB1YmxpYyByZWFkb25seSByZWxlYXNlTmFtZTogc3RyaW5nO1xuXG4gIGNvbnN0cnVjdG9yKHNjb3BlOiBDb25zdHJ1Y3QsIGlkOiBzdHJpbmcsIHByb3BzOiBIZWxtUHJvcHMpIHtcbiAgICBjb25zdCB3b3JrZGlyID0gZnMubWtkdGVtcFN5bmMocGF0aC5qb2luKG9zLnRtcGRpcigpLCAnY2RrOHMtaGVsbS0nKSk7XG5cbiAgICBjb25zdCBhcmdzID0gbmV3IEFycmF5PHN0cmluZz4oKTtcbiAgICBhcmdzLnB1c2goJ3RlbXBsYXRlJyk7XG5cbiAgICAvLyB2YWx1ZXNcbiAgICBpZiAocHJvcHMudmFsdWVzICYmIE9iamVjdC5rZXlzKHByb3BzLnZhbHVlcykubGVuZ3RoID4gMCkge1xuICAgICAgY29uc3QgdmFsdWVzUGF0aCA9IHBhdGguam9pbih3b3JrZGlyLCAnb3ZlcnJpZGVzLnlhbWwnKTtcbiAgICAgIGZzLndyaXRlRmlsZVN5bmModmFsdWVzUGF0aCwgWWFtbC5zdHJpbmdpZnkocHJvcHMudmFsdWVzKSk7XG4gICAgICBhcmdzLnB1c2goJy1mJywgdmFsdWVzUGF0aCk7XG4gICAgfVxuXG4gICAgaWYgKHByb3BzLnJlcG8pIHtcbiAgICAgIGFyZ3MucHVzaCgnLS1yZXBvJywgcHJvcHMucmVwbyk7XG4gICAgfVxuXG4gICAgaWYgKHByb3BzLnZlcnNpb24pIHtcbiAgICAgIGFyZ3MucHVzaCgnLS12ZXJzaW9uJywgcHJvcHMudmVyc2lvbik7XG4gICAgfVxuXG4gICAgaWYgKHByb3BzLm5hbWVzcGFjZSkge1xuICAgICAgYXJncy5wdXNoKCctLW5hbWVzcGFjZScsIHByb3BzLm5hbWVzcGFjZSk7XG4gICAgfVxuXG4gICAgLy8gY3VzdG9tIGZsYWdzXG4gICAgaWYgKHByb3BzLmhlbG1GbGFncykge1xuICAgICAgYXJncy5wdXNoKC4uLnByb3BzLmhlbG1GbGFncyk7XG4gICAgfVxuXG4gICAgLy8gcmVsZWFzZSBuYW1lXG4gICAgLy8gY29uc3RyYWludHM6IGh0dHBzOi8vZ2l0aHViLmNvbS9oZWxtL2hlbG0vaXNzdWVzLzYwMDZcbiAgICBjb25zdCByZWxlYXNlTmFtZSA9IHByb3BzLnJlbGVhc2VOYW1lID8/IE5hbWVzLnRvRG5zTGFiZWwoc2NvcGUsIHsgbWF4TGVuOiA1MywgZXh0cmE6IFtpZF0gfSk7XG4gICAgYXJncy5wdXNoKHJlbGVhc2VOYW1lKTtcblxuICAgIC8vIGNoYXJ0XG4gICAgYXJncy5wdXNoKHByb3BzLmNoYXJ0KTtcblxuICAgIGNvbnN0IHByb2cgPSBwcm9wcy5oZWxtRXhlY3V0YWJsZSA/PyAnaGVsbSc7XG4gICAgY29uc3Qgb3V0cHV0RmlsZSA9IHJlbmRlclRlbXBsYXRlKHdvcmtkaXIsIHByb2csIGFyZ3MpO1xuXG4gICAgc3VwZXIoc2NvcGUsIGlkLCB7IHVybDogb3V0cHV0RmlsZSB9KTtcblxuICAgIHRoaXMucmVsZWFzZU5hbWUgPSByZWxlYXNlTmFtZTtcbiAgfVxufVxuXG5mdW5jdGlvbiByZW5kZXJUZW1wbGF0ZSh3b3JrZGlyOiBzdHJpbmcsIHByb2c6IHN0cmluZywgYXJnczogc3RyaW5nW10pIHtcbiAgY29uc3QgaGVsbSA9IF9jaGlsZF9wcm9jZXNzLnNwYXduU3luYyhwcm9nLCBhcmdzLCB7XG4gICAgbWF4QnVmZmVyOiBNQVhfSEVMTV9CVUZGRVIsXG4gIH0pO1xuXG4gIGlmIChoZWxtLmVycm9yKSB7XG4gICAgY29uc3QgZXJyID0gaGVsbS5lcnJvci5tZXNzYWdlO1xuICAgIGlmIChlcnIuaW5jbHVkZXMoJ0VOT0VOVCcpKSB7XG4gICAgICB0aHJvdyBuZXcgRXJyb3IoYHVuYWJsZSB0byBleGVjdXRlICcke3Byb2d9JyB0byByZW5kZXIgSGVsbSBjaGFydC4gSXMgaXQgaW5zdGFsbGVkIG9uIHlvdXIgc3lzdGVtP2ApO1xuICAgIH1cblxuICAgIHRocm93IG5ldyBFcnJvcihgZXJyb3Igd2hpbGUgcmVuZGVyaW5nIGEgaGVsbSBjaGFydDogJHtlcnJ9YCk7XG4gIH1cblxuICBpZiAoaGVsbS5zdGF0dXMgIT09IDApIHtcbiAgICB0aHJvdyBuZXcgRXJyb3IoaGVsbS5zdGRlcnIudG9TdHJpbmcoKSk7XG4gIH1cblxuICBjb25zdCBvdXRwdXRGaWxlID0gcGF0aC5qb2luKHdvcmtkaXIsICdjaGFydC55YW1sJyk7XG4gIGNvbnN0IHN0ZG91dCA9IGhlbG0uc3Rkb3V0O1xuICBmcy53cml0ZUZpbGVTeW5jKG91dHB1dEZpbGUsIHN0ZG91dCk7XG4gIHJldHVybiBvdXRwdXRGaWxlO1xufVxuIl19