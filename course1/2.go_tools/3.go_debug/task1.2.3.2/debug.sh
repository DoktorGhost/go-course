#!/bin/bash

SOURCE_FILE="$1"
BINARY_NAME="myprogram"

echo "Debug started..."

go build -o "$BINARY_NAME" "$SOURCE_FILE"

dlv exec "$BINARY_NAME"

echo "Debug ended."