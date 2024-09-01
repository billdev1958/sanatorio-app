import { useState, useEffect } from "react";
import { Outlet } from "react-router-dom";
import NavBar from "../components/nav-bar";
import AlertConfirmation from '../components/alertConfirmation'

export default function PaginaRegistro() {
    const [isVisible, setIsVisible] = useState(() => 
    {
        const saved = localStorage.getItem('isVisible');
        return saved !== null ? JSON.parse(saved) : true;
    });

    useEffect(() => {
        localStorage.setItem('isVisible', JSON.stringify(isVisible));
    }, [isVisible]);

    const hide = () => {
        setIsVisible(false);
    };
    const show = () => {
        setIsVisible(true);
    };
    // const [isVisible2, setIsVisible]

    return (
        <>
            <NavBar />
            {isVisible && (
                <div className="center2" id="center2">
                    <h1 className="class-h1">Bienvenido</h1>
                    <h3 className="class-h3">Por favor registre un nuevo usuario administrador para poder continuar</h3>
                    <a href={`/registro/admin`}>
                        <button className="class-button" onClick={hide}>Siguiente</button>
                    </a>
                </div>
            )}
            {/* <button onClick={show}>Mostrar Div</button> */}
            <div id="admin">
                <Outlet />
            </div>
            {isVisible && (
                <div>            
                    <form className="center" action={`/`} method="post" autoComplete="on">
                    <h1 className="class-h1" >Crea tu cuenta</h1>
                    <div></div>
                        <div className="center-left">
                            <label htmlFor="name">Nombre(s):</label><br />  
                            <label htmlFor="firstName">Primer apellido:</label><br />  
                            <label htmlFor="lastName">Segundo apellido:</label><br />  
                            <label htmlFor="email">Correo electronico:</label><br /> 
                            <label htmlFor="pass">Contraseña:</label><br /> 
                            <label htmlFor="confPass">Confirmar contraseña:</label><br />  

                        </div>
                        <div className="center-right">
                            <input type="text" id="name" name="name" /><br />     
                            <input type="text" id="firstName" name="firstName" /><br />     
                            <input type="text" id="lastName" name="lastName" /> <br />   
                            <input type="text" id="email" name="email" /><br />  
                            <input type="password" id="pass" name="pass" /><br />  
                            <input type="password" id="confPass" name="confPass" /> <br />                       

                        </div>
                        <div></div>
                        <AlertConfirmation/>                         
                    </form>
                </div>
            )}
            
        </>
    );
}