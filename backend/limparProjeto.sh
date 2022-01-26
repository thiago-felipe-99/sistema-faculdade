#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

find $SCRIPT_DIR -type f -regex ".*\(coverage.out\|.log\)" -exec rm {} +
