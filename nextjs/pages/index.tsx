import { useRouter } from 'next/router'
import { useEffect } from "react";

function Index(): JSX.Element {

  let router = useRouter();

  useEffect(() => {
    router.push('/login.html');
  })


  return (
    <p>Redirecting...</p>
  );
}

export default Index;
