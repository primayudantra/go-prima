-- Create posts table
CREATE TABLE IF NOT EXISTS posts (
    id BIGSERIAL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    content TEXT NOT NULL
);
