import { Event } from 'kubernetes-models/v1';
export declare interface EventData extends QueryType {
	events: Event[];
}
