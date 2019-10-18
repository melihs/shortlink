# shortlink

[![codecov](https://codecov.io/gh/batazor/shortlink/branch/master/graph/badge.svg)](https://codecov.io/gh/batazor/shortlink)
[![Build Status](https://travis-ci.org/batazor/shortlink.svg?branch=master)](https://travis-ci.org/batazor/shortlink)
[![Go Report Card](https://goreportcard.com/badge/github.com/batazor/shortlink)](https://goreportcard.com/report/github.com/batazor/shortlink)

Shortlink service

![logo](./docs/go.png)

### Run

```
docker-compose \
    -f docker-compose.yaml \
    -f ops/docker-compose/database/[DATABASE_NAME] \
    up -d 
```

##### As example

```
docker-compose \
    -f docker-compose.yaml \
    -f ops/docker-compose/database/dgraph.yaml \
    up -d 
```

### HTTP API

+ Import [Postman link](./docs/shortlink.postman_collection.json) for test HTTP API

- GET /:URL #=> 200 or 404
- POST / {"url":"example.com"} #=> 201
- DELETE / {"url":"example.com"} #=> 200
- GET /s/:URL #=> 301

###### Support 2 HTTP REST API:
- HTTP (chi)
- gRPC-gateway

### Store provider

+ RAM
+ Redis
+ MongoDB
+ Postgres
+ DGraph
+ LevelDB
+ Badger
