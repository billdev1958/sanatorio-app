import { Link } from 'react-router-dom';
import { Patient } from '../models.tsx/users'; 

type PatientCardProps = {
  patient: Patient; 
};

const PatientCard = ({ patient }: PatientCardProps) => {
  return (
    <div className="patientCard">
      <h2 className="patientCardName">{patient.name} {patient.lastname1} {patient.lastname2}</h2>
      <p className="patientCardEmail">{patient.email}</p>
      <p className="patientCardCurp">CURP: {patient.curp}</p>
      <div className="patientCardActions">
        <Link to={`/patients/view/${patient.id}`} className="patientCardButton">Ver</Link>
        <Link to={`/patient/update/${patient.id}`} className="patientCardButton">Editar</Link>
        <button className="patientCardButton" onClick={() => handleDelete(patient.id)}>Eliminar</button>
      </div>
    </div>
  );
};

const handleDelete = (id: number) => {
  console.log(`Eliminar paciente con ID: ${id}`);
};

export default PatientCard;
