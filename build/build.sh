#!/bin/bash

BASEDIR="$(dirname "$(readlink -fm "$0")")"
SRCDIR="$(dirname "$BASEDIR")"
echo $SRCDIR

env GOOS=linux GOARCH=amd64 go build -o $SRCDIR/bin/task_manager $SRCDIR/cmd/*.go