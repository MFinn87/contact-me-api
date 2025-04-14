#!/bin/bash

docker-compose -f ./docker-compose-local.yml down
docker-compose -f ./docker-compose-local.yml rm -fsv
docker volume prune -f
docker system prune --volumes
docker image prune -a