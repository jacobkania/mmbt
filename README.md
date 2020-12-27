# mmbt

Most Minimal Budgeting Tool

## How to run locally:

1. Install `postgres` and `golang-migrate` (can `brew install` both)
2. Create databases within postgres for running locally: `mmbt_dev` and `mmbt_test`
3. Create user within postgres for running locally: `mmbt_user`

## Make commands

- `make init`: fetches dependencies and initializes project.

- `make test`: runs all tests.

- `make dev`: runs Go-server & snowpack locally. Snowpack hot-reloading enabled.

- `make bs`: **b**uild and **s**tart in a prod-similar `test` environment.

- `NAME=migration_name make new-migration`: creates new up&down migration files for db

- `make migrate-up`: migrates to newest version
