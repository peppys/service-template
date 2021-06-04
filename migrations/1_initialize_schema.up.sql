CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id uuid DEFAULT uuid_generate_v4(),
  email text NOT NULL,
  email_verified bool NOT NULL DEFAULT false,
  username text NOT NULL,
  password_hash text NOT NULL,
  given_name text NOT NULL,
  family_name text NOT NULL,
  name text NOT NULL,
  nickname text NULL DEFAULT NULL,
  picture text NULL DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY(id),
  UNIQUE(email),
  UNIQUE(username)
);

CREATE TABLE refresh_tokens (
  id uuid DEFAULT uuid_generate_v4(),
  token_hash text NOT NULL,
  user_id uuid NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  expires_at timestamp NOT NULL,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY(id),
  UNIQUE(token_hash),
  UNIQUE(user_id, deleted_at),
  CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE todos (
  id uuid DEFAULT uuid_generate_v4(),
  text text,
  user_id uuid NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY(id),
  CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id)
);
