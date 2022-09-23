# Money

Dummy project to learn web architecture and test some things

# Local environment

## Installation

Once everything is downloaded, you can run the project in local with:

    docker compose -f docker-compose.yml -f docker-compose-dev.yml up --build

## Links
The following links are available:
* [pvt.localhost](pvt.localhost): front-end address
* [api.pvt.localhost](api.pvt.localhost): back-end address. A GraphiQL server is available at this address to make direct calls to the back-end. The API URL is [api.pvt.localhost/query](api.pvt.localhost/query).
* [traefik.pvt.localhost](traefik.pvt.localhost): dashboard of traefik, to check the reverse proxy.
* [pma.pvt.localhost](pma.pvt.localhost): PhpMyAdmin server, connected to the database. You can find the credentials in the docker compose files.


## Main stack
The three main components of the project are:
* a back-end written in Golang (v1.18), which lies into the *internal* folder. The back-end uses the framework Gin (v1.8) and exposes a GraphQL API, generated with 99designs/gqlgen. Pages are served with Gin's web server.
* a front-end written with the framework VueJS (v3.2 - in particular, we're using the *composition* API) and Vite (v3.0) for bundling. It lies into the *web* folder. The front-end uses ApolloClient (v3.6) to communicate with the back-end. Pages are served with Vite's web server.
* MariaDB as the database.

## Docker
Docker (v20.10) and Docker Compose are used to get everything running. The local environment should merge the prod file (`docker-compose.yml`) with the local file (`docker-compose-dev.yml`).

The Docker files lies into the `docker/dev` folder. Hot reload is configured for both the front-end and the back-end, so you should be able to edit the code and it should be (almost) immediately repercuted on your local version.

## Other components of the stack
* PhpMyAdmin to manage the database
* Traefik as a reverse proxy

# Production environment

## Links

The website is available at http://paulvalentintalbot.com. Please note that to keep costs at a minimum, the instances are usually shut down and need to be started for the site to be available.

## Differences with the dev environment
In production, you won't have:
* PhpMyAdmin
* Access to Traefik dashboard
* Access to GraphiQL

## Docker

The `docker-compose.yml` file is used to make things run in production.

The images use their Alpine version. The front-end uses the Nginx base image, other ones their own base image. The files lies into the `docker/prod` folder.

Docker Swarm is used as an orchestrator. A manager node with the label `entry == traefik` and a worker node with the label `storage == maria` are required.

## AWS

Three EC2 instances are used.

## CircleCI

Circle is used for Continuous Deployment (Continuous Integration to come, but there are no tests for now). On every push to the master branch, it builds the images, push them to DockerHub, connects to the manager, redeploys the stack and prunes everything.

The configuration file can be found in the `.circleci` folder.