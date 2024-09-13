CREATE TABLE IF NOT EXISTS cat_rol (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS account (
    id UUID PRIMARY KEY,
    user_id INT NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    rol INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    password_change_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    lastname1 VARCHAR(255) NOT NULL,
    lastname2 VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS super_user (
    id SERIAL PRIMARY KEY,
    account_id UUID NOT NULL,
    curp VARCHAR(18) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS patient (
    id SERIAL PRIMARY KEY,
    account_id UUID NOT NULL,
    curp VARCHAR(18) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS cat_specialty (
    id SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS appointment_status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS office_status(
    id SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS office (
    id SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    specialty_id INTEGER NOT NULL,
    status_id INTEGER NOT NULL,
    doctor_account_id UUID,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS doctor(
    id SERIAL PRIMARY KEY,
    account_id UUID NOT NULL,
    specialty_id INT NOT NULL,
    medical_license VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS schedule (
    id SERIAL PRIMARY KEY,
    office_id INTEGER NOT NULL, 
    day_of_week INT NOT NULL,
    time_start TIME NOT NULL,
    time_end TIME NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT unique_office_schedule UNIQUE (office_id, day_of_week, time_start, time_end)
);

CREATE TABLE IF NOT EXISTS appointment (
    id UUID PRIMARY KEY,
    doctor_account_id UUID NOT NULL,
    patient_account_id UUID NOT NULL,
    office_id INTEGER NOT NULL,
    time_start TIMESTAMP NOT NULL,
    time_end TIMESTAMP NOT NULL,
    schedule_id INTEGER NOT NULL,
    status_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT chk_time_validity CHECK (time_start < time_end)
);
-- Foreign keys --

ALTER TABLE account
ADD CONSTRAINT fk_rol_account
FOREIGN KEY (rol) REFERENCES cat_rol(id);

ALTER TABLE users
ADD CONSTRAINT fk_account_id_users
FOREIGN KEY (user_id) REFERENCES account(id);

ALTER TABLE super_user
ADD CONSTRAINT fk_account_id_super_user
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE patient_user
ADD CONSTRAINT fk_account_id_patient_user
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE office
ADD CONSTRAINT fk_specialty_office
FOREIGN KEY (specialty_id) REFERENCES cat_specialty(id);

ALTER TABLE office
ADD CONSTRAINT fk_status_office
FOREIGN KEY (status_id) REFERENCES office_status(id);

ALTER TABLE office
ADD CONSTRAINT fk_doctor_office
FOREIGN KEY (doctor_account_id) REFERENCES doctor_user(account_id);

ALTER TABLE doctor_user
ADD CONSTRAINT fk_account_id_doctor_user
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE doctor_user
ADD CONSTRAINT fk_cat_id_specialty
FOREIGN KEY (id_specialty) REFERENCES cat_specialty(id);

ALTER TABLE schedule
ADD CONSTRAINT fk_doctor_schedule
FOREIGN KEY (doctor_account_id) REFERENCES doctor_user(account_id);

ALTER TABLE schedule
ADD CONSTRAINT fk_office_schedule
FOREIGN KEY (office_id) REFERENCES office(id);

ALTER TABLE appointment
ADD CONSTRAINT fk_schedule_appointment
FOREIGN KEY (schedule_id) REFERENCES schedule(id);

ALTER TABLE appointment
ADD CONSTRAINT fk_doctor_appointment
FOREIGN KEY (doctor_account_id) REFERENCES doctor_user(account_id);

ALTER TABLE appointment
ADD CONSTRAINT fk_patient_appointment
FOREIGN KEY (patient_account_id) REFERENCES patient_user(account_id);

ALTER TABLE appointment
ADD CONSTRAINT fk_office_appointment
FOREIGN KEY (office_id) REFERENCES office(id);

ALTER TABLE appointment
ADD CONSTRAINT fk_status_appointment
FOREIGN KEY (status_id) REFERENCES appointment_status(id);

-- Función de validación --

CREATE OR REPLACE FUNCTION validate_office_status() 
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT status_id FROM office WHERE id = NEW.office_id) != 1 THEN
        RAISE EXCEPTION 'The office is not available for scheduling';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger para validar el estado de la oficina antes de insertar o actualizar en la tabla schedule --

CREATE TRIGGER trigger_validate_office_status
BEFORE INSERT OR UPDATE ON schedule
FOR EACH ROW EXECUTE FUNCTION validate_office_status();
