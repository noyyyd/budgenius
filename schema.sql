CREATE TABLE "category" (
  "id"  INTEGER NOT NULL,
  "name"  TEXT NOT NULL,
  "type"  TEXT NOT NULL CHECK("type" IN ('income', 'expenses')),
  PRIMARY KEY("id" AUTOINCREMENT)
);

INSERT INTO "category" ("name", "type") VALUES
("Ипотека", "expenses"),
("ЖКХ", "expenses"),
("Зарплата", "income");

CREATE TABLE "budget" (
  "id"  INTEGER NOT NULL,
  "name" TEXT NOT NULL,
  "start"  INTEGER NOT NULL,
  "end"  INTEGER NOT NULL,
  PRIMARY KEY("id")
);

INSERT INTO "budget" ("name", "start", "end") VALUES 
("Июль", 1782853200, 1785531600),
("Август", 1785531600, 1788210000);

CREATE TABLE "budget_category" (
  "budget_id"  INTEGER NOT NULL,
  "amount"  INTEGER NOT NULL CHECK("amount" > 0),
  "category_id"  INTEGER NOT NULL,
  PRIMARY KEY("budget_id","category_id"),
  FOREIGN KEY("category_id") REFERENCES "category"("id") ON DELETE NO ACTION ON UPDATE NO ACTION
);

INSERT INTO "budget_category" ("budget_id", "amount", "category_id") VALUES
(1, 1782853200, 1785531600),
("Август", 1785531600, 1788210000);

CREATE TABLE "transaction" (
  "id"  INTEGER NOT NULL,
  "ts"  INTEGER NOT NULL,
  "amount"  INTEGER NOT NULL CHECK("amount" > 0),
  "comment"  TEXT,
  "category_id"  INTEGER NOT NULL,
  PRIMARY KEY("id" AUTOINCREMENT),
  FOREIGN KEY("category_id") REFERENCES "category"("id") ON DELETE NO ACTION ON UPDATE NO ACTION
);