CREATE TABLE IF NOT EXISTS messages (
  id VARCHAR(100) NOT NULL,
  sender VARCHAR(100),
  receiver VARCHAR(100),
  msg_type VARCHAR(50),
  msg_body VARCHAR(100),
  created_by VARCHAR(100),
  created_at date,
  update_at date,
  PRIMARY KEY (id)
);