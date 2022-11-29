# Requirements:
- docker v20.0+
- docker-compose v2.12+
- If you're running from Windows, you should have WSL installed (it's already installed with Docker Desktop by default) to be able to run shell scripts.

# How to run:
1) Change your working directory to project root
2) Create your own `.env` file. You should set your own `HMAC_SECRET` and `DB_PASSWORD`. Other params may also be changed. Content should be similar to:
```dotenv
LISTEN_ADDR=:8080
HMAC_SECRET=233972089023bb1838ae877063b3080c7a4fefd57a1a8125a5ff29546b0ea1f1
DB_HOST=db
DB_PORT=5432
DB_USER=companies
DB_PASSWORD=12345
DB_NAME=companies
DB_MIGRATIONS=db_migration
```
3) Run shell script that is located at `build/build_and_up.sh` without any params. Remember that your working directory must be the project root.
4) You may import Postman collection at `tests/postman-collection/companies-test-task.postman_collection.json` for testing.
5) After finishing testing run shell script `build/down.sh` to stop and delete all containers.

# Environmental variable description:
- `LISTEN_ADDR` socket address that companies service will listen at. Default is `:8080`;
- `HMAC_SECRET` is used for signing/verifying JWT tokens. Empty by default;
- `DB_HOST` must point to Postgres IP address. Default: `localhost`;
- `DB_PORT` must point to Postgres port. Default:`5432`;
- `DB_USER` is Postgres username. Default: `companies`;
- `DB_PASSWORD` is Postgres password. Empty by default;
- `DB_NAME` is the database name. Default: `companies`;
- `DB_MIGRATIONS` - relative path to migration scripts. Default: `db_migration`.

# Note about authentication:
There is a tiny auth "service" at `:8090` that has only one endpoint `GET /login` which signs new JWT token with configured `HMAC_SECRET`. This JWT token then passed to CRUD endpoints on companies service. There is no login/password prompt for simplicity purpose - just a new valid JWT token.