# Kasusastran API

Kasusastran API is the backend of Kasusastran website. Written fully in go using grpc gateway.

## Architecture

There will be at least 4 things that are running for the back-end:

- NSQ (`localhost:4171`)
- PostgreSQL (`localhost:5432`)
- gRPC server (`localhost:8080`)
- gateway server (`localhost:8081`)
- redis (`internal`)

## Getting Started

For now, clone the repo, run this following command:

```sh
make init
make setupdb
make develop
```

## Code Generation

There are three things that you can generate on this repository using codegen.

1. Protoc
2. Golang SQL Query
3. Mocks

### Protoc

This is basic on the GRPC things, you just need to run:

```sh
make apis
```

### Query

This repository use `sqlc` as query code generator.

```sh
make query
```

### Mocks

For testing purposes

```sh
make mocks
```
