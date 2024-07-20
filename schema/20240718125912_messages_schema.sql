-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS messages (
    id uuid NOT NULL,
    message TEXT NOT NULL,
    marked BOOLEAN DEFAULT FALSE,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END
$$ LANGUAGE 'plpgsql';

CREATE TRIGGER update_updated_at_column
BEFORE UPDATE ON messages
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS messages;

DROP FUNCTION update_updated_at_column
-- +goose StatementEnd

-- goose -dir schema postgres 'postgresql://Messaggio:somestrongpassword@localhost:5432/db_messages' down
