import {DropEvent, FileRejection, useDropzone} from 'react-dropzone'
import styles from "../styles/FileUpload.module.css"
import {useState} from "react";

function FileUpload(props) {

    const [list, setList] = useState<File[]>([])
    let uploadingList = [] as File[]

    const allowedTypes = [
        "text/plain",
        "text/html",
        "text/css",
        "text/javascript",
        "text/xml",
        "image/gif",
        "image/jpeg",
        "image/png",
        "image/webp",
        "image/svg+xml",
        "application/json",
        "application/pdf",
        "video/3gpp",
        "video/mp4",
        "video/webm",
    ]

    const uploadFile = async (file: File): Promise<APIResponse> => {
        let url = "http://localhost:3000/api/upload";

        let formData = new FormData();
        formData.set("fileupload", file)

        let mData = await fetch(url, {method: "POST", body: formData})
        let jsonData = await mData.json() as APIResponse

        return new Promise(((resolve, reject) => {
            if (jsonData.success) {
                resolve(jsonData)
            }else {
                reject(jsonData)
            }
        }))
    }

    const dropAccepted = (filesAccepted: File[]) => {
        setList(filesAccepted)
        uploadingList = filesAccepted

        filesAccepted.map(async (file: File) => {
            try {
                let response = await uploadFile(file) as APIResponseUpload
                uploadingList = uploadingList.filter((fileAccepted: File) => {
                    return fileAccepted.name != response.data.filename;
                })
                setList(uploadingList)
                props.onUpload()
            }catch (error) {
                props.onError(error.message)
            }
        })
    }

    const dropRejected = (filesRejected: FileRejection[], event: DropEvent) => {
        (filesRejected.map((file: FileRejection) => {
            props.onError("File type not allowed!")
        }))
    }

    const {getRootProps, getInputProps} = useDropzone({
        accept: allowedTypes,
        onDropRejected: dropRejected,
        onDropAccepted: dropAccepted,
    });



    return (
        <section className={`${styles.uploadContainer} container`}>
            <div {...getRootProps({className: `dropzone ${styles.alignCenter}`})}>
                <input {...getInputProps()} />
                <p className={`${styles.marginZero} text-center`}>Drag & drop files here, or click to select files</p>
            </div>
            <aside className={styles.alignCenter}>
                {
                    list.length > 0 ? (<><ul className={styles.marginZero}>
                        {
                            list.map((file: File) => (
                                <li key={file.name}>
                                    {file.name} - {(file.size / 1024).toFixed(2)} KBs
                                </li>
                            ))
                        }
                    </ul></>) : (<></>)
                }
            </aside>
        </section>
    );
}

export default FileUpload;