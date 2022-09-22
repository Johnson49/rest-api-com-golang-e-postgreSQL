#!/bin/bash

psql -U postgres <<-EOSQL

CREATE USER docker WITH PASSWORD 'docker';
    CREATE DATABASE dbgolang;
    
    GRANT ALL PRIVILEGES ON DATABASE dbgolang TO docker;

    \connect dbgolang docker

    BEGIN;

    CREATE TABLE livros(
            id   TEXT PRIMARY KEY NOT NULL,
            titulo      VARCHAR(255) NOT NULL,
            autor     VARCHAR(255) NOT NULL,
            genero     VARCHAR(255) NOT NULL,
            descricao     TEXT NOT NULL,
            preco      REAL NOT NULL
        );

    INSERT INTO livros (id, titulo, autor, genero, descricao, preco) values ('1', 'Um dia depois de amanhã','Ana','ação/drama','Após o fim da guerra...', 10.99);

    COMMIT;

    select * from livros;

EOSQL