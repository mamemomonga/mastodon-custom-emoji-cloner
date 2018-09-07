#!/bin/bash
set -eu
BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && cd ..&& pwd )"
echo "$$" > $PIDFILE
bin/emoji-cloner etc/config.yaml > $LOGFILE 2>&1

