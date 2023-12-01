
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    password VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    is_doctor BOOLEAN NOT NULL
);

CREATE TABLE slots (
    slot_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    date VARCHAR NOT NULL,
    hour VARCHAR NOT NULL,
    empty BOOLEAN NOT NULL
);

CREATE TABLE appointments (
    appointment_id SERIAL PRIMARY KEY,
    doctor_id INTEGER NOT NULL,
    patient_id INTEGER NOT NULL,
    slot_id INTEGER NOT NULL
);
