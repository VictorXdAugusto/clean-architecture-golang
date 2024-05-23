CREATE DATABASE gocrud;

\c gocrud;

-- Criação da tabela "users"
CREATE TABLE "users" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) UNIQUE NOT NULL,
    "username" VARCHAR(255) UNIQUE NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "age" INT NOT NULL,
    "is_active" BOOLEAN NOT NULL DEFAULT TRUE
);