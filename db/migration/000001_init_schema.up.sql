CREATE TYPE "transactiontype" AS ENUM (
  'DEBIT',
  'CREDIT',
  'TRANSFER'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "type" transactiontype NOT NULL,
  "category_id" bigint NOT NULL,
  "wallet_id" bigint
);

CREATE TABLE "category" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "name" varchar NOT NULL,
  "type" transactiontype NOT NULL,
  "icon" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "wallet" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "name" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "icon" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("full_name");

CREATE INDEX ON "transactions" ("user_id", "category_id", "type", "updated_at");

CREATE INDEX ON "category" ("user_id");

CREATE INDEX ON "wallet" ("user_id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallet" ("id");

ALTER TABLE "category" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "wallet" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
