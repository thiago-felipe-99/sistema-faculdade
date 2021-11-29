#!/usr/bin/env bash

SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

$SCRIPT_DIR/scripts/init-db-administrativo/init-db.sh

export $(grep -v '^#' .env | xargs)

mysql --user root --password=$DB_ADMINISTRATIVO_ROOT_PASSWORD --port=$DB_ADMINISTRATIVO_PORT < $SCRIPT_DIR/scripts/init-db-administrativo/init.sql

unset $(grep -v '^#' .env | sed -E 's/(.*)=.*/\1/' | xargs -d '\n')

