[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![CI](https://github.com/shankar524/go-app-invite-service/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/shankar524/go-app-invite-service/actions/workflows/test.yml)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://makeapullrequest.com)
# Go based RESTful API service Boilerplate
A boilerplate/starter project for quickly building production-ready RESTful APIs using Golang with features like:
  - API server implemented in [go-genic](https://gin-gonic.com/)
  - Handled common webserver scenarios:
    - allows both API key validation(permanent access) and token based validation(temporary access)
    - token rate limiting and caching
    - pre-configured cron setup
    - pre-implemented private and public routes
    - Swagger docs setup
    - RBAC implementation
    - GitHub Action based CI setup to run test

Additionally the following role-based functionality are already implemented:
- Private routes:
  - Generating and deactivating token(token gets cached)
- Public routes:
  - Make request on public routes with provided token
  - Request throttled for public routes

## Design
![System Design](./img/system-diagram.png "System design")

## Folder structure
```bash
├── api
│   ├── controller
│   ├── repository
│   ├── routes
│   └── service
├── bootstrap
├── cron
├── docs
│   └── swagger.yaml
├── img
├── lib
├── main.go
├── makefile
├── middlewares
└── models
```

### Setting up locally
- install [MySql](https://www.mysql.com/downloads/) Or run mysql docker image
  ```
  docker run \
    --detach \
    --name=mysql \
    --env="MYSQL_ROOT_PASSWORD=admin" \
    --env="MYSQL_USER=admin" \
    --env="MYSQL_PASSWORD=admin" \
    --env="MYSQL_DATABASE=test" \
    --publish 3306:3306 \
    mysql/mysql-server:latest
  ```
- install [Redis]() Or run Redis docker image
    ```
     docker run \
     -d --name redis \
     -p 6379:6379 \
     redis/redis-stack:latest
    ```
- setup `.env` file
  ```
  APP_PORT=8080
  DB_USERNAME=admin
  DB_PASSWORD=admin
  DB_PORT=3306
  DB_HOST=localhost
  DB_NAME=test
  REDIS_HOST=localhost
  REDIS_PORT=6379
  REDIS_PASSWORD=
  API_KEY=test
  API_RATE=10
  ```
- run app: `make run`

### API Authentication
Private routes are api key protected wth API key authentication. Add `API_KEY` value in `.env` file and supply it as `api-key` header field of each admin request
```curl
curl --location --request GET '<endpoint>/api/v1/admin/token' \
--header 'api-key: <api_key>'
```

### API Documentation
API documentation is build with [swagger](https://swagger.io/). To run api documentation execute following command:
  ```
  make doc
  ```

![API Doc](./img/swagger.png "api doc")
