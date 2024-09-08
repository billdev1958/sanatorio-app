import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import UserCard from '../components/MenuCard';

const ControlMenu = () => {
  const navigate = useNavigate(); // Hook para la navegación
  const [showUserOptions, setShowUserOptions] = useState(false); // Estado para controlar el submenú

  const handleEnter = (section: string) => {
    if (section === 'Control de Usuarios') {
      setShowUserOptions(!showUserOptions); // Alternar el estado para mostrar u ocultar opciones
    } else if (section === 'Control de Citas') {
      navigate('/appointments'); // Redirigir a la página de citas
    }
  };

  const handleUserOption = (role: string) => {
    if (role === 'SuperUsers') {
      navigate('/user'); // Redirigir a SuperUsuarios
    } else if (role === 'Doctors') {
      navigate('/user/doctors'); // Redirigir a Doctores
    } else if (role === 'Patients') {
      navigate('/user/patients'); // Redirigir a Pacientes
    }
  };

  return (
    <div className="controlMenuContainer">
      <UserCard
        title="Control de Usuarios"
        subtitle="Usuarios registrados"
        counter={34}
        showOptions={showUserOptions} // Pasar el estado del submenú al UserCard
        onEnter={() => handleEnter('Control de Usuarios')}
        onUserOptionSelect={handleUserOption} // Pasar la función para manejar las opciones
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
