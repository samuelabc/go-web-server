-- Write your migrate up statements here
CREATE TABLE tokens (
	id uuid NOT NULL,
	account_id uuid NOT NULL REFERENCES accounts(id),
	token text NOT NULL UNIQUE,
	expiry timestamp with time zone NOT NULL,
	mobile boolean NOT NULL DEFAULT FALSE,
	identifier text,

	created_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'utc'),
	updated_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'utc'),
	PRIMARY KEY (id)
);

CREATE TRIGGER update_timestamp 
	BEFORE UPDATE ON tokens
	FOR EACH ROW 
	EXECUTE PROCEDURE updatedAt_trig()

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

DROP TABLE tokens;
DROP TRIGGER IF EXISTS update_timestamp ON tokens;