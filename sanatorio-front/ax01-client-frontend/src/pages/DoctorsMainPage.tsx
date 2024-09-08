import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import DoctorCard from '../components/DoctorCard'; // Componente reutilizable para doctores
import FilterNav from '../components/FilterNav';
import { Doctor, Filters } from '../models.tsx/users'; // Importamos solo Doctor
import { getDoctors } from '../services/getDoctors'; // Servicio para obtener los doctores

function DoctorsMainPage() {
  const [doctors, setDoctors] = useState<Doctor[]>([]); // Solo aceptamos doctores
  const [filteredDoctors, setFilteredDoctors] = useState<Doctor[]>([]);
  const [showAddMenu, setShowAddMenu] = useState(false);
  const token = localStorage.getItem('token');
  const navigate = useNavigate();

  useEffect(() => {
    const fetchDoctors = async () => {
      try {
        if (!token) {
          throw new Error('Token not found');
        }
        const data = await getDoctors(token); // Obtenemos los doctores del servicio
        setDoctors(data);
        setFilteredDoctors(data); // Inicializamos con todos los doctores
      } catch (error) {
        console.error("Error al cargar los doctores:", error);
      }
    };

    fetchDoctors();
  }, [token]);

  const handleFilterChange = (filters: Filters) => {
    let filtered = [...doctors];

    // Filtrar por nombre
    if (filters.name?.trim()) {
      filtered = filtered.filter(doctor =>
        `${doctor.name} ${doctor.lastname1} ${doctor.lastname2}`
          .toLowerCase()
          .includes(filters.name!.toLowerCase())
      );
    }

    // Filtrar por fecha
    if (filters.date) {
      filtered = filtered.filter(doctor => new Date(doctor.created_at).toISOString().slice(0, 10) === filters.date);
    }

    setFilteredDoctors(filtered);
  };

  const toggleAddMenu = () => {
    setShowAddMenu(!showAddMenu);
  };

  const goToAddDoctor = () => navigate('/register/doctor'); // Navegar a la página de registro de doctores

  return (
    <div className="doctorsMainPage">
      <h1 className="pageTitle">Doctores</h1>

      <div className='filters'>
        <FilterNav onFilterChange={handleFilterChange} />
      </div>

      <div className="doctorList">
        {filteredDoctors.length > 0 ? (
          filteredDoctors.map(doctor => (
            <DoctorCard
              key={doctor.id} // Aquí agregamos la propiedad `key` única
              doctor={doctor}
            />
          ))
        ) : (
          <p className="noDoctors">No se encontraron doctores o no hay doctores disponibles.</p>
        )}
      </div>

      <div className="addButtonContainer">
        <button className="addButton" onClick={toggleAddMenu}>
          +
        </button>

        {showAddMenu && (
          <div className="addMenu">
            <button className="addMenuItem" onClick={goToAddDoctor}>
              Agregar Doctor
            </button>
          </div>
        )}
      </div>
    </div>
  );
}

export default DoctorsMainPage;
