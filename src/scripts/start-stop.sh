#!/bin/bash
set -eu
BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $BASEDIR

export PIDFILE=var/emoji-cloner.pid
export LOGFILE=var/emoji-cloner.log

if [ ! -e $PIDFILE ]; then
	echo "START"
	bin/emoji-cloner.sh &
	exit 0
fi

echo "STOP"
kill -9 $(cat $PIDFILE) || true
rm -f $PIDFILE

