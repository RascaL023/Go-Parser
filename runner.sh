#!/usr/bin/env bash

BINARY="./cmd/bin/parser"
SRC="./cmd/parser"
mode=$1
opt=$2

clear

case "$mode" in
  run-bin)
    echo "> Running.."
    "$BINARY" "$opt"
    echo "> Done."
    ;;  
  run-raw)
    echo "> Running.."
    go run "$SRC" "$opt"
    echo "> Done."
    ;;  
  build-bin)
    echo "> Building..."
    go build -o "$BINARY" "$SRC"
    echo "> Done."
    ;;  
  build-run)
    echo "> Building..."
    go build -o "$BINARY" "$SRC"
    echo "> Running binary"
    "./$BINARY" "$opt"
    echo "> Done."
    ;;
  *)
    echo "Invalid mode!"
    ;;
esac
