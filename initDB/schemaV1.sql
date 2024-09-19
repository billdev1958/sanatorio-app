CREATE TABLE IF NOT EXISTS role (
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
    role_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    password_change_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS super_user (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name1 VARCHAR(255) NOT NULL,
    last_name2 VARCHAR(255) NOT NULL,
    account_id UUID NOT NULL,
    curp VARCHAR(18) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS patient (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name1 VARCHAR(255) NOT NULL,
    last_name2 VARCHAR(255) NOT NULL,
    record_id UUID UNIQUE,
    account_id UUID NOT NULL,
    curp VARCHAR(18) NOT NULL,
    sex CHAR(1) NOT NULL,
    phone VARCHAR(10),
    address VARCHAR(50),
    occupation VARCHAR(20),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS medical_history (
    id SERIAL PRIMARY KEY,
    patient_id INT REFERENCES patient(id) ON DELETE CASCADE,
    family_history TEXT,
    personal_pathological_history TEXT,
    personal_non_pathological_history TEXT,
    immunizations TEXT,
    allergies TEXT
);

CREATE TABLE IF NOT EXISTS evolution_note (
    id SERIAL PRIMARY KEY,
    consultation_id INT REFERENCES consultation(id) ON DELETE CASCADE,
    follow_up_notes TEXT,
    condition_changes TEXT,
    diagnosis_evolution TEXT,
    treatment_adjustments TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS incapacity (
    -- Define incapacity fields here
);

CREATE TABLE IF NOT EXISTS specialty (
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

CREATE TABLE IF NOT EXISTS office_status (
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
    doctor_account_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS doctor (
    id SERIAL PRIMARY KEY,
    account_id UUID NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name1 VARCHAR(255) NOT NULL,
    last_name2 VARCHAR(255) NOT NULL,
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
    deleted_at TIMESTAMP
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

CREATE TABLE IF NOT EXISTS consultation (
    id SERIAL PRIMARY KEY,
    patient_id UUID NOT NULL,
    date DATE NOT NULL,
    time TIME NOT NULL,
    reason TEXT NOT NULL,
    symptoms TEXT NOT NULL,
    doctor_notes TEXT,
    requested_tests TEXT,
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS attachment (
    id SERIAL PRIMARY KEY,
    consultation_id INT REFERENCES consultation(id) ON DELETE CASCADE,
    patient_id INT REFERENCES patient(id) ON DELETE CASCADE,
    file_path VARCHAR NOT NULL,
    file_name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Foreign keys --

ALTER TABLE account
ADD CONSTRAINT fk_role_account
FOREIGN KEY (role_id) REFERENCES role(id);

ALTER TABLE super_user
ADD CONSTRAINT fk_account_super_user
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE patient
ADD CONSTRAINT fk_account_patient
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE office
ADD CONSTRAINT fk_specialty_office
FOREIGN KEY (specialty_id) REFERENCES specialty(id);

ALTER TABLE office
ADD CONSTRAINT fk_status_office
FOREIGN KEY (status_id) REFERENCES office_status(id);

ALTER TABLE office
ADD CONSTRAINT fk_doctor_office
FOREIGN KEY (doctor_account_id) REFERENCES doctor(account_id);

ALTER TABLE doctor
ADD CONSTRAINT fk_account_doctor
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE doctor
ADD CONSTRAINT fk_specialty_doctor
FOREIGN KEY (specialty_id) REFERENCES specialty(id);

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

-- Validation function --

CREATE OR REPLACE FUNCTION validate_office_status()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT status_id FROM office WHERE id = NEW.office_id) != 1 THEN
        RAISE EXCEPTION 'The office is not available for scheduling';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to validate office status before inserting or updating schedule --

CREATE TRIGGER trigger_validate_office_status
BEFORE INSERT OR UPDATE ON schedule
FOR EACH ROW EXECUTE FUNCTION validate_office_status();
