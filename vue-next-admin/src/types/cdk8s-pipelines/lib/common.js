"use strict";
// const paramNameRegex = new RegExp('^[a-zA-Z_][-a-zA-Z0-9_]+$');
Object.defineProperty(exports, "__esModule", { value: true });
exports.buildParam = exports.buildWorkingDir = exports.secretKeyRef = exports.TektonV1ApiVersion = void 0;
// function hasValidName(res: NamedResource): boolean {
//   let valid = false;
//   if (res.name) {
//     valid = paramNameRegex.test(res.name);
//   }
//   return valid;
// }
/**
 * The supported apiVerion for the Pipelines, Tasks, etc.
 */
exports.TektonV1ApiVersion = 'tekton.dev/v1';
function secretKeyRef(name, key) {
    return {
        name: name,
        key: key,
    };
}
exports.secretKeyRef = secretKeyRef;
/**
 * Convenience method for formatting the value of a working directory.
 * @param workspace
 */
function buildWorkingDir(workspace) {
    return `$(workspaces.${workspace}.path)`;
}
exports.buildWorkingDir = buildWorkingDir;
/**
 * Builds the correct string for referencing the parameter specified by `name`
 * that can be used when building tasks and others.
 *
 * For example, if the parameter is `foo`, the result will be `$(params.foo)`.
 * @param name The name of the parameter.
 */
function buildParam(name) {
    return `$(params.${name})`;
}
exports.buildParam = buildParam;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiY29tbW9uLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsiLi4vc3JjL2NvbW1vbi50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiO0FBQUEsa0VBQWtFOzs7QUFFbEUsdURBQXVEO0FBQ3ZELHVCQUF1QjtBQUN2QixvQkFBb0I7QUFDcEIsNkNBQTZDO0FBQzdDLE1BQU07QUFDTixrQkFBa0I7QUFDbEIsSUFBSTtBQUVKOztHQUVHO0FBQ1UsUUFBQSxrQkFBa0IsR0FBRyxlQUFlLENBQUM7QUFVbEQsU0FBZ0IsWUFBWSxDQUFDLElBQVksRUFBRSxHQUFXO0lBQ3BELE9BQU87UUFDTCxJQUFJLEVBQUUsSUFBSTtRQUNWLEdBQUcsRUFBRSxHQUFHO0tBQ1QsQ0FBQztBQUNKLENBQUM7QUFMRCxvQ0FLQztBQUVEOzs7R0FHRztBQUNILFNBQWdCLGVBQWUsQ0FBQyxTQUFpQjtJQUMvQyxPQUFPLGdCQUFnQixTQUFTLFFBQVEsQ0FBQztBQUMzQyxDQUFDO0FBRkQsMENBRUM7QUFFRDs7Ozs7O0dBTUc7QUFDSCxTQUFnQixVQUFVLENBQUMsSUFBWTtJQUNyQyxPQUFPLFlBQVksSUFBSSxHQUFHLENBQUM7QUFDN0IsQ0FBQztBQUZELGdDQUVDIiwic291cmNlc0NvbnRlbnQiOlsiLy8gY29uc3QgcGFyYW1OYW1lUmVnZXggPSBuZXcgUmVnRXhwKCdeW2EtekEtWl9dWy1hLXpBLVowLTlfXSskJyk7XG5cbi8vIGZ1bmN0aW9uIGhhc1ZhbGlkTmFtZShyZXM6IE5hbWVkUmVzb3VyY2UpOiBib29sZWFuIHtcbi8vICAgbGV0IHZhbGlkID0gZmFsc2U7XG4vLyAgIGlmIChyZXMubmFtZSkge1xuLy8gICAgIHZhbGlkID0gcGFyYW1OYW1lUmVnZXgudGVzdChyZXMubmFtZSk7XG4vLyAgIH1cbi8vICAgcmV0dXJuIHZhbGlkO1xuLy8gfVxuXG4vKipcbiAqIFRoZSBzdXBwb3J0ZWQgYXBpVmVyaW9uIGZvciB0aGUgUGlwZWxpbmVzLCBUYXNrcywgZXRjLlxuICovXG5leHBvcnQgY29uc3QgVGVrdG9uVjFBcGlWZXJzaW9uID0gJ3Rla3Rvbi5kZXYvdjEnO1xuXG5leHBvcnQgaW50ZXJmYWNlIE5hbWVkUmVzb3VyY2Uge1xuICByZWFkb25seSBuYW1lPzogc3RyaW5nO1xufVxuXG5leHBvcnQgaW50ZXJmYWNlIE5hbWVLZXlQYWlyIGV4dGVuZHMgTmFtZWRSZXNvdXJjZSB7XG4gIHJlYWRvbmx5IGtleT86IHN0cmluZztcbn1cblxuZXhwb3J0IGZ1bmN0aW9uIHNlY3JldEtleVJlZihuYW1lOiBzdHJpbmcsIGtleTogc3RyaW5nKTogTmFtZUtleVBhaXIge1xuICByZXR1cm4ge1xuICAgIG5hbWU6IG5hbWUsXG4gICAga2V5OiBrZXksXG4gIH07XG59XG5cbi8qKlxuICogQ29udmVuaWVuY2UgbWV0aG9kIGZvciBmb3JtYXR0aW5nIHRoZSB2YWx1ZSBvZiBhIHdvcmtpbmcgZGlyZWN0b3J5LlxuICogQHBhcmFtIHdvcmtzcGFjZVxuICovXG5leHBvcnQgZnVuY3Rpb24gYnVpbGRXb3JraW5nRGlyKHdvcmtzcGFjZTogc3RyaW5nKTogc3RyaW5nIHtcbiAgcmV0dXJuIGAkKHdvcmtzcGFjZXMuJHt3b3Jrc3BhY2V9LnBhdGgpYDtcbn1cblxuLyoqXG4gKiBCdWlsZHMgdGhlIGNvcnJlY3Qgc3RyaW5nIGZvciByZWZlcmVuY2luZyB0aGUgcGFyYW1ldGVyIHNwZWNpZmllZCBieSBgbmFtZWBcbiAqIHRoYXQgY2FuIGJlIHVzZWQgd2hlbiBidWlsZGluZyB0YXNrcyBhbmQgb3RoZXJzLlxuICpcbiAqIEZvciBleGFtcGxlLCBpZiB0aGUgcGFyYW1ldGVyIGlzIGBmb29gLCB0aGUgcmVzdWx0IHdpbGwgYmUgYCQocGFyYW1zLmZvbylgLlxuICogQHBhcmFtIG5hbWUgVGhlIG5hbWUgb2YgdGhlIHBhcmFtZXRlci5cbiAqL1xuZXhwb3J0IGZ1bmN0aW9uIGJ1aWxkUGFyYW0obmFtZTogc3RyaW5nKTogc3RyaW5nIHtcbiAgcmV0dXJuIGAkKHBhcmFtcy4ke25hbWV9KWA7XG59Il19