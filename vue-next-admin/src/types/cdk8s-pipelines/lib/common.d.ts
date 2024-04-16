/**
 * The supported apiVerion for the Pipelines, Tasks, etc.
 */
export declare const TektonV1ApiVersion = 'tekton.dev/v1';
export interface NamedResource {
	name?: string;
}
export interface NameKeyPair extends NamedResource {
	key?: string;
}
export declare function secretKeyRef(name: string, key: string): NameKeyPair;
/**
 * Convenience method for formatting the value of a working directory.
 * @param workspace
 */
export declare function buildWorkingDir(workspace: string): string;
/**
 * Builds the correct string for referencing the parameter specified by `name`
 * that can be used when building tasks and others.
 *
 * For example, if the parameter is `foo`, the result will be `$(params.foo)`.
 * @param name The name of the parameter.
 */
export declare function buildParam(name: string): string;
