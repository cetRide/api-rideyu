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
   "no_of_likes" bigint,
   "created_at" timestamp NOT NULL DEFAULT (now()),
   "updated_at"  timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

CREATE TABLE "parent_child_comments"(
   "id" bigserial PRIMARY KEY,
   "parent_comment_id" bigint NOT NULL,
   "child_comment_id" bigint NOT NULL,
   "created_at" timestamp NOT NULL DEFAULT (now()),
   "updated_at"  timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "parent_child_comments" ADD FOREIGN KEY ("parent_comment_id") REFERENCES "comments" ("id");
ALTER TABLE "parent_child_comments" ADD FOREIGN KEY ("child_comment_id") REFERENCES "comments" ("id");

CREATE TABLE "categories"(
   "id" bigserial PRIMARY KEY,
   "type_code" char(1) UNIQUE NOT NULL,
   "type_code_name" text UNIQUE NOT NULL,
   "created_at" timestamp NOT NULL DEFAULT (now()),
   "updated_at"  timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "likes"(
   "id" bigserial PRIMARY KEY,
   "entity_id" bigint NOT NULL,
   "user_id" bigint NOT NULL,
   "type_code" char(1) NOT NULL,
   "created_at" timestamp NOT NULL DEFAULT (now()),
   "updated_at"  timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "likes" ADD FOREIGN KEY ("entity_id") REFERENCES "comments" ("id");
ALTER TABLE "likes" ADD FOREIGN KEY ("type_code") REFERENCES "categories" ("type_code");
ALTER TABLE "likes" ADD FOREIGN KEY ("entity_id") REFERENCES "posts" ("id");
ALTER TABLE "likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");


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