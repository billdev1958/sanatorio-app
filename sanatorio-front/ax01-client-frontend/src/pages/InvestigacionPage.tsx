import softwarelogo from "../assets/software-logo.svg"
import ciencia from "../assets/cientifico.svg"


function InvestigacionPage() {
    return (
        <div className="investigacion-page">
            <header className="header-investigacion">
                <div className="header-content-investigacion">
                    <img src={softwarelogo} className="header-logo-investigacion" alt="logo" />
                </div>
            </header>

            <section className="intro-section">
                <div className="intro-content">
                    <h2>Proyectos de Investigación y Desarrollo</h2>
                    <p>
                    En nuestro núcleo, somos más que programadores; somos científicos de la computación dedicados a avanzar en el conocimiento y la tecnología. Reconocemos la importancia crítica de apoyar a la comunidad científica y nos motiva profundamente el deseo de contribuir al progreso colectivo. Creemos que el software no es solo una herramienta, sino un puente hacia nuevas fronteras de investigación y desarrollo.
                    </p>
                </div>
            </section>

            <div className="middle-section">
                <img src={ciencia} className="planet-image" alt="Planeta" />
            </div>

            <section className="opensource-section">
                <div className="opensource-content">
                    <h2>Proyectos de Código Abierto</h2>
                    <p>
                    En AX01, estamos profundamente comprometidos con el avance de la tecnología a través del código abierto. Reconocemos que muchos de los avances tecnológicos más significativos de nuestra era han sido impulsados por la colaboración y el intercambio de conocimiento en comunidades de código abierto. Por eso, nos entusiasma anunciar que contribuiremos activamente a proyectos de código abierto, tanto apoyando iniciativas existentes como lanzando nuestras propias soluciones a problemas comunes.                    </p>
                </div>
            </section>





        </div>
    )
}

export default InvestigacionPage;