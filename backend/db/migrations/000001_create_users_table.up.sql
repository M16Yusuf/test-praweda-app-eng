CREATE TABLE IF NOT EXISTS users_task2 (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);