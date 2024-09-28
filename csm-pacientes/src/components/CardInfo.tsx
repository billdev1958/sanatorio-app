import logo2 from '../assets/logo2.png'; // Asegúrate de que la ruta al logo sea correcta

const CardInfo = (props: { nombre: string; curp: string; telefono: string; cuenta: string }) => {
  return (
    <div class="card-info">
      <img src={logo2} alt="Logo" class="card-logo" />
      <div class="card-details">
        <h2>{props.nombre}</h2>
        <p>Curp: {props.curp}</p>
        <p>Teléfono: {props.telefono}</p>
        <p>Número de cuenta: {props.cuenta}</p>
      </div>
    </div>
  );
};

export default CardInfo;
