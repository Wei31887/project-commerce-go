CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "user_name" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "full_name" varchar NOT NULL,
  "gender" smallint,
  "user_rank" int NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "changed_password_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_delivers" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "address" varchar NOT NULL
);

CREATE TABLE "sellers" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "seller_name" varchar NOT NULL,
  "bank_account" varchar NOT NULL
);

ALTER TABLE "user_delivers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");