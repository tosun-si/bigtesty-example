# bigtesty-example

This project shows examples with BigQuery integration tests launched with BigTesty.\
BigTesty is a framework that allows to create Integration Tests with BigQuery on a real and short-lived Infrastructure.

## Run integration tests with BigTesty

```bash
 docker run -it \
    -e PROJECT_ID=$PROJECT_ID \
    -e SA_EMAIL=$SA_EMAIL \
    -e LOCATION=$LOCATION \
    -e IAC_BACKEND_URL=$IAC_BACKEND_URL \
    -e ROOT_TEST_FOLDER=$ROOT_TEST_FOLDER \
    -v $(pwd)/tests:/app/tests \
    -v $(pwd)/tests/tables:/app/bigtesty/infra/resource/tables \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v $HOME/.config/gcloud:/root/.config/gcloud \
    bigtesty
```

## Run integration tests with Cloud Build

```bash
gcloud builds submit \
    --project=$PROJECT_ID \
    --region=$LOCATION \
    --config run-tests-bigtesty.yaml \
    --substitutions _SA_EMAIL=$SA_EMAIL,_IAC_BACKEND_URL=$IAC_BACKEND_URL,_ROOT_TEST_FOLDER=$ROOT_TEST_FOLDER \
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