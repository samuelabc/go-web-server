-- Write your migrate up statements here
CREATE TABLE articles (
	id uuid NOT NULL,
	title text NOT NULL,
	content text NOT NULL,
	created_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'utc'),
	updated_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'utc'),
	PRIMARY KEY (id)
);

CREATE FUNCTION updatedAt_trig() RETURNS trigger
   LANGUAGE plpgsql AS
$$BEGIN
   NEW.updated_at := (now() at time zone 'utc');
   RETURN NEW;
END;$$;

CREATE TRIGGER update_timestamp 
	BEFORE UPDATE ON articles
	FOR EACH ROW 
	EXECUTE PROCEDURE updatedAt_trig()


---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

DROP TABLE articles;
DROP FUNCTION IF EXISTS updatedAt_trig;
DROP TRIGGER IF EXISTS update_timestamp ON articles;
