CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
SELECT * FROM pg_timezone_names;
ALTER DATABASE "dev_dbgofiber" SET timezone TO "ASIA/JAKARTA";

CREATE TABLE IF NOT EXISTS "role"(
  id uuid primary key default uuid_generate_v4(),
  created_at timestamp default now(),
  updated_at timestamp default now(),
  deleted_at timestamp null,
  name varchar not null
);

CREATE INDEX idx_role_name ON "role" (name);

CREATE TABLE IF NOT EXISTS "user"(
  id uuid primary key default uuid_generate_v4(),
  created_at timestamp default now(),
  updated_at timestamp default now(),
  deleted_at timestamp null,
  fullname varchar not null,
  email varchar not null,
  password varchar not null,
  phone varchar(20) null,
  token_verify text null,
  address text null,
  is_active boolean not null default false,
  is_blocked boolean not null default false,
  "role_id" uuid not null,
  foreign key ("role_id") references "role"(id),
  unique(email)
);

CREATE INDEX idx_user_fullname ON "user" (fullname);
CREATE INDEX idx_user_email ON "user" (email);
CREATE INDEX idx_user_token_verify ON "user" (token_verify);
CREATE INDEX idx_user_is_active ON "user" (is_active);
CREATE INDEX idx_user_is_blocked ON "user" (is_blocked);
CREATE INDEX idx_user_role_id ON "user" (role_id);
