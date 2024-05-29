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

INSERT INTO "users" (
    "full_name","email","password_hash"
) VALUES (
    'Super Admin', 
    'admin@gmail.com', 
    '686a717268343631376149415337333951576a6668616a73d033e22ae348aeb5660fc2140aec35850c4da997'
);

-- +goose Down
DROP TABLE IF EXISTS "users";