"use strict";
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
exports.Yaml = void 0;
const JSII_RTTI_SYMBOL_1 = Symbol.for("jsii.rtti");
const child_process_1 = require("child_process");
const fs = require("fs");
const os = require("os");
const path = require("path");
const YAML = require("yaml");
const MAX_DOWNLOAD_BUFFER = 10 * 1024 * 1024;
// Set default YAML schema to 1.1. This ensures saved YAML is backward compatible with other parsers, such as PyYAML
// It also ensures that octal numbers in the form `0775` will be parsed
// correctly on YAML load. (see https://github.com/eemeli/yaml/issues/205)
const yamlSchemaVersion = '1.1';
/**
 * YAML utilities.
 */
class Yaml {
    /**
     * @deprecated use `stringify(doc[, doc, ...])`
     */
    static formatObjects(docs) {
        return this.stringify(...docs);
    }
    /**
     * Saves a set of objects as a multi-document YAML file.
     * @param filePath The output path
     * @param docs The set of objects
     */
    static save(filePath, docs) {
        const data = this.stringify(...docs);
        fs.writeFileSync(filePath, data, { encoding: 'utf8' });
    }
    /**
     * Stringify a document (or multiple documents) into YAML
     *
     * We convert undefined values to null, but ignore any documents that are
     * undefined.
     *
     * @param docs A set of objects to convert to YAML
     * @returns a YAML string. Multiple docs are separated by `---`.
     */
    static stringify(...docs) {
        return docs.map(r => r === undefined ? '\n' : YAML.stringify(r, { keepUndefined: true, lineWidth: 0, version: yamlSchemaVersion })).join('---\n');
    }
    /**
     * Saves a set of YAML documents into a temp file (in /tmp)
     *
     * @returns the path to the temporary file
     * @param docs the set of documents to save
     */
    static tmp(docs) {
        const tmpdir = fs.mkdtempSync(path.join(os.tmpdir(), 'cdk8s-'));
        const filePath = path.join(tmpdir, 'temp.yaml');
        Yaml.save(filePath, docs);
        return filePath;
    }
    /**
     * Downloads a set of YAML documents (k8s manifest for example) from a URL or
     * a file and returns them as javascript objects.
     *
     * Empty documents are filtered out.
     *
     * @param urlOrFile a URL of a file path to load from
     * @returns an array of objects, each represents a document inside the YAML
     */
    static load(urlOrFile) {
        const body = loadurl(urlOrFile);
        const objects = YAML.parseAllDocuments(body, {
            version: yamlSchemaVersion,
        });
        const result = new Array();
        for (const obj of objects.map(x => x.toJSON())) {
            // skip empty documents
            if (obj === undefined) {
                continue;
            }
            if (obj === null) {
                continue;
            }
            if (Array.isArray(obj) && obj.length === 0) {
                continue;
            }
            if (typeof (obj) === 'object' && Object.keys(obj).length === 0) {
                continue;
            }
            result.push(obj);
        }
        return result;
    }
    /**
     * Utility class.
     */
    constructor() {
        return;
    }
}
exports.Yaml = Yaml;
_a = JSII_RTTI_SYMBOL_1;
Yaml[_a] = { fqn: "cdk8s.Yaml", version: "2.68.60" };
/**
 * Loads a url (or file) and returns the contents.
 * This method spawns a child process in order to perform an http call synchronously.
 */
function loadurl(url) {
    const script = path.join(__dirname, '_loadurl.js');
    return (0, child_process_1.execFileSync)(process.execPath, [script, url], {
        encoding: 'utf-8',
        maxBuffer: MAX_DOWNLOAD_BUFFER,
    }).toString();
}
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoieWFtbC5qcyIsInNvdXJjZVJvb3QiOiIiLCJzb3VyY2VzIjpbIi4uL3NyYy95YW1sLnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7Ozs7O0FBQUEsaURBQTZDO0FBQzdDLHlCQUF5QjtBQUN6Qix5QkFBeUI7QUFDekIsNkJBQTZCO0FBQzdCLDZCQUE2QjtBQUU3QixNQUFNLG1CQUFtQixHQUFHLEVBQUUsR0FBRyxJQUFJLEdBQUcsSUFBSSxDQUFDO0FBRTdDLG9IQUFvSDtBQUNwSCx1RUFBdUU7QUFDdkUsMEVBQTBFO0FBQzFFLE1BQU0saUJBQWlCLEdBQUcsS0FBSyxDQUFDO0FBRWhDOztHQUVHO0FBQ0gsTUFBYSxJQUFJO0lBQ2Y7O09BRUc7SUFDSSxNQUFNLENBQUMsYUFBYSxDQUFDLElBQVc7UUFDckMsT0FBTyxJQUFJLENBQUMsU0FBUyxDQUFDLEdBQUcsSUFBSSxDQUFDLENBQUM7SUFDakMsQ0FBQztJQUVEOzs7O09BSUc7SUFDSSxNQUFNLENBQUMsSUFBSSxDQUFDLFFBQWdCLEVBQUUsSUFBVztRQUM5QyxNQUFNLElBQUksR0FBRyxJQUFJLENBQUMsU0FBUyxDQUFDLEdBQUcsSUFBSSxDQUFDLENBQUM7UUFDckMsRUFBRSxDQUFDLGFBQWEsQ0FBQyxRQUFRLEVBQUUsSUFBSSxFQUFFLEVBQUUsUUFBUSxFQUFFLE1BQU0sRUFBRSxDQUFDLENBQUM7SUFDekQsQ0FBQztJQUVEOzs7Ozs7OztPQVFHO0lBQ0ksTUFBTSxDQUFDLFNBQVMsQ0FBQyxHQUFHLElBQVc7UUFDcEMsT0FBTyxJQUFJLENBQUMsR0FBRyxDQUNiLENBQUMsQ0FBQyxFQUFFLENBQUMsQ0FBQyxLQUFLLFNBQVMsQ0FBQyxDQUFDLENBQUMsSUFBSSxDQUFDLENBQUMsQ0FBQyxJQUFJLENBQUMsU0FBUyxDQUFDLENBQUMsRUFBRSxFQUFFLGFBQWEsRUFBRSxJQUFJLEVBQUUsU0FBUyxFQUFFLENBQUMsRUFBRSxPQUFPLEVBQUUsaUJBQWlCLEVBQUUsQ0FBQyxDQUNuSCxDQUFDLElBQUksQ0FBQyxPQUFPLENBQUMsQ0FBQztJQUNsQixDQUFDO0lBRUQ7Ozs7O09BS0c7SUFDSSxNQUFNLENBQUMsR0FBRyxDQUFDLElBQVc7UUFDM0IsTUFBTSxNQUFNLEdBQUcsRUFBRSxDQUFDLFdBQVcsQ0FBQyxJQUFJLENBQUMsSUFBSSxDQUFDLEVBQUUsQ0FBQyxNQUFNLEVBQUUsRUFBRSxRQUFRLENBQUMsQ0FBQyxDQUFDO1FBQ2hFLE1BQU0sUUFBUSxHQUFHLElBQUksQ0FBQyxJQUFJLENBQUMsTUFBTSxFQUFFLFdBQVcsQ0FBQyxDQUFDO1FBQ2hELElBQUksQ0FBQyxJQUFJLENBQUMsUUFBUSxFQUFFLElBQUksQ0FBQyxDQUFDO1FBQzFCLE9BQU8sUUFBUSxDQUFDO0lBQ2xCLENBQUM7SUFFRDs7Ozs7Ozs7T0FRRztJQUNJLE1BQU0sQ0FBQyxJQUFJLENBQUMsU0FBaUI7UUFDbEMsTUFBTSxJQUFJLEdBQUcsT0FBTyxDQUFDLFNBQVMsQ0FBQyxDQUFDO1FBRWhDLE1BQU0sT0FBTyxHQUFHLElBQUksQ0FBQyxpQkFBaUIsQ0FBQyxJQUFJLEVBQUU7WUFDM0MsT0FBTyxFQUFFLGlCQUFpQjtTQUMzQixDQUFDLENBQUM7UUFDSCxNQUFNLE1BQU0sR0FBRyxJQUFJLEtBQUssRUFBTyxDQUFDO1FBRWhDLEtBQUssTUFBTSxHQUFHLElBQUksT0FBTyxDQUFDLEdBQUcsQ0FBQyxDQUFDLENBQUMsRUFBRSxDQUFDLENBQUMsQ0FBQyxNQUFNLEVBQUUsQ0FBQyxFQUFFLENBQUM7WUFDL0MsdUJBQXVCO1lBQ3ZCLElBQUksR0FBRyxLQUFLLFNBQVMsRUFBRSxDQUFDO2dCQUFDLFNBQVM7WUFBQyxDQUFDO1lBQ3BDLElBQUksR0FBRyxLQUFLLElBQUksRUFBRSxDQUFDO2dCQUFDLFNBQVM7WUFBQyxDQUFDO1lBQy9CLElBQUksS0FBSyxDQUFDLE9BQU8sQ0FBQyxHQUFHLENBQUMsSUFBSSxHQUFHLENBQUMsTUFBTSxLQUFLLENBQUMsRUFBRSxDQUFDO2dCQUFDLFNBQVM7WUFBQyxDQUFDO1lBQ3pELElBQUksT0FBTSxDQUFDLEdBQUcsQ0FBQyxLQUFLLFFBQVEsSUFBSSxNQUFNLENBQUMsSUFBSSxDQUFDLEdBQUcsQ0FBQyxDQUFDLE1BQU0sS0FBSyxDQUFDLEVBQUUsQ0FBQztnQkFBQyxTQUFTO1lBQUMsQ0FBQztZQUU1RSxNQUFNLENBQUMsSUFBSSxDQUFDLEdBQUcsQ0FBQyxDQUFDO1FBQ25CLENBQUM7UUFFRCxPQUFPLE1BQU0sQ0FBQztJQUNoQixDQUFDO0lBRUQ7O09BRUc7SUFDSDtRQUNFLE9BQU87SUFDVCxDQUFDOztBQWpGSCxvQkFrRkM7OztBQUVEOzs7R0FHRztBQUNILFNBQVMsT0FBTyxDQUFDLEdBQVc7SUFDMUIsTUFBTSxNQUFNLEdBQUcsSUFBSSxDQUFDLElBQUksQ0FBQyxTQUFTLEVBQUUsYUFBYSxDQUFDLENBQUM7SUFDbkQsT0FBTyxJQUFBLDRCQUFZLEVBQUMsT0FBTyxDQUFDLFFBQVEsRUFBRSxDQUFDLE1BQU0sRUFBRSxHQUFHLENBQUMsRUFBRTtRQUNuRCxRQUFRLEVBQUUsT0FBTztRQUNqQixTQUFTLEVBQUUsbUJBQW1CO0tBQy9CLENBQUMsQ0FBQyxRQUFRLEVBQUUsQ0FBQztBQUNoQixDQUFDIiwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IHsgZXhlY0ZpbGVTeW5jIH0gZnJvbSAnY2hpbGRfcHJvY2Vzcyc7XG5pbXBvcnQgKiBhcyBmcyBmcm9tICdmcyc7XG5pbXBvcnQgKiBhcyBvcyBmcm9tICdvcyc7XG5pbXBvcnQgKiBhcyBwYXRoIGZyb20gJ3BhdGgnO1xuaW1wb3J0ICogYXMgWUFNTCBmcm9tICd5YW1sJztcblxuY29uc3QgTUFYX0RPV05MT0FEX0JVRkZFUiA9IDEwICogMTAyNCAqIDEwMjQ7XG5cbi8vIFNldCBkZWZhdWx0IFlBTUwgc2NoZW1hIHRvIDEuMS4gVGhpcyBlbnN1cmVzIHNhdmVkIFlBTUwgaXMgYmFja3dhcmQgY29tcGF0aWJsZSB3aXRoIG90aGVyIHBhcnNlcnMsIHN1Y2ggYXMgUHlZQU1MXG4vLyBJdCBhbHNvIGVuc3VyZXMgdGhhdCBvY3RhbCBudW1iZXJzIGluIHRoZSBmb3JtIGAwNzc1YCB3aWxsIGJlIHBhcnNlZFxuLy8gY29ycmVjdGx5IG9uIFlBTUwgbG9hZC4gKHNlZSBodHRwczovL2dpdGh1Yi5jb20vZWVtZWxpL3lhbWwvaXNzdWVzLzIwNSlcbmNvbnN0IHlhbWxTY2hlbWFWZXJzaW9uID0gJzEuMSc7XG5cbi8qKlxuICogWUFNTCB1dGlsaXRpZXMuXG4gKi9cbmV4cG9ydCBjbGFzcyBZYW1sIHtcbiAgLyoqXG4gICAqIEBkZXByZWNhdGVkIHVzZSBgc3RyaW5naWZ5KGRvY1ssIGRvYywgLi4uXSlgXG4gICAqL1xuICBwdWJsaWMgc3RhdGljIGZvcm1hdE9iamVjdHMoZG9jczogYW55W10pOiBzdHJpbmcge1xuICAgIHJldHVybiB0aGlzLnN0cmluZ2lmeSguLi5kb2NzKTtcbiAgfVxuXG4gIC8qKlxuICAgKiBTYXZlcyBhIHNldCBvZiBvYmplY3RzIGFzIGEgbXVsdGktZG9jdW1lbnQgWUFNTCBmaWxlLlxuICAgKiBAcGFyYW0gZmlsZVBhdGggVGhlIG91dHB1dCBwYXRoXG4gICAqIEBwYXJhbSBkb2NzIFRoZSBzZXQgb2Ygb2JqZWN0c1xuICAgKi9cbiAgcHVibGljIHN0YXRpYyBzYXZlKGZpbGVQYXRoOiBzdHJpbmcsIGRvY3M6IGFueVtdKSB7XG4gICAgY29uc3QgZGF0YSA9IHRoaXMuc3RyaW5naWZ5KC4uLmRvY3MpO1xuICAgIGZzLndyaXRlRmlsZVN5bmMoZmlsZVBhdGgsIGRhdGEsIHsgZW5jb2Rpbmc6ICd1dGY4JyB9KTtcbiAgfVxuXG4gIC8qKlxuICAgKiBTdHJpbmdpZnkgYSBkb2N1bWVudCAob3IgbXVsdGlwbGUgZG9jdW1lbnRzKSBpbnRvIFlBTUxcbiAgICpcbiAgICogV2UgY29udmVydCB1bmRlZmluZWQgdmFsdWVzIHRvIG51bGwsIGJ1dCBpZ25vcmUgYW55IGRvY3VtZW50cyB0aGF0IGFyZVxuICAgKiB1bmRlZmluZWQuXG4gICAqXG4gICAqIEBwYXJhbSBkb2NzIEEgc2V0IG9mIG9iamVjdHMgdG8gY29udmVydCB0byBZQU1MXG4gICAqIEByZXR1cm5zIGEgWUFNTCBzdHJpbmcuIE11bHRpcGxlIGRvY3MgYXJlIHNlcGFyYXRlZCBieSBgLS0tYC5cbiAgICovXG4gIHB1YmxpYyBzdGF0aWMgc3RyaW5naWZ5KC4uLmRvY3M6IGFueVtdKSB7XG4gICAgcmV0dXJuIGRvY3MubWFwKFxuICAgICAgciA9PiByID09PSB1bmRlZmluZWQgPyAnXFxuJyA6IFlBTUwuc3RyaW5naWZ5KHIsIHsga2VlcFVuZGVmaW5lZDogdHJ1ZSwgbGluZVdpZHRoOiAwLCB2ZXJzaW9uOiB5YW1sU2NoZW1hVmVyc2lvbiB9KSxcbiAgICApLmpvaW4oJy0tLVxcbicpO1xuICB9XG5cbiAgLyoqXG4gICAqIFNhdmVzIGEgc2V0IG9mIFlBTUwgZG9jdW1lbnRzIGludG8gYSB0ZW1wIGZpbGUgKGluIC90bXApXG4gICAqXG4gICAqIEByZXR1cm5zIHRoZSBwYXRoIHRvIHRoZSB0ZW1wb3JhcnkgZmlsZVxuICAgKiBAcGFyYW0gZG9jcyB0aGUgc2V0IG9mIGRvY3VtZW50cyB0byBzYXZlXG4gICAqL1xuICBwdWJsaWMgc3RhdGljIHRtcChkb2NzOiBhbnlbXSk6IHN0cmluZyB7XG4gICAgY29uc3QgdG1wZGlyID0gZnMubWtkdGVtcFN5bmMocGF0aC5qb2luKG9zLnRtcGRpcigpLCAnY2RrOHMtJykpO1xuICAgIGNvbnN0IGZpbGVQYXRoID0gcGF0aC5qb2luKHRtcGRpciwgJ3RlbXAueWFtbCcpO1xuICAgIFlhbWwuc2F2ZShmaWxlUGF0aCwgZG9jcyk7XG4gICAgcmV0dXJuIGZpbGVQYXRoO1xuICB9XG5cbiAgLyoqXG4gICAqIERvd25sb2FkcyBhIHNldCBvZiBZQU1MIGRvY3VtZW50cyAoazhzIG1hbmlmZXN0IGZvciBleGFtcGxlKSBmcm9tIGEgVVJMIG9yXG4gICAqIGEgZmlsZSBhbmQgcmV0dXJucyB0aGVtIGFzIGphdmFzY3JpcHQgb2JqZWN0cy5cbiAgICpcbiAgICogRW1wdHkgZG9jdW1lbnRzIGFyZSBmaWx0ZXJlZCBvdXQuXG4gICAqXG4gICAqIEBwYXJhbSB1cmxPckZpbGUgYSBVUkwgb2YgYSBmaWxlIHBhdGggdG8gbG9hZCBmcm9tXG4gICAqIEByZXR1cm5zIGFuIGFycmF5IG9mIG9iamVjdHMsIGVhY2ggcmVwcmVzZW50cyBhIGRvY3VtZW50IGluc2lkZSB0aGUgWUFNTFxuICAgKi9cbiAgcHVibGljIHN0YXRpYyBsb2FkKHVybE9yRmlsZTogc3RyaW5nKTogYW55W10ge1xuICAgIGNvbnN0IGJvZHkgPSBsb2FkdXJsKHVybE9yRmlsZSk7XG5cbiAgICBjb25zdCBvYmplY3RzID0gWUFNTC5wYXJzZUFsbERvY3VtZW50cyhib2R5LCB7XG4gICAgICB2ZXJzaW9uOiB5YW1sU2NoZW1hVmVyc2lvbixcbiAgICB9KTtcbiAgICBjb25zdCByZXN1bHQgPSBuZXcgQXJyYXk8YW55PigpO1xuXG4gICAgZm9yIChjb25zdCBvYmogb2Ygb2JqZWN0cy5tYXAoeCA9PiB4LnRvSlNPTigpKSkge1xuICAgICAgLy8gc2tpcCBlbXB0eSBkb2N1bWVudHNcbiAgICAgIGlmIChvYmogPT09IHVuZGVmaW5lZCkgeyBjb250aW51ZTsgfVxuICAgICAgaWYgKG9iaiA9PT0gbnVsbCkgeyBjb250aW51ZTsgfVxuICAgICAgaWYgKEFycmF5LmlzQXJyYXkob2JqKSAmJiBvYmoubGVuZ3RoID09PSAwKSB7IGNvbnRpbnVlOyB9XG4gICAgICBpZiAodHlwZW9mKG9iaikgPT09ICdvYmplY3QnICYmIE9iamVjdC5rZXlzKG9iaikubGVuZ3RoID09PSAwKSB7IGNvbnRpbnVlOyB9XG5cbiAgICAgIHJlc3VsdC5wdXNoKG9iaik7XG4gICAgfVxuXG4gICAgcmV0dXJuIHJlc3VsdDtcbiAgfVxuXG4gIC8qKlxuICAgKiBVdGlsaXR5IGNsYXNzLlxuICAgKi9cbiAgcHJpdmF0ZSBjb25zdHJ1Y3RvcigpIHtcbiAgICByZXR1cm47XG4gIH1cbn1cblxuLyoqXG4gKiBMb2FkcyBhIHVybCAob3IgZmlsZSkgYW5kIHJldHVybnMgdGhlIGNvbnRlbnRzLlxuICogVGhpcyBtZXRob2Qgc3Bhd25zIGEgY2hpbGQgcHJvY2VzcyBpbiBvcmRlciB0byBwZXJmb3JtIGFuIGh0dHAgY2FsbCBzeW5jaHJvbm91c2x5LlxuICovXG5mdW5jdGlvbiBsb2FkdXJsKHVybDogc3RyaW5nKTogc3RyaW5nIHtcbiAgY29uc3Qgc2NyaXB0ID0gcGF0aC5qb2luKF9fZGlybmFtZSwgJ19sb2FkdXJsLmpzJyk7XG4gIHJldHVybiBleGVjRmlsZVN5bmMocHJvY2Vzcy5leGVjUGF0aCwgW3NjcmlwdCwgdXJsXSwge1xuICAgIGVuY29kaW5nOiAndXRmLTgnLFxuICAgIG1heEJ1ZmZlcjogTUFYX0RPV05MT0FEX0JVRkZFUixcbiAgfSkudG9TdHJpbmcoKTtcbn1cbiJdfQ==