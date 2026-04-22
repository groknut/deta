CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email TEXT UNIQUE
);

INSERT INTO users (name, email) 
VALUES 
    ('Alice Johnson', 'alice@com');