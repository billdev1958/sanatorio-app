export interface LoginUser {
    email: string;
    password: string;
  }
  
  export interface RegisterPatientRequest {
    dependency_id: number;
    name: string;
    lastname1: string;
    lastname2: string;
    curp: string;
    sex: string;
    phone: string;
    email: string;
    password: string;
  }
  
  export interface RegisterBeneficiaryRequest {
    name: string;
    lastname1: string;
    lastname2: string;
    curp: string;
    sex: string;
  }
  
