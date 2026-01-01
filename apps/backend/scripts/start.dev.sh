#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
BACKEND_DIR="$(dirname "$SCRIPT_DIR")"

docker stop where_light_api_container
docker rm where_light_api_container
docker build --file "$BACKEND_DIR/Dockerfile.dev" --tag where_light_api_image "$BACKEND_DIR"
docker run --publish 8080:8080 --name where_light_api_container \
  --mount type=bind,source="$BACKEND_DIR",target=/app \
  --env-file "$BACKEND_DIR/.env.development" \
  where_light_api_image
