/*
============================================================
01 — CONFIGURATION
Problem 7 — Wire Config to HTTP Server
============================================================

Start HTTP server using config Port.

Requirements:

    1. LoadConfig with PORT (default 8080)

    2. Create http.Server:

           Addr: ":" + cfg.Port

    3. Register GET /health → return "ok"

    4. Log startup message with port (no secrets)

    5. Run server (Ctrl+C to stop)

    6. Test: curl http://localhost:8080/health

============================================================
Goal
============================================================

Practice:

    - configurable server port

============================================================
*/

package main

func main() {}
