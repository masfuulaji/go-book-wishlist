CREATE TABLE IF NOT EXISTS directors (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  birthday DATE,
  place_of_birth VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
