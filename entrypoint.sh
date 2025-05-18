#!/bin/sh
set -e

go mod download

exec "$@"
