import * as React from "react";
import { useState, useEffect } from "react";
import { EventBookingForm } from "./event_booking_form";
import { Event } from "../model/event";
import { useParams } from 'react-router-dom';

interface EventBookingFormContainerProps {
    eventServiceURL: string;
    bookingServiceURL: string;
}

interface EventBookingFormContainerState {
    state: "loading" | "ready" | "saving" | "done" | "error";
    event?: Event;
}

export const EventBookingFormContainer: React.FC<EventBookingFormContainerProps> = (props) => {
    const [state, setState] = useState<EventBookingFormContainerState>({
        state: "loading"
    });

    const { id } = useParams();
    
    useEffect(() => {
        const eventId = id || "654cf01d9635f821f4318d54";
        
        fetch(`${props.eventServiceURL}/events/${eventId}`, { method: "GET" })
            .then<Event>(response => response.json())
            .then(event => {
                setState({
                    state: "ready",
                    event: event
                });
            })
            .catch(error => {
                setState({
                    state: "error"
                });
            });
    }, [id, props.eventServiceURL]);

    const handleSubmit = (seats: number) => {
        //const bookingUrl = `${props.bookingServiceURL}/events/${id}/bookings`;
        const bookingUrl = `${props.bookingServiceURL}/bookings/${id}`;
        const payload = { seats: seats };

        setState({
            event: state.event,
            state: "saving"
        });

        fetch(bookingUrl, {
            method: "POST",
            body: JSON.stringify(payload),
        })
            .then(response => {
                setState({
                    event: state.event,
                    state: response.ok ? "done" : "error"
                });
            });
    };

    if (state.state === "loading") {
        return <div>Loading...</div>;
    }

    if (state.state === "saving") {
        return <div>Saving...</div>;
    }

    if (state.state === "done") {
        return <div className="alert alert-success">Booking completed! Thank you!</div>;
    }

    if (state.state === "error" || !state.event) {
        return <div className="alert alert-danger">Unknown error!</div>;
    }

    if (state.state === "ready" && state.event) {
        console.log('Received event data:', state.event);

        return (
            <EventBookingForm
                event={state.event}
                onSubmit={seats => handleSubmit(seats)}
            />
        );
    }

    return <EventBookingForm event={state.event} onSubmit={seats => handleSubmit(seats)} />;
};

//export default EventBookingFormContainer;