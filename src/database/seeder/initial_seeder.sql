INSERT INTO public."role" (id,created_at,updated_at,deleted_at,"name") VALUES
	 ('03ba326e-f9ed-410a-818f-eaa409c13622',now(),now(),NULL,'Super Admin'),
	 ('9dc8b32b-aefe-44d3-bf19-6dc088d13174',now(),now(),NULL,'Admin'),
	 ('d7efa7e9-3c97-4217-a6bd-59e2eba53068',now(),now(),NULL,'User'),
	 ('be8482c9-7410-45eb-8c28-4dfd508a0de6',now(),now(),NULL,'Guest');

INSERT INTO public."user" (id,created_at,updated_at,deleted_at,fullname,email,"password",phone,token_verify,address,is_active,is_blocked,role_id) VALUES
	 (uuid_generate_v4(),now(),now(),NULL,'Super Admin','super.admin@example.com','$argon2id$v=19$m=65536,t=3,p=2$hXwlaW+1NCwqKWDySLUk4g$ftx5ZLF5QjKLi50RW6qxPKZVDAPOvs6DxCY0L+GZz6A',NULL,NULL,NULL,true,false,'03ba326e-f9ed-410a-818f-eaa409c13622'),
	 (uuid_generate_v4(),now(),now(),NULL,'Admin','admin@example.com','$argon2id$v=19$m=65536,t=3,p=2$ssShjR+1zMucGwSWI1p7rw$vTHTwnKQejOrxC4SlirCsJ7NfA1IC9pHonRAzBqKOUA',NULL,NULL,NULL,true,false,'9dc8b32b-aefe-44d3-bf19-6dc088d13174'),
	 (uuid_generate_v4(),now(),now(),NULL,'User','user@example.com','$argon2id$v=19$m=65536,t=3,p=2$wnMuSBm5Fbw6mo5p4f3I6A$FzqhdZTYyklKziq506MM7cA2Cm7n4ud7GoSXMw6VVnc',NULL,NULL,NULL,true,false,'d7efa7e9-3c97-4217-a6bd-59e2eba53068');
