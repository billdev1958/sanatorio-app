type UserCardProps = {
  title: string;
  subtitle: string;
  counter: number;
  showOptions?: boolean; // Nueva propiedad para controlar si se muestra el submenú
  onEnter: () => void;
  onUserOptionSelect?: (role: string) => void; // Propiedad para manejar la selección de roles
};

const UserCard = ({ title, subtitle, counter, showOptions = false, onEnter, onUserOptionSelect }: UserCardProps) => {
  return (
    <div className="appleCard">
      <h2 className="appleCardTitle">{title}</h2>
      <p className="appleCardSubtitle">{subtitle}</p>
      <p className="appleCardCounter">Total: {counter}</p>

      <div className="appleCardActions">
        <button className="appleCardButton" onClick={onEnter}>
          Ingresar
        </button>
      </div>

      {/* Mostrar las opciones dentro de la tarjeta si showOptions es true */}
      {showOptions && (
        <div className="userOptionsMenu">
          <button className="userOptionButton" onClick={() => onUserOptionSelect?.('SuperUsers')}>
            SuperUsuarios
          </button>
          <button className="userOptionButton" onClick={() => onUserOptionSelect?.('Doctors')}>
            Doctores
          </button>
          <button className="userOptionButton" onClick={() => onUserOptionSelect?.('Patients')}>
            Pacientes
          </button>
        </div>
      )}
    </div>
  );
};

export default UserCard;
