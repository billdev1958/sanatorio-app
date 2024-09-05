import { useNavigate } from 'react-router-dom';
import UserCard from '../components/MenuCard';

const ControlMenu = () => {
  const navigate = useNavigate(); // Hook para la navegación

  const handleEnter = (section: string) => {
    if (section === 'Control de Usuarios') {
      navigate('/user'); // Redirigir a la página de usuarios
    } else if (section === 'Control de Citas') {
      navigate('/appointments'); // Redirigir a la página de citas (puedes cambiar la ruta si es otra)
    }
  };

  return (
    <div className="controlMenuContainer">
      <UserCard
        title="Control de Usuarios"
        subtitle="Usuarios registrados"
        counter={34}
        onEnter={() => handleEnter('Control de Usuarios')}
      />
      <UserCard
        title="Control de Citas"
        subtitle="Citas agendadas"
        counter={15}
        onEnter={() => handleEnter('Control de Citas')}
      />
    </div>
  );
};

export default ControlMenu;
