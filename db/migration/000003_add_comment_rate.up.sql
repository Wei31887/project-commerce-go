CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "content" varchar(2000) NOT NULL DEFAULT '',
  "status" smallint NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL
);

CREATE TABLE "product_rates" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "product_id" bigint,
  "rate" smallint NOT NULL DEFAULT 0
);


CREATE TABLE "seller_rates" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "rate" smallint NOT NULL
);


ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "product_rates" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "product_rates" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "seller_rates" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");