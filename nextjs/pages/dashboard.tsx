import {useEffect} from "react";
import { useRouter } from 'next/router'
import checkAuth from "../utils/authCheck";

function Dashboard(): JSX.Element {

    let router = useRouter()

    useEffect(()=> {
        checkAuth().then((auth) => {
            console.log(auth);
            if (!auth) {
                router.push("/login")
            }
        })
    })

    return (
        <>

        </>
    );
}


export default Dashboard;