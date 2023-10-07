import {Event} from "../model/event"
import {EventListItem} from "./event_list_item"
import * as React from "react"

export interface EventListProps {
    events: Event[];
}

export class EventList extends React.Component<EventListProps,{}>{
    render() {
        const items = this.props.events.map(e =>
            <EventListItem event={e} />
        );
        return <table className="table">
            <thead>
                <tr>
                    <th>Event</th>
                    <th>Where</th>
                    <th colSpan={2}>When (start/end)</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                {items}
            </tbody>
        </table>
    }
}