import * as React from "react";
import {Event} from "../model/event";
import {FormRow} from "./form_row";

export interface EventBookingFormProps {
    event: Event;
    onSubmit: (seats: number) => any;
}

export interface EventBookingFormState {
    seats: number;
}

export class EventBookingForm extends React.Component<EventBookingFormProps, EventBookingFormState>{
    constructor(p: EventBookingFormProps) {
        super(p);
        this.state = {
            seats: 1
        };
    }

    private handleNewAmount(event: React.ChangeEvent<HTMLSelectElement>) {
        const state: EventBookingFormState = {
            seats: parseInt(event.target.value)
        }
        this.setState(state);
    }

    render() {
        return <div>
            <h2>Book tickets for {this.props.event.Name}</h2>
            <form className='form-horizontal'>
                <FormRow label="Event">
                    <p className="form-control-static">{this.props.event.Name}</p>
                </FormRow>
                <FormRow label="Number of tickets">
                    <select className="form-control" value={this.state.seats} onChange={event => this.handleNewAmount(event)}>
                        <option value="1">1</option>
                        <option value="2">2</option>
                        <option value="3">3</option>
                        <option value="4">4</option>
                    </select>
                </FormRow>
                <FormRow>
                    <button className="btn btn-primary" onClick={() => this.props.onSubmit(this.state.seats)}>Submit order</button>
                </FormRow>
            </form>
        </div>
            
    }
}