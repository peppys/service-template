CREATE TABLE todos (
  id SERIAL,
  text VARCHAR(255) NOT NULL,
  author VARCHAR(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id)
);
