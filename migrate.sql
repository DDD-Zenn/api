CREATE TABLE IF NOT EXISTS users (
    uid VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO users (uid, name) VALUES ('1', 'Alice');
INSERT INTO users (uid, name) VALUES ('2', 'Bob');
INSERT INTO users (uid, name) VALUES ('3', 'Charlie');