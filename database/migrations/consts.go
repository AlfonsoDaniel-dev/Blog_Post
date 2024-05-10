package migrations

const psqlUser = `CREATE TABLE IF NOT EXISTS users(
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	name varchar(100) NOT NULL,
	email varchar(50) NOT NULL,
	password varchar(150) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	Updated_at TIMESTAMP,
	CONSTRAINT users_id_pk PRIMARY KEY (id),
	CONSTRAINT users_email_uq UNIQUE (email)
)`

const psqlProduct = `CREATE TABLE IF NOT EXISTS products(
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	name varchar(60) NOT NULL,
	price INT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	Upadated_at TIMESTAMP,
	CONSTRAINT products_id_pk PRIMARY KEY (id)
)`

const psqlOrder = `
CREATE TABLE IF NOT EXISTS orders(
id uuid NOT NULL DEFAULT gen_random_uuid(),
user_id uuid NOT NULL,
details varchar(240) NOT NULL,
created_at TIMESTAMP NOT NULL DEFAULT now(),
Upadated_at TIMESTAMP,
CONSTRAINT orders_id_pk PRIMARY KEY (id),
CONSTRAINT orders_users_id_fk FOREIGN KEY (user_id)
REFERENCES users (id) ON DELETE RESTRICT ON UPDATE CASCADE
)
`

const psql16_04_24Migration = `ALTER TABLE users ADD COLUMN is_admin BOOL NOT NULL DEFAULT false`

const psql16_04_24Migrationproduct = `ALTER TABLE products ADD COLUMN creator id uuid NOT NULL FOREIGN KEY REFERENCES users (id) on UPDATE CASCADE on DELETE RESTRICT`
