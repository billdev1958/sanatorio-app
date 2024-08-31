import logo from '../assets/logo.jpeg'
import { useState, useEffect } from "react";

export default function navBar(){
    const [isVisible, setIsVisible] = useState(() => 
    {
        const saved = localStorage.getItem('isVisible');
        return saved !== null ? JSON.parse(saved) : true;
    });

    useEffect(() => {
        localStorage.setItem('isVisible', JSON.stringify(isVisible));
    }, [isVisible]);
    const show = () => {
        setIsVisible(true);
    };
    return(
        <>
        <div className='nav-bar'>
            <div></div>
            <div className='logo'>
                <img src={logo} alt=""/>    
                <a href={`/`}><button onClick={show}><h3 className='class-h3'>Sanatorio App</h3></button></a>
            </div>            
            <div></div>
            <div></div>
            <div>
                <a href={`/`}><button id='login'><h3 className='class-h3'>Regresar</h3></button></a>
            </div>
        </div>
    </>
    );
}