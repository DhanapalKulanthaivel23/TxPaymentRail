# Collection and MCP interface

API Collections, MCP Tools and the applicable file, graph and cited-answer views share the same client assets in Pages and the local Go server. The client reads `client-config.json` and resolves the published registry when a local clone has no endpoint override. It calls only the configured core catalog, execution, MCP and upload routes. No service credentials, private implementation or collection export is embedded in this directory.

The ivory, charcoal, sky-blue and copper palette is retained. Connection progress is shown while the catalog loads; a failed connection replaces the workspace controls with its explicit status.
