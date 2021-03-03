#!/bin/bash
# Bash script used to load environment variables of our development machine
# and then start running this app. When running on production machine please
# do not use the same values as it is a security risk!
export SERVER_API_DOMAIN=127.0.0.1; \
export SERVER_API_PROTOCOL=http; \
export SERVER_PORT=5000; \
export DATABASE_HOST=127.0.0.1; \
export DATABASE_PORT=5433; \
export DATABASE_PASSWORD=Jiby@1998; \
export DATABASE_USER=postgres; \
export DATABASE_NAME=pillshare_db;
exec make run
