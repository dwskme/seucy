CREATE TABLE users (
    uuid      VARCHAR(36) PRIMARY KEY,
    firstname VARCHAR(255),
    lastname  VARCHAR(255),
    email     VARCHAR(255) UNIQUE,
    username  VARCHAR(255) UNIQUE,
    password  VARCHAR(255),
    role      VARCHAR(20)
);
