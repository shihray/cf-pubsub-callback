#!/bin/bash
set -x

source .cloudbuild/script/loadenv.sh

GITHUB_TOKEN=${1}

go env -w GOPRIVATE="github.com/lctech-tw,github.com/lctech-lucy,github.com/jkforum"
echo "machine github.com login irir password ${GITHUB_TOKEN}" >> ~/.netrc

echo "go mod vendor"
go mod vendor

rm -Rf ~/.netrc
