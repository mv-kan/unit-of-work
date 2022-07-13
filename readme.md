# Unit of work in golang 

This is my simple showcase of how unit of work pattern can be implemented in go.
This is a really simple showcase of this [repo](https://github.com/rafael-piovesan/go-rocket-ride) particularly unit of work pattern

## Project structure
```
    /migrate # this is where migrations are 
    /db # dockerfile of postgresdb
    /src # source code
        /entity # contains all entities of this project 
        /repo # contains interfaces and implementations of repositories
        /uow # Unit of work pattern the meat and potatoes of this whole project
```
## How to run it?
**Prerequisite**:
1. docker
```
// up all containers
docker-compose up -d
// enter db to check tables and result of golang code
make enter-db
```
In main.go it prints out result of queries so you can see it in the container log

## Dependencies
1. I use migration tool called [migrate](https://github.com/golang-migrate/migrate) the cli version for manual migrations. You can use makefile commands to ease the pain of a new unknown tool. But in main.go there is a piece of code that automatically run migrations so you don't need to install this tool if you are to run this project 
2. Postgres database runs as docker container. Dockerfile is located in db folder 
