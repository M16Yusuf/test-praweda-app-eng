CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nama VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    telepon VARCHAR(15)
);

-- INSERT INTO users (id, nama, email, telepon) VALUES
-- ('b899acb0-ec0e-427b-992c-f03924397bc0', 'Imron', NULL, 081234567890),
-- ('be5ae506-2461-4004-80f4-94dd5b0b7634','Juli', 'sammy@mail.com', '0987654321' ),
-- (NULL, 'Gajah Mada', NULL, NULL);