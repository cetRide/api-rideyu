-- migrate -path migration -database "postgresql://rideyu:password@localhost:5432/pexs?sslmode=disable" -verbose up

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE "users"(
   "id" bigserial PRIMARY KEY,
   "username" varchar (300) UNIQUE NOT NULL,
   "phone" varchar (300) UNIQUE NOT NULL,
   "email" varchar (300) UNIQUE NOT NULL,
   "firstname" varchar,
   "lastname" varchar,
   "password" varchar (512) NOT NULL,
   "cover_photo" varchar,
   "profile_picture" varchar,
   "created_at" timestamp NOT NULL DEFAULT (now()),
   "updated_at"  timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("username");
CREATE INDEX ON "users" ("email");
CREATE INDEX ON "users" ("phone");

CREATE TABLE "followers"(
   "id" bigserial PRIMARY KEY,
   "following" bigserial NOT NULL,
   "follower" bigserial NOT NULL,
   "created_at" timestamp NOT NULL DEFAULT (now()),
   "updated_at"  timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "followers" ADD FOREIGN KEY ("following") REFERENCES "users" ("id");
ALTER TABLE "followers" ADD FOREIGN KEY ("followers") REFERENCES "users" ("id");



-- ON UPDATE - set current_timestamp (trigger).
DO $$
DECLARE
    t record;
BEGIN
    FOR t IN 
        SELECT * FROM information_schema.columns
        WHERE column_name = 'updated_at'
    LOOP
        EXECUTE format('CREATE TRIGGER set_updated_at
                        BEFORE INSERT ON %I.%I
                        FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp()',
                        t.table_schema, t.table_name);
    END LOOP;
END;
$$ LANGUAGE plpgsql;