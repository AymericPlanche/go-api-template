# Go API template

Replace `myapp` with your app name
It provides sample CRUD operations around a `Thing` entity.

## Local dev environment

Run `task` to get a list of possible commands

The API is available on port `8080` by default.

The local image supports hot reloading, so the go application is rebuilt every time you make a change to a go file.

## Tests

### Unit tests

Unit tests are run with `task test-unit`
Unit tests should be stateless and can be run in parallel.

### Integration tests

Integration tests are run with `task test-integration`
Integration tests rely on an `integration` database that is reset between tests with `integration_tests.ResetDatabase()`
Tests are run sequentially to avoid race condition to the database
