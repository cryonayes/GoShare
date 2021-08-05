import React, {ForwardedRef, PropsWithChildren} from "react";
import { Dropdown } from 'react-bootstrap';
import {DropdownToggleProps} from "reactstrap";

function DropdownMenu(props): JSX.Element {

    let CustomDropdown = React.forwardRef<HTMLButtonElement, DropdownToggleProps>( (props: PropsWithChildren<DropdownToggleProps>, ref: ForwardedRef<HTMLButtonElement>) => {
        return(
            <button className={"btn btn-primary text-primary bg-white border rounded-pill d-flex justify-content-center align-items-center"} ref={ref} onClick={event => {event.preventDefault(); props.onClick(event)}}>
                <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 16 16" fill="currentColor" className="bi bi-three-dots" style={{pointerEvents: "none"}}>
                    <path fillRule="evenodd" d="M3 9.5a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3zm5 0a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3zm5 0a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3z"/>
                </svg>
            </button>
        )
    })

    const clickedShare   = () => props.onShare(props.fileAccessCode)
    const clickedUnshare = () => props.onUnshare(props.fileAccessCode)
    const clickedDelete  = () => props.onDeleteFile(props.fileAccessCode)


    return(
        <Dropdown>
            <Dropdown.Toggle as={CustomDropdown} id={"dropdown"}/>
            <Dropdown.Menu className={"p-0"}>
                <Dropdown.Item href={"http://localhost:3000/api/download/" + props.fileAccessCode} download className={"font-weight-bold"}>
                    <i className="fas fa-save pr-2"/>Download
                </Dropdown.Item>

                <Dropdown.Item onClick={props.shared ? (clickedUnshare) : (clickedShare) } className={props.shared ? ("bg-warning text-white font-weight-bold") : "font-weight-bold"}>
                    {props.shared ? (<><i className="fas fa-reply pr-2"/>Unshare</>) : (<><i className="fas fa-share pr-2"/>Share</>)}
                </Dropdown.Item>

                <Dropdown.Item onClick={clickedDelete} className={"bg-danger text-white font-weight-bold"}>
                    <i className="fas fa-trash pr-2"/>Delete
                </Dropdown.Item>
            </Dropdown.Menu>
        </Dropdown>
    )
}

export default DropdownMenu