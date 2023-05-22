import { Event } from 'kubernetes-types/core/v1';
declare interface EventData extends QueryType {
	events: Event[];
}
