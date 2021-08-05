import React, {ChangeEvent, useEffect, useRef, useState} from "react";
import {useRouter} from 'next/router'
import checkAuth from "../utils/authCheck";
import Navbar from "../components/Navbar";
import Header from "../components/Header";
import File from "../components/File";
import FileContainer from "../components/FileContainer";
import FileUpload from "../components/FileUpload";
import useSWR from "swr";
import SweetAlert from "react-bootstrap-sweetalert";


function Dashboard(): JSX.Element {

    const [isError, setIsError] = useState<string>(null)
    const [userLoggedin, setUserLoggedin] = useState<boolean>(false)
    const [filtered, setFiltered] = useState<UserFileModel[]>([])
    const [shareLink, setShareLink] = useState<string>("")
    const [timeInput, setTimeInput] = useState<number>(24)
    const [showShareLink, setShowShareLink] = useState<string>("")
    let router = useRouter()

    const showErrorMessage = (message) => setIsError(message)

    useEffect(()=> {
        checkAuth().then((auth: boolean) => {
            if (!auth) {
                router.push("/login")
            }else {
                setUserLoggedin(true)
            }
        })
    }, [])

    const fetcher = async () : Promise<APIResponseFiles> => {
        const res = await fetch('http://localhost:3000/api/files', {
                credentials: "include",
                method: 'POST'
            }
        )
        return await res.json()
    }
    const dataSWR = useSWR<APIResponseFiles>('/api/files', fetcher)
    let userData = dataSWR.data

    const filterFiles = (navbarKeyEvent: React.KeyboardEvent<HTMLInputElement>) => {
        let searchKeyword = navbarKeyEvent.currentTarget.value
        let newFiles = userData.data.files.filter((file: UserFileModel) => {
            return file.filename.includes(searchKeyword);
        })
        setFiltered(newFiles)
    }

    const getUserFiles = (): JSX.Element[] => {
        if (filtered.length > 0) {
            return filtered.map((file: UserFileModel, index: number) => {
                return(<File onDeleteFile={onDeleteFile} onShareFile={onShareFile} onUnshareFile={onUnshareFile}
                             filename={file.filename} fileAccessCode={file.access_code}
                             refresh={dataSWR.revalidate} shared={file.shared} />)
            })
        }
        return userData.data.files.map((file: UserFileModel, index: number) => {
            return(<File onDeleteFile={onDeleteFile} onShareFile={onShareFile} onUnshareFile={onUnshareFile}
                         filename={file.filename} fileAccessCode={file.access_code}
                         refresh={dataSWR.revalidate} shared={file.shared} />)
        })
    }

    const onShareFile = (accessCode) => setShareLink(accessCode)

    const onClickShare = async (accessCode) => {
        let dateTime = new Date();
        dateTime.setHours(dateTime.getHours() + timeInput);
        let timeStamp = (dateTime.getTime() / 1000).toFixed(0)

        const res = await fetch('http://localhost:3000/api/share', {
            body: JSON.stringify({
                accesscode: accessCode,
                sharetime: timeStamp.toString()
            }),
            headers: {
                'Content-Type': 'application/json'
            },
            method: 'POST'
        })

        const result = await res.json() as APIResponseShare

        if (result.success) {
            setShowShareLink(result.data.accesslink)
            await dataSWR.revalidate()
        }else {
            setIsError(result.message)
        }
    }

    const onUnshareFile = async (accessCode) => {
        const res = await fetch('http://localhost:3000/api/unshare', {
            body: JSON.stringify({
                accesscode: accessCode,
            }),
            headers: {
                'Content-Type': 'application/json'
            },
            method: 'POST'
        })

        const result = await res.json() as APIResponseShare

        if (result.success) {
            await dataSWR.revalidate()
        } else {
            setIsError(result.message)
        }
    }

    const onDeleteFile = async (accessCode) => {
        const res = await fetch('http://localhost:3000/api/delete', {
            body: JSON.stringify({
                accesscode: accessCode,
            }),
            headers: {
                'Content-Type': 'application/json'
            },
            method: 'POST'
        })

        const result = await res.json() as APIResponseShare

        if (result.success) {
            await dataSWR.revalidate()
        } else {
            setIsError(result.message)
        }
    }

    return (
        (userLoggedin && userData) ? (
            <>
                <Header title={"Dashboard"}/>
                <div className={"wrapper"} style={{height: "100%"}}>
                    <div className={"content-wrapper"}>
                        <div className={"content"}>
                            <Navbar name={userData.data.name} lastname={userData.data.lastname} onSearch={filterFiles}/>
                            <div className={"container-fluid"}>
                                <FileContainer>
                                    {
                                        getUserFiles()
                                    }
                                </FileContainer>
                                <FileUpload onUpload={dataSWR.revalidate} onError={showErrorMessage} />
                            </div>
                        </div>
                    </div>
                </div>

                <SweetAlert onConfirm={()=>{setIsError(null)}} title={isError} danger show={isError!==null}
                            openAnim={{name: 'showSweetAlert', duration: 200}}
                            closeAnim={{name: 'hideSweetAlert', duration: 200}}/>

                <SweetAlert onConfirm={()=>{onClickShare(shareLink); setShareLink("");}} title={"Set time limit"} show={shareLink !== ""}
                            openAnim={{name: 'showSweetAlert', duration: 200}}
                            closeAnim={{name: 'hideSweetAlert', duration: 200}}>
                    <input type="text" className="border rounded-pill border-primary shadow-sm" inputMode="numeric" placeholder="Hours"
                           style={{padding: "5px 20px", width: "100px",textAlign: "center", outlineStyle: "none"}}
                           onChange={(event: ChangeEvent<HTMLInputElement>)=>{setTimeInput(parseInt(event.target.value))}}/>
                </SweetAlert>

                <SweetAlert onConfirm={()=>{setShowShareLink("");}} title={"Copy Link"} show={showShareLink !== ""}
                            openAnim={{name: 'showSweetAlert', duration: 200}}
                            closeAnim={{name: 'hideSweetAlert', duration: 200}}>
                    <a href={document.location.origin+"/api/download/"+showShareLink}>{document.location.origin+"/api/download/"+showShareLink}</a>
                </SweetAlert>
            </>
        ) : (<></>)
    );
}


export default Dashboard;