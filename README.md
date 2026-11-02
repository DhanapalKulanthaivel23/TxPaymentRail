# TxPaymentRail

[**TxPaymentRail-demo**](https://dhanapalkulanthaivel23.github.io/TxPaymentRail/) · [API Collections](https://dhanapalkulanthaivel23.github.io/TxPaymentRail/#/collections) · [MCP Tools](https://dhanapalkulanthaivel23.github.io/TxPaymentRail/#/mcp-tools)

SWIFT MT / ISO 20022 MX field inspection, address correction, CBPR+, SEPA SCT/SDD/Inst, RTGS, CHAPS, Bacs/FPS lifecycle evidence and reconciliation.

## Open and run

Use the links above, or clone and start the same interface locally:

```sh
git clone https://github.com/DhanapalKulanthaivel23/TxPaymentRail.git
cd TxPaymentRail
go run .
```

Open http://127.0.0.1:8080/. To use another port, run `LISTEN_ADDR=127.0.0.1:8090 go run .`.

The local client automatically reads the published connection configuration and calls the same hosted core as GitHub Pages. Visitors need no private-repository access or service credentials. The Go application serves the public interface; private implementation and provider keys stay on the backend.

## Test the workflow

1. Choose a collection, then a flow, and inspect its request body.
2. Run a request and inspect the response, request sent and captured identifiers. Follow-up requests reuse those identifiers within the workspace.
3. Open MCP Tools to discover tools and execute a JSON-RPC tool call.
4. Use Files where available, Knowledge Graph for cited relationships, and Questions & Answers for source-backed retrieval.

Technology tags describe the project components. Optional Gemini and external-provider actions report their availability and execution evidence in the response. A retrieved reference answer is distinct from a model-generated result.

## Hosting and connection

GitHub Pages serves the HTML/CSS/JavaScript. The private core is hosted on Render; this project's API base is `https://banking-project-cores.onrender.com/payments`. The free core can pause during inactivity, so its first connection can take longer. Runtime state resets after a core restart.

The public interface is identical in the Go server and Pages deployment. Hash-based navigation preserves repository subpaths. Request definitions, file processing and business operations are supplied by the core; no private source, provider keys or full collection export is bundled here.

Maintainers set the GitHub Actions repository variable `CORE_API_URL` and run **Publish project interface** to update the published connection. Local clones resolve that configuration automatically. For owner-side local integration tests, `CORE_API_URL` can point to the corresponding loopback core instead.

## Checks

```sh
go test -race ./...
node --test docs/*.test.mjs
```

The embedded frontend and Pages assets are checked for equality. Hosted release checks exercise native API responses, MCP discovery/calls, PDF intake, graph evidence and cited retrieval, plus source-route denial. Browser checks cover all four Pages sites and fresh public clones using the same hosted endpoint.
