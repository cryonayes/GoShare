interface resp {
    success: boolean,
    message: string,
    data:    any,
}

let checkAuth = async () => {
    const res = await fetch('http://localhost:3000/api/authcheck',
        {
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