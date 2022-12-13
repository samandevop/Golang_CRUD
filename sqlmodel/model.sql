

CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid(),
    first_name varchar(24) NOT NULL,
    last_name varchar(36) NOT NULL
);

INSERT INTO users (first_name, last_name) VALUES
('Samandar', 'Foziljonov');

