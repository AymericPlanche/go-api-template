# Go API template

Replace `myapp` with your app name
It provides sample CRUD operations around a `Thing` entity.

## Local dev environment

Run `make` to get a list of possible commands

`make up` builds the start the local dev environment.
The API is available on port `8080` by default.

The local image supports hot reloading, so the go application is rebuilt every time you make a change to a go file.

## Tests






## Work with the production docker image

**Note: You will rarely have to do this** Basicaly only when you make changes to the `Dockerfile`.

The prod image is a very lightweight image (approx 20Mb at the moment) It is great for production but not very convenient to work with.
That's why we use a different docker image in dev (with hot reloading, etc)
If you want to run the app against the prod image, use `make run-prod-image`

Note that both instances can live side by side if installed on different ports (`8080` by default for dev, `8081` by default for prod)
