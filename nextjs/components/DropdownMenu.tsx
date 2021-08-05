import React from "react";
import { Dropdown } from 'react-bootstrap';
import {DropdownToggleProps} from "reactstrap";

function DropdownMenu(props): JSX.Element {

    let CustomDropdown = React.forwardRef<HTMLButtonElement, DropdownToggleProps>( (props, ref) => {
        return(
            <button className={"btn btn-primary text-primary bg-white border rounded-pill d-flex justify-content-center align-items-center"} ref={ref} onClick={event => {event.preventDefault(); props.onClick(event)}}>
                <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 16 16" fill="currentColor" className="bi bi-three-dots" style={{pointerEvents: "none"}}>
                    <path fillRule="evenodd" d="M3 9.5a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3zm5 0a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3zm5 0a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3z"/>
                </svg>
            </button>
        )
    })

    return(
        <Dropdown>
            <Dropdown.Toggle as={CustomDropdown} id={"dropdown"}/>
            <Dropdown.Menu>
                <Dropdown.Item href={"http://localhost:3000/api/download/" + props.downloadLink} download>
                    <i className="fas fa-save pr-2"/>Download
                </Dropdown.Item>

                <Dropdown.Item>
                    <i className="fas fa-share pr-2"/>Share
                </Dropdown.Item>

                <Dropdown.Item href={"#"} className={"bg-danger text-white"}>
                    <i className="fas fa-trash pr-2"/>Delete
                </Dropdown.Item>
            </Dropdown.Menu>
        </Dropdown>
    )
}

export default DropdownMenu

/*
<div className="dropdown d-inline-block">
    <button onClick={onClickHandler} className="btn btn-primary text-primary bg-white border rounded-pill d-flex justify-content-center align-items-center" aria-expanded="false" data-toggle="dropdown" type="button">
        <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 16 16" fill="currentColor" className="bi bi-three-dots" style={{pointerEvents: "none"}}>
            <path fillRule="evenodd" d="M3 9.5a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3zm5 0a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3zm5 0a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3z"/>
        </svg>
    </button>
    <div className="dropdown-menu">
        <a className="dropdown-item" href="#">First Item</a>
        <a className="dropdown-item" href="#">Second Item</a>
        <a className="dropdown-item" href="#">Third Item</a>
    </div>
</div>
 */