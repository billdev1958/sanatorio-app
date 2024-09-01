import '../styles/PaginaRegistro.css'
import Alert from '../components/alertAutorizacion';

export default function PRSUser(){
    return(
        <>        
        
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
                <Alert/>                         
            </form>
        </div>
        
            
        </>
    );
}