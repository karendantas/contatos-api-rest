CREATE TABLE IF NOT EXISTS contact (
    contact_id SERIAL PRIMARY KEY,
    contact_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL
);

-- Dados de exemplo
INSERT INTO contact (contact_name, email) VALUES ('Karen', 'karen@gmail.com');
INSERT INTO contact (contact_name, email) VALUES ('Mirla', 'mirla@gmail.com');
