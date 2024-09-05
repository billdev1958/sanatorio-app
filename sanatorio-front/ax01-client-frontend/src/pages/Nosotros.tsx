import worldlogo from "../assets/logoworld.svg"
import logo from "../assets/logo-ax01-500px.svg"
import openLogo from "../assets/opensource.webp"


function Nosotros() {
    return (

        <div className="header-nosotros">
            <div className="header-content-nosotros">

                <div className="text-container-nosotros">
                    <h1 >Somos AX01 una empresa mexicana</h1>
                    <p >Desarrollamos software veloz y de calidad donde nuestra prioridad es desarrollar software eficiente, escalable y seguro
                        tambien hacemos consultoria para ayudarte a encontrar la mejor manera de evolucionar tu negocio, nos inspiramos en el axolote mexicano
                        ya que creemos que el buen software es el que dura en el tiempo y aunque le quiten una parte puede adaptarse al cambio y regenerarse
                        mediante la facilidad de actualizarse a pesar de la antiguedad que tenga nuestro software, a eso nos comprometemos.
                    </p>
                </div>
                <img src={logo} className="header-logo-nosotros" alt="logo" />
            </div>

            <div className="header-content-nosotros">

                <div className="text-container-nosotros">
                    <h1 >Queremos ayudar al mundo con software</h1>
                    <p >
                        Creemos que la investigacion cientifica en nuestra era sera vital para poder ayudar a nuestro planeta y algo que nos preocupa es la poca atencion
                        que reciben los cientificos y para esto queremos aportar nuestro grano de arena apoyando a estos cientificos en lo que necesiten y que este en 
                        nuestras manos
                    </p>
                </div>
                <img src={worldlogo} className="header-logo-world " alt="logo" />
            </div>

            <div className="header-content-nosotros">

                <div className="text-container-nosotros">
                    <h1 >Opensource</h1>
                    <p >
                        Otra inspiracion ademas de el axolotl y ayudar al planeta es el codigo abierto, pensamos que la verdadera inovacion esta en el conjunto de
                        cientos de personas que siempre quieren mejorar algo y llegamos a la conclusion de que esto nos define como empresa y aunque vendamos software
                        apoyaremos como sea a la comunidad opensource.
                    </p>
                </div>
                <img src={openLogo} className="header-logo-nosotros " alt="logo" />
            </div>

        </div>
    )
}

export default Nosotros