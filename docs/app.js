const element = id => document.getElementById(id);
const config = await fetch(new URL("config.json", import.meta.url), {cache: "no-store"}).then(r => {if (!r.ok) throw Error(); return r.json();}).catch(() => ({}));
element("gateway").value = config.gatewayURL || (["localhost", "127.0.0.1", "[::1]"].includes(location.hostname) ? config.localGatewayURL || "" : "");
element("path").value = config.healthPath || element("path").value;
element("clear").addEventListener("click", () => {
  for (const id of ["token", "authorization", "serviceKey"]) element(id).value = "";
  element("result").textContent = "";
  element("status").textContent = "Credentials cleared";
});
element("connection").addEventListener("submit", async event => {
  event.preventDefault(); element("send").disabled = true; element("result").textContent = "";
  try {
    const gateway = new URL(element("gateway").value);
    if (gateway.protocol !== "https:" && !(gateway.protocol === "http:" && ["localhost", "127.0.0.1", "[::1]"].includes(gateway.hostname))) throw Error("Use HTTPS, or HTTP on localhost.");
    if (gateway.username || gateway.password || gateway.search || gateway.hash || gateway.pathname !== "/") throw Error("Enter only the gateway origin.");
    const path = element("path").value;
    if (!path.startsWith("/") || path.startsWith("//") || path.includes("#")) throw Error("Enter a backend path starting with /.");
    const method = element("method").value;
    const headers = {Authorization: "Bearer " + element("token").value};
    if (element("authorization").value) headers["X-Core-Authorization"] = element("authorization").value;
    if (element("serviceKey").value) headers["X-Core-Service-Key"] = element("serviceKey").value;
    let body;
    if (method !== "GET" && element("body").value.trim()) {body = JSON.stringify(JSON.parse(element("body").value)); headers["Content-Type"] = "application/json";}
    element("status").textContent = "Connecting…";
    const response = await fetch(new URL("/core" + path, gateway), {method, headers, body, credentials: "omit", redirect: "error", cache: "no-store", signal: AbortSignal.timeout(15000)});
    element("status").textContent = response.ok ? "Connected · HTTP " + response.status : "Request rejected · HTTP " + response.status;
    const text = await response.text();
    try {element("result").textContent = JSON.stringify(JSON.parse(text), null, 2);} catch {element("result").textContent = text;}
  } catch (error) {element("status").textContent = "Connection failed"; element("result").textContent = error.message + " Check the gateway address, token and allowed UI origin.";}
  finally {element("send").disabled = false;}
});
