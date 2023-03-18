CREATE TABLE IF NOT EXISTS todo_list
(
    id         SERIAL UNIQUE NOT NULL,
    title      TEXT          NOT NULL DEFAULT '',
    created_at TIMESTAMP     NOT NULL DEFAULT now(),
    updated_at TIMESTAMP     NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS task
(
    id        SERIAL UNIQUE,
    title     TEXT    NOT NULL DEFAULT '',
    list_id   INT REFERENCES todo_list (id) ON DELETE CASCADE,
    completed BOOLEAN NOT NULL DEFAULT 'false'
);


