import { TektonTask } from './tasks';
export declare class TektonHubLink {
    /**
   * Grabs all of the Task defined on Tekton Hub.
   * @link https://hub.tekton.dev
   * @returns Promise<TektonTask[]>
   */
    fetchTekonHubTasks(): Promise<TektonTask[]>;
    private CreateHubTaskFileContent;
    build(): Promise<boolean>;
    /**
     * Helper function to camel case the Task names
     * @param str string to camel case
     * @returns string
     */
    private camelize;
}
