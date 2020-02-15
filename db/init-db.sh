#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 <<-EOSQL
  DROP DATABASE IF EXISTS goplants;
  CREATE DATABASE goplants;
EOSQL

psql -v ON_ERROR_STOP=1 --dbname goplants <<-EOSQL
  CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created TIMESTAMP NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,

    CONSTRAINT users_uc_email UNIQUE (email)
  );

  CREATE TABLE plants(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    owner INTEGER,
    FOREIGN KEY (owner) REFERENCES users (id)
  );
EOSQL
