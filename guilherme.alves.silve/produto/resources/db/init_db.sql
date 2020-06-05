CREATE DATABASE alura;

CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255),
    descricao VARCHAR(255),
    preco DECIMAL,
    quantidade INTEGER
);

INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ('Camisa 1', 'Descrição 1', 50.5, 13);
INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ('Caneta 2', 'Descrição 2', 20, 542);
INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ('Bola de Futebol 3', 'Descrição 3', 15.95, 10);
INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ('Laranja 4', 'Descrição 4', 1.99, 18);
