# bigtesty-example

This project shows examples with BigQuery integration tests launched with BigTesty.\
BigTesty is a framework that allows to create Integration Tests with BigQuery on a real and short-lived Infrastructure.

## Run integration tests with BigTesty

```bash
docker run -it \
    -e PROJECT_ID=$PROJECT_ID \
    -e SA_EMAIL=$SA_EMAIL \
    -e LOCATION=$LOCATION \
    -e TF_VAR_project_id=$PROJECT_ID \
    -e TF_STATE_BUCKET=$TF_STATE_BUCKET \
    -e TF_STATE_PREFIX=$TF_STATE_PREFIX \
    -e GOOGLE_PROVIDER_VERSION=$GOOGLE_PROVIDER_VERSION \
    -e ROOT_TEST_FOLDER=$ROOT_TEST_FOLDER \
    -v $(pwd)/tests:/app/tests \
    -v $(pwd)/tests/tables:/app/infra/bigquery/tables \
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
    --substitutions _TF_STATE_BUCKET=$TF_STATE_BUCKET,_TF_STATE_PREFIX=$TF_STATE_PREFIX,_GOOGLE_PROVIDER_VERSION=$GOOGLE_PROVIDER_VERSION,_ROOT_TEST_FOLDER=$ROOT_TEST_FOLDER \
    --verbosity="debug" .
```