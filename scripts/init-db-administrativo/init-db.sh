#!/usr/bin/env bash

SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

export NEW_USER='Administrativo' \
       USER_PASSWORD='Administrativo' \
       DB_NAME='Administrativo' \
       UUID='UUID' \
       STRING='VARCHAR(255)' \
       TEXT='TEXT' \
       CPF='VARCHAR(11)' \
       DATE='DATE' \
       TIME='TIME' \
       SENHA='VARCHAR(255)' \
       ID='VARCHAR(11)'

envsubst < $SCRIPT_DIR/template.sql > $SCRIPT_DIR/init.sql

export DB_NAME='Teste' \
       NEW_USER='Teste' \
       USER_PASSWORD='Teste'

envsubst <  $SCRIPT_DIR/template.sql >> $SCRIPT_DIR/init.sql


