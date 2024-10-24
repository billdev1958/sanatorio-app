-- ====================================
-- Tablas principales
-- ====================================

-- Tabla de dependencias
CREATE TABLE IF NOT EXISTS cat_dependencies(
    id SERIAL PRIMARY KEY,
    name VARCHAR(70),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tabla de cuentas de usuario
CREATE TABLE IF NOT EXISTS account (
    id UUID PRIMARY KEY,
    dependency_id INTEGER NOT NULL,
    phone VARCHAR(10) NOT NULL,
    email VARCHAR(75) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    password_change_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- ====================================
-- Tablas de roles y permisos
-- ====================================

-- Tabla de roles
CREATE TABLE IF NOT EXISTS cat_role (
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
CREATE TABLE IF NOT EXISTS role_permission (
    role_id INTEGER NOT NULL,
    permission_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (role_id, permission_id) -- Clave primaria compuesta
);

-- Tabla para asociar roles a usuarios
CREATE TABLE IF NOT EXISTS user_roles (
    id SERIAL PRIMARY KEY,
    account_id UUID NOT NULL, -- Relación con la cuenta/usuario
    role_id INTEGER NOT NULL, -- Relación con roles
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE (account_id, role_id) -- Clave única para evitar duplicados
);

-- ====================================
-- Tablas relacionadas con pacientes, médicos, beneficiarios y super-usuarios
-- ====================================

-- Tabla de pacientes
CREATE TABLE IF NOT EXISTS patient (
    account_id UUID NOT NULL PRIMARY KEY UNIQUE,
    medical_history_id VARCHAR(12) NOT NULL, -- id de la tabla medical_history
    legacy_id INTEGER,
    first_name VARCHAR(50) NOT NULL,
    last_name1 VARCHAR(50) NOT NULL,
    last_name2 VARCHAR(50) NOT NULL,
    curp CHAR(18) NOT NULL,
    sex CHAR(1) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tabla de médicos
CREATE TABLE IF NOT EXISTS doctor (
    account_id UUID NOT NULL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name1 VARCHAR(50) NOT NULL,
    last_name2 VARCHAR(50) NOT NULL,
    specialty_id INT NOT NULL,
    medical_license VARCHAR(25) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tabla de beneficiarios
CREATE TABLE IF NOT EXISTS beneficiary (
    id UUID PRIMARY KEY,
    account_holder UUID NOT NULL, -- id de la cuenta principal
    medical_history_id VARCHAR(12) NOT NULL, -- id de la tabla medical_history
    first_name VARCHAR(50) NOT NULL,
    last_name1 VARCHAR(50) NOT NULL,
    last_name2 VARCHAR(50) NOT NULL,
    curp CHAR(18) NOT NULL,
    sex CHAR(1) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tabla de recepcionista
CREATE TABLE IF NOT EXISTS recepcionista (
    account_id UUID NOT NULL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name1 VARCHAR(50) NOT NULL,
    last_name2 VARCHAR(50) NOT NULL,
    curp CHAR(18) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tabla de super admin
CREATE TABLE IF NOT EXISTS super_admin (
    account_id UUID NOT NULL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name1 VARCHAR(50) NOT NULL,
    last_name2 VARCHAR(50) NOT NULL,
    curp CHAR(18) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- ====================================
-- Tablas relacionadas con historiales médicos
-- ====================================

-- Tabla de historiales médicos
CREATE TABLE IF NOT EXISTS medical_history (
    id VARCHAR(12) PRIMARY KEY NOT NULL,
    date_of_record DATE, -- 'Fecha'
    time_of_record TIME, -- 'Hora'
    patient_name VARCHAR(50) NOT NULL, -- 'Nombre'
    curp CHAR(18) NOT NULL, -- 'CURP'
    birth_date DATE , -- 'Fecha nacimiento'
    age VARCHAR(3) , -- 'Edad'
    gender VARCHAR(10) NOT NULL, -- 'Sexo'
    place_of_origin VARCHAR(10) , -- 'Procedencia'
    ethnic_group VARCHAR(20) , -- 'Grupo étnico'
    phone_number VARCHAR(10) , -- 'Teléfono'
    address VARCHAR(50) , -- 'Domicilio'
    occupation VARCHAR(20) , -- 'Ocupación'
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
    body_temperature VARCHAR(10), -- 'Temperatura'
    weight VARCHAR(5), -- 'Peso'
    height VARCHAR(10), -- 'Talla'
    bmi VARCHAR(10), -- 'IMC'
    heart_rate VARCHAR(10), -- 'Frecuencia cardiaca'
    respiratory_rate VARCHAR(10), -- 'Frecuencia respiratoria'
    blood_pressure VARCHAR(10), -- 'T/A'
    physical VARCHAR(100), -- 'Habitus exterior'
    head VARCHAR(100), -- 'Cabeza'
    neck_and_chest VARCHAR(100), -- 'Cuello tórax'
    abdomen VARCHAR(100), -- 'Abdomen'
    genital VARCHAR(100), -- 'Genitales'
    extremities VARCHAR(100), -- 'Extremidades'
    previous_results VARCHAR(100), -- 'Resultados previos y actuales'
    diagnoses VARCHAR(100), -- 'Diagnósticos o problemas'
    pharmacological_treatment VARCHAR(100), -- 'Tratamiento farmacológico'
    prognosis VARCHAR(100), -- 'Pronóstico'
    doctor_name VARCHAR(50), -- 'Nombre médico'
    medical_license VARCHAR(10), -- 'Cédula profesional'
    specialty_license VARCHAR(10) -- 'Cédula especialidad'
);

-- Tabla de relaciones de historial médico
CREATE TABLE IF NOT EXISTS medical_history_relation (
    id UUID PRIMARY KEY,
    medical_history_id VARCHAR(12) NOT NULL,
    patient_id UUID,
    beneficiary_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    CHECK (patient_id IS NOT NULL OR beneficiary_id IS NOT NULL)
);

-- ====================================
-- Tablas relacionadas con citas y consultas
-- ====================================

-- Tabla de citas
CREATE TABLE IF NOT EXISTS appointment (
    id UUID PRIMARY KEY,
    doctor_id UUID NOT NULL,
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

-- Tabla de consultas
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

-- ====================================
-- Tablas relacionadas con oficinas, horarios y servicios
-- ====================================

-- Tabla de especialidades
CREATE TABLE IF NOT EXISTS specialty (
    id SERIAL PRIMARY KEY,
    name VARCHAR(75) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tabla de oficinas
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

-- Tabla de horarios
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

-- ====================================
-- Otras tablas auxiliares
-- ====================================

-- Tabla de estatus de citas
CREATE TABLE IF NOT EXISTS appointment_status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tabla de estatus de oficinas
CREATE TABLE IF NOT EXISTS office_status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tabla de servicios
CREATE TABLE IF NOT EXISTS services (
    id SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tabla de notas de evolución
CREATE TABLE IF NOT EXISTS evolution_note (
    folio SERIAL PRIMARY KEY, -- folio
    date DATE NOT NULL, -- fecha
    name VARCHAR(50) NOT NULL, -- nombre
    curp CHAR(18) NOT NULL, -- curp
    department VARCHAR(15), -- dependencia
    affiliation VARCHAR(15), -- afiliacion
    age VARCHAR(3) NOT NULL, -- edad
    weight VARCHAR(6) NOT NULL, -- peso
    height VARCHAR(6) NOT NULL, -- estatura
    heart_rate VARCHAR(6) NOT NULL, -- fc
    respiratory_rate VARCHAR(6) NOT NULL, -- fr
    blood_pressure VARCHAR(6) NOT NULL, -- ta
    temperature VARCHAR(6) NOT NULL, -- temperatura
    spo2 VARCHAR(6) NOT NULL, -- spo2
    glucose VARCHAR(6) NOT NULL, -- glucosa
    notes VARCHAR(6) NOT NULL -- notas
);

-- Tabla de incapacidades
CREATE TABLE IF NOT EXISTS incapacity (
    folio SERIAL PRIMARY KEY, -- folio
    date DATE NOT NULL, -- fecha
    name VARCHAR(50) NOT NULL, -- nombre
    curp CHAR(18) NOT NULL, -- curp
    department VARCHAR(15), -- dependencia
    assigned_to VARCHAR(15), -- adscrito
    total_days VARCHAR(3) NOT NULL, -- totaldias
    start_incapacity DATE NOT NULL, -- inicio
    end_incapacity DATE NOT NULL, -- fin
    doctor VARCHAR(50) NOT NULL, -- medico
    service VARCHAR(20) NOT NULL, -- servicio
    key_code VARCHAR(10) NOT NULL -- clave
);

-- ====================================
-- Claves foráneas
-- ====================================

-- Foreign keys para la tabla account
ALTER TABLE account
ADD CONSTRAINT fk_role_account
FOREIGN KEY (role_id) REFERENCES cat_role(id);

ALTER TABLE account 
ADD CONSTRAINT fk_dependency_id
FOREIGN KEY (dependency_id) REFERENCES cat_dependencies(id);  -- Relaciona las cuentas directamente con las dependencias

-- Foreign keys para la tabla patient
ALTER TABLE patient
ADD CONSTRAINT fk_account_patient
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE patient
ADD CONSTRAINT fk_record_patient
FOREIGN KEY (medical_history_id) REFERENCES medical_history(id);

-- Foreign keys para la tabla doctor
ALTER TABLE doctor
ADD CONSTRAINT fk_account_doctor
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE doctor
ADD CONSTRAINT fk_specialty_doctor
FOREIGN KEY (specialty_id) REFERENCES specialty(id);

-- Foreign keys para la tabla office
ALTER TABLE office
ADD CONSTRAINT fk_specialty_office
FOREIGN KEY (specialty_id) REFERENCES specialty(id);

ALTER TABLE office
ADD CONSTRAINT fk_status_office
FOREIGN KEY (status_id) REFERENCES office_status(id);

-- **Corrección aplicada aquí**
ALTER TABLE office
ADD CONSTRAINT fk_doctor_office
FOREIGN KEY (doctor_id) REFERENCES doctor(account_id);

-- Foreign keys para la tabla schedule
ALTER TABLE schedule
ADD CONSTRAINT fk_office_schedule
FOREIGN KEY (office_id) REFERENCES office(id);

-- Foreign keys para la tabla appointment
ALTER TABLE appointment
ADD CONSTRAINT fk_schedule_appointment
FOREIGN KEY (schedule_id) REFERENCES schedule(id);

-- **Corrección aplicada aquí**
ALTER TABLE appointment
ADD CONSTRAINT fk_doctor_appointment
FOREIGN KEY (doctor_id) REFERENCES doctor(account_id);

ALTER TABLE appointment
ADD CONSTRAINT fk_patient_appointment
FOREIGN KEY (patient_account_id) REFERENCES patient(account_id);

ALTER TABLE appointment
ADD CONSTRAINT fk_office_appointment
FOREIGN KEY (office_id) REFERENCES office(id);

ALTER TABLE appointment
ADD CONSTRAINT fk_status_appointment
FOREIGN KEY (status_id) REFERENCES appointment_status(id);

-- Foreign keys para la tabla beneficiary
ALTER TABLE beneficiary
ADD CONSTRAINT fk_account_holder_beneficiary
FOREIGN KEY (account_holder) REFERENCES patient(account_id);

ALTER TABLE beneficiary
ADD CONSTRAINT fk_record_beneficiary
FOREIGN KEY (medical_history_id) REFERENCES medical_history(id);

-- Foreign keys para la tabla medical_history_relation
ALTER TABLE medical_history_relation
ADD CONSTRAINT fk_medical_history_relation_medical_history
FOREIGN KEY (medical_history_id) REFERENCES medical_history(id);

-- **Corrección aplicada aquí**
ALTER TABLE medical_history_relation
ADD CONSTRAINT fk_medical_history_relation_patient
FOREIGN KEY (patient_id) REFERENCES patient(account_id);

ALTER TABLE medical_history_relation
ADD CONSTRAINT fk_medical_history_relation_beneficiary
FOREIGN KEY (beneficiary_id) REFERENCES beneficiary(id);

-- Foreign keys para la tabla role_permission
ALTER TABLE role_permission
ADD CONSTRAINT fk_role_permission_role
FOREIGN KEY (role_id) REFERENCES cat_role(id);

ALTER TABLE role_permission
ADD CONSTRAINT fk_role_permission_permission
FOREIGN KEY (permission_id) REFERENCES permissions(id);

-- Foreign keys para la tabla user_roles
ALTER TABLE user_roles
ADD CONSTRAINT fk_user_roles_account
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE user_roles
ADD CONSTRAINT fk_user_roles_role
FOREIGN KEY (role_id) REFERENCES cat_role(id);

-- Foreign keys para la tabla consultation
-- **Corrección aplicada aquí**
ALTER TABLE consultation
ADD CONSTRAINT fk_patient_consultation
FOREIGN KEY (patient_id) REFERENCES patient(account_id);

-- Foreign keys para la tabla super_admin
-- **Corrección aplicada aquí**
ALTER TABLE super_admin
ADD CONSTRAINT fk_account_super_admin
FOREIGN KEY (account_id) REFERENCES account(id);

-- ====================================
-- Funciones y triggers
-- ====================================

-- Función para validar el estado de la oficina antes de programar un horario
CREATE OR REPLACE FUNCTION validate_office_status()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT status_id FROM office WHERE id = NEW.office_id) != 1 THEN
        RAISE EXCEPTION 'The office is not available for scheduling';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger para validar el estado de la oficina antes de insertar o actualizar un horario
CREATE TRIGGER trigger_validate_office_status
BEFORE INSERT OR UPDATE ON schedule
FOR EACH ROW EXECUTE FUNCTION validate_office_status();
