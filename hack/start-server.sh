#!/bin/bash
#
# Start server and mongodb.
#

set -eu

if [[ $(docker network ls | grep -o "test" | wc -l) -eq 0  ]]; then
  docker network create test
fi

# start mongodb
docker run -d -p 27017:27017 \
	  --network test \
	  -e MONGO_INITDB_ROOT_USERNAME=root \
	  -e MONGO_INITDB_ROOT_PASSWORD=passw0rd \
	  -v $(pwd)/.db:/data/db \
	  --name mgo mongo:3.6 

# start FR server
docker run -d -p 8080:8080 \
	--network test \
	-e MONGODB_HOST="mgo.test:27017" \
	-e MONGODB_SOURCE=admin \
	-e MONGODB_USER=root \
	-e MONGODB_PASSWORD=passw0rd \
	-e MONGODB_DB=CGMH \
	--name cgmh-forms \
	inwinstack/cgmh-forms:v0.1.0
