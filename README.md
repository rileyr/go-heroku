# go-heroku

This repository is a boilerplate golang project for deployment on heroku, using the heroku container repository.

The build/deploy process for any project built from this boilerplate is:
  1. compile the contents of `/public` into a go package at `/statik`.
  2. compile the application targeted for an alpine docker container.
  3. build the docker image for heroku.
  4. push to heroku container repository.

---

## installation / usage

  1. install rakyll/statik: `go get github.com/rakyll/statik`
  2. clone this repository.
  3. delete the `.git` folder.
  4. write a sweet web app
  5. set heroku app name env var `HEROKU_APP=coolapp`
  6. deploy to heroku: `make heroku`
  7. ???
  8. profit
