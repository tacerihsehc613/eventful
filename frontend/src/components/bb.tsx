import React = require("react");

interface Props<T> {
    data: T[];
}

function List<T>(props: Props<number>)  {
    // Render a list of items of type T
    const listItems = props.data.map((item, index) => (
        <li key={index}>{item}</li>
    ));

    return <ul>{listItems}</ul>;
}

class NList extends React.Component<Props<number>> {
    render() {
        return <List<number> data={[1, 2, 3]} />;
        //const data: number[] = [1, 2, 3];
        //return <List<number> data={data} />;
    }
}

export default NList;





