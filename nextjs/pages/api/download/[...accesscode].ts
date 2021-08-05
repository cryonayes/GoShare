import {NextApiRequest, NextApiResponse} from "next";
import httpProxyMiddleware from "next-http-proxy-middleware";


// For preventing Content-Length corruption
export const config = {
    api: {
        bodyParser: false,
    },
}

export default (req: NextApiRequest, res: NextApiResponse) => {
    httpProxyMiddleware(req, res, {
        target: 'http://localhost:21942'
    })
}
