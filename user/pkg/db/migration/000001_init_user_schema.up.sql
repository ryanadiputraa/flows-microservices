CREATE TABLE users (
  id varchar(100) PRIMARY KEY,
  first_name varchar(100) NOT NULL,
  last_name varchar(100) NOT NULL,
  email varchar(100) UNIQUE NOT NULL,
  picture varchar(200),
  currency varchar(20) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  is_verified boolean NOT NULL
);