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
	-v $(PWD)/conf/config.yml:/etc/cgmh/config.yml \
	--name cgmh-apiserver \
	registry.gitlab.com/inwinstack/cgmh/apiserver:v0.2.1 \
	  --config /etc/cgmh/config.yml
