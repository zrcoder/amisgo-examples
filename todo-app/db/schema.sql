-- create table users
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL
);

-- create table todos
CREATE TABLE IF NOT EXISTS todos (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    detail TEXT,
    priority INTEGER DEFAULT 1,
    due_date DATE,
    is_completed INTEGER DEFAULT 0,
    user_id INTEGER,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- create trigger to update updated_at field when update todos
CREATE TRIGGER IF NOT EXISTS update_todos_updated_at AFTER
UPDATE ON todos FOR EACH ROW BEGIN
UPDATE todos
SET
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = NEW.id;

END;
