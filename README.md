[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
# go-app-invite-service
This service allows generating invite token by admin. Invite token can validated.

# Design
![System Design](./img/system_diagram.png "System design")

# Spinning up locally
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
  ```
- run app: `make run`

# API Authentication
Admin api have authentication. API key based authentication for these routes. Add `API_KEY` value in `.env` file and supply it as `api-key` header field of each admin request
```curl
curl --location --request GET '<endpoint>/api/v1/admin/token' \
--header 'api-key: <api_key>'
```

# API Authentication
API documentation is build with [swagger](https://swagger.io/). To run api documentation execute following command:
  ```
  make doc
  ```

![API Doc](./img/swagger.png "api doc")