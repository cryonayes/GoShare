import styles from '../styles/Login.module.css'
import {useRouter} from "next/router";
import SweetAlert from "react-bootstrap-sweetalert";
import {useEffect, useState} from "react";
import checkAuth from "../utils/authCheck";
import Header from "../components/Header";

function Register(): JSX.Element {

    const [successAlert, setSuccessAlert] = useState(false)
    const [errorAlert, setErrorAlert] = useState(null)
    const mRouter = useRouter()

    const showSuccessAlert = () => {
        setSuccessAlert(true)
    }
    const showErrorAlert = (errorMSG) => {
        setErrorAlert(errorMSG)
    }
    const hideErrorAlert = () => {
        setErrorAlert(null)
    }

    function successAlertConfirm() {
        mRouter.push("/dashboard")
    }

    const loginUser = async event => {
        event.preventDefault()

        const res = await fetch('http://localhost:3000/api/login',
            {
                body: JSON.stringify(
                    {
                        email: event.target.email.value,
                        password: event.target.password.value,
                    }),
                headers: {
                    'Content-Type': 'application/json'
                },
                method: 'POST'
            }
        )
        const result = await res.json()

        if (result.success) {
            showSuccessAlert()
        } else {
            showErrorAlert(result.message)
        }
    }
    let router = useRouter()

    useEffect(() => {
        checkAuth().then((auth) => {
            if (auth) {
                router.push("/dashboard")
            }
        })
    }, [])

    return (
        <>
            <Header title={"Login"}/>
            <div className={`${styles.loginContainer} bg-gradient-primary`}>
                <div className="container my-auto">
                    <div className="row justify-content-center">
                        <div className="col-md-9 col-lg-12 col-xl-10">
                            <div className="card shadow-lg o-hidden border-0 my-5">
                                <div className="card-body p-0">
                                    <div className="row">
                                        <div className="col-lg-6 d-none d-lg-flex">
                                            <div className={`${styles.imageContainer} flex-grow-1 bg-login-image`}/>
                                        </div>
                                        <div className="col-lg-6">
                                            <div className="p-5">
                                                <div className="text-center">
                                                    <h4 className="text-dark mb-4">Welcome Back!</h4>
                                                </div>
                                                <form onSubmit={loginUser} className="user">
                                                    <div className="form-group"><input
                                                        className="form-control form-control-user" type="email"
                                                        id="exampleInputEmail" aria-describedby="emailHelp"
                                                        placeholder="Enter Email Address..." name="email"/></div>
                                                    <div className="form-group"><input
                                                        className="form-control form-control-user" type="password"
                                                        id="exampleInputPassword" placeholder="Password"
                                                        name="password"/>
                                                    </div>
                                                    <div className="form-group">
                                                        <div className="custom-control custom-checkbox small">
                                                            <div className="form-check"><input
                                                                className="form-check-input custom-control-input"
                                                                type="checkbox" id="formCheck-1"/><label
                                                                className="form-check-label custom-control-label"
                                                                htmlFor="formCheck-1">Remember Me</label></div>
                                                        </div>
                                                    </div>
                                                    <button className="btn btn-primary btn-block text-white btn-user"
                                                            type="submit">Login
                                                    </button>
                                                    <br/>
                                                </form>
                                                <hr/>
                                                <div className="text-center"><a className="small"
                                                                                href="/forgot-password">Forgot
                                                    Password?</a></div>
                                                <div className="text-center"><a className="small" href="/register">Create
                                                    an Account!</a></div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    {
                        successAlert && <SweetAlert onConfirm={successAlertConfirm} confirmBtnText={"Go to dashboard"}
                                                    title={"Logged in!"}
                                                    success
                                                    openAnim={{name: 'showSweetAlert', duration: 200}}
                                                    closeAnim={{name: 'hideSweetAlert', duration: 200}}
                        />
                    }
                    {
                        errorAlert && <SweetAlert onConfirm={hideErrorAlert} title={errorAlert}
                                                  danger
                                                  openAnim={{name: 'showSweetAlert', duration: 200}}
                                                  closeAnim={{name: 'hideSweetAlert', duration: 200}}
                        />
                    }
                </div>
            </div>
        </>
    )
}

export default Register;
