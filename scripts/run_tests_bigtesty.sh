#!/usr/bin/env bash

set -e
set -o pipefail
set -u

echo "#######Running tests with BigTesty"

docker run -it \
  -e PROJECT_ID="$PROJECT_ID" \
  -e LOCATION="$LOCATION" \
  -e SA_EMAIL="$SA_EMAIL" \
  -e IAC_BACKEND_URL="$IAC_BACKEND_URL" \
  -e ROOT_TEST_FOLDER="$ROOT_TEST_FOLDER" \
  -v tests:/app/tests \
  -v tests/tables:/app/bigtesty/infra/resource/tables \
  -v /var/run/docker.sock:/var/run/docker.sock \
  bigtesty
