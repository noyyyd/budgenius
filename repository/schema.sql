CREATE TABLE IF NOT EXISTS "category" (
  "id"  INTEGER NOT NULL,
  "name"  TEXT NOT NULL UNIQUE,
  "type"  TEXT NOT NULL CHECK("type" IN ('income', 'expenses')),
  PRIMARY KEY("id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "budget" (
  "id"  INTEGER NOT NULL,
  "name" TEXT NOT NULL,
  "start"  INTEGER NOT NULL,
  "end"  INTEGER NOT NULL,
  PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "budget_category" (
  "budget_id"  INTEGER NOT NULL,
  "amount"  INTEGER NOT NULL CHECK("amount" > 0),
  "category_id"  INTEGER NOT NULL,
  PRIMARY KEY("budget_id","category_id"),
  FOREIGN KEY("category_id") REFERENCES "category"("id") ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS "money_transaction" (
  "id"  INTEGER NOT NULL,
  "ts"  INTEGER NOT NULL,
  "amount"  INTEGER NOT NULL CHECK("amount" > 0),
  "comment"  TEXT,
  "category_id"  INTEGER NOT NULL,
  PRIMARY KEY("id" AUTOINCREMENT),
  FOREIGN KEY("category_id") REFERENCES "category"("id") ON DELETE NO ACTION ON UPDATE NO ACTION
);