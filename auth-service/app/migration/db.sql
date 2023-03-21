CREATE TABLE "users"
(
    "id"           SERIAL PRIMARY KEY,
    "username"     VARCHAR(255) UNIQUE NOT NULL,
    "password"     VARCHAR(255)        NOT NULL,
    "email"        VARCHAR(255) UNIQUE NOT NULL,
    "phone_number" VARCHAR(11) UNIQUE  NOT NULL,
    "full_name"    VARCHAR(255),
    "created_at"   timestamptz         NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    "updated_at"   timestamptz         NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    "deleted_at"   timestamptz
);

CREATE TABLE "roles"
(
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(255) UNIQUE NOT NULL
);


CREATE TABLE "user_roles"
(
    "user_id" SERIAL PRIMARY KEY,
    "role_id" INT,
    UNIQUE (user_id,role_id)
);