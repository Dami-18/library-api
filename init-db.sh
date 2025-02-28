#!/bin/bash

[ ! -f .env ] || export $(grep -v '^#' .env | xargs)

mysql -h"$DB_HOST" -u"$DB_USER" -p"$DB_PASSWORD" -e "CREATE DATABASE IF NOT EXISTS $DB_NAME;"

echo "Database $DB_NAME created"