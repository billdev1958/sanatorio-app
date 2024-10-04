-- Tabla de roles
CREATE TABLE IF NOT EXISTS role (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tabla de permisos
CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(60) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tabla intermedia para rol y permiso (relación muchos a muchos)
CREATE TABLE IF NOT EXISTS role_permission(
    id_role INTEGER NOT NULL,
    id_permission INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id_role, id_permission) -- Clave primaria compuesta
);

-- Tabla para asociar roles a usuarios
CREATE TABLE IF NOT EXISTS user_roles(
    id SERIAL PRIMARY KEY, 
    account_id UUID NOT NULL, -- Relación con la cuenta/usuario
    role_id INTEGER NOT NULL,    -- Relación con roles
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE (account_id, role_id) -- Clave única para evitar duplicados
);

CREATE TABLE IF NOT EXISTS account (
    id UUID PRIMARY KEY,
    telefono VARCHAR(10) NOT NULL,
    email VARCHAR(75) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    password_change_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS super_user (
    id UUID PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name1 VARCHAR(50) NOT NULL,
    last_name2 VARCHAR(50) NOT NULL,
    account_id UUID NOT NULL,
    curp CHAR(18) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS patient (
    id UUID PRIMARY KEY,
    medical_history_id VARCHAR(12) NOT NULL, -- id de la tabla medical_history
    legacy_id INTEGER,
    account_id UUID NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name1 VARCHAR(50) NOT NULL,
    last_name2 VARCHAR(50) NOT NULL,
    curp CHAR(18) NOT NULL,
    sex CHAR(1) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

--beneficiarios
CREATE TABLE IF NOT EXISTS beneficiary (
    id UUID PRIMARY KEY,
    account_holder UUID NOT NULL, -- id de la cuenta principal
    medical_history_id VARCHAR(12) NOT NULL, -- id de la tabla medical_history
    first_name VARCHAR(50) NOT NULL,
    last_name1 VARCHAR(50) NOT NULL,
    last_name2 VARCHAR(50) NOT NULL,
    sex CHAR(1) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS medical_history_relation (
    id UUID PRIMARY KEY,
    medical_history_id VARCHAR(12) NOT NULL,
    patient_id UUID,
    beneficiary_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    CHECK (patient_id IS NOT NULL OR beneficiary_id IS NOT NULL)
);


CREATE TABLE IF NOT EXISTS medical_history (
    -- Para un futuras probabilidades de escalabilidad el identificador de expediente cambiara a UUID por ahora sera el personalizado
    id VARCHAR(12) NOT NULL,
    date_of_record DATE NOT NULL, -- 'Fecha'
    time_of_record TIME NOT NULL, -- 'Hora'
    patient_name VARCHAR(50) NOT NULL, -- 'Nombre'
    curp CHAR(18) NOT NULL, -- 'CURP'
    birth_date DATE NOT NULL, -- 'Fecha nacimiento'
    age VARCHAR(3) NOT NULL, -- 'Edad'
    gender VARCHAR(10) NOT NULL, -- 'Sexo'
    place_of_origin VARCHAR(10) NOT NULL, -- 'Procedencia'
    ethnic_group VARCHAR(20) NOT NULL, -- 'Grupo étnico'
    phone_number VARCHAR(10) NOT NULL, -- 'Teléfono'
    address VARCHAR(50) NOT NULL, -- 'Domicilio'
    occupation VARCHAR(20) NOT NULL, -- 'Ocupación'
    guardian_name VARCHAR(50), -- 'Nombre tutor'
    family_medical_history VARCHAR(100), -- 'Antecedentes heredofamiliares'
    non_pathological_history VARCHAR(100), -- 'Antecedentes personales no patológicos'
    pathological_history VARCHAR(100), -- 'Antecedentes personales patológicos'
    gynec_obstetric_history VARCHAR(100), -- 'Antecendentes gineco-obstétricos'
    current_condition VARCHAR(100), -- 'Padecimiento actual'
    cardiovascular VARCHAR(100), -- 'Cardiovascular'
    respiratory VARCHAR(100), -- 'Respiratorio'
    gastrointestinal VARCHAR(100), -- 'Gastrointestinal'
    genitourinary VARCHAR(100), -- 'Genitourinario'
    hematic_lymphatic VARCHAR(100), -- 'Hemático linfático'
    endocrine VARCHAR(100), -- 'Endocrino'
    nervous_system VARCHAR(100), -- 'Nervioso'
    musculoskeletal VARCHAR(100), -- 'Musculo esquelético'
    skin VARCHAR(100), -- 'Piel'
    body_temperature VARCHAR(10) NOT NULL, -- 'Temperatura'
    weight VARCHAR(5) NOT NULL, -- 'Peso'
    height VARCHAR(10) NOT NULL, -- 'Talla'
    bmi VARCHAR(10) NOT NULL, -- 'IMC'
    heart_rate VARCHAR(10) NOT NULL, -- 'Frecuencia cardiaca'
    respiratory_rate VARCHAR(10) NOT NULL, -- 'Frecuencia respiratoria'
    blood_pressure VARCHAR(10) NOT NULL, -- 'T/A'
    physical VARCHAR(100), -- 'Habitus exterior'
    head VARCHAR(100), -- 'Cabeza'
    neck_and_chest VARCHAR(100), -- 'Cuello tórax'
    abdomen VARCHAR(100), -- 'Abdomen'
    genital VARCHAR(100), -- 'Genitales'
    extremities VARCHAR(100), -- 'Extremidades'
    previous_results VARCHAR(100) NOT NULL, -- 'Resultados previos y actuales'
    diagnoses VARCHAR(100) NOT NULL, -- 'Diagnósticos o problemas'
    pharmacological_treatment VARCHAR(100) NOT NULL, -- 'Tratamiento farmacológico'
    prognosis VARCHAR(100) NOT NULL, -- 'Pronóstico'
    doctor_name VARCHAR(50) NOT NULL, -- 'Nombre médico'
    medical_license VARCHAR(10) NOT NULL, -- 'Cédula profesional'
    specialty_license VARCHAR(10) NOT NULL -- 'Cédula especialidad'
);

CREATE TABLE IF NOT EXISTS evolution_note (
    folio SERIAL PRIMARY KEY,
    fecha DATE NOT NULL,
    nombre VARCHAR(50) NOT NULL,
    curp CHAR(18) NOT NULL,
    dependencia VARCHAR(15),
    afiliacion VARCHAR(15),
    edad VARCHAR(3) NOT NULL,
    peso VARCHAR(6) NOT NULL,
    estatura VARCHAR(6) NOT NULL,
    fc VARCHAR(6) NOT NULL,
    fr VARCHAR(6) NOT NULL,
    ta VARCHAR(6) NOT NULL,
    temperatura VARCHAR(6) NOT NULL,
    spo2 VARCHAR(6) NOT NULL,
    glucosa VARCHAR(6) NOT NULL,
    notas VARCHAR(6) NOT NULL
);

CREATE TABLE IF NOT EXISTS incapacity (
    folio SERIAL PRIMARY KEY,
    fecha DATE NOT NULL,
    nombre VARCHAR(50) NOT NULL,
    curp CHAR(18) NOT NULL,
    dependencia VARCHAR(15),
    adscrito VARCHAR(15),
    totaldias VARCHAR(3) NOT NULL,
    inicio DATE NOT NULL,
    fin DATE NOT NULL,
    medico VARCHAR(50) NOT NULL,
    servicio VARCHAR(20) NOT NULL,
    clave VARCHAR(10) NOT NULL
);

CREATE TABLE IF NOT EXISTS services (
    id SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS appointment_status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS office_status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS office (
    id SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    specialty_id INTEGER NOT NULL,
    status_id INTEGER NOT NULL,
    doctor_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS specialty (
    id SERIAL PRIMARY KEY,
    name VARCHAR(75) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS doctor (
    id UUID PRIMARY KEY,
    account_id UUID NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name1 VARCHAR(50) NOT NULL,
    last_name2 VARCHAR(50) NOT NULL,
    specialty_id INT NOT NULL,
    medical_license VARCHAR(25) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS schedule (
    id SERIAL PRIMARY KEY,
    office_id INTEGER NOT NULL,
    day_of_week INT NOT NULL,
    time_start TIME NOT NULL,
    time_end TIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
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
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
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
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS attachment (
    id SERIAL PRIMARY KEY,
    consultation_id INT REFERENCES consultation(id) ON DELETE CASCADE,
    patient_id UUID REFERENCES patient(id) ON DELETE CASCADE,
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
FOREIGN KEY (doctor_id) REFERENCES doctor(id);

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
FOREIGN KEY (doctor_id) REFERENCES doctor(id);

ALTER TABLE appointment
ADD CONSTRAINT fk_patient_appointment
FOREIGN KEY (patient_account_id) REFERENCES patient(account_id);

ALTER TABLE appointment
ADD CONSTRAINT fk_office_appointment
FOREIGN KEY (office_id) REFERENCES office(id);

ALTER TABLE appointment
ADD CONSTRAINT fk_status_appointment
FOREIGN KEY (status_id) REFERENCES appointment_status(id);

ALTER TABLE beneficiary
ADD CONSTRAINT fk_account_holder_beneficiary
FOREIGN KEY (account_holder) REFERENCES account(id);

ALTER TABLE consultation
ADD CONSTRAINT fk_patient_consultation
FOREIGN KEY (patient_id) REFERENCES patient(id);


-- Roles foreing keys 
ALTER TABLE role_permission
    ADD CONSTRAINT fk_role_permission_role
    FOREIGN KEY (id_role) REFERENCES role(id);

ALTER TABLE role_permission
    ADD CONSTRAINT fk_role_permission_permission
    FOREIGN KEY (id_permission) REFERENCES permissions(id);

ALTER TABLE user_roles
    ADD CONSTRAINT fk_user_roles_account
    FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE user_roles
    ADD CONSTRAINT fk_user_roles_role
    FOREIGN KEY (role_id) REFERENCES role(id);

ALTER TABLE patient
ADD CONSTRAINT fk_record_patient
FOREIGN KEY (medical_history_id) REFERENCES medical_history(id);

ALTER TABLE beneficiary
ADD CONSTRAINT fk_record_beneficiary
FOREIGN KEY (medical_history_id) REFERENCES medical_history(id);

-- Foreign keys for medical_history_relation

ALTER TABLE medical_history_relation
ADD CONSTRAINT fk_medical_history_relation_medical_history
FOREIGN KEY (medical_history_id) REFERENCES medical_history(id);

ALTER TABLE medical_history_relation
ADD CONSTRAINT fk_medical_history_relation_patient
FOREIGN KEY (patient_id) REFERENCES patient(id);

ALTER TABLE medical_history_relation
ADD CONSTRAINT fk_medical_history_relation_beneficiary
FOREIGN KEY (beneficiary_id) REFERENCES beneficiary(id);



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
