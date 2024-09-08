import { Link } from 'react-router-dom';
import { Doctor } from '../models.tsx/users'; 

type DoctorCardProps = {
  doctor: Doctor; 
};

const DoctorCard = ({ doctor }: DoctorCardProps) => {
  return (
    <div className="doctorCard">
      <h2 className="doctorCardName">{doctor.name} {doctor.lastname1} {doctor.lastname2}</h2>
      <p className="doctorCardEmail">{doctor.email}</p>
      <p className="doctorCardMedicalLicense">Licencia médica: {doctor.medical_license}</p>
      <p className="doctorCardSpecialty">Especialidad: {doctor.specialty}</p>
      <div className="doctorCardActions">
        <Link to={`/doctors/view/${doctor.id}`} className="doctorCardButton">Ver</Link>
        <Link to={`/doctor/update/${doctor.id}`} className="doctorCardButton">Editar</Link>
        <button className="doctorCardButton" onClick={() => handleDelete(doctor.id)}>Eliminar</button>
      </div>
    </div>
  );
};

const handleDelete = (id: number) => {
  console.log(`Eliminar doctor con ID: ${id}`);
};

export default DoctorCard;
