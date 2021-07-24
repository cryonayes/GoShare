import styles from '../styles/Register.module.css'
import Head from 'next/head'

function Register(): JSX.Element {
    return(
            <>
            <Head>
            <title>Register - Brand</title>
            <meta charSet="utf-8" />
            <meta name="viewport" content="width=device-width, initial-scale=1.0, shrink-to-fit=no" />
            <link rel="stylesheet" href="assets/bootstrap/css/bootstrap.min.css" />
            <link rel="stylesheet"
                href="https://fonts.googleapis.com/css?family=Nunito:200,200i,300,300i,400,400i,600,600i,700,700i,800,800i,900,900i" />
            <link rel="stylesheet" href="assets/fonts/fontawesome-all.min.css" />
            <link rel="stylesheet" href="assets/fonts/font-awesome.min.css" />
            <link rel="stylesheet" href="assets/fonts/fontawesome5-overrides.min.css" />
        </Head>

        <div className={`${styles.registerContainer} bg-gradient-primary`}>
            <div className="container my-auto">
                <div className="card shadow-lg o-hidden border-0">
                    <div className="card-body p-0">
                        <div className="row">
                            <div className="col-lg-5 d-none d-lg-flex">
                                <div className={`${styles.imageContainer} flex-grow-1 bg-register-image`}></div>
                            </div>
                            <div className="col-lg-7">
                                <div className="p-5">
                                    <div className="text-center">
                                        <h4 className="text-dark mb-4">Create an Account!</h4>
                                    </div>
                                    <form className="user">
                                        <div className="form-group row">
                                            <div className="col-sm-6 mb-3 mb-sm-0">
                                                <input className="form-control form-control-user" type="text"
                                                    id="exampleFirstName" placeholder="First Name" name="first_name" />
                                            </div>
                                            <div className="col-sm-6">
                                                <input className="form-control form-control-user" type="text"
                                                    id="exampleFirstName" placeholder="Last Name" name="last_name" />
                                            </div>
                                        </div>
                                        <div className="form-group">
                                            <input className="form-control form-control-user" type="email"
                                                id="exampleInputEmail" aria-describedby="emailHelp" placeholder="Email Address"
                                                name="email" />
                                        </div>
                                        <div className="form-group row">
                                            <div className="col-sm-6 mb-3 mb-sm-0">
                                                <input className="form-control form-control-user" type="password"
                                                    id="examplePasswordInput" placeholder="Password" name="password" />
                                            </div>
                                            <div className="col-sm-6">
                                                <input className="form-control form-control-user" type="password"
                                                    id="exampleRepeatPasswordInput" placeholder="Repeat Password"
                                                    name="password_repeat" />
                                            </div>
                                        </div>
                                        <button className="btn btn-primary btn-block text-white btn-user" type="submit">Register
                                            Account</button>
                                        <hr />
                                    </form>
                                    <div className="text-center"><a className="small" href="forgot-password.html">Forgot
                                            Password?</a></div>
                                    <div className="text-center"><a className="small" href="login.html">Already have an account?
                                            Login!</a></div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        </>
    )
}

export default Register;
