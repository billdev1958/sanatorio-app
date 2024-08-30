import logo from '../assets/logo.jpeg'
import imgPrin from '../assets/imgPrin.jpg'
import '../styles/PaginaPrincipal.css'



export default function PaginaPrincipal() {
    return (
      <>
        <div className='nav-bar' id='mainPage'>
                <div></div> 
                <div className='logo'>
                    <img src={logo} alt=""/>    
                    <button><h3 className='class-h3'>Sanatorio App</h3></button>
                </div>
                <div></div>
                <a href={``}><button id='login'><h3 className='class-h3'>Inicio de seccion</h3></button></a>
                <a href={`/registro`}><button id='registro'><h3 className='class-h3'>Registrar</h3></button></a>                
                <div></div>
                
            </div>
            
            <div className='body'>
                <img src={imgPrin} alt="" />
                <div className='texto'>
                    <h1 className='class-h1'>Atencion Medica</h1>
                    <p>Hacemos que sea fácil recibir atención médica.</p>
                    
                </div>
            </div>
      </>
    );
  }