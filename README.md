# bigtesty-example

This project shows examples with BigQuery integration tests launched with BigTesty.\
BigTesty is a framework that allows to create Integration Tests with BigQuery on a real and short-lived Infrastructure.

## Run integration tests with BigTesty

```bash
docker run -it \
    -e GOOGLE_PROJECT=$PROJECT_ID \
    -e SA_EMAIL=$SA_EMAIL \
    -e GOOGLE_REGION=$LOCATION \
    -e PULUMI_BACKEND_URL=$IAC_BACKEND_URL \
    -e ROOT_TEST_FOLDER=$ROOT_TEST_FOLDER \
    -e ROOT_TABLES_FOLDER="$ROOT_TABLES_FOLDER" \
    -e TABLES_CONFIG_FILE_PATH="$TABLES_CONFIG_FILE_PATH" \
    -e BIGTESTY_STACK_NAME=bigtesty \
    -e PULUMI_CONFIG_PASSPHRASE=gcp_fake_passphrase \
    -v $(pwd)/tests:/opt/bigtesty/tests \
    -v $(pwd)/tests/tables:/opt/bigtesty/tests/tables \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v $HOME/.config/gcloud:/opt/bigtesty/.config/gcloud
```

## Run integration tests with Cloud Build

```bash
gcloud builds submit \
    --project=$PROJECT_ID \
    --region=$LOCATION \
    --config run-tests-bigtesty.yaml \
    --substitutions _IAC_BACKEND_URL=$IAC_BACKEND_URL \
    --verbosity="debug" .
```

```bash
gcloud beta builds triggers create manual \
    --project=$PROJECT_ID \
    --region=$LOCATION \
    --name="run-tests-bigtesty" \
    --repo="https://github.com/tosun-si/bigtesty-example" \
    --repo-type="GITHUB" \
    --branch="main" \
    --build-config="run-tests-bigtesty.yaml" \
    --substitutions _SA_EMAIL=$SA_EMAIL,_IAC_BACKEND_URL=$IAC_BACKEND_URL,_ROOT_TEST_FOLDER=$ROOT_TEST_FOLDER \
    --verbosity="debug"
```