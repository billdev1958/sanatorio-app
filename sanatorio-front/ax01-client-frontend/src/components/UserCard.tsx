import { Link } from 'react-router-dom';
import { Users, Doctor } from '../models.tsx/users'; 

type UserCardProps = {
  user: Users | Doctor; 
};

const UserCard = ({ user }: UserCardProps) => {
  return (
    <div className="userCard">
      <h2 className="userCardName">{user.name} {user.lastname1} {user.lastname2}</h2>
      <p className="userCardEmail">{user.email}</p>

      {user.role === 2 && 'medical_license' in user ? (
        <>
          <p className="userCardMedicalLicense">Licencia médica: {(user as Doctor).medical_license}</p>
          <p className="userCardSpecialty">Especialidad: {(user as Doctor).specialty}</p>
        </>
      ) : (
        <p className="userCardCurp">CURP: {(user as Users).curp || 'N/A'}</p>
      )}

      <div className="userCardActions">
        <Link to={`/users/view/${user.id}`} className="userCardButton">Ver</Link>
        <Link to={`/user/update/${user.id}`} className="userCardButton">Editar</Link>
        <button className="userCardButton" onClick={() => handleDelete(user.id)}>Eliminar</button>
      </div>
    </div>
  );
};

const handleDelete = (id: number) => {
  console.log(`Eliminar usuario con ID: ${id}`);
};

export default UserCard;
