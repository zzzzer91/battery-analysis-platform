#!/bin/bash

ABSOLUTE_PROJECT_PATH="$(cd "$(dirname $0)/..";pwd)"
cd $ABSOLUTE_PROJECT_PATH

docker-compose -f docker-compose.release.yml stop