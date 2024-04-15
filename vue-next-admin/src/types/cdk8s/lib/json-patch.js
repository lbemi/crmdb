"use strict";
var _a;
Object.defineProperty(exports, "__esModule", { value: true });
exports.JsonPatch = void 0;
const JSII_RTTI_SYMBOL_1 = Symbol.for("jsii.rtti");
const fast_json_patch_1 = require("fast-json-patch");
/**
 * Utility for applying RFC-6902 JSON-Patch to a document.
 *
 * Use the the `JsonPatch.apply(doc, ...ops)` function to apply a set of
 * operations to a JSON document and return the result.
 *
 * Operations can be created using the factory methods `JsonPatch.add()`,
 * `JsonPatch.remove()`, etc.
 *
 * @example
 *
 *const output = JsonPatch.apply(input,
 *  JsonPatch.replace('/world/hi/there', 'goodbye'),
 *  JsonPatch.add('/world/foo/', 'boom'),
 *  JsonPatch.remove('/hello'));
 *
 */
class JsonPatch {
    /**
     * Applies a set of JSON-Patch (RFC-6902) operations to `document` and returns the result.
     * @param document The document to patch
     * @param ops The operations to apply
     * @returns The result document
     */
    static apply(document, ...ops) {
        const result = (0, fast_json_patch_1.applyPatch)(document, (0, fast_json_patch_1.deepClone)(ops.map(o => o._toJson())));
        return result.newDocument;
    }
    /**
     * Adds a value to an object or inserts it into an array. In the case of an
     * array, the value is inserted before the given index. The - character can be
     * used instead of an index to insert at the end of an array.
     *
     * @example JsonPatch.add('/biscuits/1', { "name": "Ginger Nut" })
     */
    static add(path, value) { return new JsonPatch({ op: 'add', path, value }); }
    /**
     * Removes a value from an object or array.
     *
     * @example JsonPatch.remove('/biscuits')
     * @example JsonPatch.remove('/biscuits/0')
     */
    static remove(path) { return new JsonPatch({ op: 'remove', path }); }
    /**
     * Replaces a value. Equivalent to a “remove” followed by an “add”.
     *
     * @example JsonPatch.replace('/biscuits/0/name', 'Chocolate Digestive')
     */
    static replace(path, value) { return new JsonPatch({ op: 'replace', path, value }); }
    /**
     * Copies a value from one location to another within the JSON document. Both
     * from and path are JSON Pointers.
     *
     * @example JsonPatch.copy('/biscuits/0', '/best_biscuit')
     */
    static copy(from, path) { return new JsonPatch({ op: 'copy', from, path }); }
    /**
     * Moves a value from one location to the other. Both from and path are JSON Pointers.
     *
     * @example JsonPatch.move('/biscuits', '/cookies')
     */
    static move(from, path) { return new JsonPatch({ op: 'move', from, path }); }
    /**
     * Tests that the specified value is set in the document. If the test fails,
     * then the patch as a whole should not apply.
     *
     * @example JsonPatch.test('/best_biscuit/name', 'Choco Leibniz')
     */
    static test(path, value) { return new JsonPatch({ op: 'test', path, value }); }
    constructor(operation) {
        this.operation = operation;
    }
    /**
     * Returns the JSON representation of this JSON patch operation.
     *
     * @internal
     */
    _toJson() {
        return this.operation;
    }
}
exports.JsonPatch = JsonPatch;
_a = JSII_RTTI_SYMBOL_1;
JsonPatch[_a] = { fqn: "cdk8s.JsonPatch", version: "2.68.60" };
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoianNvbi1wYXRjaC5qcyIsInNvdXJjZVJvb3QiOiIiLCJzb3VyY2VzIjpbIi4uL3NyYy9qc29uLXBhdGNoLnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7Ozs7O0FBQUEscURBQW1FO0FBRW5FOzs7Ozs7Ozs7Ozs7Ozs7O0dBZ0JHO0FBQ0gsTUFBYSxTQUFTO0lBQ3BCOzs7OztPQUtHO0lBQ0ksTUFBTSxDQUFDLEtBQUssQ0FBQyxRQUFhLEVBQUUsR0FBRyxHQUFnQjtRQUNwRCxNQUFNLE1BQU0sR0FBRyxJQUFBLDRCQUFVLEVBQUMsUUFBUSxFQUFFLElBQUEsMkJBQVMsRUFBQyxHQUFHLENBQUMsR0FBRyxDQUFDLENBQUMsQ0FBQyxFQUFFLENBQUMsQ0FBQyxDQUFDLE9BQU8sRUFBRSxDQUFDLENBQUMsQ0FBQyxDQUFDO1FBQzFFLE9BQU8sTUFBTSxDQUFDLFdBQVcsQ0FBQztJQUM1QixDQUFDO0lBRUQ7Ozs7OztPQU1HO0lBQ0ksTUFBTSxDQUFDLEdBQUcsQ0FBQyxJQUFZLEVBQUUsS0FBVSxJQUFJLE9BQU8sSUFBSSxTQUFTLENBQUMsRUFBRSxFQUFFLEVBQUUsS0FBSyxFQUFFLElBQUksRUFBRSxLQUFLLEVBQUUsQ0FBQyxDQUFDLENBQUMsQ0FBQztJQUVqRzs7Ozs7T0FLRztJQUNJLE1BQU0sQ0FBQyxNQUFNLENBQUMsSUFBWSxJQUFJLE9BQU8sSUFBSSxTQUFTLENBQUMsRUFBRSxFQUFFLEVBQUUsUUFBUSxFQUFFLElBQUksRUFBRSxDQUFDLENBQUMsQ0FBQyxDQUFDO0lBRXBGOzs7O09BSUc7SUFDSSxNQUFNLENBQUMsT0FBTyxDQUFDLElBQVksRUFBRSxLQUFVLElBQUksT0FBTyxJQUFJLFNBQVMsQ0FBQyxFQUFFLEVBQUUsRUFBRSxTQUFTLEVBQUUsSUFBSSxFQUFFLEtBQUssRUFBRSxDQUFDLENBQUMsQ0FBQyxDQUFDO0lBRXpHOzs7OztPQUtHO0lBQ0ksTUFBTSxDQUFDLElBQUksQ0FBQyxJQUFZLEVBQUUsSUFBWSxJQUFJLE9BQU8sSUFBSSxTQUFTLENBQUMsRUFBRSxFQUFFLEVBQUUsTUFBTSxFQUFFLElBQUksRUFBRSxJQUFJLEVBQUUsQ0FBQyxDQUFDLENBQUMsQ0FBQztJQUVwRzs7OztPQUlHO0lBQ0ksTUFBTSxDQUFDLElBQUksQ0FBQyxJQUFZLEVBQUUsSUFBWSxJQUFJLE9BQU8sSUFBSSxTQUFTLENBQUMsRUFBRSxFQUFFLEVBQUUsTUFBTSxFQUFFLElBQUksRUFBRSxJQUFJLEVBQUUsQ0FBQyxDQUFDLENBQUMsQ0FBQztJQUVwRzs7Ozs7T0FLRztJQUNJLE1BQU0sQ0FBQyxJQUFJLENBQUMsSUFBWSxFQUFFLEtBQVUsSUFBSSxPQUFPLElBQUksU0FBUyxDQUFDLEVBQUUsRUFBRSxFQUFFLE1BQU0sRUFBRSxJQUFJLEVBQUUsS0FBSyxFQUFFLENBQUMsQ0FBQyxDQUFDLENBQUM7SUFFbkcsWUFBcUMsU0FBb0I7UUFBcEIsY0FBUyxHQUFULFNBQVMsQ0FBVztJQUFHLENBQUM7SUFFN0Q7Ozs7T0FJRztJQUNJLE9BQU87UUFDWixPQUFPLElBQUksQ0FBQyxTQUFTLENBQUM7SUFDeEIsQ0FBQzs7QUFwRUgsOEJBcUVDIiwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IHsgYXBwbHlQYXRjaCwgZGVlcENsb25lLCBPcGVyYXRpb24gfSBmcm9tICdmYXN0LWpzb24tcGF0Y2gnO1xuXG4vKipcbiAqIFV0aWxpdHkgZm9yIGFwcGx5aW5nIFJGQy02OTAyIEpTT04tUGF0Y2ggdG8gYSBkb2N1bWVudC5cbiAqXG4gKiBVc2UgdGhlIHRoZSBgSnNvblBhdGNoLmFwcGx5KGRvYywgLi4ub3BzKWAgZnVuY3Rpb24gdG8gYXBwbHkgYSBzZXQgb2ZcbiAqIG9wZXJhdGlvbnMgdG8gYSBKU09OIGRvY3VtZW50IGFuZCByZXR1cm4gdGhlIHJlc3VsdC5cbiAqXG4gKiBPcGVyYXRpb25zIGNhbiBiZSBjcmVhdGVkIHVzaW5nIHRoZSBmYWN0b3J5IG1ldGhvZHMgYEpzb25QYXRjaC5hZGQoKWAsXG4gKiBgSnNvblBhdGNoLnJlbW92ZSgpYCwgZXRjLlxuICpcbiAqIEBleGFtcGxlXG4gKlxuICpjb25zdCBvdXRwdXQgPSBKc29uUGF0Y2guYXBwbHkoaW5wdXQsXG4gKiAgSnNvblBhdGNoLnJlcGxhY2UoJy93b3JsZC9oaS90aGVyZScsICdnb29kYnllJyksXG4gKiAgSnNvblBhdGNoLmFkZCgnL3dvcmxkL2Zvby8nLCAnYm9vbScpLFxuICogIEpzb25QYXRjaC5yZW1vdmUoJy9oZWxsbycpKTtcbiAqXG4gKi9cbmV4cG9ydCBjbGFzcyBKc29uUGF0Y2gge1xuICAvKipcbiAgICogQXBwbGllcyBhIHNldCBvZiBKU09OLVBhdGNoIChSRkMtNjkwMikgb3BlcmF0aW9ucyB0byBgZG9jdW1lbnRgIGFuZCByZXR1cm5zIHRoZSByZXN1bHQuXG4gICAqIEBwYXJhbSBkb2N1bWVudCBUaGUgZG9jdW1lbnQgdG8gcGF0Y2hcbiAgICogQHBhcmFtIG9wcyBUaGUgb3BlcmF0aW9ucyB0byBhcHBseVxuICAgKiBAcmV0dXJucyBUaGUgcmVzdWx0IGRvY3VtZW50XG4gICAqL1xuICBwdWJsaWMgc3RhdGljIGFwcGx5KGRvY3VtZW50OiBhbnksIC4uLm9wczogSnNvblBhdGNoW10pOiBhbnkge1xuICAgIGNvbnN0IHJlc3VsdCA9IGFwcGx5UGF0Y2goZG9jdW1lbnQsIGRlZXBDbG9uZShvcHMubWFwKG8gPT4gby5fdG9Kc29uKCkpKSk7XG4gICAgcmV0dXJuIHJlc3VsdC5uZXdEb2N1bWVudDtcbiAgfVxuXG4gIC8qKlxuICAgKiBBZGRzIGEgdmFsdWUgdG8gYW4gb2JqZWN0IG9yIGluc2VydHMgaXQgaW50byBhbiBhcnJheS4gSW4gdGhlIGNhc2Ugb2YgYW5cbiAgICogYXJyYXksIHRoZSB2YWx1ZSBpcyBpbnNlcnRlZCBiZWZvcmUgdGhlIGdpdmVuIGluZGV4LiBUaGUgLSBjaGFyYWN0ZXIgY2FuIGJlXG4gICAqIHVzZWQgaW5zdGVhZCBvZiBhbiBpbmRleCB0byBpbnNlcnQgYXQgdGhlIGVuZCBvZiBhbiBhcnJheS5cbiAgICpcbiAgICogQGV4YW1wbGUgSnNvblBhdGNoLmFkZCgnL2Jpc2N1aXRzLzEnLCB7IFwibmFtZVwiOiBcIkdpbmdlciBOdXRcIiB9KVxuICAgKi9cbiAgcHVibGljIHN0YXRpYyBhZGQocGF0aDogc3RyaW5nLCB2YWx1ZTogYW55KSB7IHJldHVybiBuZXcgSnNvblBhdGNoKHsgb3A6ICdhZGQnLCBwYXRoLCB2YWx1ZSB9KTsgfVxuXG4gIC8qKlxuICAgKiBSZW1vdmVzIGEgdmFsdWUgZnJvbSBhbiBvYmplY3Qgb3IgYXJyYXkuXG4gICAqXG4gICAqIEBleGFtcGxlIEpzb25QYXRjaC5yZW1vdmUoJy9iaXNjdWl0cycpXG4gICAqIEBleGFtcGxlIEpzb25QYXRjaC5yZW1vdmUoJy9iaXNjdWl0cy8wJylcbiAgICovXG4gIHB1YmxpYyBzdGF0aWMgcmVtb3ZlKHBhdGg6IHN0cmluZykgeyByZXR1cm4gbmV3IEpzb25QYXRjaCh7IG9wOiAncmVtb3ZlJywgcGF0aCB9KTsgfVxuXG4gIC8qKlxuICAgKiBSZXBsYWNlcyBhIHZhbHVlLiBFcXVpdmFsZW50IHRvIGEg4oCccmVtb3Zl4oCdIGZvbGxvd2VkIGJ5IGFuIOKAnGFkZOKAnS5cbiAgICpcbiAgICogQGV4YW1wbGUgSnNvblBhdGNoLnJlcGxhY2UoJy9iaXNjdWl0cy8wL25hbWUnLCAnQ2hvY29sYXRlIERpZ2VzdGl2ZScpXG4gICAqL1xuICBwdWJsaWMgc3RhdGljIHJlcGxhY2UocGF0aDogc3RyaW5nLCB2YWx1ZTogYW55KSB7IHJldHVybiBuZXcgSnNvblBhdGNoKHsgb3A6ICdyZXBsYWNlJywgcGF0aCwgdmFsdWUgfSk7IH1cblxuICAvKipcbiAgICogQ29waWVzIGEgdmFsdWUgZnJvbSBvbmUgbG9jYXRpb24gdG8gYW5vdGhlciB3aXRoaW4gdGhlIEpTT04gZG9jdW1lbnQuIEJvdGhcbiAgICogZnJvbSBhbmQgcGF0aCBhcmUgSlNPTiBQb2ludGVycy5cbiAgICpcbiAgICogQGV4YW1wbGUgSnNvblBhdGNoLmNvcHkoJy9iaXNjdWl0cy8wJywgJy9iZXN0X2Jpc2N1aXQnKVxuICAgKi9cbiAgcHVibGljIHN0YXRpYyBjb3B5KGZyb206IHN0cmluZywgcGF0aDogc3RyaW5nKSB7IHJldHVybiBuZXcgSnNvblBhdGNoKHsgb3A6ICdjb3B5JywgZnJvbSwgcGF0aCB9KTsgfVxuXG4gIC8qKlxuICAgKiBNb3ZlcyBhIHZhbHVlIGZyb20gb25lIGxvY2F0aW9uIHRvIHRoZSBvdGhlci4gQm90aCBmcm9tIGFuZCBwYXRoIGFyZSBKU09OIFBvaW50ZXJzLlxuICAgKlxuICAgKiBAZXhhbXBsZSBKc29uUGF0Y2gubW92ZSgnL2Jpc2N1aXRzJywgJy9jb29raWVzJylcbiAgICovXG4gIHB1YmxpYyBzdGF0aWMgbW92ZShmcm9tOiBzdHJpbmcsIHBhdGg6IHN0cmluZykgeyByZXR1cm4gbmV3IEpzb25QYXRjaCh7IG9wOiAnbW92ZScsIGZyb20sIHBhdGggfSk7IH1cblxuICAvKipcbiAgICogVGVzdHMgdGhhdCB0aGUgc3BlY2lmaWVkIHZhbHVlIGlzIHNldCBpbiB0aGUgZG9jdW1lbnQuIElmIHRoZSB0ZXN0IGZhaWxzLFxuICAgKiB0aGVuIHRoZSBwYXRjaCBhcyBhIHdob2xlIHNob3VsZCBub3QgYXBwbHkuXG4gICAqXG4gICAqIEBleGFtcGxlIEpzb25QYXRjaC50ZXN0KCcvYmVzdF9iaXNjdWl0L25hbWUnLCAnQ2hvY28gTGVpYm5peicpXG4gICAqL1xuICBwdWJsaWMgc3RhdGljIHRlc3QocGF0aDogc3RyaW5nLCB2YWx1ZTogYW55KSB7IHJldHVybiBuZXcgSnNvblBhdGNoKHsgb3A6ICd0ZXN0JywgcGF0aCwgdmFsdWUgfSk7IH1cblxuICBwcml2YXRlIGNvbnN0cnVjdG9yKHByaXZhdGUgcmVhZG9ubHkgb3BlcmF0aW9uOiBPcGVyYXRpb24pIHt9XG5cbiAgLyoqXG4gICAqIFJldHVybnMgdGhlIEpTT04gcmVwcmVzZW50YXRpb24gb2YgdGhpcyBKU09OIHBhdGNoIG9wZXJhdGlvbi5cbiAgICpcbiAgICogQGludGVybmFsXG4gICAqL1xuICBwdWJsaWMgX3RvSnNvbigpOiBhbnkge1xuICAgIHJldHVybiB0aGlzLm9wZXJhdGlvbjtcbiAgfVxufSJdfQ==