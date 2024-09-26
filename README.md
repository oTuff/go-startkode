# go-startkode

Scaffolding for a go api project.

# Prerequisites

- Golang installed
- Make installed
- [direnv](https://direnv.net/) installed and configured
- a postgres database running

# setup:

- run the install-tools.sh scripts if you don't already have the tools
- create a .env file with the following fields:

```
DBSTRING="host=localhost port=5432 user=postgres password=test dbname=todo sslmode=disable"
GOOSE_DRIVER=postgres
GOOSE_DBSTRING="$DBSTRING"
GOOSE_MIGRATION_DIR="db/migrations"
```

- run `direnv allow` to enable the activation of environment variables by direnv.
