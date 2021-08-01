import styles from '../styles/Register.module.css'
import {useEffect, useState} from "react";
import SweetAlert from "react-bootstrap-sweetalert";
import {useRouter} from "next/router";
import checkAuth from "../utils/authCheck";
import Header from "../components/Header";


function Register(): JSX.Element {

    const [successAlert, setSuccessAlert] = useState(false)
    const [errorAlert, setErrorAlert] = useState(null)

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
        window.location.replace("/login")
    }

    const registerUser = async event => {
        event.preventDefault()
        
        const res = await fetch('http://localhost:3000/api/register',
            {
              body: JSON.stringify(
                {
                    name: event.target.first_name.value,
                    lastname: event.target.last_name.value,
                    email: event.target.email.value,
                    password: event.target.password.value,
                    passwordRepeat: event.target.password_repeat.value
                }),
                headers: {
                    'Content-Type': 'application/json'
                },
                method: 'POST'
            }
          )
        const result = await res.json() as APIResponse

        if(result.success) {
            showSuccessAlert()
        }else {
            showErrorAlert(result.message)
        }
    }

    let router = useRouter()

    useEffect(()=> {
        checkAuth().then((auth) => {
            console.log(auth);
            if (auth) {
                router.push("/dashboard")
            }
        })
    })

    return(
        <>
        <Header title={"Register"}/>
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
                                    <form onSubmit={registerUser} method="post" className="user">
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
                                    <div className="text-center"><a className="small" href="/forgot-password">Forgot
                                            Password?</a></div>
                                    <div className="text-center"><a className="small" href="/login">Already have an account?
                                            Login!</a></div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                {
                    successAlert && <SweetAlert onConfirm={successAlertConfirm} confirmBtnText={"Go to login"} title={"Registered!"}
                                                success
                                                openAnim={{ name: 'showSweetAlert', duration: 200 }}
                                                closeAnim={{ name: 'hideSweetAlert', duration: 200 }}
                    />
                }
                {
                    errorAlert && <SweetAlert onConfirm={hideErrorAlert} title={errorAlert}
                                              danger
                                              openAnim={{ name: 'showSweetAlert', duration: 200 }}
                                              closeAnim={{ name: 'hideSweetAlert', duration: 200 }}
                    />
                }
            </div>
        </div>
        </>
    )
}

export default Register;
