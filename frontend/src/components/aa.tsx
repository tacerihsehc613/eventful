import * as React from "react"

interface Props<T> {
    data: T[];
}

function List<T>(props: Props<number>) {
    // Render a list of items of type T
    
}

//<List<number> data={[1, 2, 3]} />

// export class NList extends React.Component<Props<number>>{
//     render(){
//         return <List<number> data={[1, 2, 3]} />;
//     }
// }