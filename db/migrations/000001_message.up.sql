CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    message VARCHAR(5000) NOT NULL,
    status VARCHAR(100) NOT NULL,
    processed_time VARCHAR(100),
    date_create TIMESTAMP    
);

CREATE INDEX idx_message ON messages(message);
CREATE INDEX idx_status ON messages(status);
CREATE INDEX idx_processed_time ON messages(processed_time);
CREATE INDEX idx_date_create ON messages(date_create);
