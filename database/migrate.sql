CREATE TABLE IF NOT EXISTS roles (
	id SERIAL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	name VARCHAR(64) UNIQUE NOT NULL
);
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	username VARCHAR(64) UNIQUE NOT NULL,
	email VARCHAR(255) UNIQUE NOT NULL,
	password VARCHAR(60) NOT NULL,
	first_name VARCHAR(64),
	last_name VARCHAR(64),
	role_id INT,
	CONSTRAINT fk_role FOREIGN KEY(role_id) REFERENCES roles(id) ON DELETE
	SET NULL
);
CREATE OR REPLACE FUNCTION update_timestamp() RETURNS TRIGGER AS $$ BEGIN NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;
DO $$
DECLARE t text;
BEGIN FOR t IN
SELECT table_name
FROM information_schema.columns
WHERE column_name = 'updated_at' LOOP EXECUTE format(
		'CREATE TRIGGER trigger_update_timestamp
		BEFORE UPDATE ON %I
		FOR EACH ROW EXECUTE PROCEDURE update_timestamp()',
		t,
		t
	);
END loop;
END;
$$ language 'plpgsql';