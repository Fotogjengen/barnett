#!/usr/bin/env bash

docker exec -it my_postgres psql -U postgres -c "create database barnett_db"