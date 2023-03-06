CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL DEFAULT '',
  "category_id" bigint NOT NULL DEFAULT 0,
  "stock" bigint NOT NULL DEFAULT 0,
  "sell" bigint NOT NULL DEFAULT 0,
  "price" decimal(12,2) NOT NULL DEFAULT 0,
  "on_sell" decimal(3,2) NOT NULL DEFAULT 1,
  "description" text NOT NULL DEFAULT '',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product_categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(200) NOT NULL DEFAULT '',
  "sort" int NOT NULL DEFAULT 0
);

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "product_categories" ("id");


