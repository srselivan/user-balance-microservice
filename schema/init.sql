CREATE TABLE users (
  	id bigserial NOT NULL,
	balance float8,
  	PRIMARY KEY(id)
);

CREATE TABLE reserve (
  	user_id bigint REFERENCES users(id) NOT NULL,
	order_id bigint NOT NULL,
	service_id bigint NOT NULL,
    amount float8
);

CREATE TABLE deals (
  	user_id bigint REFERENCES users(id) NOT NULL,
	order_id bigint NOT NULL,
	service_id bigint NOT NULL,
    amount float8,
    date date
);

INSERT INTO users
  (balance)
VALUES
  (56.99),
  (63.99),
  (17.99),
  (34.98);