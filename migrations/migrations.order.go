package migrations

const psqlOrder = `
	CREATE TABLE IF NOT EXISTS order(
	id uuid NOT NULL DEFAULT gen_random_uuid()
	details varchar(240) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	Upadated_at TIMESTAMP,
	)
`
