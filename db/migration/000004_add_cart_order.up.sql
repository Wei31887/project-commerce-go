CREATE TABLE "carts" (
  "id" bigint PRIMARY KEY NOT NULL
);

CREATE TABLE "cart_items" (
  "id" bigserial PRIMARY KEY,
  "cart_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "count" bigint NOT NULL
);


CREATE TABLE "orders" (
  "id" uuid PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "total_amount" decimal(12,2) NOT NULL DEFAULT 0,
  "total_count" bigint NOT NULL DEFAULT 0,
  "is_paied" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "order_items" (
  "id" bigserial PRIMARY KEY,
  "order_id" uuid NOT NULL,
  "count" bigint NOT NULL,
  "product_id" bigint NOT NULL
);

ALTER TABLE "carts" ADD FOREIGN KEY ("id") REFERENCES "users" ("id");

ALTER TABLE "cart_items" ADD FOREIGN KEY ("cart_id") REFERENCES "carts" ("id");

ALTER TABLE "cart_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
