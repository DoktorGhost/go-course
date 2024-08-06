#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: gofmt.sh <filename>"
  exit 1
fi

FILE="$1"

go fmt "$FILE"