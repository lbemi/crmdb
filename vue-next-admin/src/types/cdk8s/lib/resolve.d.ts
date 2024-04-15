import { ApiObject } from './api-object';
/**
 * Context object for a specific resolution process.
 */
export declare class ResolutionContext {
    /**
     * Which ApiObject is currently being resolved.
     */
    readonly obj: ApiObject;
    /**
     * Which key is currently being resolved.
     */
    readonly key: string[];
    /**
     * The value associated to the key currently being resolved.
     */
    readonly value: any;
    /**
     * The replaced value that was set via the `replaceValue` method.
     */
    replacedValue: any;
    /**
     * Whether or not the value was replaced by invoking the `replaceValue` method.
     */
    replaced: boolean;
    constructor(
    /**
     * Which ApiObject is currently being resolved.
     */
    obj: ApiObject, 
    /**
     * Which key is currently being resolved.
     */
    key: string[], 
    /**
     * The value associated to the key currently being resolved.
     */
    value: any);
    /**
     * Replaces the original value in this resolution context
     * with a new value. The new value is what will end up in the manifest.
     */
    replaceValue(newValue: any): void;
}
/**
 * Contract for resolver objects.
 */
export interface IResolver {
    /**
     * This function is invoked on every property during cdk8s synthesis.
     * To replace a value, implementations must invoke `context.replaceValue`.
     */
    resolve(context: ResolutionContext): void;
}
/**
 * Resolvers instanecs of `Lazy`.
 */
export declare class LazyResolver implements IResolver {
    resolve(context: ResolutionContext): void;
}
/**
 * Resolves implicit tokens.
 */
export declare class ImplicitTokenResolver implements IResolver {
    resolve(context: ResolutionContext): void;
}
/**
 * Resolves union types that allow using either number or string (as generated by the CLI).
 *
 * E.g IntOrString, Quantity, ...
 */
export declare class NumberStringUnionResolver implements IResolver {
    private static readonly TYPES;
    resolve(context: ResolutionContext): void;
}
/**
 * Resolves any value attached to a specific ApiObject.
 */
export declare function resolve(key: string[], value: any, apiObject: ApiObject): any;
