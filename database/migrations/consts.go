package migrations

const SqlCreateUuidExtension = `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`

const SqlCreateRollTable = `CREATE TABLE IF NOT EXISTS rolls (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    roll varchar(10) NOT NULL 
)`

const SqlCreateUserTable = `CREATE TABLE IF NOT EXISTS users (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    name varchar(40) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    roll_name varchar(10) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    CONSTRAINT users_id_pk PRIMARY KEY (id),
    CONSTRAINT users_id_uq unique (email),
    CONSTRAINT users_roll_name_fk FOREIGN KEY (roll_name)
    REFERENCES rolls (roll)
)`

const SqlCreatePostTable = `CREATE TABLE IF NOT EXISTS posts (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    owner_id uuid NOT NULL,
    title varchar(255) NOT NULL,
    body varchar(255) NOT NULL,
    CONSTRAINT posts_id_pk PRIMARY KEY (id),
    CONSTRAINT posts.owner_id_fk FOREIGN KEY (owner_id)
    REFERENCES users (id) ON DELETE RESTRICT ON UPDATE RESTRICT,
)`

const SqlCreateAdminUsersTable = `CREATE TABLE IF NOT EXISTS admin_users (
    id uuid NOT NULL,
    name varchar(40) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    CreatedAt TIMESTAMP NOT NULL DEFAULT now(),
    updatedAt TIMESTAMP NOT NULL DEFAULT now(),
    CONSTRAINT admin_users_id_pk PRIMARY KEY (id),
    CONSTRAINT admin_users_name_uq UNIQUE (name),
    CONSTRAINT admin_users_email_uq UNIQUE (email)
)`
