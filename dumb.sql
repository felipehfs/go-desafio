CREATE DATABASE godesafio;
CREATE TABLE users ( 
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    cpf VARCHAR(50) NOT NULL,
    lastName VARCHAR(50) NOT NULL,
    password TEXT NOT NULL,
    avatarURL VARCHAR(50),
    UUID uuid,
    datastart DATE default CURRENT_DATE
);
