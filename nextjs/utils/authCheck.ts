interface resp {
    success: boolean,
    message: string,
    data:    any,
}

let checkAuth = async () => {
    const res = await fetch('http://localhost:21942/api/authcheck',
        {
            body: JSON.stringify({token: localStorage.getItem("token")}),
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'POST'
        }
    )
    return await res.json().then((response : resp) => {
        return !!response.success;
    })
};

export default checkAuth;