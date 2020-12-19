CREATE TABLE IF NOT EXISTS products (
  id INT PRIMARY KEY AUTO_INCREMENT,
  shopid BIGINT,
  itemid BIGINT,
  price_max BIGINT,
  price_min BIGINT,
  name VARCHAR(255),
  images VARCHAR(2048),
  historical_sold INT,
  rating VARCHAR(255)
);
