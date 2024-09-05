import linuxLogo from "../assets/linuxLogo.png"
import cloudLogo from "../assets/google-cloud.svg"
import openLogo from "../assets/opensource.webp"
import databaselogo from "../assets/databaselogo.png"



const SectionUtils = () => {

    return (
        <div className="Section-Utils">
            <div className="Section-Container">
                <div className="H2-Container">
                    <h2>Desarrollamos para tu empresa y para la comunidad opensource </h2>

                </div>
                <div className="Logos-Container">
                    <img className="section-logo" src={linuxLogo} alt="Linux Logo"></img>
                    <img className="section-logo" src={cloudLogo} alt="Google Cloud Logo"></img>
                    <img className="section-logo" src={openLogo} alt="Open Source Logo"></img>
                    <img className="section-logo" src={databaselogo} alt="Database Logo"></img>
                </div>
            </div>
        </div>
    )

}

export default SectionUtils;