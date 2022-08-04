#!/bin/bash

echo 'Init script running. Creating DB: '"${DB_NAME}"

psql --username "${POSTGRES_USER}" <<EOF
  CREATE USER ${DB_USER} WITH PASSWORD '${DB_PASS}';
  CREATE DATABASE ${DB_NAME};
  GRANT ALL PRIVILEGES ON DATABASE ${DB_NAME} TO ${DB_USER};
EOF