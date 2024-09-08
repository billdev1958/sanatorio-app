import Card from './Card';
import { Users } from '../models.tsx/users';

type UserCardProps = {
  user: Users;
};

const UserCard = ({ user }: UserCardProps) => {
  return (
    <Card
      title={`${user.name} ${user.lastname1} ${user.lastname2}`}
      email={user.email}
      id={user.id}
      role={user.role}
      onDelete={handleDelete}
      extraInfo={<p className="cardCurp">CURP: {user.curp || 'N/A'}</p>}
    />
  );
};

const handleDelete = (id: number) => {
  console.log(`Eliminar usuario con ID: ${id}`);
};

export default UserCard;
