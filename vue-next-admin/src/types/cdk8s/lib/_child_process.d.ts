/// <reference types="node" />
/****************************************************************************************
 * Expose `child_process` via our own object that can be easily patched by jest for tests.
 * Consumers of the `child_process` module should add functions to this object and import it
 * wherever needed.
 */
import { spawnSync } from 'child_process';
export declare const _child_process: {
    spawnSync: typeof spawnSync;
};
