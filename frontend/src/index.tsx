import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { Hello } from './components/hello';
import { EventListContainer } from './components/event_list_container';

// ReactDOM.render (
//     <div className="container">
//         <h1>MyEvents</h1>
//         <Hello name="World" />
//         <EventListContainer eventListURL="http://localhost:8181" />
//     </div>,
//     document.getElementById('pEvents-app')

// )

class App extends React.Component<{},{}> {
    render() {
        return <div className="container">
            <h1>MyEvents</h1>
            <EventListContainer eventListURL="http://localhost:8181" />
        </div>
    }
}

ReactDOM.render(
    <App />,
    document.getElementById('pEvents-app')
);