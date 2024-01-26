CREATE TABLE "users"
(
  "id" bigserial PRIMARY KEY,
  "username" VARCHAR UNIQUE NOT NULL,
  "hash_password" VARCHAR NOT NULL,
  "full_name" VARCHAR NOT NULL,
  "email" VARCHAR UNIQUE NOT NULL,
  "password_change_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
  "created_at" timestamptz NOT NULL DEFAULT(NOW())
);


ALTER TABLE "cars" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
