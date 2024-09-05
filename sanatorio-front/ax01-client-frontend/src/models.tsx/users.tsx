export interface RegisterUserByAdminRequest {
    name: string;
    lastname1: string;
    lastname2: string;
    email: string;
    password: string;
    rol: number;
    curp: string;
    admin_password: string;
}
export interface RegisterDoctorByAdminRequest {
    name: string;
    lastname1: string;
    lastname2: string;
    email: string;
    password: string;
    rol: number;
    medical_license: string;
    specialty: number;
    admin_password: string;

}


export interface LoginUser {

    email: string;
    password: string;
    

}

export interface Users {
    account_id: string;
    id: number;
    name: string;
    lastname1: string;
    lastname2: string;
    email: string;
    curp: string;
    role: number
    created_at: string
  }

  export interface Doctor {
    account_id: string;
    id: number;
    name: string;
    lastname1: string;
    lastname2: string;
    email: string;
    medical_license: string;
    specialty: number;
    role: number;
    created_at: string;
}

  export interface UpdateUserRequest {
    account_id: string; // UUID en formato string
    name?: string;      // Campo opcional (puede estar vacío si no se modifica)
    lastname1?: string; // Campo opcional
    lastname2?: string; // Campo opcional
    email?: string;     // Campo opcional
    password?: string;  // Campo opcional
    curp?: string;      // Campo opcional
    admin_password: string; // Campo obligatorio para la contraseña del admin
}

  export interface Filters {
    name?: string;  // Filtrar por nombre o apellido
    role?: number;  // Filtrar por número de rol
    date?: string;  // Filtrar por fecha de creación
  }