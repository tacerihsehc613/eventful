import * as React from "react";
import {EventBookingForm} from "./event_booking_form";
import {Event} from "../model/event";
import { useParams } from 'react-router-dom';

export interface EventBookingFormContainerProps {
    //eventID: string;
    eventServiceURL: string;
    bookingServiceURL: string;
}

export interface EventBookingFormContainerState {
    state: "loading" | "ready" | "saving" | "done" | "error";
    event?: Event;
}

export class EventBookingFormContainer extends React.Component<EventBookingFormContainerProps, EventBookingFormContainerState>{
    constructor(p: EventBookingFormContainerProps) {
        super(p);
        this.state = {
            state: "loading"
        };
        //const { id } = useParams(); 
        let params = useParams();
        const id = params.id || "654cf01d9635f821f4318d54";
        alert('EventBookingFormContainer id:'+ id);
        //fetch(p.eventServiceURL+"/events/"+p.eventID)
        alert("ssss")
        fetch(p.eventServiceURL+"/events/"+id, {method: "GET"})
            .then<Event>(response=> response.json())
            .then(event=>{
                alert('Received event data:'+ event);
                this.setState({
                    state: "ready",
                    event: event
                });
            })
            .catch(error => {
                alert('Error fetching event data:'+ error); // Log any errors
                this.setState({
                    state: "error"
                });
            });
    }
    render() {
        if (this.state.state === "loading") {
            return <div>Loading...</div>;
        }
        if (this.state.state === "saving") {
            return <div>Saving...</div>;
        }
        if (this.state.state === "done") {
            return <div className="alert alert-success">Booking completed! Thank you!</div>;
        }
        if (this.state.state === "error" || !this.state.event) {
            return <div className="alert alert-danger">Unknown error!</div>;
        }
        if (this.state.state === "ready" && this.state.event) {
            // Print the received event data in the console
            console.log('Received event data:', this.state.event);

            return (
                <EventBookingForm
                    event={this.state.event}
                    onSubmit={seats => this.handleSubmit(seats)}
                />
            );
        }
        return <EventBookingForm event={this.state.event} onSubmit={seats => this.handleSubmit(seats)} />;
    }

    private handleSubmit(seats: number) {
        //const url = this.props.bookingServiceURL+"/events/"+this.props.eventID+"/bookings";
        const { id } = useParams(); 
        const url = this.props.bookingServiceURL+"/events/"+id+"/bookings";
        const payload = {seats: seats};

        this.setState({
            event: this.state.event,
            state: "saving"
        });
        fetch(url, {
            method: "POST",
            body: JSON.stringify(payload),
        })
            .then(response => {
                this.setState({
                    event: this.state.event,
                    state: response.ok ? "done" : "error"
                });
            });
    }
}