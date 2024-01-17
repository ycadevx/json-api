

### JSON API PROJECT 

This is a bank API

## Quick Start (How to run ?)

Open a terminal and run the following command (start a postgres instance) :

```
docker run --name local-postgres -e POSTGRES_PASSWORD=gobank22 -p 5432:5432 -d postgres
```

 PostgresSQL with docker ===> [PostgreSQL](https://hub.docker.com/_/postgres)

For a list of all containers in Docker :

```
docker ps -a 
```

Lists all Docker images :



```
docker images
```
To run a ready container. The ID in the list should be written in the containerID field.

```
docker start containerID
```

###  Start program



```
go run ./
```

// TO DO YCA :  Will add request and response

The  JSON API server running. API is now available on your host at http://localhost:3000



