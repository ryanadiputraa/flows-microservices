CREATE TABLE users (
  id varchar(100) PRIMARY KEY,
  email varchar(100) UNIQUE NOT NULL,
  password varchar(100) NOT NULL,
  first_name varchar(100) NOT NULL,
  last_name varchar(100) NOT NULL,
  currency varchar(20) NOT NULL,
  picture varchar(200),
  created_at timestamptz NOT NULL DEFAULT NOW()
);