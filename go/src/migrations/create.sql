CREATE TABLE IF NOT EXISTS sessions (
    username VARCHAR(255) NOT NULL,
    token VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    PRIMARY KEY (username)
);

INSERT INTO sessions (username, token, expires_at)
VALUES (?, ?, ?)
ON CONFLICT (username) DO UPDATE
SET username = ?, token = ?, expires_at = ?
