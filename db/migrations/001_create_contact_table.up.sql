CREATE TABLE IF NOT EXISTS contact (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(15) NOT NULL
);

-- Dados de exemplo
INSERT INTO contact (name, email, phone) VALUES ('Karen', 'karen@gmail.com', '900000001');
INSERT INTO contact (name, email, phone) VALUES ('Mirla', 'mirla@gmail.com', '900000000');
