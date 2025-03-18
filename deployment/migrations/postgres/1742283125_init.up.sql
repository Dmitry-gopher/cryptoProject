BEGIN;
CREATE TABLE coins (
    title VARCHAR(255) NOT NULL,
    current_rate REAL NOT NULL,
    timestamp TIMESTAMP NOT NULL DEFAULT NOW()
);
END;
