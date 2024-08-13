#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: debug.sh <filename>"
  exit 1

fi

SOURCE_FILE="$1"
BINARY_NAME="myprogram"

echo "Debug started..."

go build -o "$BINARY_NAME" "$SOURCE_FILE"

dlv exec "$BINARY_NAME"

echo "Debug ended."