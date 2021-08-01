let checkAuth = async () => {
    const res = await fetch('http://localhost:3000/api/authcheck',
        {
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'POST'
        }
    )
    return await res.json().then((response : APIResponse) => {
        return !!response.success;
    })
};

export default checkAuth;