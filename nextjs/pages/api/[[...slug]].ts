import {NextApiRequest, NextApiResponse} from "next";

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

function handleRegister(req: NextApiRequest, res : NextApiResponse) {}

function handleAuthCheck(req: NextApiRequest, res : NextApiResponse) {

    let apiUrl = "http://localhost:21942/api/authcheck";

    req.headers["X-TOKEN"] = req.cookies.token;

    fetchDefault(apiUrl, "", req.headers, "POST").then(async r => {
        let json = await r.json()
        res.json(json)
        res.end()
    })
}

function handleUserFiles(req : NextApiRequest, res : NextApiResponse) {
    let apiUrl = "http://localhost:21942/api/files";

    req.headers["X-TOKEN"] = req.cookies.token;

    fetchDefault(apiUrl, "", req.headers, "POST").then(async r => {
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

        case "/api/register":
            handleRegister(req, res);
            break;

        case "/api/authcheck":
            handleAuthCheck(req, res);
            break;

        case "/api/files":
            handleUserFiles(req, res);
            break;
    }

}

export default handler