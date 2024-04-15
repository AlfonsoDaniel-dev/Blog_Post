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
CREATE TABLE IF NOT EXISTS order(
id uuid NOT NULL DEFAULT gen_random_uuid()
user_id uuid NOT NULL NOT NULL
details varchar(240) NOT NULL,
created_at TIMESTAMP NOT NULL DEFAULT now(),
Upadated_at TIMESTAMP,
CONSTRAINT orders_id_pk PRIMARY KEY (id),
CONSTRAINT orders_users_is_fk FOREIGN KEY (users_id)
REFERENCES users (id) ON DELETE RESTRICT ON UPDATE CASCADE
)
`
