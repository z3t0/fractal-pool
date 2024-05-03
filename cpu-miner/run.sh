#!/bin/bash
set -euxo pipefail

docker run --network host cpu-miner
