CREATE TABLE IF NOT EXISTS cat_rol (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS account (
    id UUID PRIMARY KEY,
    user_id INT NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    rol INT NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    account_id UUID,
    name VARCHAR(255) NOT NULL,
    lastname1 VARCHAR(255) NOT NULL,
    lastname2 VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS super_user (
    account_id UUID PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS patient_user (
    account_id UUID PRIMARY KEY,
    curp VARCHAR(18)
);

-- Cites --

CREATE TABLE IF NOT EXISTS cat_specialty (
    id SERIAL PRIMARY KEY,
    name VARCHAR(60)
);

CREATE TABLE IF NOT EXISTS appointment_status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS office_status(
    id SERIAL PRIMARY KEY,
    name VARCHAR(60)
);

CREATE TABLE IF NOT EXISTS office (
    id SERIAL PRIMARY KEY,
    name VARCHAR(60),
    specialty_id INTEGER,
    status_id INTEGER,
    doctor_account_id UUID
);

CREATE TABLE IF NOT EXISTS doctor_user (
    account_id UUID PRIMARY KEY,
    id_specialty INT,
    medical_license VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS schedule (
    id SERIAL PRIMARY KEY,
    doctor_account_id UUID,
    office_id INTEGER,
    day_of_week INT NOT NULL, -- 0 = Sunday, 1 = Monday, ..., 6 = Saturday
    time_start TIME NOT NULL,
    time_end TIME NOT NULL,
    CONSTRAINT chk_office_status CHECK (
        (SELECT status_id FROM office WHERE id = office_id) = 1 -- Status 1 podría representar 'Disponible'
    )
);

CREATE TABLE IF NOT EXISTS appointment (
    id UUID PRIMARY KEY,
    doctor_account_id UUID,
    patient_account_id UUID,
    office_id INTEGER,
    time_start TIMESTAMP,
    time_end TIMESTAMP,
    schedule_id INTEGER,
    status_id INTEGER,
    CONSTRAINT chk_time_validity CHECK (time_start < time_end)
);

-- Foreign keys --

ALTER TABLE account
ADD CONSTRAINT fk_rol_account
FOREIGN KEY (rol) REFERENCES cat_rol(id);

ALTER TABLE users
ADD CONSTRAINT fk_account_id_users
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE super_user
ADD CONSTRAINT fk_account_id_super_user
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE patient
ADD CONSTRAINT fk_account_id_patient
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE office
ADD CONSTRAINT fk_specialty_office
FOREIGN KEY (specialty_id) REFERENCES cat_specialty(id);

ALTER TABLE office
ADD CONSTRAINT fk_status_office
FOREIGN KEY (status_id) REFERENCES office_status(id);

ALTER TABLE office
ADD CONSTRAINT fk_doctor_office
FOREIGN KEY (doctor_account_id) REFERENCES doctor(account_id);

ALTER TABLE doctor
ADD CONSTRAINT fk_account_id_doctor
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE doctor
ADD CONSTRAINT fk_cat_id_specialty
FOREIGN KEY (id_specialty) REFERENCES cat_specialty(id);

ALTER TABLE schedule
ADD CONSTRAINT fk_doctor_schedule
FOREIGN KEY (doctor_account_id) REFERENCES doctor(account_id);

ALTER TABLE schedule
ADD CONSTRAINT fk_office_schedule
FOREIGN KEY (office_id) REFERENCES office(id);

ALTER TABLE appointment
ADD CONSTRAINT fk_schedule_appointment
FOREIGN KEY (schedule_id) REFERENCES schedule(id);

ALTER TABLE appointment
ADD CONSTRAINT fk_doctor_appointment
FOREIGN KEY (doctor_account_id) REFERENCES doctor(account_id);

ALTER TABLE appointment
ADD CONSTRAINT fk_patient_appointment
FOREIGN KEY (patient_account_id) REFERENCES patient(account_id);

ALTER TABLE appointment
ADD CONSTRAINT fk_office_appointment
FOREIGN KEY (office_id) REFERENCES office(id);

ALTER TABLE appointment
ADD CONSTRAINT fk_status_appointment
FOREIGN KEY (status_id) REFERENCES appointment_status(id);
