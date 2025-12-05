#!/bin/bash

# Check if folder name is provided
if [ -z "$1" ]; then
    echo "Usage: ./build.sh [folder_name]"
    echo "Example: ./build.sh demo"
    exit 1
fi

TARGET="$1"
BASE_DIR="$(cd "$(dirname "$0")" && pwd)"
CMD_DIR="$BASE_DIR/cmd/$TARGET"

# Check main.go exists
if [ ! -f "$CMD_DIR/main.go" ]; then
    echo "Error: $CMD_DIR/main.go not found."
    exit 1
fi

echo "Building $TARGET ..."

OUTPUT="$CMD_DIR/$TARGET"   # Linux/macOS 默认没扩展名
go build -o "$OUTPUT" "$CMD_DIR/main.go"

if [ $? -ne 0 ]; then
    echo "Build failed."
    exit 1
fi

echo "Build succeeded: $OUTPUT"
