#!/usr/bin/env bash

SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

DB_ADMINISTRATIVO_FINAL_DIR=$SCRIPT_DIR/../data/administrativo/mariadb-init
DB_ADMINISTRATIVO_TEMPLATE=$SCRIPT_DIR/db-administrativo-template.sql

export DB_NAME='Administrativo' \
  UUID='UUID' \
  STRING='VARCHAR(255)' \
  TEXT='TEXT' \
  CPF='VARCHAR(11)' \
  DATE='DATE' \
  TIME='TIME' \
  SENHA='VARCHAR(255)' \
  ID='VARCHAR(11)'

if [[ ! -d $DB_ADMINISTRATIVO_FINAL_DIR ]]
then
  mkdir -p $DB_ADMINISTRATIVO_FINAL_DIR
fi

envsubst < $DB_ADMINISTRATIVO_TEMPLATE > $DB_ADMINISTRATIVO_FINAL_DIR/init.sql
