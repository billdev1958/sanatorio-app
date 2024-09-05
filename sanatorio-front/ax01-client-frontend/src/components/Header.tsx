import logo from "../assets/ax01.png"
const scrollDown = () => {
    // Obtiene el ancho de la ventana del navegador
    const width = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth;

    // Determina el valor de 'top' basado en el tamaño de la pantalla
    const scrollToValue = width > 768 ? 600 : 450; // 600 para pantallas grandes, 450 para pequeñas

    window.scrollTo({
        top: scrollToValue,
        behavior: "smooth"
    });
}

const Header = () => {

    return (
        <header className="header">
            <div className="header-content">

                <div className="text-container">
                    <h1 className="header-title">Soluciones TI a la medida</h1>
                    <h2 className="header-subtitle-h2">Desarrollamos software veloz y de calidad</h2>
                    <h2 className="header-subtitle-h2">Somos una consultora que busca depurar el software lento en mexico y en mas lados </h2>
                    <button className="header-button" onClick={scrollDown}>Ver mas</button>
                </div>
                <img src={logo} className="header-logo" alt="logo" />
            </div>
        </header>
    );
}

export default Header;