BEGIN;

CREATE TABLE IF NOT EXISTS accounts (
  id SERIAL PRIMARY KEY,
  document_number INT NOT NULL
);

CREATE TABLE IF NOT EXISTS operation_types (
  id SERIAL PRIMARY KEY,
  description VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
  id SERIAL PRIMARY KEY,
  account_id INT NOT NULL,
  operation_type_id INT NOT NULL,
  amount DECIMAL NOT NULL,
  event_date TIMESTAMP default (now() at time zone 'utc') NOT NULL
);

ALTER TABLE transactions
  ADD CONSTRAINT fk_account_id FOREIGN KEY (account_id) REFERENCES accounts (id),
  ADD CONSTRAINT fk_operation_type_id FOREIGN KEY (operation_type_id) REFERENCES operation_types (id);

COMMIT;