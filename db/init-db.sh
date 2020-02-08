#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 <<-EOSQL
  DROP DATABASE IF EXISTS goplants;
  CREATE DATABASE goplants;
EOSQL

psql -v ON_ERROR_STOP=1 --dbname goplants <<-EOSQL
  CREATE TABLE plants(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100)
  );
EOSQL