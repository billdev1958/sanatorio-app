import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import UserCard from '../components/UserCard'; // Componente reutilizable
import FilterNav from '../components/FilterNav';
import { Users, Filters } from '../models.tsx/users'; // Importamos solo Users
import { getUsers } from '../services/getUsers';

function UsersMainPage() {
  const [users, setUsers] = useState<Users[]>([]); // Solo aceptamos Users
  const [filteredUsers, setFilteredUsers] = useState<Users[]>([]);
  const [showAddMenu, setShowAddMenu] = useState(false);
  const token = localStorage.getItem('token');
  const navigate = useNavigate();

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        if (!token) {
          throw new Error('Token not found');
        }
        const data = await getUsers(token);
        const normalUsers = data.filter((user: Users) => user.role !== 2); // Excluir doctores
        setUsers(normalUsers);
        setFilteredUsers(normalUsers); // Inicializar con todos los usuarios normales
      } catch (error) {
        console.error("Error al cargar los usuarios:", error);
      }
    };

    fetchUsers();
  }, [token]);

  const handleFilterChange = (filters: Filters) => {
    let filtered = [...users];

    // Filtrar por nombre
    if (filters.name?.trim()) {
      filtered = filtered.filter(user =>
        `${user.name} ${user.lastname1} ${user.lastname2}`
          .toLowerCase()
          .includes(filters.name!.toLowerCase())
      );
    }

    // Filtrar por fecha
    if (filters.date) {
      filtered = filtered.filter(user => new Date(user.created_at).toISOString().slice(0, 10) === filters.date);
    }

    setFilteredUsers(filtered);
  };

  const toggleAddMenu = () => {
    setShowAddMenu(!showAddMenu);
  };

  const goToAddPatient = () => navigate('/register/user');

  return (
    <div className="usersMainPage">
      <h1 className="pageTitle">Usuarios</h1>

      <div className='filters'>
        <FilterNav onFilterChange={handleFilterChange} />
      </div>

      <div className="userList">
        {filteredUsers.length > 0 ? (
          filteredUsers.map(user => (
            <UserCard
              key={user.id} // Aquí agregamos la propiedad `key` única
              user={user}
            />
          ))
        ) : (
          <p className="noUsers">No se encontraron usuarios o no hay usuarios disponibles.</p>
        )}
      </div>

      <div className="addButtonContainer">
        <button className="addButton" onClick={toggleAddMenu}>
          +
        </button>

        {showAddMenu && (
          <div className="addMenu">
            <button className="addMenuItem" onClick={goToAddPatient}>
              Agregar Paciente o SuperUsuario
            </button>
          </div>
        )}
      </div>
    </div>
  );
}

export default UsersMainPage;
