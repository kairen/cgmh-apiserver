# CGMH API server
The CGMH API server services REST operations. This server also validates and configures data for the API objects which include user, forms and others.

## Building from Source
Clone repo into your go path under $GOPATH/src:
```sh
$ git clone https://gitlab.com/inwinstack/cgmh/apiserver.git $GOPATH/src/inwinstack/cgmh/apiserver
$ cd $GOPATH/src/inwinstack/cgmh/apiserver
$ dep ensure
$ make
```

## Setup dev environment
The following command will launch a MongoDB and run API server as debug mode:
```sh
$ docker run -d -p 27017:27017 \
    -e MONGO_INITDB_ROOT_USERNAME=root \
    -e MONGO_INITDB_ROOT_PASSWORD=passw0rd \
    -v $(pwd)/.db:/data/db \
    --name mgo \
    mongo:3.6

$ go run cmd/main.go --config $(PWD)/conf/config.yml
```