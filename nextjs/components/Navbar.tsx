function Navbar() : JSX.Element {
    return (
        <nav className="navbar navbar-light navbar-expand bg-white shadow mb-4 topbar static-top">
            <div className="container-fluid">
                <form className="form-inline d-sm-inline-block mr-auto ml-md-3 my-2 my-md-0 mw-100 navbar-search">
                    <div className="input-group"><input className="bg-light form-control border-0 small" type="text"
                                                        placeholder="Search for ..."/>
                        <div className="input-group-append">
                            <button className="btn btn-primary py-0" type="button"><i className="fas fa-search"/>
                            </button>
                        </div>
                    </div>
                </form>
                <ul className="navbar-nav flex-nowrap ml-auto">
                    <li className="nav-item dropdown no-arrow mx-1">
                        <a className="nav-link" style={{textAlign: "center"}} href="#">Ayberk ESER</a>
                    </li>
                </ul>
            </div>
        </nav>
    )
}


export default Navbar;