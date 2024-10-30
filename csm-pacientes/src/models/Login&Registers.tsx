export interface LoginUser {
    email: string;
    password: string;
  }
  
export interface RegisterUser {
    AfiliationID: number;
    Name: string;
    Lastname1: string;
    Lastname2: string;
    Curp: string;
    Sex: string;
    PhoneNumber: string;
    Email: string;
    Password: string;
}

export interface RegisterBeneficiaryRequest {
  name: string;
  lastname1: string;
  lastname2: string;
  curp: string;
  sex: string;
}