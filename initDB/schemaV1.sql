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
    specialty_license VARCHAR(25) NOT NULL,
    medical_license VARCHAR(25) NOT NULL,
    sex VARCHAR(1) NOT NULL,
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
CREATE TABLE IF NOT EXISTS receptionist (
    account_id UUID NOT NULL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name1 VARCHAR(50) NOT NULL,
    last_name2 VARCHAR(50) NOT NULL,
    curp CHAR(18) NOT NULL,
    sex VARCHAR(1),
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
    sex VARCHAR(1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tabla de admin
CREATE TABLE IF NOT EXISTS admin (
    account_id UUID NOT NULL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name1 VARCHAR(50) NOT NULL,
    last_name2 VARCHAR(50) NOT NULL,
    curp CHAR(18) NOT NULL,
    sex VARCHAR(1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);


-- ====================================
-- Tablas relacionadas con historiales médicos
-- ====================================

-- Updates 
-- Procedencias  = 1: dependencias,  ALUMNOS, FAAPA ,SUTES Y CONFIANZA AL REGISTRO DE PACIENTES
-- DerechoHabiencia: IMSS, ISSSTE, ISSEMYM, SEDENA, PEMEX, LLENADO DE HISTORIAL, NOTA DE EVOLUCION Y HOJA DE REFERENCIA

CREATE TABLE IF NOT EXISTS cat_medical_institutions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tabla de historiales médicos
CREATE TABLE IF NOT EXISTS medical_history (
    id UUID PRIMARY KEY NOT NULL,
    medical_history_id VARCHAR(12) NOT NULL UNIQUE,
    date_of_record DATE, -- 'Fecha'
    time_of_record TIME, -- 'Hora'
    patient_name VARCHAR(50) NOT NULL, -- 'Nombre'
    lastname_1 VARCHAR (50) NOT NULL,
    lastname_2 VARCHAR (50) NOT NULL,
    curp CHAR(18) NOT NULL, -- 'CURP'
    birth_date DATE , -- 'Fecha nacimiento'
    age VARCHAR(3) , -- 'Edad'
    gender VARCHAR(1) NOT NULL, -- 'Sexo'
    adress VARCHAR(100), -- 'Domicilio'
    dependency_id INTEGER, -- 'Procedencia'
    derecho_habiencia_id INTEGER,
    ethnic_group VARCHAR(20) , -- 'Grupo étnico'
    phone_number VARCHAR(10) , -- 'Teléfono'
    address VARCHAR(50) , -- 'Domicilio'
    occupation VARCHAR(20) , -- 'Ocupación'
    guardian_name VARCHAR(50), -- 'Nombre tutor'
    family_medical_history VARCHAR(100), -- 'Antecedentes heredofamiliares'
    non_pathological_history VARCHAR(100), -- 'Antecedentes personales no patológicos'
    pathological_history VARCHAR(100), -- 'Antecedentes personales patológicos'
    gynec_obstetric_history VARCHAR(100), -- 'Antecendentes gineco-obstétricos / androgenos'
    current_condition VARCHAR(100), -- 'Padecimiento actual'
    cardiovascular VARCHAR(100), -- 'Cardiovascular'
    respiratory VARCHAR(100), -- 'Respiratorio'
    gastrointestinal VARCHAR(100), -- 'Gastrointestinal'
    genitourinary VARCHAR(100), -- 'Genitourinario'
    hematic_lymphatic VARCHAR(100), -- 'Hemático linfático'
    endocrine VARCHAR(100), -- 'Endocrino'
    nervous_system VARCHAR(100), -- 'Nervioso'
    musculoskeletal VARCHAR(100), -- 'Musculo esquelético'
    skin_mucous_appendages VARCHAR(100), -- 'Piel, Mucosas y Anexos'
    body_temperature VARCHAR(10), -- 'Temperatura'
    weight VARCHAR(5), -- 'Peso'
    height VARCHAR(10), -- 'Talla'
    bmi VARCHAR(10), -- 'IMC'
    heart_rate VARCHAR(10), -- 'Frecuencia cardiaca'
    respiratory_rate VARCHAR(10), -- 'Frecuencia respiratoria'
    blood_pressure VARCHAR(10), -- 'T/A'
    physical VARCHAR(100), -- 'Habitus exterior'
    head VARCHAR(100), -- 'Cabeza'
    neck VARCHAR(100),
    chest VARCHAR(100),
    abdomen VARCHAR(100), -- 'Abdomen'
    genital VARCHAR(100), -- 'Genitales'
    extremities VARCHAR(100), -- 'Extremidades'
    skin VARCHAR(100), -- 'Piel'
    previous_results VARCHAR(100), -- 'Resultados previos y actuales'
    diagnoses VARCHAR(100), -- 'Diagnósticos o problemas'
    pharmacological_treatment VARCHAR(100), -- 'Tratamiento farmacológico'
    prognosis VARCHAR(100), -- 'Pronóstico'
    doctor_name VARCHAR(50), -- 'Nombre médico'
    medical_license VARCHAR(10), -- 'Cédula profesional'
    specialty_license VARCHAR(10), -- 'Cédula especialidad'
    status_md BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

-- Tabla de relaciones de historial médico
CREATE TABLE IF NOT EXISTS medical_history_relation (
    id UUID PRIMARY KEY,
    patient_id UUID,
    beneficiary_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    CHECK (patient_id IS NOT NULL OR beneficiary_id IS NOT NULL)
);

-- ====================================
-- Tablas relacionadas con citas y consultas
-- ====================================

CREATE TABLE IF NOT EXISTS days (
    id SERIAL PRIMARY KEY,
    day_of_week INT NOT NULL UNIQUE,
    name VARCHAR(10) NOT NULL
);

-- Tabla de citas
CREATE TABLE IF NOT EXISTS appointment (
    id UUID PRIMARY KEY,
    account_id UUID NOT NULL,
    schedule_id INTEGER NOT NULL, -- ID del horario en office_schedule
    patient_id UUID NOT NULL,
    beneficiary_id UUID,
    time_start TIMESTAMP NOT NULL,
    time_end TIMESTAMP NOT NULL,
    status_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT chk_time_validity CHECK (time_start < time_end)
);


-- Tabla de consultas
CREATE TABLE IF NOT EXISTS consultation (
    id SERIAL PRIMARY KEY,
    appointment_id UUID NOT NULL,
    reason TEXT NOT NULL,
    symptoms TEXT NOT NULL,
    doctor_notes TEXT,
    requested_tests TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- ====================================
-- Tablas relacionadas con oficinas, horarios y servicios
-- ====================================

-- Tabla de especialidades
CREATE TABLE IF NOT EXISTS cat_specialty (
    id SERIAL PRIMARY KEY,
    name VARCHAR(75) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tabla de oficinas
CREATE TABLE IF NOT EXISTS office (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tabla de horarios


CREATE TABLE IF NOT EXISTS cat_shift(
    id SERIAL PRIMARY KEY,
    name VARCHAR (10) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS office_schedule (
    id SERIAL PRIMARY KEY,
    office_id INTEGER NOT NULL,
    shift_id INTEGER NOT NULL,
    service_id INTEGER NOT NULL,
    doctor_id UUID NOT NULL,
    status_id INTEGER NOT NULL,
    day_of_week INT NOT NULL, -- Día de la semana (1 = Lunes, 7 = Domingo)
    time_start TIME NOT NULL, -- Hora de inicio
    time_end TIME NOT NULL,   -- Hora de fin
    time_duration INTERVAL NOT NULL, -- Duración
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE (office_id, shift_id, service_id, doctor_id, day_of_week, time_start)
);

CREATE TABLE IF NOT EXISTS schedule_block (
    id SERIAL PRIMARY KEY,
    office_schedule_id INTEGER NOT NULL,
    block_date DATE NOT NULL,
    time_start TIME,
    time_end TIME,
    reason VARCHAR(255),
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
    id SERIAL PRIMARY KEY,
    folio VARCHAR(10), -- folio
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
    id SERIAL PRIMARY KEY,
    folio VARCHAR(10), -- folio
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
FOREIGN KEY (medical_history_id) REFERENCES medical_history(medical_history_id);

-- Foreign keys para la tabla doctor
ALTER TABLE doctor
ADD CONSTRAINT fk_account_doctor
FOREIGN KEY (account_id) REFERENCES account(id);

-- Foreign keys para la tabla office

ALTER TABLE office_schedule
ADD CONSTRAINT fk_status_office
FOREIGN KEY (status_id) REFERENCES office_status(id);


-- Foreign keys para la tabla office_schedule
ALTER TABLE office_schedule
ADD CONSTRAINT fk_office_schedule_office_id
FOREIGN KEY (office_id) REFERENCES office(id);

ALTER TABLE office_schedule
ADD CONSTRAINT fk_office_schedule_shift
FOREIGN KEY (shift_id) REFERENCES cat_shift(id);

ALTER TABLE office_schedule
ADD CONSTRAINT fk_office_schedule_service
FOREIGN KEY (service_id) REFERENCES services(id);

ALTER TABLE office_schedule
ADD CONSTRAINT fk_office_schedule_doctor_id
FOREIGN KEY (doctor_id) REFERENCES doctor(account_id);

ALTER TABLE office_schedule
ADD CONSTRAINT fk_schedule_day
FOREIGN KEY (day_of_week) REFERENCES days(day_of_week);


-- Foreign keys para la tabla schedule_block
ALTER TABLE schedule_block
ADD CONSTRAINT fk_schedule_block_office_schedule
FOREIGN KEY (office_schedule_id) REFERENCES office_schedule(id);

-- Foreign keys para la tabla appointment

ALTER TABLE appointment
ADD CONSTRAINT fk_patient_id
FOREIGN KEY (patient_id) REFERENCES patient(account_id);

ALTER TABLE appointment
ADD CONSTRAINT fk_account_id_appointment
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE appointment
ADD CONSTRAINT fk_status_appointment
FOREIGN KEY (status_id) REFERENCES appointment_status(id);

-- Relación con office_schedule
ALTER TABLE appointment
ADD CONSTRAINT fk_schedule
FOREIGN KEY (schedule_id) REFERENCES office_schedule (id);

-- Relación con appointment_status
ALTER TABLE appointment
ADD CONSTRAINT fk_status
FOREIGN KEY (status_id) REFERENCES appointment_status (id);


-- Foreign keys para la tabla beneficiary
ALTER TABLE beneficiary
ADD CONSTRAINT fk_account_holder_beneficiary
FOREIGN KEY (account_holder) REFERENCES patient(account_id);

ALTER TABLE beneficiary
ADD CONSTRAINT fk_record_beneficiary
FOREIGN KEY (medical_history_id) REFERENCES medical_history(medical_history_id);

-- Foreign keys para la tabla medical_history
ALTER TABLE medical_history
ADD CONSTRAINT fk_dependency_medical_history
FOREIGN KEY (dependency_id) REFERENCES cat_dependencies(id); 

ALTER TABLE medical_history
ADD CONSTRAINT fk_derecho_habiencia_medical_history
FOREIGN KEY (derecho_habiencia_id) REFERENCES cat_medical_institutions(id); 

-- Foreign keys para la tabla medical_history_relation

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

-- Foreign keys para la tabla consultation
ALTER TABLE consultation
ADD CONSTRAINT fk_appointment_consultation
FOREIGN KEY (appointment_id) REFERENCES appointment(id);

-- Foreign keys para la tabla super_admin
ALTER TABLE super_admin
ADD CONSTRAINT fk_account_super_admin
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE admin
ADD CONSTRAINT fk_account_admin
FOREIGN KEY (account_id) REFERENCES account(id);

-- ====================================
-- Funciones y triggers
-- ====================================

CREATE OR REPLACE FUNCTION validate_appointment()
RETURNS TRIGGER AS $$
BEGIN
  -- Verificar que el status_id del schedule sea 1
  IF NOT EXISTS (
    SELECT 1
    FROM office_schedule
    WHERE id = NEW.schedule_id AND status_id = 1
  ) THEN
    RAISE EXCEPTION 'El schedule no está disponible para registrar appointments.';
  END IF;

  -- Verificar conflictos de horarios en la misma fecha y horario
  IF EXISTS (
    SELECT 1
    FROM appointment
    WHERE schedule_id = NEW.schedule_id
      AND DATE(time_start) = DATE(NEW.time_start) -- Misma fecha
      AND (
        time_start < NEW.time_end AND time_end > NEW.time_start -- Solapamiento
      )
  ) THEN
    RAISE EXCEPTION 'Ya existe un appointment en el mismo horario y fecha.';
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
