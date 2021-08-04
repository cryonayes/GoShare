function File(props) : JSX.Element {

    const toggleDropdown = (event) => {
        let id = event.currentTarget.id.split('-')[1]
        let dropDown = document.getElementById('dropdown-'+id) || undefined
        if (dropDown == undefined) { return }

        if (dropDown.style.display == "block") {
            dropDown.style.display = "none"
        }else {
            dropDown.style.display = "block"
        }
    }

    return (
        <>
            <div className="col-md-6 col-xl-3 mb-4">
                <div className="card shadow border-left-primary py-2">
                    <div className="card-body">
                        <div className="row align-items-center no-gutters">
                            <div className="col-auto"><i className="far fa-file fa-2x text-gray-800"/></div>
                            <div className="col ml-2">
                                <div className="text-dark font-weight-bold h5 mb-0"><span>{props.filename}</span></div>
                            </div>
                            <div className="col-auto">
                                <div className={`dropdown no-arrow my-dropdown`}>
                                    <button className="btn btn-primary" type="button" id={`button-${props.dropdownID}`} onClick={toggleDropdown} style={{background: "transparent", color: "black", padding: "0px", width: "1.5rem", height: "1.5rem", border: "none", outline: "none" }}>
                                        <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" fill="none" style={{height: "100%", width: "100%", outline: "none"}}>
                                            <path fillRule="evenodd" clipRule="evenodd" d="M5 15C6.65685 15 8 13.6569 8 12C8 10.3431 6.65685 9 5 9C3.34315 9 2 10.3431 2 12C2 13.6569 3.34315 15 5 15ZM5 13C5.55228 13 6 12.5523 6 12C6 11.4477 5.55228 11 5 11C4.44772 11 4 11.4477 4 12C4 12.5523 4.44772 13 5 13Z" fill="currentColor"/>
                                            <path fillRule="evenodd" clipRule="evenodd" d="M12 15C13.6569 15 15 13.6569 15 12C15 10.3431 13.6569 9 12 9C10.3431 9 9 10.3431 9 12C9 13.6569 10.3431 15 12 15ZM12 13C12.5523 13 13 12.5523 13 12C13 11.4477 12.5523 11 12 11C11.4477 11 11 11.4477 11 12C11 12.5523 11.4477 13 12 13Z" fill="currentColor"/>
                                            <path fillRule="evenodd" clipRule="evenodd" d="M22 12C22 13.6569 20.6569 15 19 15C17.3431 15 16 13.6569 16 12C16 10.3431 17.3431 9 19 9C20.6569 9 22 10.3431 22 12ZM20 12C20 12.5523 19.5523 13 19 13C18.4477 13 18 12.5523 18 12C18 11.4477 18.4477 11 19 11C19.5523 11 20 11.4477 20 12Z" fill="currentColor"/>
                                        </svg>
                                    </button>
                                    <div id={`dropdown-${props.dropdownID}`} className="dropdown-menu flex-column" style={{padding: "0px", overflow: "hidden"}}>
                                        <a className="dropdown-item" href={"http://localhost:3000/api/download/" + props.downloadLink} download><i className="fas fa-save pr-2"/>Download</a>
                                        <a className="dropdown-item" href="#"><i className="fas fa-share pr-2"/>Share</a>
                                        <a className="dropdown-item bg-danger text-white" href="#"><i className="fas fa-trash pr-2"/>Delete</a>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}

export default File;