CREATE TABLE "players" (
  "id_players" bigserial PRIMARY KEY,
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
  "idActions" int,
  "ISIN" varchar(20) UNIQUE NOT NULL,
  "WKN" varchar(20) UNIQUE NOT NULL,
  "currentValue" numeric(10, 2) NOT NULL,
  "BID" numeric(10, 2) NOT NULL,
  "ASK" numeric(10, 2) NOT NULL,
  "spread" numeric(10, 2) NOT NULL,
  "timeOfLastRefresh" timestamp NOT NULL DEFAULT (now()),
  "changePercentage" numeric(5, 2) NOT NULL,
  "changeAbsolute" numeric(10, 2) NOT NULL,
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
  "actionIdBuy" bigint NOT NULL,
  "profileId" bigint NOT NULL,
  "numberStocks" int NOT NULL,
  "limit" numeric(10, 2) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "portfolio" (
  "id" bigserial PRIMARY KEY,
  "playerId" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "portfolioActions" (
  "id" int PRIMARY KEY,
  "portfolioId" bigint NOT NULL,
  "actionId" bigint NOT NULL,
  "playerId" bigint NOT NULL,
  "quantity" int NOT NULL,
  "purchasePrice" numeric(10, 2) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "buy" ("profileId");

CREATE INDEX ON "buy" ("actionIdBuy");

CREATE UNIQUE INDEX ON "buy" ("profileId", "numberStocks");

CREATE INDEX ON "portfolio" ("playerId");

CREATE INDEX ON "portfolioActions" ("portfolioId", "actionId");

ALTER TABLE "buy" ADD FOREIGN KEY ("actionIdBuy") REFERENCES "actions" ("id");

ALTER TABLE "buy" ADD FOREIGN KEY ("profileId") REFERENCES "players" ("id_players");

ALTER TABLE "portfolio" ADD FOREIGN KEY ("playerId") REFERENCES "players" ("id_players");

ALTER TABLE "portfolioActions" ADD FOREIGN KEY ("portfolioId") REFERENCES "portfolio" ("id");

ALTER TABLE "portfolioActions" ADD FOREIGN KEY ("actionId") REFERENCES "actions" ("id");

ALTER TABLE "portfolioActions" ADD FOREIGN KEY ("playerId") REFERENCES "players" ("id_players");
