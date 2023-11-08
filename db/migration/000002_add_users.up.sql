CREATE TABLE "users" (
                         "username" varchar PRIMARY KEY,
                         "hashed_password" varchar,
                         "full_name" varchar NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "password_changed_at" timestamptz NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");
CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");