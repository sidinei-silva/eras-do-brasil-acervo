type Ev = { seq: number; tipo: "dano" }
Bun.serve({ port: 7777, fetch(req, server) { if (server.upgrade(req)) return; return Response.json({seq:1,tipo:"dano"} satisfies Ev) }, websocket: { message(ws, m) { ws.send(m) } } })
