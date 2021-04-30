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
   "phone" varchar (300) UNIQUE,
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
   "following" bigint NOT NULL,
   "followed" bigint NOT NULL,
   "created_at" timestamp NOT NULL DEFAULT (now()),
   "updated_at"  timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "followers" ADD FOREIGN KEY ("following") REFERENCES "users" ("id");
ALTER TABLE "followers" ADD FOREIGN KEY ("followed") REFERENCES "users" ("id");

CREATE TABLE "posts"(
   "id" bigserial PRIMARY KEY,
   "user_id" bigint NOT NULL,
   "description" text,
   "no_of_likes" bigint,
   "no_of_comments" bigint,
   "latitude_longitude" varchar,
   "created_at" timestamp NOT NULL DEFAULT (now()),
   "updated_at"  timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE TABLE "posts_media"(
   "id" bigserial PRIMARY KEY,
   "post_id" bigint NOT NULL,
   "file_url" varchar NOT NULL,
   "created_at" timestamp NOT NULL DEFAULT (now()),
   "updated_at"  timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "posts_media" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

CREATE TABLE "comments"(
   "id" bigserial PRIMARY KEY,
   "user_id" bigint NOT NULL,
   "post_id" bigint NOT NULL,
   "comment" text NOT NULL,
   "parent_comment_id" bigint NOT NULL DEFAULT (0)
   "created_at" timestamp NOT NULL DEFAULT (now()),
   "updated_at"  timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

CREATE TABLE "post_likes"(
   "id" bigserial PRIMARY KEY,
   "post_id" bigint NOT NULL,
   "user_id" bigint NOT NULL,
   "created_at" timestamp NOT NULL DEFAULT (now()),
   "updated_at"  timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "post_likes" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");
ALTER TABLE "post_likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE TABLE "comment_likes"(
   "id" bigserial PRIMARY KEY,
   "comment_id" bigint NOT NULL,
   "user_id" bigint NOT NULL,
   "created_at" timestamp NOT NULL DEFAULT (now()),
   "updated_at"  timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "comment_likes" ADD FOREIGN KEY ("comment_id") REFERENCES "comments" ("id");
ALTER TABLE "comment_likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");


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