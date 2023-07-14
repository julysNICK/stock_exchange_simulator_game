CREATE TABLE "players" (
  "id_player" bigserial PRIMARY KEY,
  "username" varchar UNIQUE,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "cash" numeric(10, 2) NOT NULL DEFAULT 1000000.00,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamp NOT NULL DEFAULT '0001-01-01 00:00:00',
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "actions" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "id_actions" int,
  "isin" varchar(20) UNIQUE NOT NULL,
  "wkn" varchar(20) UNIQUE NOT NULL,
  "current_value" numeric(10, 2) NOT NULL,
  "bid" numeric(10, 2) NOT NULL,
  "ask" numeric(10, 2) NOT NULL,
  "spread" numeric(10, 2) NOT NULL,
  "time_of_last_refresh" timestamp NOT NULL DEFAULT (now()),
  "change_percentage" numeric(5, 2) NOT NULL,
  "change_absolute" numeric(10, 2) NOT NULL,
  "peak24h" numeric(10, 2) NOT NULL,
  "low24h" numeric(10, 2) NOT NULL,
  "peak7d" numeric(10, 2) NOT NULL,
  "low7d" numeric(10, 2) NOT NULL,
  "peak30d" numeric(10, 2) NOT NULL,
  "low30d" numeric(10, 2) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "buy" (
  "id" bigserial PRIMARY KEY,
  "action_id_buy" bigint NOT NULL,
  "profile_id" bigint NOT NULL,
  "number_stocks" int NOT NULL,
  "limit" numeric(10, 2) NOT NULL,
  "status" varchar(255) NOT NULL DEFAULT 'pending',
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "portfolio" (
  "id" bigserial PRIMARY KEY,
  "player_id" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "portfolioActions" (
  "id" bigserial PRIMARY KEY,
  "portfolio_id" bigint NOT NULL,
  "action_id" bigint NOT NULL,
  "player_id" bigint NOT NULL,
  "quantity" int NOT NULL,
  "purchase_price" numeric(10, 2) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "buy" ("profile_id");

CREATE INDEX ON "buy" ("action_id_buy");

CREATE UNIQUE INDEX ON "buy" ("profile_id", "number_stocks");

CREATE INDEX ON "portfolio" ("player_id");

CREATE INDEX ON "portfolioActions" ("portfolio_id", "action_id");

ALTER TABLE "buy" ADD FOREIGN KEY ("action_id_buy") REFERENCES "actions" ("id");

ALTER TABLE "buy" ADD FOREIGN KEY ("profile_id") REFERENCES "players" ("id_player");

ALTER TABLE "portfolio" ADD FOREIGN KEY ("player_id") REFERENCES "players" ("id_player");

ALTER TABLE "portfolioActions" ADD FOREIGN KEY ("portfolio_id") REFERENCES "portfolio" ("id");

ALTER TABLE "portfolioActions" ADD FOREIGN KEY ("action_id") REFERENCES "actions" ("id");

ALTER TABLE "portfolioActions" ADD FOREIGN KEY ("player_id") REFERENCES "players" ("id_player");
