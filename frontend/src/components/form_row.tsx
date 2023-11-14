import * as React from "react";

export interface FormRowProps {
    label?: string;
    children?: React.ReactNode;
}

export class FormRow extends React.Component<FormRowProps, {}> {
    render() {
        return <div className="form-group">
            <label className="col-sm-2 control-label">
                {this.props.label}
            </label>
            <div className="col-sm-10">
                {this.props.children}
            </div>
        </div>
    }
}