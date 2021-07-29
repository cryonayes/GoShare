import {useDropzone} from 'react-dropzone'
import styles from "../styles/FileUpload.module.css"

function FileUpload() {
    const {acceptedFiles, getRootProps, getInputProps} = useDropzone();

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