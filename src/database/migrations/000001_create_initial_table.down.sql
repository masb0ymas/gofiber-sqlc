DROP EXTENSION IF EXISTS "uuid-ossp";

DROP INDEX IF EXISTS idx_role_name;
DROP INDEX IF EXISTS idx_user_fullname, idx_user_email, idx_user_token_verify, idx_user_is_active, idx_user_is_blocked, idx_user_role_id;

DROP TABLE IF EXISTS "role";
DROP TABLE IF EXISTS "user";
