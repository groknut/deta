CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email TEXT UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (name, email) 
VALUES 
    ('Alice Johnson', 'alice@com'),
    ('Bob Mil', 'Bob@com'),
    ('Max', 'Max@com');