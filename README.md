# go-startkode

Scaffolding for a go api project.

# Prerequisites:

- [Golang](https://go.dev/doc/install)
- [Make](https://www.gnu.org/software/make/)
- [docker](https://www.docker.com/) & [docker compose](https://docs.docker.com/compose/)

# Setup:

- Run the `install-tools.sh` scripts if you don't already have the tools. (Ensure your `PATH` is set up correctly).

- Rename the `.example.env` file to `.env` and populate the fields as needed.

- Run `docker compose up` to start the PostgreSQL Docker container or `docker compose up -d` to run it in the background as a daemon on every boot.

- Run `make migrate-up` to setup the database with the tables etc. specified in the `db/migration/` folder.

# Usage:

The Makefile defines different task to run. When in doubt you can just run `make` to find out which tasks are available or read the `Makefile` for more information
