CREATE TABLE "purchaseSchedule" (
  "id" bigserial PRIMARY KEY,
  "buyId" bigint NOT NULL,
  "stage" varchar(255) NOT NULL DEFAULT 'waiting',
  "created_order_buy" timestamp NOT NULL DEFAULT '0001-01-01 00:00:00',
   "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "purchaseSchedule" ADD FOREIGN KEY ("buyId") REFERENCES "buy" ("id");