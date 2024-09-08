import { Link } from 'react-router-dom';

type CardProps = {
  title: string;
  email: string;
  extraInfo?: JSX.Element;
  id: number;
  role: number;
  onDelete: (id: number) => void;
};

const Card = ({ title, email, extraInfo, id, role, onDelete }: CardProps) => {
  return (
    <div className="card">
      <h2 className="cardTitle">{title}</h2>
      <p className="cardEmail">{email}</p>

      {extraInfo}

      <div className="cardActions">
        {/* Diferentes rutas dependiendo del rol */}
        <Link to={`/${role === 2 ? 'doctors' : role === 3 ? 'patients' : 'users'}/view/${id}`} className="cardButton">
          Ver
        </Link>
        <Link to={`/${role === 2 ? 'doctor' : role === 3 ? 'patient' : 'user'}/update/${id}`} className="cardButton">
          Editar
        </Link>
        <button className="cardButton" onClick={() => onDelete(id)}>Eliminar</button>
      </div>
    </div>
  );
};

export default Card;
