"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.TektonHubLink = void 0;
const fs = require("fs");
const axios_1 = require("axios");
class TektonHubLink {
    /**
   * Grabs all of the Task defined on Tekton Hub.
   * @link https://hub.tekton.dev
   * @returns Promise<TektonTask[]>
   */
    async fetchTekonHubTasks() {
        try {
            const endpoint = 'https://api.hub.tekton.dev/v1/resources';
            const response = await axios_1.default.get(endpoint);
            if (response.status !== 200) {
                console.log('Request was not successful.');
                return [];
            }
            const jsonData = response.data;
            if (!jsonData.hasOwnProperty('data')) {
                console.log("'data' property not found in the JSON.");
                return [];
            }
            const data = jsonData.data;
            return data;
        }
        catch (error) {
            throw Error(`Error creating Tekton hub tasks file error: ${error}`);
        }
    }
    CreateHubTaskFileContent() {
        return this.fetchTekonHubTasks().then((task) => {
            // Make sure we have some tasks
            let fileContents;
            fileContents = [];
            if (task.length === 0) {
                return fileContents;
            }
            // Let's create the tektonHubTask.ts
            // The following code is creating the imports for the Tekton Hub Task
            fileContents.push("import { TektonHubTask } from './tektonHubTasksResolver';");
            fileContents.push("import { Construct } from 'constructs';");
            fileContents.push("import { TaskBuilder } from 'cdk8s-pipelines';");
            task.forEach((item) => {
                if (item.kind !== 'Task') {
                    return;
                }
                let name = item.name;
                name = name.replace(/-/g, '_');
                name = name.replace(/[0-9]/g, '');
                name = this.camelize(name);
                const url = item.latestVersion.rawURL;
                // We want to init an instance of tektonHubTaskResolver
                // let str = `export const ${name} = new TektonHubTask('${url}').build();`;
                let str = `export const ${name} = function(scope: Construct, id: string) : TaskBuilder { return new TektonHubTask(scope, id, '${url}').build(); };`;
                fileContents.push(str);
            });
            return fileContents;
        });
    }
    build() {
        const fileContents = this.CreateHubTaskFileContent();
        const success = fileContents.then((data => {
            try {
                const fp = `${__dirname}/tektonHubTasks.ts`;
                const numberOfTasks = data.length;
                fs.writeFile(fp, data.join('\n'), (err) => {
                    if (err) {
                        throw Error(`Error creating Tekton hub tasks file error: ${err}`);
                    }
                });
                console.log(`${numberOfTasks} Tekton Hub task created.`);
                return true;
            }
            catch (err) {
                console.error(err);
                return false;
            }
        }));
        return success;
    }
    /**
     * Helper function to camel case the Task names
     * @param str string to camel case
     * @returns string
     */
    camelize(str) {
        return str.replace(/(?:^\w|[A-Z]|\b\w)/g, function (word, index) {
            return index === 0 ? word.toLowerCase() : word.toUpperCase();
        }).replace(/\s+/g, '');
    }
}
exports.TektonHubLink = TektonHubLink;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiVGVrdG9uSHViTGluay5qcyIsInNvdXJjZVJvb3QiOiIiLCJzb3VyY2VzIjpbIi4uLy4uL3NyYy90ZWt0b25IdWIvVGVrdG9uSHViTGluay50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiOzs7QUFBQSx5QkFBeUI7QUFDekIsaUNBQTBCO0FBRzFCLE1BQWEsYUFBYTtJQUN4Qjs7OztLQUlDO0lBQ00sS0FBSyxDQUFDLGtCQUFrQjtRQUM3QixJQUFJO1lBQ0YsTUFBTSxRQUFRLEdBQUcseUNBQXlDLENBQUM7WUFDM0QsTUFBTSxRQUFRLEdBQUcsTUFBTSxlQUFLLENBQUMsR0FBRyxDQUFDLFFBQVEsQ0FBQyxDQUFDO1lBQzNDLElBQUksUUFBUSxDQUFDLE1BQU0sS0FBSyxHQUFHLEVBQUU7Z0JBQzNCLE9BQU8sQ0FBQyxHQUFHLENBQUMsNkJBQTZCLENBQUMsQ0FBQztnQkFDM0MsT0FBTyxFQUFFLENBQUM7YUFDWDtZQUNELE1BQU0sUUFBUSxHQUFHLFFBQVEsQ0FBQyxJQUFJLENBQUM7WUFDL0IsSUFBSSxDQUFDLFFBQVEsQ0FBQyxjQUFjLENBQUMsTUFBTSxDQUFDLEVBQUU7Z0JBQ3BDLE9BQU8sQ0FBQyxHQUFHLENBQUMsd0NBQXdDLENBQUMsQ0FBQztnQkFDdEQsT0FBTyxFQUFFLENBQUM7YUFDWDtZQUNELE1BQU0sSUFBSSxHQUFpQixRQUFRLENBQUMsSUFBSSxDQUFDO1lBQ3pDLE9BQU8sSUFBSSxDQUFDO1NBQ2I7UUFBQyxPQUFPLEtBQUssRUFBRTtZQUNkLE1BQU0sS0FBSyxDQUFDLCtDQUErQyxLQUFLLEVBQUUsQ0FBQyxDQUFDO1NBQ3JFO0lBQ0gsQ0FBQztJQUVPLHdCQUF3QjtRQUM5QixPQUFPLElBQUksQ0FBQyxrQkFBa0IsRUFBRSxDQUFDLElBQUksQ0FBQyxDQUFDLElBQUksRUFBRSxFQUFFO1lBQzdDLCtCQUErQjtZQUMvQixJQUFJLFlBQXNCLENBQUM7WUFDM0IsWUFBWSxHQUFHLEVBQUUsQ0FBQztZQUNsQixJQUFJLElBQUksQ0FBQyxNQUFNLEtBQUssQ0FBQyxFQUFFO2dCQUNyQixPQUFPLFlBQVksQ0FBQzthQUNyQjtZQUNELG9DQUFvQztZQUNwQyxxRUFBcUU7WUFDckUsWUFBWSxDQUFDLElBQUksQ0FBQywyREFBMkQsQ0FBQyxDQUFDO1lBQy9FLFlBQVksQ0FBQyxJQUFJLENBQUMseUNBQXlDLENBQUMsQ0FBQztZQUM3RCxZQUFZLENBQUMsSUFBSSxDQUFDLGdEQUFnRCxDQUFDLENBQUM7WUFDcEUsSUFBSSxDQUFDLE9BQU8sQ0FBQyxDQUFDLElBQWdCLEVBQUUsRUFBRTtnQkFDaEMsSUFBSSxJQUFJLENBQUMsSUFBSSxLQUFLLE1BQU0sRUFBRTtvQkFDeEIsT0FBTztpQkFDUjtnQkFDRCxJQUFJLElBQUksR0FBRyxJQUFJLENBQUMsSUFBSSxDQUFDO2dCQUNyQixJQUFJLEdBQUcsSUFBSSxDQUFDLE9BQU8sQ0FBQyxJQUFJLEVBQUUsR0FBRyxDQUFDLENBQUM7Z0JBQy9CLElBQUksR0FBRyxJQUFJLENBQUMsT0FBTyxDQUFDLFFBQVEsRUFBRSxFQUFFLENBQUMsQ0FBQztnQkFDbEMsSUFBSSxHQUFHLElBQUksQ0FBQyxRQUFRLENBQUMsSUFBSSxDQUFDLENBQUM7Z0JBQzNCLE1BQU0sR0FBRyxHQUFHLElBQUksQ0FBQyxhQUFhLENBQUMsTUFBTSxDQUFDO2dCQUN0Qyx1REFBdUQ7Z0JBQ3ZELDJFQUEyRTtnQkFDM0UsSUFBSSxHQUFHLEdBQUcsZ0JBQWdCLElBQUksa0dBQWtHLEdBQUcsZ0JBQWdCLENBQUM7Z0JBQ3BKLFlBQVksQ0FBQyxJQUFJLENBQUMsR0FBRyxDQUFDLENBQUM7WUFDekIsQ0FBQyxDQUFDLENBQUM7WUFDSCxPQUFPLFlBQVksQ0FBQztRQUN0QixDQUFDLENBQUMsQ0FBQztJQUNMLENBQUM7SUFFTSxLQUFLO1FBQ1YsTUFBTSxZQUFZLEdBQUcsSUFBSSxDQUFDLHdCQUF3QixFQUFFLENBQUM7UUFDckQsTUFBTSxPQUFPLEdBQUcsWUFBWSxDQUFDLElBQUksQ0FBQyxDQUFDLElBQUksQ0FBQyxFQUFFO1lBQ3hDLElBQUk7Z0JBQ0YsTUFBTSxFQUFFLEdBQUcsR0FBRyxTQUFTLG9CQUFvQixDQUFDO2dCQUM1QyxNQUFNLGFBQWEsR0FBRyxJQUFJLENBQUMsTUFBTSxDQUFDO2dCQUNsQyxFQUFFLENBQUMsU0FBUyxDQUFDLEVBQUUsRUFBRSxJQUFJLENBQUMsSUFBSSxDQUFDLElBQUksQ0FBQyxFQUFFLENBQUMsR0FBUSxFQUFFLEVBQUU7b0JBQzdDLElBQUksR0FBRyxFQUFFO3dCQUNQLE1BQU0sS0FBSyxDQUFDLCtDQUErQyxHQUFHLEVBQUUsQ0FBQyxDQUFDO3FCQUNuRTtnQkFDSCxDQUFDLENBQUMsQ0FBQztnQkFDSCxPQUFPLENBQUMsR0FBRyxDQUFDLEdBQUcsYUFBYSwyQkFBMkIsQ0FBQyxDQUFDO2dCQUN6RCxPQUFPLElBQUksQ0FBQzthQUNiO1lBQUMsT0FBTyxHQUFRLEVBQUU7Z0JBQ2pCLE9BQU8sQ0FBQyxLQUFLLENBQUMsR0FBRyxDQUFDLENBQUM7Z0JBQ25CLE9BQU8sS0FBSyxDQUFDO2FBQ2Q7UUFDSCxDQUFDLENBQUMsQ0FBQyxDQUFDO1FBQ0osT0FBTyxPQUFPLENBQUM7SUFDakIsQ0FBQztJQUNEOzs7O09BSUc7SUFDSyxRQUFRLENBQUMsR0FBVztRQUMxQixPQUFPLEdBQUcsQ0FBQyxPQUFPLENBQUMscUJBQXFCLEVBQUUsVUFBVSxJQUFZLEVBQUUsS0FBYTtZQUM3RSxPQUFPLEtBQUssS0FBSyxDQUFDLENBQUMsQ0FBQyxDQUFDLElBQUksQ0FBQyxXQUFXLEVBQUUsQ0FBQyxDQUFDLENBQUMsSUFBSSxDQUFDLFdBQVcsRUFBRSxDQUFDO1FBQy9ELENBQUMsQ0FBQyxDQUFDLE9BQU8sQ0FBQyxNQUFNLEVBQUUsRUFBRSxDQUFDLENBQUM7SUFDekIsQ0FBQztDQUNGO0FBdkZELHNDQXVGQyIsInNvdXJjZXNDb250ZW50IjpbImltcG9ydCAqIGFzIGZzIGZyb20gJ2ZzJztcbmltcG9ydCBheGlvcyBmcm9tICdheGlvcyc7XG5pbXBvcnQgeyBUZWt0b25UYXNrIH0gZnJvbSAnLi90YXNrcyc7XG5cbmV4cG9ydCBjbGFzcyBUZWt0b25IdWJMaW5rIHtcbiAgLyoqXG4gKiBHcmFicyBhbGwgb2YgdGhlIFRhc2sgZGVmaW5lZCBvbiBUZWt0b24gSHViLlxuICogQGxpbmsgaHR0cHM6Ly9odWIudGVrdG9uLmRldlxuICogQHJldHVybnMgUHJvbWlzZTxUZWt0b25UYXNrW10+XG4gKi9cbiAgcHVibGljIGFzeW5jIGZldGNoVGVrb25IdWJUYXNrcygpOiBQcm9taXNlPFRla3RvblRhc2tbXT4ge1xuICAgIHRyeSB7XG4gICAgICBjb25zdCBlbmRwb2ludCA9ICdodHRwczovL2FwaS5odWIudGVrdG9uLmRldi92MS9yZXNvdXJjZXMnO1xuICAgICAgY29uc3QgcmVzcG9uc2UgPSBhd2FpdCBheGlvcy5nZXQoZW5kcG9pbnQpO1xuICAgICAgaWYgKHJlc3BvbnNlLnN0YXR1cyAhPT0gMjAwKSB7XG4gICAgICAgIGNvbnNvbGUubG9nKCdSZXF1ZXN0IHdhcyBub3Qgc3VjY2Vzc2Z1bC4nKTtcbiAgICAgICAgcmV0dXJuIFtdO1xuICAgICAgfVxuICAgICAgY29uc3QganNvbkRhdGEgPSByZXNwb25zZS5kYXRhO1xuICAgICAgaWYgKCFqc29uRGF0YS5oYXNPd25Qcm9wZXJ0eSgnZGF0YScpKSB7XG4gICAgICAgIGNvbnNvbGUubG9nKFwiJ2RhdGEnIHByb3BlcnR5IG5vdCBmb3VuZCBpbiB0aGUgSlNPTi5cIik7XG4gICAgICAgIHJldHVybiBbXTtcbiAgICAgIH1cbiAgICAgIGNvbnN0IGRhdGE6IFRla3RvblRhc2tbXSA9IGpzb25EYXRhLmRhdGE7XG4gICAgICByZXR1cm4gZGF0YTtcbiAgICB9IGNhdGNoIChlcnJvcikge1xuICAgICAgdGhyb3cgRXJyb3IoYEVycm9yIGNyZWF0aW5nIFRla3RvbiBodWIgdGFza3MgZmlsZSBlcnJvcjogJHtlcnJvcn1gKTtcbiAgICB9XG4gIH1cblxuICBwcml2YXRlIENyZWF0ZUh1YlRhc2tGaWxlQ29udGVudCgpOiBQcm9taXNlPHN0cmluZ1tdPiB7XG4gICAgcmV0dXJuIHRoaXMuZmV0Y2hUZWtvbkh1YlRhc2tzKCkudGhlbigodGFzaykgPT4ge1xuICAgICAgLy8gTWFrZSBzdXJlIHdlIGhhdmUgc29tZSB0YXNrc1xuICAgICAgbGV0IGZpbGVDb250ZW50czogc3RyaW5nW107XG4gICAgICBmaWxlQ29udGVudHMgPSBbXTtcbiAgICAgIGlmICh0YXNrLmxlbmd0aCA9PT0gMCkge1xuICAgICAgICByZXR1cm4gZmlsZUNvbnRlbnRzO1xuICAgICAgfVxuICAgICAgLy8gTGV0J3MgY3JlYXRlIHRoZSB0ZWt0b25IdWJUYXNrLnRzXG4gICAgICAvLyBUaGUgZm9sbG93aW5nIGNvZGUgaXMgY3JlYXRpbmcgdGhlIGltcG9ydHMgZm9yIHRoZSBUZWt0b24gSHViIFRhc2tcbiAgICAgIGZpbGVDb250ZW50cy5wdXNoKFwiaW1wb3J0IHsgVGVrdG9uSHViVGFzayB9IGZyb20gJy4vdGVrdG9uSHViVGFza3NSZXNvbHZlcic7XCIpO1xuICAgICAgZmlsZUNvbnRlbnRzLnB1c2goXCJpbXBvcnQgeyBDb25zdHJ1Y3QgfSBmcm9tICdjb25zdHJ1Y3RzJztcIik7XG4gICAgICBmaWxlQ29udGVudHMucHVzaChcImltcG9ydCB7IFRhc2tCdWlsZGVyIH0gZnJvbSAnY2RrOHMtcGlwZWxpbmVzJztcIik7XG4gICAgICB0YXNrLmZvckVhY2goKGl0ZW06IFRla3RvblRhc2spID0+IHtcbiAgICAgICAgaWYgKGl0ZW0ua2luZCAhPT0gJ1Rhc2snKSB7XG4gICAgICAgICAgcmV0dXJuO1xuICAgICAgICB9XG4gICAgICAgIGxldCBuYW1lID0gaXRlbS5uYW1lO1xuICAgICAgICBuYW1lID0gbmFtZS5yZXBsYWNlKC8tL2csICdfJyk7XG4gICAgICAgIG5hbWUgPSBuYW1lLnJlcGxhY2UoL1swLTldL2csICcnKTtcbiAgICAgICAgbmFtZSA9IHRoaXMuY2FtZWxpemUobmFtZSk7XG4gICAgICAgIGNvbnN0IHVybCA9IGl0ZW0ubGF0ZXN0VmVyc2lvbi5yYXdVUkw7XG4gICAgICAgIC8vIFdlIHdhbnQgdG8gaW5pdCBhbiBpbnN0YW5jZSBvZiB0ZWt0b25IdWJUYXNrUmVzb2x2ZXJcbiAgICAgICAgLy8gbGV0IHN0ciA9IGBleHBvcnQgY29uc3QgJHtuYW1lfSA9IG5ldyBUZWt0b25IdWJUYXNrKCcke3VybH0nKS5idWlsZCgpO2A7XG4gICAgICAgIGxldCBzdHIgPSBgZXhwb3J0IGNvbnN0ICR7bmFtZX0gPSBmdW5jdGlvbihzY29wZTogQ29uc3RydWN0LCBpZDogc3RyaW5nKSA6IFRhc2tCdWlsZGVyIHsgcmV0dXJuIG5ldyBUZWt0b25IdWJUYXNrKHNjb3BlLCBpZCwgJyR7dXJsfScpLmJ1aWxkKCk7IH07YDtcbiAgICAgICAgZmlsZUNvbnRlbnRzLnB1c2goc3RyKTtcbiAgICAgIH0pO1xuICAgICAgcmV0dXJuIGZpbGVDb250ZW50cztcbiAgICB9KTtcbiAgfVxuXG4gIHB1YmxpYyBidWlsZCgpIHtcbiAgICBjb25zdCBmaWxlQ29udGVudHMgPSB0aGlzLkNyZWF0ZUh1YlRhc2tGaWxlQ29udGVudCgpO1xuICAgIGNvbnN0IHN1Y2Nlc3MgPSBmaWxlQ29udGVudHMudGhlbigoZGF0YSA9PiB7XG4gICAgICB0cnkge1xuICAgICAgICBjb25zdCBmcCA9IGAke19fZGlybmFtZX0vdGVrdG9uSHViVGFza3MudHNgO1xuICAgICAgICBjb25zdCBudW1iZXJPZlRhc2tzID0gZGF0YS5sZW5ndGg7XG4gICAgICAgIGZzLndyaXRlRmlsZShmcCwgZGF0YS5qb2luKCdcXG4nKSwgKGVycjogYW55KSA9PiB7XG4gICAgICAgICAgaWYgKGVycikge1xuICAgICAgICAgICAgdGhyb3cgRXJyb3IoYEVycm9yIGNyZWF0aW5nIFRla3RvbiBodWIgdGFza3MgZmlsZSBlcnJvcjogJHtlcnJ9YCk7XG4gICAgICAgICAgfVxuICAgICAgICB9KTtcbiAgICAgICAgY29uc29sZS5sb2coYCR7bnVtYmVyT2ZUYXNrc30gVGVrdG9uIEh1YiB0YXNrIGNyZWF0ZWQuYCk7XG4gICAgICAgIHJldHVybiB0cnVlO1xuICAgICAgfSBjYXRjaCAoZXJyOiBhbnkpIHtcbiAgICAgICAgY29uc29sZS5lcnJvcihlcnIpO1xuICAgICAgICByZXR1cm4gZmFsc2U7XG4gICAgICB9XG4gICAgfSkpO1xuICAgIHJldHVybiBzdWNjZXNzO1xuICB9XG4gIC8qKlxuICAgKiBIZWxwZXIgZnVuY3Rpb24gdG8gY2FtZWwgY2FzZSB0aGUgVGFzayBuYW1lc1xuICAgKiBAcGFyYW0gc3RyIHN0cmluZyB0byBjYW1lbCBjYXNlXG4gICAqIEByZXR1cm5zIHN0cmluZ1xuICAgKi9cbiAgcHJpdmF0ZSBjYW1lbGl6ZShzdHI6IHN0cmluZyk6IHN0cmluZyB7XG4gICAgcmV0dXJuIHN0ci5yZXBsYWNlKC8oPzpeXFx3fFtBLVpdfFxcYlxcdykvZywgZnVuY3Rpb24gKHdvcmQ6IHN0cmluZywgaW5kZXg6IG51bWJlcikge1xuICAgICAgcmV0dXJuIGluZGV4ID09PSAwID8gd29yZC50b0xvd2VyQ2FzZSgpIDogd29yZC50b1VwcGVyQ2FzZSgpO1xuICAgIH0pLnJlcGxhY2UoL1xccysvZywgJycpO1xuICB9XG59XG5cbiJdfQ==