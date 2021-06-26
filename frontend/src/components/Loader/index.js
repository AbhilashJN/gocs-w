import React from 'react'
import logo from '../../assets/logo.svg'
import './Loader.css'

const Loader=()=>{
    return (
        <div className="Loader-container">
            <img src={logo} alt="loading" className="Loader-logo"/>
        </div>
    )
}

export default Loader