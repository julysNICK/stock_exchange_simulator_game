CREATE TABLE "players" (
  "id" bigserial,
  "username" varchar UNIQUE,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "cash" "decimal(10, 2)" NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestampz NOT NULL DEFAULT (0001-01-01 00:00:00Z),
  "created_at" timestampz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("id", "email")
);

CREATE TABLE "actions" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "idActions" int,
  "ISIN" varchar(20) UNIQUE NOT NULL,
  "WKN" varchar(20) UNIQUE NOT NULL,
  "currentValue" "decimal(10, 2)" NOT NULL,
  "BID" "decimal(10, 2)" NOT NULL,
  "ASK" "decimal(10, 2)" NOT NULL,
  "spread" "decimal(10, 2)" NOT NULL,
  "timeOfLastRefresh" datetime NOT NULL,
  "changePercentage" "decimal(5, 2)" NOT NULL,
  "changeAbsolute" "decimal(10, 2)" NOT NULL,
  "peak24h" "decimal(10, 2)" NOT NULL,
  "low24h" "decimal(10, 2)" NOT NULL,
  "peak7d" "decimal(10, 2)" NOT NULL,
  "low7d" "decimal(10, 2)" NOT NULL,
  "peak30d" "decimal(10, 2)" NOT NULL,
  "low30d" "decimal(10, 2)" NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now())
);

CREATE TABLE "buy" (
  "id" bigserial PRIMARY KEY,
  "actionIdBuy" bigint NOT NULL,
  "profileId" bigint NOT NULL,
  "numberStocks" int NOT NULL,
  "limit" "decimal(10, 2)" NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now())
);

CREATE TABLE "portfolio" (
  "id" bigserial PRIMARY KEY,
  "playerId" bigint NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now())
);

CREATE TABLE "portfolioActions" (
  "id" int PRIMARY KEY,
  "portfolioId" bigint NOT NULL,
  "actionId" bigint NOT NULL,
  "playerId" bigint NOT NULL,
  "quantity" int NOT NULL,
  "purchasePrice" "decimal(10, 2)" NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "buy" ("profileId");

CREATE INDEX ON "buy" ("actionIdBuy");

CREATE UNIQUE INDEX ON "buy" ("profileId", "numberStocks");

CREATE INDEX ON "portfolio" ("playerId");

CREATE INDEX ON "portfolioActions" ("portfolioId", "actionId");

ALTER TABLE "buy" ADD FOREIGN KEY ("actionIdBuy") REFERENCES "actions" ("id");

ALTER TABLE "buy" ADD FOREIGN KEY ("profileId") REFERENCES "players" ("id");

ALTER TABLE "portfolio" ADD FOREIGN KEY ("playerId") REFERENCES "players" ("id");

ALTER TABLE "portfolioActions" ADD FOREIGN KEY ("portfolioId") REFERENCES "portfolio" ("id");

ALTER TABLE "portfolioActions" ADD FOREIGN KEY ("actionId") REFERENCES "actions" ("id");

ALTER TABLE "portfolioActions" ADD FOREIGN KEY ("playerId") REFERENCES "players" ("id");
