#!/usr/bin/env bash

SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

$SCRIPT_DIR/scripts/init-db-administrativo/init-db.sh

export $(grep -v '^#' env-docker-compose | xargs)

docker exec -i db-administrativo mysql --user root --password=$DB_ADMINISTRATIVO_ROOT_PASSWORD < $SCRIPT_DIR/scripts/init-db-administrativo/init.sql

unset $(grep -v '^#' env-docker-compose | sed -E 's/(.*)=.*/\1/' | xargs -d '\n')

