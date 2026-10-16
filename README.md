# bigtesty-example

This project shows BigQuery integration tests launched with [BigTesty](https://github.com/tosun-si/bigtesty).
BigTesty creates isolated, short-lived datasets for each scenario, loads the test data, runs the queries and compares
the results with the expected rows.

Docs: https://tosun-si.github.io/bigtesty/

## Project layout

```text
bigtesty.yaml                          # datasets and tables created for each scenario
schema/monitoring/job_failure.json     # BigQuery JSON schema of the table
tests/monitoring/
├── find_failures_by_feature_name.bigtesty.yaml
├── find_failures_by_job_name.bigtesty.yaml
├── find_failures_with_functions.bigtesty.yaml
├── given/                             # test data
├── when/                              # SQL queries
└── then/                              # expected rows and assertion functions
```

## Setup

BigTesty is a dev dependency of the project, managed with [uv](https://docs.astral.sh/uv/):

```bash
uv sync
gcloud auth application-default login
```

The user or service account needs the `roles/bigquery.user` role on the test project.

## Run the tests

```bash
uv run bigtesty test --project $PROJECT_ID
# or with pytest, one test per scenario
uv run pytest --bigtesty-project $PROJECT_ID
```

`GOOGLE_CLOUD_PROJECT` can be used instead of `--project`.

## CI

All the pipelines install the dependencies locked in `uv.lock` with `uv sync --frozen`:

- GitHub Actions: [`.github/workflows/run-tests-bigtesty-github.yaml`](.github/workflows/run-tests-bigtesty-github.yaml),
  Workload Identity Federation and the uv cache of `astral-sh/setup-uv`.
- GitLab CI: [`.gitlab-ci.yml`](.gitlab-ci.yml), uv image and uv cache keyed on `uv.lock`.
- Cloud Build: [`run-tests-bigtesty.yaml`](run-tests-bigtesty.yaml).

```bash
gcloud builds submit \
    --project=$PROJECT_ID \
    --region=$LOCATION \
    --config run-tests-bigtesty.yaml .
```
