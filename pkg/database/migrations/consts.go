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

const addCreatedAndUpdatedAtToPosts = `ALTER TABLE IF EXISTS posts 
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP NOT NULL DEFAULT now(),
    ADD COLUMN IF NOT EXISTS Updated_at TIMESTAMP;`

const hacer_is_admin_Default_false = `ALTER TABLE users ALTER COLUMN is_admin SET DEFAULT false`
const change_owner_id_to_owner_email = `ALTER TABLE posts DROP COLUMN owner_id, ADD COLUMN owner_email varchar(255) NOT NULL`
const Add_email_constraint = `ALTER TABLE posts ADD CONSTRAINT posts_owner_id_fk FOREIGN KEY (owner_email) REFERENCES users(email) ON UPDATE CASCADE`

const add_unique_titleAndOwner_email = `ALTER TABLE posts ADD CONSTRAINT IF NOT EXISTS posts_unique_title_email_per_post_uq UNIQUE (title, owner_email)`
