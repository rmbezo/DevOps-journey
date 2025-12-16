CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT
);
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    product TEXT
);
