# Canvas GraphQL instructions

## Installation

This tutorial assumes you have git , go environment and docker with docker-compose
 properly set up.

First clone the repository as follows:

```
$ git clone https://github.com/juanfgs/canvas-graphql-server.git
```


## Fetch the dependencies
cd to the directory where you downloaded the project then run

```
$ go get -v 
```
to fetch all the dependencies

## Docker configuration

The project has a docker-compose.yml for easy setup. Just run 
```
docker-compose up -d 
```
to bring up the postgresql servers.
The servers are set up to run by default on non standard ports to not interfere 
with currently running postgresql servers.

## Configuration
The database configuration is defined in the ```env.example``` file. The program
should run with the provided docker configuration, but it's advised to make a 
 copy of this file ```cp env.example .env``` and edit the 
resulting .env file.

Afterwards you have to load the file with ```source .env``` to properly load the
database credentials

## Running the project

To run the project use

```
$ go run ./server.go
```

The API should be browseable from http://localhost:8080
## Running tests
You can run all the tests in the project with the following command:
```
$ GO_ENV=test go test ./... 
```

Make sure to specify GO_ENV=test to avoid running the integration tests in the
main database.
