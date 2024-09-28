
import Calendario from '../components/Calendario';
import CardInfo from '../components/CardInfo';

const HomePage = () => {
    const curp = "ABCD123456EF789GH1";  // CURP aleatorio generado
    const accountNumber = "d79b9797-bbfd-4c50-8db8-e7f3c8b1d5f1";  // UUID aleatorio generado
  
    return (
      <>
        <CardInfo 
          nombre="Billy Rivera Salinas" 
          curp={curp} 
          telefono="555-123-4567" 
          cuenta={accountNumber}
        />
        <Calendario/>
      </>
    );
  };

export default HomePage;
