import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import PatientCard from '../components/PatientCard'; // Componente reutilizable para pacientes
import FilterNav from '../components/FilterNav';
import { Patient, Filters } from '../models.tsx/users'; // Importamos solo Pacientes
import { getPatients } from '../services/getPatients'; // Servicio para obtener los pacientes

function PatientsMainPage() {
  const [patients, setPatients] = useState<Patient[]>([]); // Estado para almacenar la lista de pacientes
  const [filteredPatients, setFilteredPatients] = useState<Patient[]>([]);
  const [showAddMenu, setShowAddMenu] = useState(false); // Control para mostrar el menú de agregar pacientes
  const token = localStorage.getItem('token');
  const navigate = useNavigate();

  useEffect(() => {
    const fetchPatients = async () => {
      try {
        if (!token) {
          throw new Error('Token not found');
        }
        const data = await getPatients(token); // Obtenemos la lista de pacientes desde el servicio
        setPatients(data);
        setFilteredPatients(data); // Inicializamos el estado con todos los pacientes
      } catch (error) {
        console.error("Error al cargar los pacientes:", error);
      }
    };

    fetchPatients();
  }, [token]);

  const handleFilterChange = (filters: Filters) => {
    let filtered = [...patients];

    // Filtrar por nombre
    if (filters.name?.trim()) {
      filtered = filtered.filter(patient =>
        `${patient.name} ${patient.lastname1} ${patient.lastname2}`
          .toLowerCase()
          .includes(filters.name!.toLowerCase())
      );
    }

    // Filtrar por fecha
    if (filters.date) {
      filtered = filtered.filter(patient => new Date(patient.created_at).toISOString().slice(0, 10) === filters.date);
    }

    setFilteredPatients(filtered);
  };

  const toggleAddMenu = () => {
    setShowAddMenu(!showAddMenu); // Alterna el menú para agregar pacientes
  };

  const goToAddPatient = () => navigate('/register/patient'); // Navegar a la página de registro de pacientes

  return (
    <div className="patientsMainPage">
      <h1 className="pageTitle">Pacientes</h1>

      <div className='filters'>
        <FilterNav onFilterChange={handleFilterChange} />
      </div>

      <div className="patientList">
        {filteredPatients.length > 0 ? (
          filteredPatients.map(patient => (
            <PatientCard
              key={patient.id} // Agregamos una clave única para cada paciente
              patient={patient}
            />
          ))
        ) : (
          <p className="noPatients">No se encontraron pacientes o no hay pacientes disponibles.</p>
        )}
      </div>

      <div className="addButtonContainer">
        <button className="addButton" onClick={toggleAddMenu}>
          +
        </button>

        {showAddMenu && (
          <div className="addMenu">
            <button className="addMenuItem" onClick={goToAddPatient}>
              Agregar Paciente
            </button>
          </div>
        )}
      </div>
    </div>
  );
}

export default PatientsMainPage;
