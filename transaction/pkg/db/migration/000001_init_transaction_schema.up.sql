CREATE TABLE transactions (
  "id" varchar(100) PRIMARY KEY,
  "user_id" varchar(100) NOT NULL,
  "title" varchar(50) NOT NULL,
  "description" varchar(255),
  "amount" int   NOT NULL,
  "date" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT(now()),
  "updated_at" timestamptz NOT NULL DEFAULT(now())
);

CREATE INDEX idx_transactions_title ON transactions("title");
CREATE INDEX idx_transactions_date ON transactions("date");