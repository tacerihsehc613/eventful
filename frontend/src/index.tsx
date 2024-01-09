import * as React from 'react';
import * as ReactDOM from 'react-dom';
import {HashRouter as Router, Route, Routes} from "react-router-dom";
import { Hello } from './components/hello';
import { EventListContainer } from './components/event_list_container';
import { Navigation } from './components/navigation';
import { EventBookingFormContainer } from './components/event_booking_form_container';

// ReactDOM.render (
//     <div className="container">
//         <h1>MyEvents</h1>
//         <Hello name="World" />
//         <EventListContainer eventListURL="http://localhost:8181" />
//     </div>,
//     document.getElementById('pEvents-app')

// )

// class App extends React.Component<{},{}> {
//     render() {
//         //const eventList = () => <EventListContainer eventServiceURL="http://localhost:8181" />;
//         return <Router>
//             <Navigation brandName="MyEvents" />
//             <div className="container">
//                 <h1>MyEvents</h1>
//                 <Route path="/" element={<EventListContainer eventServiceURL="http://localhost:8181" />} />
//             </div>
//         </Router>
//     }
// }

//<Route path="/events/:id/book" element={<EventBookingFormContainer eventServiceURL="http://localhost:8181" bookingServiceURL="http://localhost:8181" />} />

/*class App extends React.Component<{},{}> {
    render() {
        //const eventList = () => <EventListContainer eventServiceURL="http://localhost:8181" />;
        return <Router>
            <Navigation brandName="MyEvents" />
            <div className="container">
                <h1>MyEvents</h1>
                <Routes>
                <Route path="/" element={<EventListContainer eventServiceURL="http://localhost:8181"/>} />
                <Route path="/events/:id/book" element={<EventBookingFormContainer eventServiceURL="http://localhost:8181" bookingServiceURL="http://localhost:8282" />} />
                </Routes>
            </div>
        </Router>
    }
}*/

class App extends React.Component<{},{}> {
    render() {
        //const eventList = () => <EventListContainer eventServiceURL="http://localhost:8181" />;
        return <Router>
            <Navigation brandName="MyEvents" />
            <div className="container">
                <h1>MyEvents</h1>
                <Routes>
                <Route path="/" element={<EventListContainer eventServiceURL="http://0.0.0.0:8181"/>} />
                <Route path="/events/:id/book" element={<EventBookingFormContainer eventServiceURL="http://0.0.0.0:8181" bookingServiceURL="http://0.0.0.0:8282" />} />
                </Routes>
            </div>
        </Router>
    }
} 
/*class App extends React.Component<{},{}> {
    render() {
        //const eventList = () => <EventListContainer eventServiceURL="http://localhost:8181" />;
        return <Router>
            <Navigation brandName="MyEvents" />
            <div className="container">
                <h1>MyEvents</h1>
                <Routes>
                <Route path="/" element={<EventListContainer eventServiceURL="http://api.myevents.example"/>} />
                <Route path="/events/:id/book" element={<EventBookingFormContainer eventServiceURL="http://api.myevents.example" bookingServiceURL="http://api.myevents.example" />} />
                </Routes>
            </div>
        </Router>
    }
}*/

ReactDOM.render(
    <App />,
    document.getElementById('pEvents-app')
);