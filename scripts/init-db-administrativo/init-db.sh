#!/usr/bin/env bash

SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

NEW_USER='Administrativo'
USER_PASSWORD='Administrativo'
DB_NAME='Administrativo' 
UUID='UUID' 
STRING='VARCHAR(255)'
TEXT='TEXT'
CPF='VARCHAR(11)'
DATE='DATE'
TIME='TIME'
SENHA='VARCHAR(255)'
ID='VARCHAR(11)'

SCRIPT=$(eval "echo \"$(cat $SCRIPT_DIR/template* | sed 's/`/***/g')\"" | sed 's/\*\*\*/`/g')
DB_NAME='Teste' 
NEW_USER='Teste'
USER_PASSWORD='Teste'
SCRIPT+=$(eval "echo -e \"\n$(cat $SCRIPT_DIR/template* | sed 's/`/***/g')\"" | sed 's/\*\*\*/`/g')

echo "$SCRIPT" | mysql --user root --password=$MARIADB_ROOT_PASSWORD
