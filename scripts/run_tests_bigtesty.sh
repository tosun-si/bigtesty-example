#!/usr/bin/env bash

set -e
set -o pipefail
set -u

echo "#######Running tests with BigTesty"

docker run -it \
  -e PROJECT_ID="$PROJECT_ID" \
  -e LOCATION="$LOCATION" \
  -e TF_VAR_project_id="$PROJECT_ID" \
  -e TF_STATE_BUCKET="$TF_STATE_BUCKET" \
  -e TF_STATE_PREFIX="$TF_STATE_PREFIX" \
  -e GOOGLE_PROVIDER_VERSION="$GOOGLE_PROVIDER_VERSION" \
  -e ROOT_TEST_FOLDER="$ROOT_TEST_FOLDER" \
  -v tests:/app/tests \
  -v $HOME/.config/gcloud:/root/.config/gcloud \
  bigtesty
