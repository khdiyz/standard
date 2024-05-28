-- +goose Up
CREATE TABLE IF NOT EXISTS "users" (
    "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "full_name" VARCHAR(64) NOT NULL,
    "email" VARCHAR(64) NOT NULL,
    "password_hash" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NULL,
    "deleted_at" TIMESTAMP NULL
);

-- +goose Down
DROP TABLE IF EXISTS "users";
