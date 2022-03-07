-- Write your migrate up statements here
CREATE TABLE users (
	id uuid NOT NULL,
	name text NOT NULL,
	password_hash text NOT NULL,
	email text,
	created_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'utc'),
	updated_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'utc'),
	PRIMARY KEY (id)
);

CREATE TRIGGER update_timestamp 
	BEFORE UPDATE ON users
	FOR EACH ROW 
	EXECUTE PROCEDURE updatedAt_trig()

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

DROP TABLE users;
DROP TRIGGER IF EXISTS update_timestamp ON users;