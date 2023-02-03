#!/bin/sh

set -e

while true;
do
    watchexec --restart --exts "go,html" --watch . "go run *.go"
    echo "restarting process"
done