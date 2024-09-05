export interface Users {
    ID: number;
    Name: string;
    Lastname1: string;
    Lastname2: string;
    Email: string;
    Rol: number; // 1 = SuperUser, 3 = Paciente
    Curp?: string; // Solo para usuarios que no sean doctores
    Created_At: string; // Fecha de creación del usuario
    AccountID: string;
}

export interface Doctor {
    ID: number;
    Name: string;
    Lastname1: string;
    Lastname2: string;
    Email: string;
    Rol: number; // 2 = Doctor
    MedicalLicense: string; // Licencia médica
    Specialty: number; // Especialidad (1 = Cardiologo, 2 = Dermatologo, etc.)
    AccountID: string;
    Created_At: string;
}

export interface Filters {
    name?: string;  // Filtrar por nombre o apellido
    role?: number;  // Filtrar por número de rol
    date?: string;  // Filtrar por fecha de creación
  }