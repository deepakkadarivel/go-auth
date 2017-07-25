
-- +goose Up
CREATE TABLE IF NOT EXISTS account
(
  user_id SERIAL PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(50) NOT NULL,
  password_hash VARCHAR(50),
  password_hash_algorithm VARCHAR(50),
  created_on TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  last_login TIMESTAMP WITHOUT TIME ZONE

);


-- +goose Down
DROP TABLE IF EXISTS account CASCADE;

