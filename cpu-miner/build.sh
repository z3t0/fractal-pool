#!/bin/bash
set -euxo pipefail

export DOCKER_BUILDKIT=1

SCRIPTPATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"
PATH_TO_PACKAGE=$SCRIPTPATH
PATH_TO_REPO=$(dirname $PATH_TO_PACKAGE)

cd $PATH_TO_PACKAGE

docker build . -t cpu-miner