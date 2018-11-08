#!/bin/bash
#
# Convert swagger spec to go bindata.
# 

set -eu

which go-bindata || GOBIN=$(GOPATH)/bin go get github.com/jteeuwen/go-bindata/...
go-bindata -pkg swagger -o ./api/swagger_gen.go ./api/swagger-spec