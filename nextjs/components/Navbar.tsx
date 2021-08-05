import {useRouter} from "next/router";

function Navbar(props) : JSX.Element {

    let mRouter = useRouter()

    const itemStyle = {
        display: "inline",
        padding: "0 .75rem",
        "text-align": "center",
        "line-height": "1",
        "text-decoration": "none"
    }

    const handleKeyPress = (event) => {
        props.onSearch(event)
    }

    const onLogout = async () => {
        await fetch('http://localhost:3000/api/logout', {
            headers: {
                'Content-Type': 'application/json'
            },
            method: 'POST'
        })

        await mRouter.push("/login")
    }

    return (
        <nav className="navbar navbar-light navbar-expand bg-white shadow mb-4 topbar static-top">
            <div className="container-fluid">
                <form className="form-inline d-sm-inline-block mr-auto ml-md-3 my-2 my-md-0 mw-100 navbar-search">
                    <div className="input-group"><input className="bg-light form-control border-0 small" type="text"
                                                        placeholder="Search for ..." onChange={handleKeyPress}/>
                        <div className="input-group-append">
                            <button className="btn btn-primary py-0" type="button"><i className="fas fa-search"/>
                            </button>
                        </div>
                    </div>
                </form>
                <ul className="navbar-nav flex-nowrap ml-auto">
                    <li className="nav-item no-arrow mx-1 d-flex" style={{alignItems: "center", justifyContent:"center"}}>
                        <a style={itemStyle} href="#">{props.name} {props.lastname}</a>
                        <a style={itemStyle} href="#" onClick={onLogout}>Logout</a>
                    </li>
                </ul>
            </div>
        </nav>
    )
}


export default Navbar;