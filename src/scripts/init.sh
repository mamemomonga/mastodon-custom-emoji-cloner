#!/bin/bash
set -eu
BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && cd ..&& pwd )"

cd $BASEDIR
mkdir -p var
mkdir -p etc

