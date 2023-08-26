CREATE TABLE IF NOT EXISTS users
(
    id       serial PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    password varchar(255)        NOT NULL
);

CREATE TABLE IF NOT EXISTS notes
(
    id      serial PRIMARY KEY,
    title   varchar(255) UNIQUE                             NOT NULL,
    info    text,
    user_id integer REFERENCES users (id) ON DELETE CASCADE NOT NULL
);