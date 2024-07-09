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

CREATE TABLE IF NOT EXISTS "user"(
  id uuid primary key default uuid_generate_v4(),
  created_at timestamp default now(),
  updated_at timestamp default now(),
  deleted_at timestamp null,
  fullname varchar not null,
  email varchar not null,
  password varchar null,
  phone varchar(20) null,
  token_verify text null,
  address text null,
  is_active boolean not null default false,
  is_blocked boolean not null default false,
  "role_id" uuid not null,
  foreign key ("role_id") references "role"(id)
);