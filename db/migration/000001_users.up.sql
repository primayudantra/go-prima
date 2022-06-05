-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    email TEXT NOT NULL,
    password_hash TEXT NOT NULL
);
