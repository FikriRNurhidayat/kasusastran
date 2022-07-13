# Kasusastran API

Kasusastran API is the backend of Kasusastran website. Written fully in go using grpc gateway.

## Architecture

There will be at least 4 things that are running for the back-end:

- NSQ (`localhost:4171`)
- PostgreSQL (`localhost:5432`)
- gRPC server (`localhost:8080`)
- gateway server (`localhost:8081`)

## Getting Started

For now, clone the repo, run this following command:

```sh
make init
make setupdb
make develop
```

## Contributing

**TODO**
