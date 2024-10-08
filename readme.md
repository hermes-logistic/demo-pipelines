#### System Requirements

| Name | Version |
| :---: | :---: |
| postgres | => 14 |
| golang | => 1.21 |
| docker | => 20.10.16 |
| docker-compose | => 1.29.2 |
| mongo-db | => 7.0.5 |

#### List of Required Packages

| Name | Version |
| :---: | :---: | 
| bash | ~> 5.1.4 |

## Project Structure
```sh
.
├── api
│   ├── tasks
│   │   ├── application
│   │   │   └── empty
│   │   │       ├── empty_test.go
│   │   │       └── empty_validation.go
│   │   ├── domain
│   │   │   ├── tasks.go
│   │   │   └── tasks_interface.go
│   │   └── infrastructure
│   │       ├── create
│   │       │   ├── data_create.go
│   │       │   ├── data_test.go
│   │       │   └── db
│   │       │       ├── db_mongo_insert.go
│   │       │       ├── db_sql_insert.go
│   │       │       ├── db_structs.go
│   │       │       └── db_test.go
│   │       ├── delete
│   │       │   ├── data_delete.go
│   │       │   └── db
│   │       │       └── db_sql_delete.go
│   │       ├── read
│   │       │   ├── data_read.go
│   │       │   └── db
│   │       │       ├── db_mongo_read.go
│   │       │       └── db_sql_read.go
│   │       ├── response
│   │       │   ├── create_response.go
│   │       │   ├── delete_response.go
│   │       │   ├── read_response.go
│   │       │   └── update_response.go
│   │       └── update
│   │           ├── data_update.go
│   │           └── db
│   │               └── db_sql_update.go
│   └── users
│       ├── application
│       │   └── empty
│       │       └── empty_validation.go
│       ├── domain
│       │   ├── users.go
│       │   └── users_interface.go
│       └── infrastructure
│           ├── create
│           │   ├── data_create.go
│           │   └── db
│           │       ├── db_sql_insert.go
│           │       └── db_structs.go
│           ├── delete
│           ├── read
│           ├── response
│           │   ├── create_response.go
│           │   ├── delete_response.go
│           │   └── update_response.go
│           └── update
├── db
│   ├── mongodb
│   │   ├── database_connection.go
│   │   └── database_main.go
│   └── sql
│       ├── database_connection.go
│       ├── database_main.go
│       └── database_migration.go
├── ddosify-data.json
├── docker-compose.yml
├── Dockerfile
├── dockerfile.develop
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── helpers
│   └── structs
│       └── structs_main.go
├── logs
│   └── logs_main.go
├── main.go
├── payload
│   ├── tasks.json
│   └── users.json
├── pkg
│   └── mod
│       └── cache
│           └── lock
├── readme.md
├── router
│   ├── router_main.go
│   ├── router_task.go
│   ├── router_test.go
│   └── router_users.go
└── test.sh
```
### main 

- .gitignore: determine which files and directories to ignore in your Git repository

- .dockerignore: specifies which files and directories to exclude from the Docker build context

- .env: enviorment variables

- docker-compose.yml: used to define multi-container Docker applications, for this project we set [api, postgres db, test postgres db, mongo db]

- Dockerfile: contains instructions for building a Docker image

- main.go: kick off the execution of the program, from main(), you can call other functions, import packages, and do everything necessary to make the program work correctly

- readme.md: project documentation

### api

Contains all the entities that the API will manage, we create a dir for each entity 

-   #### entity_application
    Contains the specific validations for the entity

-   #### entity_domain
    Contains the entity base struct, this should not be modified

-   #### entity_infrastructure
    In root communicates with the http protocols and in the db subdirectory it connects to the database depending on the context

### db
Create a directory for each database context, then create a context_main.go file where the context struct will be & a context_connection.go file 

### logs
In file logs_main.go create the layers for logs handlers

### router
Contains the api routes, in file router_main.go put the general routes, create a router_entity.go file for each entity, the methods POST, PUT, DELETE, GET will be in those files

## Packages

| Name | Version | go pkg |
| :---: | :---: | :---: |
| github.com/bytedance/sonic | v1.10.2 | [Link](https://pkg.go.dev/github.com/bytedance/sonic@v1.10.2) |
| github.com/chenzhuoyu/base64x | v0.0.0-20230717121745-296ad89f973d | [Link](https://pkg.go.dev/github.com/chenzhuoyu/base64x@v0.0.0-20230717121745-296ad89f973d) | 
| github.com/chenzhuoyu/iasm | v0.9.1 | [Link](https://pkg.go.dev/github.com/chenzhuoyu/iasm@v0.9.1) |
| github.com/gabriel-vasile/mimetype | v1.4.3 | [Link](https://pkg.go.dev/github.com/gabriel-vasile/mimetype@v1.4.3) |
| github.com/gin-contrib/sse | v0.1.0 | [Link](https://pkg.go.dev/github.com/gin-contrib/sse@v0.1.0) |
| github.com/gin-gonic/gin | v1.9.1 | [Link](https://pkg.go.dev/github.com/gin-gonic/gin@v1.9.1) |
| github.com/go-playground/locales | v0.14.1 | [Link](https://pkg.go.dev/github.com/go-playground/locales@v0.14.1) |
| github.com/go-playground/universal-translator | v0.18.1        | [Link](https://pkg.go.dev/github.com/go-playground/universal-translator@v0.18.1) |
| github.com/go-playground/validator/v10 | v10.17.0 | [Link](https://pkg.go.dev/github.com/go-playground/validator/v10@v10.17.0) |
| github.com/goccy/go-json | v0.10.2 | [Link](https://pkg.go.dev/github.com/goccy/go-json@v0.10.2) |
| github.com/json-iterator/go | v1.1.12 | [Link](https://pkg.go.dev/github.com/json-iterator/go@v1.1.12) |
| github.com/klauspost/cpuid/v2 | v2.2.6 | [Link](https://pkg.go.dev/github.com/klauspost/cpuid/v2@v2.2.6) |
| github.com/leodido/go-urn | v1.3.0 | [Link](https://pkg.go.dev/github.com/leodido/go-urn@v1.3.0) |
| github.com/lib/pq | v1.10.9 | [Link](https://pkg.go.dev/github.com/lib/pq@v1.10.9) |
| github.com/mattn/go-isatty | v0.0.20 | [Link](https://pkg.go.dev/github.com/mattn/go-isatty@v0.0.20) |
| github.com/modern-go/concurrent | v0.0.0-20180306012644-bacd9c7ef1dd | [Link](https://pkg.go.dev/github.com/modern-go/concurrent@v0.0.0-20180306012644-bacd9c7ef1dd) | 
| github.com/modern-go/reflect2 | v1.0.2 | [Link](https://pkg.go.dev/github.com/modern-go/reflect2@v1.0.2) |
| github.com/pelletier/go-toml/v2 | v2.1.1 | [Link](https://pkg.go.dev/github.com/pelletier/go-toml/v2@v2.1.1) |
| github.com/twitchyliquid64/golang-asm | v0.15.1 | [Link](https://pkg.go.dev/github.com/twitchyliquid64/golang-asm@v0.15.1) |
| github.com/ugorji/go/codec | v1.2.12 | [Link](https://pkg.go.dev/github.com/ugorji/go/codec@v1.2.12) |
| golang.org/x/arch | v0.7.0 | [Link](https://pkg.go.dev/golang.org/x/arch@v0.7.0) |
| golang.org/x/crypto | v0.18.0 | [Link](https://pkg.go.dev/golang.org/x/crypto@v0.18.0) |
| golang.org/x/net | v0.20.0 | [Link](https://pkg.go.dev/golang.org/x/net@v0.20.0) |
| golang.org/x/sys | v0.16.0 | [Link](https://pkg.go.dev/golang.org/x/sys@v0.16.0) |
| golang.org/x/text | v0.14.0 | [Link](https://pkg.go.dev/golang.org/x/text@v0.14.0) |
| google.golang.org/protobuf | v1.32.0 | [Link](https://pkg.go.dev/google.golang.org/protobuf@v1.32.0) |
| gopkg.in/yaml.v3 | v3.0.1 | [Link](https://pkg.go.dev/gopkg.in/yaml.v3@v3.0.1) |

## Environment

| Name | type | Value |
| :---: | :---: | :---: |
| DB_STRING | string | host=postgres user=postgres password=1234 dbname=backend_golang_gin port=5434 sslmode=disable |
| DB_NAME | string | backend_golang_gin |
| USER_DB | string | postgres |
| PASSWORD | string | 1234 |
| TEST_USER_DB | string | postgres |
| TEST_PASSWORD | string | test |
| TEST_DB_NAME | string | api_db_test |
| TEST_DB_STRING | string | host=host=postgres user=postgres password=test dbname=api_db_test port=5433 sslmode=disable |
| MONGO_URI | string | mongodb://admin:password@mongo:27017 |
| CONTEXT | string | [postgres, mongo] |
| MONGO_USER | string | admin |
| MONGO_PASSWORD | string | password |

## Container Operative System

Alpine

## How to use

#### Start services

```bash
$ docker-compose up --build -d
```

#### Postgres Container Access

```bash
$ docker exec -it postgres-db /bin/bash

$ psql -U postgres -d backend_golang_gin
```

#### Api logs

```bash
$ docker logs go-api
```

#### Stop services

```bash
$ docker-compose down
```

#### Stop services & delete database volumes

```bash
$ docker-compose down -v
```

#### test

```bash
$ go test ./... -coverprofile=c.out
$ go tool cover -html="c.out"
```

## Load tests
- go to ddsoify-data.json
- create a file with the structure of the petition [POST only] in the payload directory
- add the new endpoint 
  ```json
    {
        "id": "$ID",
        "method": "$METHOD",
        "url": "$URL/$ENDPOINT",
        "headers": {
          "Content-Type": "application/json"
        },
        "payload_file": "./payload/$ENDPOINT.json"
    }
  ```
    ### Replace the following
     - $ID: The number of the last ID in the file plus one
     - $METHOD: The http method of the endpoint [GET, POST, PUT, DELETE]
     - $URL: Address of the service ej. https://api.com
     - $ENDPOINT: Route of the endpoint ej, tasks

#### Example
```JSON
{
    "iteration_count": 100,
    "load_type": "waved",
    "duration": 10,
    "steps": [
      {
        "id": 1,
        "method": "POST",
        "url": "https://api-gateway.hermesv.dev/v1/tasks",
        "headers": {
          "Content-Type": "application/json"
        },
        "payload_file": "./payload/tasks.json"
      }
    ]
}
  
```
#### ./payload/ENDPOINT.json Example
```json

{
    "Name": "{{_randomString}}",
    "Status": "{{_randomString}}"
}

```
