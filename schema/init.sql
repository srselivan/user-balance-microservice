DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id            INT AUTO_INCREMENT NOT NULL,
  balance       DECIMAL(9,2) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO users
  (balance)
VALUES
  (56.99),
  (63.99),
  (17.99),
  (34.98);