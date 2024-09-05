
type UserCardProps = {
  title: string;
  subtitle: string; // Agregar subtítulo
  counter: number;
  onEnter: () => void;
};

const UserCard = ({ title, subtitle, counter, onEnter }: UserCardProps) => {
  return (
    <div className="appleCard">
      <h2 className="appleCardTitle">{title}</h2>
      <p className="appleCardSubtitle">{subtitle}</p> {/* Subtítulo */}
      <p className="appleCardCounter">Total: {counter}</p>

      <div className="appleCardActions">
        <button className="appleCardButton" onClick={onEnter}>
          Ingresar
        </button>
      </div>
    </div>
  );
};

export default UserCard;
