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

CREATE TABLE IF NOT EXISTS admin (
    account_id UUID PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS doctor (
    account_id UUID PRIMARY KEY,
    medical_license VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS patient (
    account_id UUID PRIMARY KEY,
    curp VARCHAR(18)
);

ALTER TABLE account
ADD CONSTRAINT fk_rol_account
FOREIGN KEY (rol) REFERENCES cat_rol(id);

ALTER TABLE users
ADD CONSTRAINT fk_account_id_users
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE admin
ADD CONSTRAINT fk_account_id_admin
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE doctor
ADD CONSTRAINT fk_account_id_doctor
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE patient
ADD CONSTRAINT fk_account_id_patient
FOREIGN KEY (account_id) REFERENCES account(id);

