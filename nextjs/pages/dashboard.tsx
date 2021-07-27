import {useEffect, useState} from "react";
import { useRouter } from 'next/router'
import checkAuth from "../utils/authCheck";
import Navbar from "../components/Navbar";
import Header from "../components/Header";
import styles from "../styles/Dashboard.module.css"
import File from "../components/File";
import FileContainer from "../components/FileContainer";

function Dashboard(): JSX.Element {

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
                            <Navbar/>
                            <div className={"container-fluid"}>
                                <FileContainer>
                                    <File dropdownID={1}/>
                                    <File dropdownID={2}/>
                                    <File dropdownID={3}/>
                                </FileContainer>
                            </div>
                        </div>
                    </div>
                </div>

            </>
        ) : (<></>)
    );
}


export default Dashboard;