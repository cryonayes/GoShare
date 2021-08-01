import {useEffect, useState} from "react";
import {useRouter} from 'next/router'
import checkAuth from "../utils/authCheck";
import Navbar from "../components/Navbar";
import Header from "../components/Header";
import File from "../components/File";
import FileContainer from "../components/FileContainer";
import FileUpload from "../components/FileUpload";
import useSWR from "swr";

async function fetcher() : Promise<APIResponseFiles> {

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

    const dataSWR = useSWR<APIResponseFiles>('/api/files', fetcher)
    let userData = dataSWR.data

    return (
        (userLoggedin && userData) ? (
            <>
                <Header title={"Dashboard"}/>
                <div className={"wrapper"} style={{height: "100%"}}>
                    <div className={"content-wrapper"}>
                        <div className={"content"}>
                            <Navbar name={userData.data.name} lastname={userData.data.lastname}/>
                            <div className={"container-fluid"}>
                                <FileContainer>
                                    {
                                        (userData.data.files).map((file, idx) => {
                                            return (<File filename={file.filename} dropdownID={idx}/>)
                                        })
                                    }
                                </FileContainer>
                                <FileUpload onUpload={dataSWR.revalidate}/>
                            </div>
                        </div>
                    </div>
                </div>
            </>
        ) : (<></>)
    );
}


export default Dashboard;