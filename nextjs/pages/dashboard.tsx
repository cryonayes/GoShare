import {useEffect, useState} from "react";
import { useRouter } from 'next/router'
import checkAuth from "../utils/authCheck";
import Navbar from "../components/Navbar";
import Header from "../components/Header";
import File from "../components/File";
import FileContainer from "../components/FileContainer";
import FileUpload from "../components/FileUpload";

// TODO(Get user's name, lastname & files from server)
export async function getServerSideProps() {
    return {
        props : {
            name: "Ayberk",
            lastname: "ESER",
            files : []
        }
    }
}

function Dashboard(props): JSX.Element {

    const [userLoggedin, setUserLoggedin] = useState(false)
    let router = useRouter()

    useEffect(()=> {
        checkAuth().then((auth) => {
            if (!auth) {
                router.push("/login")
            }else {
                setUserLoggedin(true)
            }
        })
    })

    return (
        userLoggedin ? (
            <>
                <Header title={"Dashboard"}/>
                <script src="assets/js/jquery.min.js"/>

                <div className={"wrapper"} style={{height: "100%"}}>
                    <div className={"content-wrapper"}>
                        <div className={"content"}>
                            <Navbar name={props.name} lastname={props.lastname}/>
                            <div className={"container-fluid"}>
                                <FileContainer>
                                    <File dropdownID={1}/>
                                    <File dropdownID={2}/>
                                    <File dropdownID={3}/>
                                </FileContainer>
                                <FileUpload/>
                            </div>
                        </div>
                    </div>
                </div>

            </>
        ) : (<></>)
    );
}


export default Dashboard;