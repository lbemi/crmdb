"use strict";
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
exports.Names = void 0;
const JSII_RTTI_SYMBOL_1 = Symbol.for("jsii.rtti");
const crypto = require("crypto");
const MAX_LEN = 63;
const VALIDATE = /^[0-9a-z-]+$/;
const VALIDATE_LABEL_VALUE = /^(([0-9a-zA-Z][0-9a-zA-Z-_.]*)?[0-9a-zA-Z])?$/;
const HASH_LEN = 8;
/**
 * Utilities for generating unique and stable names.
 */
class Names {
    /**
     * Generates a unique and stable name compatible DNS_LABEL from RFC-1123 from
     * a path.
     *
     * The generated name will:
     *  - contain at most 63 characters
     *  - contain only lowercase alphanumeric characters or ‘-’
     *  - start with an alphanumeric character
     *  - end with an alphanumeric character
     *
     * The generated name will have the form:
     *  <comp0>-<comp1>-..-<compN>-<short-hash>
     *
     * Where <comp> are the path components (assuming they are is separated by
     * "/").
     *
     * Note that if the total length is longer than 63 characters, we will trim
     * the first components since the last components usually encode more meaning.
     *
     * @link https://tools.ietf.org/html/rfc1123
     *
     * @param scope The construct for which to render the DNS label
     * @param options Name options
     * @throws if any of the components do not adhere to naming constraints or
     * length.
     */
    static toDnsLabel(scope, options = {}) {
        const maxLen = options.maxLen ?? MAX_LEN;
        const delim = options.delimiter ?? '-';
        const include_hash = options.includeHash ?? true;
        if (maxLen < HASH_LEN && include_hash) {
            throw new Error(`minimum max length for object names is ${HASH_LEN} (required for hash)`);
        }
        const node = scope.node;
        let components = node.path.split('/');
        components.push(...options.extra ?? []);
        // special case: if we only have one component in our path and it adheres to DNS_NAME, we don't decorate it
        if (components.length === 1 && VALIDATE.test(components[0]) && components[0].length <= maxLen) {
            return components[0];
        }
        // okay, now we need to normalize all components to adhere to DNS_NAME and append the hash of the full path.
        components = components.map(c => normalizeToDnsName(c, maxLen));
        if (include_hash) {
            components.push(calcHash(node, HASH_LEN));
        }
        return toHumanForm(components, delim, maxLen);
    }
    /**
     * Generates a unique and stable name compatible label key name segment and
     * label value from a path.
     *
     * The name segment is required and must be 63 characters or less, beginning
     * and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-),
     * underscores (_), dots (.), and alphanumerics between.
     *
     * Valid label values must be 63 characters or less and must be empty or
     * begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes
     * (-), underscores (_), dots (.), and alphanumerics between.
     *
     * The generated name will have the form:
     *  <comp0><delim><comp1><delim>..<delim><compN><delim><short-hash>
     *
     * Where <comp> are the path components (assuming they are is separated by
     * "/").
     *
     * Note that if the total length is longer than 63 characters, we will trim
     * the first components since the last components usually encode more meaning.
     *
     * @link https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set
     *
     * @param scope The construct for which to render the DNS label
     * @param options Name options
     * @throws if any of the components do not adhere to naming constraints or
     * length.
     */
    static toLabelValue(scope, options = {}) {
        const maxLen = options.maxLen ?? MAX_LEN;
        const delim = options.delimiter ?? '-';
        const include_hash = options.includeHash ?? true;
        if (maxLen < HASH_LEN && include_hash) {
            throw new Error(`minimum max length for label is ${HASH_LEN} (required for hash)`);
        }
        if (/[^0-9a-zA-Z-_.]/.test(delim)) {
            throw new Error('delim should not contain "[^0-9a-zA-Z-_.]"');
        }
        const node = scope.node;
        let components = node.path.split('/');
        components.push(...options.extra ?? []);
        // special case: if we only have one component in our path and it adheres to DNS_NAME, we don't decorate it
        if (components.length === 1 && VALIDATE_LABEL_VALUE.test(components[0]) && components[0].length <= maxLen) {
            return components[0];
        }
        // okay, now we need to normalize all components to adhere to label and append the hash of the full path.
        components = components.map(c => normalizeToLabelValue(c, maxLen));
        if (include_hash) {
            components.push(calcHash(node, HASH_LEN));
        }
        const result = toHumanForm(components, delim, maxLen);
        // slicing might let '-', '_', '.' be in the start of the result.
        return result.replace(/^[^0-9a-zA-Z]+/, '');
    }
    /* istanbul ignore next */
    constructor() {
        return;
    }
}
exports.Names = Names;
_a = JSII_RTTI_SYMBOL_1;
Names[_a] = { fqn: "cdk8s.Names", version: "2.68.60" };
function omitDuplicates(value, index, components) {
    return value !== components[index - 1];
}
function omitDefaultChild(value, _, __) {
    return value.toLowerCase() !== 'resource' && value.toLowerCase() !== 'default';
}
function toHumanForm(components, delim, maxLen) {
    return components.reverse()
        .filter(omitDuplicates)
        .join('/')
        .slice(0, maxLen)
        .split('/')
        .reverse()
        .filter(x => x)
        .join(delim)
        .split(delim)
        .filter(x => x)
        .filter(omitDefaultChild)
        .join(delim);
}
function normalizeToDnsName(c, maxLen) {
    return c
        .toLocaleLowerCase() // lower case
        .replace(/[^0-9a-zA-Z-_.]/g, '') // remove non-allowed characters
        .substr(0, maxLen); // trim to maxLength
}
function calcHash(node, maxLen) {
    if (process.env.CDK8S_LEGACY_HASH) {
        const hash = crypto.createHash('sha256');
        hash.update(node.path);
        return hash.digest('hex').slice(0, maxLen);
    }
    return node.addr.substring(0, HASH_LEN);
}
function normalizeToLabelValue(c, maxLen) {
    return c
        .replace(/[^0-9a-zA-Z-_.]/g, '') // remove non-allowed characters
        .substr(0, maxLen); // trim to maxLength
}
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoibmFtZXMuanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyIuLi9zcmMvbmFtZXMudHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6Ijs7Ozs7QUFBQSxpQ0FBaUM7QUFHakMsTUFBTSxPQUFPLEdBQUcsRUFBRSxDQUFDO0FBQ25CLE1BQU0sUUFBUSxHQUFHLGNBQWMsQ0FBQztBQUNoQyxNQUFNLG9CQUFvQixHQUFHLCtDQUErQyxDQUFDO0FBQzdFLE1BQU0sUUFBUSxHQUFHLENBQUMsQ0FBQztBQStCbkI7O0dBRUc7QUFDSCxNQUFhLEtBQUs7SUFDaEI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7T0F5Qkc7SUFDSSxNQUFNLENBQUMsVUFBVSxDQUFDLEtBQWdCLEVBQUUsVUFBdUIsRUFBRztRQUNuRSxNQUFNLE1BQU0sR0FBRyxPQUFPLENBQUMsTUFBTSxJQUFJLE9BQU8sQ0FBQztRQUN6QyxNQUFNLEtBQUssR0FBRyxPQUFPLENBQUMsU0FBUyxJQUFJLEdBQUcsQ0FBQztRQUN2QyxNQUFNLFlBQVksR0FBRyxPQUFPLENBQUMsV0FBVyxJQUFJLElBQUksQ0FBQztRQUVqRCxJQUFJLE1BQU0sR0FBRyxRQUFRLElBQUksWUFBWSxFQUFFLENBQUM7WUFDdEMsTUFBTSxJQUFJLEtBQUssQ0FBQywwQ0FBMEMsUUFBUSxzQkFBc0IsQ0FBQyxDQUFDO1FBQzVGLENBQUM7UUFFRCxNQUFNLElBQUksR0FBRyxLQUFLLENBQUMsSUFBSSxDQUFDO1FBRXhCLElBQUksVUFBVSxHQUFHLElBQUksQ0FBQyxJQUFJLENBQUMsS0FBSyxDQUFDLEdBQUcsQ0FBQyxDQUFDO1FBQ3RDLFVBQVUsQ0FBQyxJQUFJLENBQUMsR0FBRyxPQUFPLENBQUMsS0FBSyxJQUFJLEVBQUUsQ0FBQyxDQUFDO1FBRXhDLDJHQUEyRztRQUMzRyxJQUFJLFVBQVUsQ0FBQyxNQUFNLEtBQUssQ0FBQyxJQUFJLFFBQVEsQ0FBQyxJQUFJLENBQUMsVUFBVSxDQUFDLENBQUMsQ0FBQyxDQUFDLElBQUksVUFB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QUssQ0FBQyxDQUFDLEtBQUssQ0FBQyxDQUFDLEVBQUUsTUFBTSxDQUFDLENBQUM7SUFDN0MsQ0FBQztJQUVELE9BQU8sSUFBSSxDQUFDLElBQUksQ0FBQyxTQUFTLENBQUMsQ0FBQyxFQUFFLFFBQVEsQ0FBQyxDQUFDO0FBQzFDLENBQUM7QUFFRCxTQUFTLHFCQUFxQixDQUFDLENBQVMsRUFBRSxNQUFjO0lBQ3RELE9BQU8sQ0FBQztTQUNMLE9BQU8sQ0FBQyxrQkFBa0IsRUFBRSxFQUFFLENBQUMsQ0FBQyxnQ0FBZ0M7U0FDaEUsTUFBTSxDQUFDLENBQUMsRUFBRSxNQUFNLENBQUMsQ0FBQyxDQUFDLG9CQUFvQjtBQUM1QyxDQUFDIiwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0ICogYXMgY3J5cHRvIGZyb20gJ2NyeXB0byc7XG5pbXBvcnQgeyBDb25zdHJ1Y3QsIE5vZGUgfSBmcm9tICdjb25zdHJ1Y3RzJztcblxuY29uc3QgTUFYX0xFTiA9IDYzO1xuY29uc3QgVkFMSURBVEUgPSAvXlswLTlhLXotXSskLztcbmNvbnN0IFZBTElEQVRFX0xBQkVMX1ZBTFVFID0gL14oKFswLTlhLXpBLVpdWzAtOWEtekEtWi1fLl0qKT9bMC05YS16QS1aXSk/JC87XG5jb25zdCBIQVNIX0xFTiA9IDg7XG5cbi8qKlxuICogT3B0aW9ucyBmb3IgbmFtZSBnZW5lcmF0aW9uLlxuICovXG5leHBvcnQgaW50ZXJmYWNlIE5hbWVPcHRpb25zIHtcbiAgLyoqXG4gICAqIE1heGltdW0gYWxsb3dlZCBsZW5ndGggZm9yIHRoZSBuYW1lLlxuICAgKiBAZGVmYXVsdCA2M1xuICAgKi9cbiAgcmVhZG9ubHkgbWF4TGVuPzogbnVtYmVyO1xuXG4gIC8qKlxuICAgKiBFeHRyYSBjb21wb25lbnRzIHRvIGluY2x1ZGUgaW4gdGhlIG5hbWUuXG4gICAqIEBkZWZhdWx0IFtdIHVzZSB0aGUgY29uc3RydWN0IHBhdGggY29tcG9uZW50c1xuICAgKi9cbiAgcmVhZG9ubHkgZXh0cmE/OiBzdHJpbmdbXTtcblxuICAvKipcbiAgICogRGVsaW1pdGVyIHRvIHVzZSBiZXR3ZWVuIGNvbXBvbmVudHMuXG4gICAqIEBkZWZhdWx0IFwiLVwiXG4gICAqL1xuICByZWFkb25seSBkZWxpbWl0ZXI/OiBzdHJpbmc7XG5cbiAgLyoqXG4gICAqIEluY2x1ZGUgYSBzaG9ydCBoYXNoIGFzIGxhc3QgcGFydCBvZiB0aGUgbmFtZS5cbiAgICogQGRlZmF1bHQgdHJ1ZVxuICAgKi9cbiAgcmVhZG9ubHkgaW5jbHVkZUhhc2g/OiBib29sZWFuO1xufVxuXG4vKipcbiAqIFV0aWxpdGllcyBmb3IgZ2VuZXJhdGluZyB1bmlxdWUgYW5kIHN0YWJsZSBuYW1lcy5cbiAqL1xuZXhwb3J0IGNsYXNzIE5hbWVzIHtcbiAgLyoqXG4gICAqIEdlbmVyYXRlcyBhIHVuaXF1ZSBhbmQgc3RhYmxlIG5hbWUgY29tcGF0aWJsZSBETlNfTEFCRUwgZnJvbSBSRkMtMTEyMyBmcm9tXG4gICAqIGEgcGF0aC5cbiAgICpcbiAgICogVGhlIGdlbmVyYXRlZCBuYW1lIHdpbGw6XG4gICAqICAtIGNvbnRhaW4gYXQgbW9zdCA2MyBjaGFyYWN0ZXJzXG4gICAqICAtIGNvbnRhaW4gb25seSBsb3dlcmNhc2UgYWxwaGFudW1lcmljIGNoYXJhY3RlcnMgb3Ig4oCYLeKAmVxuICAgKiAgLSBzdGFydCB3aXRoIGFuIGFscGhhbnVtZXJpYyBjaGFyYWN0ZXJcbiAgICogIC0gZW5kIHdpdGggYW4gYWxwaGFudW1lcmljIGNoYXJhY3RlclxuICAgKlxuICAgKiBUaGUgZ2VuZXJhdGVkIG5hbWUgd2lsbCBoYXZlIHRoZSBmb3JtOlxuICAgKiAgPGNvbXAwPi08Y29tcDE+LS4uLTxjb21wTj4tPHNob3J0LWhhc2g+XG4gICAqXG4gICAqIFdoZXJlIDxjb21wPiBhcmUgdGhlIHBhdGggY29tcG9uZW50cyAoYXNzdW1pbmcgdGhleSBhcmUgaXMgc2VwYXJhdGVkIGJ5XG4gICAqIFwiL1wiKS5cbiAgICpcbiAgICogTm90ZSB0aGF0IGlmIHRoZSB0b3RhbCBsZW5ndGggaXMgbG9uZ2VyIHRoYW4gNjMgY2hhcmFjdGVycywgd2Ugd2lsbCB0cmltXG4gICAqIHRoZSBmaXJzdCBjb21wb25lbnRzIHNpbmNlIHRoZSBsYXN0IGNvbXBvbmVudHMgdXN1YWxseSBlbmNvZGUgbW9yZSBtZWFuaW5nLlxuICAgKlxuICAgKiBAbGluayBodHRwczovL3Rvb2xzLmlldGYub3JnL2h0bWwvcmZjMTEyM1xuICAgKlxuICAgKiBAcGFyYW0gc2NvcGUgVGhlIGNvbnN0cnVjdCBmb3Igd2hpY2ggdG8gcmVuZGVyIHRoZSBETlMgbGFiZWxcbiAgICogQHBhcmFtIG9wdGlvbnMgTmFtZSBvcHRpb25zXG4gICAqIEB0aHJvd3MgaWYgYW55IG9mIHRoZSBjb21wb25lbnRzIGRvIG5vdCBhZGhlcmUgdG8gbmFtaW5nIGNvbnN0cmFpbnRzIG9yXG4gICAqIGxlbmd0aC5cbiAgICovXG4gIHB1YmxpYyBzdGF0aWMgdG9EbnNMYWJlbChzY29wZTogQ29uc3RydWN0LCBvcHRpb25zOiBOYW1lT3B0aW9ucyA9IHsgfSkge1xuICAgIGNvbnN0IG1heExlbiA9IG9wdGlvbnMubWF4TGVuID8/IE1BWF9MRU47XG4gICAgY29uc3QgZGVsaW0gPSBvcHRpb25zLmRlbGltaXRlciA/PyAnLSc7XG4gICAgY29uc3QgaW5jbHVkZV9oYXNoID0gb3B0aW9ucy5pbmNsdWRlSGFzaCA/PyB0cnVlO1xuXG4gICAgaWYgKG1heExlbiA8IEhBU0hfTEVOICYmIGluY2x1ZGVfaGFzaCkge1xuICAgICAgdGhyb3cgbmV3IEVycm9yKGBtaW5pbXVtIG1heCBsZW5ndGggZm9yIG9iamVjdCBuYW1lcyBpcyAke0hBU0hfTEVOfSAocmVxdWlyZWQgZm9yIGhhc2gpYCk7XG4gICAgfVxuXG4gICAgY29uc3Qgbm9kZSA9IHNjb3BlLm5vZGU7XG5cbiAgICBsZXQgY29tcG9uZW50cyA9IG5vZGUucGF0aC5zcGxpdCgnLycpO1xuICAgIGNvbXBvbmVudHMucHVzaCguLi5vcHRpb25zLmV4dHJhID8/IFtdKTtcblxuICAgIC8vIHNwZWNpYWwgY2FzZTogaWYgd2Ugb25seSBoYXZlIG9uZSBjb21wb25lbnQgaW4gb3VyIHBhdGggYW5kIGl0IGFkaGVyZXMgdG8gRE5TX05BTUUsIHdlIGRvbid0IGRlY29yYXRlIGl0XG4gICAgaWYgKGNvbXBvbmVudHMubGVuZ3RoID09PSAxICYmIFZBTElEQVRFLnRlc3QoY29tcG9uZW50c1swXSkgJiYgY29tcG9uZW50c1swXS5sZW5ndGggPD0gbWF4TGVuKSB7XG4gICAgICByZXR1cm4gY29tcG9uZW50c1swXTtcbiAgICB9XG5cbiAgICAvLyBva2F5LCBub3cgd2UgbmVlZCB0byBub3JtYWxpemUgYWxsIGNvbXBvbmVudHMgdG8gYWRoZXJlIHRvIEROU19OQU1FIGFuZCBhcHBlbmQgdGhlIGhhc2ggb2YgdGhlIGZ1bGwgcGF0aC5cbiAgICBjb21wb25lbnRzID0gY29tcG9uZW50cy5tYXAoYyA9PiBub3JtYWxpemVUb0Ruc05hbWUoYywgbWF4TGVuKSk7XG4gICAgaWYgKGluY2x1ZGVfaGFzaCkge1xuICAgICAgY29tcG9uZW50cy5wdXNoKGNhbGNIYXNoKG5vZGUsIEhBU0hfTEVOKSk7XG4gICAgfVxuXG4gICAgcmV0dXJuIHRvSHVtYW5Gb3JtKGNvbXBvbmVudHMsIGRlbGltLCBtYXhMZW4pO1xuICB9XG5cbiAgLyoqXG4gICAqIEdlbmVyYXRlcyBhIHVuaXF1ZSBhbmQgc3RhYmxlIG5hbWUgY29tcGF0aWJsZSBsYWJlbCBrZXkgbmFtZSBzZWdtZW50IGFuZFxuICAgKiBsYWJlbCB2YWx1ZSBmcm9tIGEgcGF0aC5cbiAgICpcbiAgICogVGhlIG5hbWUgc2VnbWVudCBpcyByZXF1aXJlZCBhbmQgbXVzdCBiZSA2MyBjaGFyYWN0ZXJzIG9yIGxlc3MsIGJlZ2lubmluZ1xuICAgKiBhbmQgZW5kaW5nIHdpdGggYW4gYWxwaGFudW1lcmljIGNoYXJhY3RlciAoW2EtejAtOUEtWl0pIHdpdGggZGFzaGVzICgtKSxcbiAgICogdW5kZXJzY29yZXMgKF8pLCBkb3RzICguKSwgYW5kIGFscGhhbnVtZXJpY3MgYmV0d2Vlbi5cbiAgICpcbiAgICogVmFsaWQgbGFiZWwgdmFsdWVzIG11c3QgYmUgNjMgY2hhcmFjdGVycyBvciBsZXNzIGFuZCBtdXN0IGJlIGVtcHR5IG9yXG4gICAqIGJlZ2luIGFuZCBlbmQgd2l0aCBhbiBhbHBoYW51bWVyaWMgY2hhcmFjdGVyIChbYS16MC05QS1aXSkgd2l0aCBkYXNoZXNcbiAgICogKC0pLCB1bmRlcnNjb3JlcyAoXyksIGRvdHMgKC4pLCBhbmQgYWxwaGFudW1lcmljcyBiZXR3ZWVuLlxuICAgKlxuICAgKiBUaGUgZ2VuZXJhdGVkIG5hbWUgd2lsbCBoYXZlIHRoZSBmb3JtOlxuICAgKiAgPGNvbXAwPjxkZWxpbT48Y29tcDE+PGRlbGltPi4uPGRlbGltPjxjb21wTj48ZGVsaW0+PHNob3J0LWhhc2g+XG4gICAqXG4gICAqIFdoZXJlIDxjb21wPiBhcmUgdGhlIHBhdGggY29tcG9uZW50cyAoYXNzdW1pbmcgdGhleSBhcmUgaXMgc2VwYXJhdGVkIGJ5XG4gICAqIFwiL1wiKS5cbiAgICpcbiAgICogTm90ZSB0aGF0IGlmIHRoZSB0b3RhbCBsZW5ndGggaXMgbG9uZ2VyIHRoYW4gNjMgY2hhcmFjdGVycywgd2Ugd2lsbCB0cmltXG4gICAqIHRoZSBmaXJzdCBjb21wb25lbnRzIHNpbmNlIHRoZSBsYXN0IGNvbXBvbmVudHMgdXN1YWxseSBlbmNvZGUgbW9yZSBtZWFuaW5nLlxuICAgKlxuICAgKiBAbGluayBodHRwczovL2t1YmVybmV0ZXMuaW8vZG9jcy9jb25jZXB0cy9vdmVydmlldy93b3JraW5nLXdpdGgtb2JqZWN0cy9sYWJlbHMvI3N5bnRheC1hbmQtY2hhcmFjdGVyLXNldFxuICAgKlxuICAgKiBAcGFyYW0gc2NvcGUgVGhlIGNvbnN0cnVjdCBmb3Igd2hpY2ggdG8gcmVuZGVyIHRoZSBETlMgbGFiZWxcbiAgICogQHBhcmFtIG9wdGlvbnMgTmFtZSBvcHRpb25zXG4gICAqIEB0aHJvd3MgaWYgYW55IG9mIHRoZSBjb21wb25lbnRzIGRvIG5vdCBhZGhlcmUgdG8gbmFtaW5nIGNvbnN0cmFpbnRzIG9yXG4gICAqIGxlbmd0aC5cbiAgICovXG4gIHB1YmxpYyBzdGF0aWMgdG9MYWJlbFZhbHVlKHNjb3BlOiBDb25zdHJ1Y3QsIG9wdGlvbnM6IE5hbWVPcHRpb25zID0ge30pIHtcbiAgICBjb25zdCBtYXhMZW4gPSBvcHRpb25zLm1heExlbiA/PyBNQVhfTEVOO1xuICAgIGNvbnN0IGRlbGltID0gb3B0aW9ucy5kZWxpbWl0ZXIgPz8gJy0nO1xuICAgIGNvbnN0IGluY2x1ZGVfaGFzaCA9IG9wdGlvbnMuaW5jbHVkZUhhc2ggPz8gdHJ1ZTtcblxuICAgIGlmIChtYXhMZW4gPCBIQVNIX0xFTiAmJiBpbmNsdWRlX2hhc2gpIHtcbiAgICAgIHRocm93IG5ldyBFcnJvcihgbWluaW11bSBtYXggbGVuZ3RoIGZvciBsYWJlbCBpcyAke0hBU0hfTEVOfSAocmVxdWlyZWQgZm9yIGhhc2gpYCk7XG4gICAgfVxuXG4gICAgaWYgKC9bXjAtOWEtekEtWi1fLl0vLnRlc3QoZGVsaW0pKSB7XG4gICAgICB0aHJvdyBuZXcgRXJyb3IoJ2RlbGltIHNob3VsZCBub3QgY29udGFpbiBcIlteMC05YS16QS1aLV8uXVwiJyk7XG4gICAgfVxuXG4gICAgY29uc3Qgbm9kZSA9IHNjb3BlLm5vZGU7XG4gICAgbGV0IGNvbXBvbmVudHMgPSBub2RlLnBhdGguc3BsaXQoJy8nKTtcbiAgICBjb21wb25lbnRzLnB1c2goLi4ub3B0aW9ucy5leHRyYSA/PyBbXSk7XG5cbiAgICAvLyBzcGVjaWFsIGNhc2U6IGlmIHdlIG9ubHkgaGF2ZSBvbmUgY29tcG9uZW50IGluIG91ciBwYXRoIGFuZCBpdCBhZGhlcmVzIHRvIEROU19OQU1FLCB3ZSBkb24ndCBkZWNvcmF0ZSBpdFxuICAgIGlmIChjb21wb25lbnRzLmxlbmd0aCA9PT0gMSAmJiBWQUxJREFURV9MQUJFTF9WQUxVRS50ZXN0KGNvbXBvbmVudHNbMF0pICYmIGNvbXBvbmVudHNbMF0ubGVuZ3RoIDw9IG1heExlbikge1xuICAgICAgcmV0dXJuIGNvbXBvbmVudHNbMF07XG4gICAgfVxuXG4gICAgLy8gb2theSwgbm93IHdlIG5lZWQgdG8gbm9ybWFsaXplIGFsbCBjb21wb25lbnRzIHRvIGFkaGVyZSB0byBsYWJlbCBhbmQgYXBwZW5kIHRoZSBoYXNoIG9mIHRoZSBmdWxsIHBhdGguXG4gICAgY29tcG9uZW50cyA9IGNvbXBvbmVudHMubWFwKGMgPT4gbm9ybWFsaXplVG9MYWJlbFZhbHVlKGMsIG1heExlbikpO1xuICAgIGlmIChpbmNsdWRlX2hhc2gpIHtcbiAgICAgIGNvbXBvbmVudHMucHVzaChjYWxjSGFzaChub2RlLCBIQVNIX0xFTikpO1xuICAgIH1cblxuICAgIGNvbnN0IHJlc3VsdCA9IHRvSHVtYW5Gb3JtKGNvbXBvbmVudHMsIGRlbGltLCBtYXhMZW4pO1xuXG4gICAgLy8gc2xpY2luZyBtaWdodCBsZXQgJy0nLCAnXycsICcuJyBiZSBpbiB0aGUgc3RhcnQgb2YgdGhlIHJlc3VsdC5cbiAgICByZXR1cm4gcmVzdWx0LnJlcGxhY2UoL15bXjAtOWEtekEtWl0rLywgJycpO1xuICB9XG5cbiAgLyogaXN0YW5idWwgaWdub3JlIG5leHQgKi9cbiAgcHJpdmF0ZSBjb25zdHJ1Y3RvcigpIHtcbiAgICByZXR1cm47XG4gIH1cbn1cblxuZnVuY3Rpb24gb21pdER1cGxpY2F0ZXModmFsdWU6IHN0cmluZywgaW5kZXg6IG51bWJlciwgY29tcG9uZW50czogc3RyaW5nW10pIHtcbiAgcmV0dXJuIHZhbHVlICE9PSBjb21wb25lbnRzW2luZGV4LTFdO1xufVxuXG5mdW5jdGlvbiBvbWl0RGVmYXVsdENoaWxkKHZhbHVlOiBzdHJpbmcsIF86IG51bWJlciwgX186IHN0cmluZ1tdKSB7XG4gIHJldHVybiB2YWx1ZS50b0xvd2VyQ2FzZSgpICE9PSAncmVzb3VyY2UnICYmIHZhbHVlLnRvTG93ZXJDYXNlKCkgIT09ICdkZWZhdWx0Jztcbn1cblxuZnVuY3Rpb24gdG9IdW1hbkZvcm0oY29tcG9uZW50czogc3RyaW5nW10sIGRlbGltOiBzdHJpbmcsIG1heExlbjogbnVtYmVyKSB7XG4gIHJldHVybiBjb21wb25lbnRzLnJldmVyc2UoKVxuICAgIC5maWx0ZXIob21pdER1cGxpY2F0ZXMpXG4gICAgLmpvaW4oJy8nKVxuICAgIC5zbGljZSgwLCBtYXhMZW4pXG4gICAgLnNwbGl0KCcvJylcbiAgICAucmV2ZXJzZSgpXG4gICAgLmZpbHRlcih4ID0+IHgpXG4gICAgLmpvaW4oZGVsaW0pXG4gICAgLnNwbGl0KGRlbGltKVxuICAgIC5maWx0ZXIoeCA9PiB4KVxuICAgIC5maWx0ZXIob21pdERlZmF1bHRDaGlsZClcbiAgICAuam9pbihkZWxpbSk7XG5cbn1cblxuZnVuY3Rpb24gbm9ybWFsaXplVG9EbnNOYW1lKGM6IHN0cmluZywgbWF4TGVuOiBudW1iZXIpIHtcbiAgcmV0dXJuIGNcbiAgICAudG9Mb2NhbGVMb3dlckNhc2UoKSAvLyBsb3dlciBjYXNlXG4gICAgLnJlcGxhY2UoL1teMC05YS16QS1aLV8uXS9nLCAnJykgLy8gcmVtb3ZlIG5vbi1hbGxvd2VkIGNoYXJhY3RlcnNcbiAgICAuc3Vic3RyKDAsIG1heExlbik7IC8vIHRyaW0gdG8gbWF4TGVuZ3RoXG59XG5cbmZ1bmN0aW9uIGNhbGNIYXNoKG5vZGU6IE5vZGUsIG1heExlbjogbnVtYmVyKSB7XG4gIGlmIChwcm9jZXNzLmVudi5DREs4U19MRUdBQ1lfSEFTSCkge1xuICAgIGNvbnN0IGhhc2ggPSBjcnlwdG8uY3JlYXRlSGFzaCgnc2hhMjU2Jyk7XG4gICAgaGFzaC51cGRhdGUobm9kZS5wYXRoKTtcbiAgICByZXR1cm4gaGFzaC5kaWdlc3QoJ2hleCcpLnNsaWNlKDAsIG1heExlbik7XG4gIH1cblxuICByZXR1cm4gbm9kZS5hZGRyLnN1YnN0cmluZygwLCBIQVNIX0xFTik7XG59XG5cbmZ1bmN0aW9uIG5vcm1hbGl6ZVRvTGFiZWxWYWx1ZShjOiBzdHJpbmcsIG1heExlbjogbnVtYmVyKSB7XG4gIHJldHVybiBjXG4gICAgLnJlcGxhY2UoL1teMC05YS16QS1aLV8uXS9nLCAnJykgLy8gcmVtb3ZlIG5vbi1hbGxvd2VkIGNoYXJhY3RlcnNcbiAgICAuc3Vic3RyKDAsIG1heExlbik7IC8vIHRyaW0gdG8gbWF4TGVuZ3RoXG59XG4iXX0=