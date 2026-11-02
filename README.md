# TxPaymentRail — Go REST API example

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
docker build -t txpaymentrail-ai-example .
docker run --rm -p 8080:8080 txpaymentrail-ai-example
```

## Backend connection

Open http://localhost:8080 after `go run .` to use the connection page. Enter your private gateway HTTPS address (or localhost during development) and a gateway token. The page sends requests to that gateway, which calls the separately running backend. Additional backend credentials are entered only when required; tokens are not stored. The local `/records` example remains synthetic and separate.

For GitHub Pages, publish the `docs/` folder and set `docs/config.json` gatewayURL to your hosted gateway origin. Leave credentials out of that file. The backend owner must allow the exact UI origin and requested routes. No private source is imported or compiled into this UI. In-memory backend data resets on restart.

## Domain API workbench

Open `/workbench.html` for preloaded REST requests, Postman export, applicable file/MCP tabs, a functional walkthrough and explicit technology/status explanations. Sample mode performs no network/model calls. Live mode needs a separately configured private gateway; the public shell does not implement private endpoints. No credentials are exported. The existing IAM request/ticket studio is preserved.

## Read-only MCP contract update

The MCP tab and Postman collection now include the private backend’s read-only evidence/status tools. The public shell does not execute these tools itself; use the separately running core and correct credentials. Sample mode remains illustrative.

## Private workspace boundary (supersedes public export instructions)

The workbench now loads authorized collections and MT/MX context locally from its private core. Public download/export controls and embedded request fixtures have been retired. Keep workspace.private.json out of public web roots. Reference replay is distinct from live execution; the interface does not claim a live banking connection or conformance. Earlier export instructions describe the previous UI only.

## Current collection and MCP interface

The latest generic interface source is in [docs/project-pages](docs/project-pages/README.md). It provides separate collection and MCP pages plus applicable file, graph and cited-question views. The private project host supplies request metadata and executes the existing services; no private collections or implementation are bundled here.
