import DropdownMenu from "./DropdownMenu";

function File(props) : JSX.Element {


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
                                <DropdownMenu downloadLink={props.downloadLink}/>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}

export default File;