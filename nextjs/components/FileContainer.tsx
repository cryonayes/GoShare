function FileContainer(props) : JSX.Element {
    return (
        <>
            <div className="row">
                {props.children}
            </div>
        </>
    )
}

export default FileContainer;