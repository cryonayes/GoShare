import { useRouter } from 'next/router'
import { useEffect } from "react";
import checkAuth from "../utils/authCheck";

function Index(): JSX.Element {

  let router = useRouter()

  useEffect(()=> {
    checkAuth().then((auth) => {
      console.log(auth);
      if (auth) {
        router.push("/dashboard")
      }else{
        router.push("/login")
      }
    })
  })

  return (
    <p>Redirecting...</p>
  );
}

export default Index;
