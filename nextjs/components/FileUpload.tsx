import {DropEvent, FileRejection, useDropzone} from 'react-dropzone'
import styles from "../styles/FileUpload.module.css"

function FileUpload() {

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

    const uploadFile = (file : File) => {
        let url = "http://localhost:3000/api/upload";

        let formData = new FormData();
        formData.set("testFile", file)

        fetch(url, {
            method: "POST",
            body: formData,
        }).then(r => {
            console.log(r);
        })
    }

    const dropAccepted = (files: File[], event: DropEvent) => {
        (files.map((file) => {
            console.log("Uploading: " + file.name)
            uploadFile(file)
        }))
    }

    const dropRejected = (files: FileRejection[], event: DropEvent) => {
        (files.map((file) => {
            console.log("Rejected: " + JSON.stringify(file))
        }))
    }

    const {acceptedFiles, getRootProps, getInputProps} = useDropzone({
        accept: allowedTypes,
        onDropRejected: dropRejected,
        onDropAccepted: dropAccepted,

    });


    const files = acceptedFiles.map(file => (
        <li key={file.name}>
            {file.name} - {(file.size / 1024).toFixed(2)} KBs
        </li>
    ));

    return (
        <section className={`${styles.uploadContainer} container`}>
            <div {...getRootProps({className: `dropzone ${styles.alignCenter}`})}>
                <input {...getInputProps()} />
                <p className={styles.marginZero}>Drag & drop files here, or click to select files</p>
            </div>
            <aside className={styles.alignCenter}>
                {
                    files.length > 0 ? (<><ul className={styles.marginZero}>{files}</ul></>) : (<></>)
                }
            </aside>
        </section>
    );
}

export default FileUpload;