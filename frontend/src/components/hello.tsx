import * as React from "react";
import { useParams } from 'react-router-dom';
import { useEffect } from "react";

export interface HelloProps {
    name: string;
}

// export class Hello extends React.Component<HelloProps, {}> {
//     render() {
//         return <div>Hello {this.props.name}!</div>
//     }
// }

// export class Hello extends React.Component<HelloProps, {}> {
//     private logId() {
//         let params = useParams();
//         const id = params.id || "654cf01d9635f821f4318d54";
//         alert('EventBookingFormContainer id:' + id);
//     }

//     render() {
//         this.logId(); // Call the method here or wherever needed
//         return <div>Hello {this.props.name}!</div>
//     }
// }

// export const Hello = (props: HelloProps) => {
export const Hello: React.FC<HelloProps> = (props) => {
    const params = useParams();
    const id = params.id || "654cf01d9635f821f4318d54";

    useEffect(() => {
        console.log('even inside useEffect', params.id);
    }, []);

    return <div>Hello {props.name}!</div>;
}









