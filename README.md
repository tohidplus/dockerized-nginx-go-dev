# Dockerized nginx go dev
A simple dockerized nginx-go web app project for development environment.

### Run
Rename the `.env.example` file to `.env` and set the environment variables as you wish.

Go inside the `docker` directory and run `docker-compose up -d.

### Notice
I've used [codegangsta/gin](https://github.com/codegangsta/gin) to have a live server while developing which runs on port `3000`
inside the `app` container.

To see `go` logs inside the container just run:

`docker-compose logs -t -f --tail="all" app`