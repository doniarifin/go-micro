CREATE TABLE IF NOT EXISTS orders (
  id VARCHAR(100) NOT NULL,
  product_id VARCHAR(100) NOT NULL,
  product_name VARCHAR(100),
  order_by VARCHAR(100) NOT NULL,
  date date,
  created_at date,
  PRIMARY KEY (id)
);