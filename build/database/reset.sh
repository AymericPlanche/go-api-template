#!/bin/bash

DIR=`dirname "$0"`
URL="postgres://postgres:postgres@db:5432"

psql -c "DROP DATABASE IF EXISTS $1" -c "CREATE DATABASE $1" $URL
psql -f "$DIR/schema.sql" "$URL/$1"