import {NextApiRequest, NextApiResponse} from "next";

const apiURLs = {
    "/api/files" : "http://localhost:21942/api/files",
    "/api/authcheck" : "http://localhost:21942/api/authcheck",
    "/api/register" : "http://localhost:21942/api/register", // OK
    "/api/share" : "http://localhost:21942/api/share"
}

async function fetchDefault(url : string, body : string, headers, method : "GET" | "POST") {
    return await fetch(url,
        {
            body: body,
            headers: headers,
            method: method,
            credentials: "include"
        }
    )
}

function handleLogin(req: NextApiRequest, res : NextApiResponse) {
    let apiUrl = "http://localhost:21942/api/login";

    fetchDefault(apiUrl, JSON.stringify(req.body), req.headers, "POST").then(async r => {
        let cookie = r.headers.get("set-cookie");
        res.setHeader("set-cookie", cookie);

        let json = await r.json()
        res.json(json)
        res.end()
    })
}

function handleLogout(req: NextApiRequest, res : NextApiResponse) {
    res.setHeader("set-cookie", "token=deleted; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT");
    res.end()
}

function handleAPIs(req : NextApiRequest, res : NextApiResponse) {
    let apiUrl = apiURLs[req.url] || undefined

    if (apiUrl === undefined) {
        res.json({success: false, message: "Invalid endpoint", data: null})
        res.end()
        return
    }

    fetchDefault(apiUrl, JSON.stringify(req.body), req.headers, "POST").then(async r => {
        let json = await r.json()
        res.json(json)
        res.end()
    })
}

const handler = (req: NextApiRequest, res : NextApiResponse) => {

    switch (req.url) {
        case "/api/login":
            handleLogin(req, res);
            break;

        case "/api/logout":
            handleLogout(req, res);
            break;

        default:
            handleAPIs(req, res);
            break;
    }

}

export default handler