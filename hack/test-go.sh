#!/bin/bash
#
# Run go test.
#

set -e

REPO_PATH="inwinstack/cgmh/apiserver"

if [[ -f $(which python 2>&1) ]]; then
	PYTHON="python"
else
	PYTHON="docker run --rm -it -v $(pwd):/tmp -w /tmp python python"
fi

COV_FILE=coverage.txt
COV_TMP_FILE=coverage_tmp.txt

echo "Running go tests..."
cd ${GOPATH}/src/${REPO_PATH}
rm -f out/$COV_FILE || true
echo "mode: count" > out/$COV_FILE
for pkg in $(go list -f '{{ if .TestGoFiles }} {{.ImportPath}} {{end}}' ./...); do
    go test -tags "container_image_ostree_stub containers_image_openpgp" -v $pkg -covermode=count -coverprofile=out/$COV_TMP_FILE
    tail -n +2 out/$COV_TMP_FILE >> out/$COV_FILE || (echo "Unable to append coverage for $pkg" && exit 1)
done
rm -rf out/$COV_TMP_FILE

ignore="vendor\|\_gopath\|assets.go\|\_test.go\|out\|api"

echo "Checking gofmt..."
set +e
files=$(gofmt -l -s . | grep -v ${ignore})
set -e
if [[ $files ]]; then
  echo "Gofmt errors in files: $files"
  exit 1
fi
