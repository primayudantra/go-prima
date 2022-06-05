-- Create transfers table
CREATE TABLE IF NOT EXISTS transfers (
    id BIGSERIAL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    transaction_id TEXT NOT NULL,
    transferer_user_id TEXT NOT NULL,
    transferrer_account_user_id TEXT NOT NULL,
    receiver_user_id TEXT NOT NULL,
    receiver_account_user_id TEXT NOT NULL,
    amount TEXT NOT NULL,
    currency TEXT NOT NULL,
    status TEXT NOT NULL
);
