import Head from "next/head";

function Header(props) : JSX.Element {
    return (
        <Head>
            <title>{props.title} - GoShare</title>
            <meta charSet="utf-8" />
            <meta name="viewport" content="width=device-width, initial-scale=1.0, shrink-to-fit=no" />
            <link rel="stylesheet" href="assets/bootstrap/css/bootstrap.min.css" />
            <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Nunito:200,200i,300,300i,400,400i,600,600i,700,700i,800,800i,900,900i" />
            <link rel="stylesheet" href="assets/fonts/fontawesome-all.min.css" />
            <link rel="stylesheet" href="assets/fonts/font-awesome.min.css" />
            <link rel="stylesheet" href="assets/fonts/fontawesome5-overrides.min.css" />
        </Head>
    )
}

export default Header;