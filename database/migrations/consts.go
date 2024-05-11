package migrations

const SqlCreateUuidExtension = `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`

const SqlCreateUserTable = `CREATE TABLE IF NOT EXISTS users (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    name varchar(40) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    is_admin bool NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    CONSTRAINT users_id_pk PRIMARY KEY (id),
    CONSTRAINT users_id_uq unique (email)
)`

const SqlCreatePostTable = `CREATE TABLE IF NOT EXISTS posts (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    owner_id uuid NOT NULL,
    title varchar(255) NOT NULL,
    body varchar(255) NOT NULL,
    CONSTRAINT posts_id_pk PRIMARY KEY (id),
    CONSTRAINT posts_owner_id_fk FOREIGN KEY (owner_id)
    REFERENCES users (id) ON DELETE RESTRICT ON UPDATE RESTRICT
)`
