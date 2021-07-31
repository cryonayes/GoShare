import {useEffect, useState} from "react";
import {useRouter} from 'next/router'
import checkAuth from "../utils/authCheck";
import Navbar from "../components/Navbar";
import Header from "../components/Header";
import File from "../components/File";
import FileContainer from "../components/FileContainer";
import FileUpload from "../components/FileUpload";
import useSWR from "swr";

async function fetcher() {

    const res = await fetch('http://localhost:3000/api/files', {
            credentials: "include",
            method: 'POST'
        }
    )
    return await res.json()
}

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
    }, [])

    const { data } = useSWR('/api/files', fetcher)

    return (
        (userLoggedin && data) ? (
            <>
                <Header title={"Dashboard"}/>
                <div className={"wrapper"} style={{height: "100%"}}>
                    <div className={"content-wrapper"}>
                        <div className={"content"}>
                            <Navbar name={data.data.name} lastname={data.data.lastname}/>
                            <div className={"container-fluid"}>
                                <FileContainer>
                                    {
                                        (data.data.files).map((file, idx) => {
                                            return (<File filename={file.filename} dropdownID={idx}/>)
                                        })
                                    }
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