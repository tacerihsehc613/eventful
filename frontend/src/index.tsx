import * as React from 'react';
import * as ReactDOM from 'react-dom';
import {HashRouter as Router, Route, Routes} from "react-router-dom";
import { Hello } from './components/hello';
import { EventListContainer } from './components/event_list_container';
import { Navigation } from './components/navigation';

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

class App extends React.Component<{},{}> {
    render() {
        //const eventList = () => <EventListContainer eventServiceURL="http://localhost:8181" />;
        return <Router>
            <Navigation brandName="MyEvents" />
            <div className="container">
                <h1>MyEvents</h1>
                <Routes>
                <Route path="/" element={<EventListContainer eventServiceURL="http://localhost:8181" />} />
                </Routes>
            </div>
        </Router>
    }
}


ReactDOM.render(
    <App />,
    document.getElementById('pEvents-app')
);