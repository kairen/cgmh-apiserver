#!/bin/bash
#
# Start server and mongodb.
#

set -eu

if [[ $(docker network ls | grep -o "test" | wc -l) -eq 0  ]]; then
  docker network create test
fi

# Start mongodb
docker run -d -p 27017:27017 \
	  --network test \
	  -e MONGO_INITDB_ROOT_USERNAME=root \
	  -e MONGO_INITDB_ROOT_PASSWORD=passw0rd \
	  -v $(pwd)/.db:/data/db \
	  --name mgo mongo:3.6 

# Start API server
docker run -d -p 8080:8080 \
	--network test \
	-e INIT_ADMIN_EMAIL="admin@inwinstack.com" \
	-e INIT_ADMIN_PASSWORD="r00tme" \
	--name cgmh-apiserver \
	registry.gitlab.com/inwinstack/cgmh/apiserver:v0.1.0 \
	  --db-host=mgo.test:27017 \
	  --db-password=passw0rd \
	  --init
