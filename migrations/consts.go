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
const psqladmin = `
CREATE TABLE IF NOT EXISTS admin_users(
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	name varchar(50) NOT NULL,
	email varchar(100) NOT NULL,
	roll varchar(40) NOT NULL,
	password varchar(200) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP,
	CONSTRAINT admin_user_id_pk PRIMARY KEY (id),
	CONSTRAINT admin_user_email_uq UNIQUE (email)
)
`
