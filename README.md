# AddressRail — Go REST API example

A standalone API example using synthetic records and in-memory storage.

- Create, list, read, update and delete records.
- Validate JSON input.
- Run HTTP handler tests.
- Build a small container image.

## Structure

```text
base/
controller/
libs/
main.go
Dockerfile
```

## Run and test

```sh
go test -race ./...
go run .
```

The API listens on `127.0.0.1:8080`. Set `LISTEN_ADDR` to override it. Data resets when the process stops.

## API

| Method | Path |
| --- | --- |
| GET | /health |
| POST | /records |
| GET | /records |
| GET | /records/{id} |
| PUT | /records/{id} |
| DELETE | /records/{id} |

Create or update body:

```json
{"name":"sample"}
```

## Docker

```sh
docker build -t addressrail-ai-example .
docker run --rm -p 8080:8080 addressrail-ai-example
```
