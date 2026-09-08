package main

import "fmt"

/*
================================================================================
01 — ARCHITECTURE DECISIONS
Phase 10 · Production Template
================================================================================

HOW TO READ THIS FILE
---------------------
Read section by section. Each section is one idea.
Skip the code blocks on first pass if you want — come back later.

After this lecture you should know:
  • Why we use layers
  • Where each type of code goes
  • How one HTTP request flows through the app
  • How to add a new endpoint (checklist at the end)

Reference project: conning-configuration (your CTO's codebase)


TABLE OF CONTENTS
-----------------
 1.  Why a template?
 2.  Before vs after
 3.  The big picture (4 layers)
 4.  What each layer does
 5.  One request: POST /register
 6.  Import rules
 7.  Three types of validation
 8.  Errors and DTOs
 9.  Where do I put X?
10.  Folder tree
11.  Wiring, runners, health
12.  CTO project map
13.  Add a new endpoint (checklist)
14.  Testing & common mistakes
15.  Phase 10 vs 11
16.  Quick Q&A
================================================================================
*/

// =============================================================================
// 1. WHY A TEMPLATE?
// =============================================================================

/*
In go-crash you learned skills in single files — that was correct for learning.

A production template answers one question:

    "Where do I put this code?"

Without it:
  • SQL ends up in handlers
  • Business rules hide in validator tags
  • Every new developer asks the team lead

With it:
  • Fixed layers + fixed checklist
  • Clone → run → add feature

Phase 10  → build the template (skeleton + User example)
Phase 11  → build e-commerce ON the template
*/

// =============================================================================
// 2. BEFORE VS AFTER
// =============================================================================

/*
BEFORE — everything in main.go:

    routes + handlers + SQL + JWT + config

AFTER — separated by job:

    cmd/api/main.go                 start app, wire dependencies
    domain/user/                    User entity + Repository interface
    application/usecase/            RegisterUser, LoginUser
    infra/httpserver/user/          routes, handler, JSON
    infra/adapters/repository/      PostgreSQL

Same app. Clear boundaries. Easy to test. Easy to onboard.
*/

// =============================================================================
// 3. THE BIG PICTURE (4 LAYERS)
// =============================================================================

/*
                         HTTP Request
                              │
                              ▼
              ┌───────────────────────────────┐
              │  infra/httpserver             │  JSON, routes, middleware
              └───────────────┬───────────────┘
                              │
                              ▼
              ┌───────────────────────────────┐
              │  application/usecase          │  business flow
              └───────────────┬───────────────┘
                              │
                              ▼
              ┌───────────────────────────────┐
              │  domain                       │  entities, repo interfaces
              └───────────────┬───────────────┘
                              ▲
                              │ implements
              ┌───────────────┴───────────────┐
              │  infra/adapters               │  PostgreSQL, Kafka, JWT
              └───────────────────────────────┘

    cmd/api/main.go  →  wires everything (composition root)


Golden rule: dependencies point INWARD.
  • domain never imports infra
  • use case never imports net/http
*/

// =============================================================================
// 4. WHAT EACH LAYER DOES
// =============================================================================

/*
┌─────────────┬────────────────────────────────────────────────────────────┐
│ LAYER       │ JOB                                                        │
├─────────────┼────────────────────────────────────────────────────────────┤
│ domain/     │ WHAT exists in the business (User, Order)                  │
│             │ Repository INTERFACE (not SQL)                             │
│             │ Domain errors (ErrEmailTaken)                              │
│             │ No HTTP, no SQL, no json tags                              │
├─────────────┼────────────────────────────────────────────────────────────┤
│ application/│ WHAT the app DOES per action (RegisterUser.Exec)           │
│             │ Business rules (email exists? admin only?)                 │
│             │ application/errors (Kind → HTTP later)                     │
│             │ application/contracts (MessagePublisher, RunInTx)          │
├─────────────┼────────────────────────────────────────────────────────────┤
│ infra/      │ HOW things talk to the world                               │
│             │ HTTP handlers, middleware, validator                       │
│             │ PostgreSQL repos, config, Wire bootstrap                   │
│             │ runtime (Runners), health checks                           │
├─────────────┼────────────────────────────────────────────────────────────┤
│ cmd/        │ START the app — only place that knows all concrete types   │
└─────────────┴────────────────────────────────────────────────────────────┘


Use case pattern (one folder per feature):

    application/usecase/register-user/
        register-user.go   → UseCase interface + Exec()
        contract.go        → Request / Result types


Handler depends on UseCase interface — not on postgres package.
*/

// =============================================================================
// 5. ONE REQUEST: POST /register
// =============================================================================

/*
Trace this until you can draw it from memory.


  Client
    │  POST /register  {"email":"...","password":"..."}
    ▼
  middlewares          logging, recovery (optional auth)
    ▼
  handler              1. decode JSON
                       2. validate tags  → 400
                       3. call useCase.Exec()
    ▼
  use case             1. email exists?  → 409
                       2. hash password
                       3. repo.Create(user)
    ▼
  repository           INSERT INTO users ...
    ▼
  handler              map result → JSON → 201


Who does what:

    Handler     → HTTP only (decode, validate, respond)
    Use case    → business rules
    Repository  → SQL only
    Domain      → entity + invariants


Nothing skips a layer.
*/

// =============================================================================
// 6. IMPORT RULES
// =============================================================================

/*
┌──────────────────────────┬─────────────────────────────────────────────┐
│ Package                  │ Can import                                  │
├──────────────────────────┼─────────────────────────────────────────────┤
│ domain/                  │ stdlib only                                 │
│ application/             │ domain/                                     │
│ infra/adapters/          │ domain/, database/sql                       │
│ infra/httpserver/        │ application/, domain/, net/http             │
│ infra/bootstrap/, cmd/   │ everything (wiring only)                    │
└──────────────────────────┴─────────────────────────────────────────────┘

NEVER:
  handler  → database/sql
  use case → net/http
  domain   → infra/


If these rules break, the architecture breaks.
*/

// =============================================================================
// 7. THREE TYPES OF VALIDATION
// =============================================================================

/*
Three types. Three places. Do not mix them.


┌──────┬─────────────────┬──────────────────────────┬─────────────────────┐
│ Type │ When            │ Where                    │ Example             │
├──────┼─────────────────┼──────────────────────────┼─────────────────────┤
│ 1    │ App startup     │ infra/config/            │ DATABASE_URL req    │
│      │                 │                          │ Fail fast — no boot │
├──────┼─────────────────┼──────────────────────────┼─────────────────────┤
│ 2    │ HTTP request    │ httpserver/.../requests  │ validate:"email"    │
│      │                 │                          │ → 400 Bad Request   │
├──────┼─────────────────┼──────────────────────────┼─────────────────────┤
│ 3    │ Business logic  │ application/usecase/     │ email already exists│
│      │                 │                          │ → 409 Conflict    │
└──────┴─────────────────┴──────────────────────────┴─────────────────────┘


Handler order:  decode → validate tags → useCase.Exec()
*/

// =============================================================================
// 8. ERRORS AND DTOs
// =============================================================================

/*
ERRORS — change shape per layer, map once at HTTP boundary:

    domain       → ErrEmailTaken (sentinel)
    use case     → apperror.Error with Kind (InvalidInput, NotFound, ...)
    handler      → WriteAppError(err) → 400 / 401 / 404 / 500

One WriteAppError function. Not switch in every handler.


THREE STRUCT TYPES — do not reuse one struct for everything:

    HTTP Request     infra/httpserver/user/requests.go
                     json tags + validate tags
                     {"email","password"} from client

    Use case contract application/usecase/register-user/contract.go
                     plain fields, no json tags
                     already validated at HTTP layer

    Domain entity    domain/user/entity.go
                     business object, passwordHash never in JSON

    Flow:  JSON → Request DTO → contract → entity → SQL
*/

// =============================================================================
// 9. WHERE DO I PUT X?
// =============================================================================

/*
┌─────────────────────────────┬──────────────────────────────────────────┐
│ I need to…                  │ Go to…                                   │
├─────────────────────────────┼──────────────────────────────────────────┤
│ Start app                   │ cmd/api/main.go                          │
│ Env config                  │ infra/config/                            │
│ DB connection               │ infra/database/                          │
│ SQL queries                 │ infra/adapters/repository/               │
│ Entity + repo interface     │ domain/<feature>/                        │
│ Business logic              │ application/usecase/<feature>/         │
│ Shared ports (Kafka, Tx)    │ application/contracts/                   │
│ App error kinds             │ application/errors/                      │
│ Routes + handler + JSON     │ infra/httpserver/<feature>/              │
│ Middleware                  │ infra/httpserver/middlewares/              │
│ Health checks               │ infra/httpserver/health.go               │
│ Wire / DI                   │ infra/bootstrap/                         │
│ Graceful shutdown           │ infra/runtime/ + httpserver/server.go    │
│ Migrations                  │ migrations/                              │
│ Docs for new devs           │ ARCHITECTURE.md                          │
└─────────────────────────────┴──────────────────────────────────────────┘
*/

// =============================================================================
// 10. FOLDER TREE
// =============================================================================

/*
go-backend-template/

  cmd/api/main.go

  domain/user/
      entity.go · errors.go · repository.go

  application/
      contracts/ · errors/
      usecase/register-user/

  infra/
      config/ · database/ · bootstrap/ · runtime/
      adapters/repository/
      httpserver/
          router.go · health.go · middlewares/ · user/

  migrations/
  docker-compose.yml · .env.example · Makefile
  README.md · ARCHITECTURE.md

Topic 02 → create this tree.
Topics 03–14 → fill with code.
Topic 15 → template ready to clone.
*/

// =============================================================================
// 11. WIRING, RUNNERS, HEALTH
// =============================================================================

/*
COMPOSITION ROOT (cmd/api/main.go or Wire):

    config → db → repo → use case → handler → router → server → app.Run()


Manual wiring chain:

    db
     └── userRepo
          └── registerUC
               └── userHandler
                    └── router → server


RUNNER — anything that runs until shutdown:

    type Runner interface {
        Start(ctx context.Context) error
        Stop(ctx context.Context) error
        Name() string
    }

    Today:  HTTP server is a Runner
    Later:  Kafka consumer in cmd/worker/ (same use cases)


HEALTH (Kubernetes):

    GET /livez   → process alive?        always 200
    GET /readyz  → DB connected?         503 if postgres down

    Indicator interface: Name(), Ready(ctx)
*/

// =============================================================================
// 12. CTO PROJECT MAP
// =============================================================================

/*
Study conning-configuration side-by-side:

┌─────────────────────────────┬──────────────────────────────────────────┐
│ CTO folder                  │ Your template                            │
├─────────────────────────────┼──────────────────────────────────────────┤
│ cmd/main/main.go            │ cmd/api/main.go                          │
│ domain/conningconfig/       │ domain/<feature>/                        │
│ application/usecase/.../    │ application/usecase/<feature>/           │
│ infra/httpserver/.../       │ infra/httpserver/<feature>/              │
│ infra/adapters/repository/  │ infra/adapters/repository/               │
│ infra/bootstrap/wire_gen.go │ infra/bootstrap/wire_gen.go              │
│ infra/runtime/app.go        │ infra/runtime/app.go                     │
│ MongoDB                     │ PostgreSQL (same layers, different DB)   │
└─────────────────────────────┴──────────────────────────────────────────┘

Study order: main → wire_gen → routes → handler → usecase → domain → adapter
*/

// =============================================================================
// 13. ADD A NEW ENDPOINT (CHECKLIST)
// =============================================================================

/*
Example: POST /products

  1. domain/product/           entity, errors, repository.go (interface)
  2. application/usecase/      create-product/ with UseCase + Exec()
  3. infra/adapters/repository/ product_repo.go (PostgreSQL)
  4. infra/httpserver/product/ routes, handler, requests, responses
  5. infra/httpserver/router.go  register product routes
  6. cmd or bootstrap/         wire repo → use case → handler
  7. migrations/               SQL schema
  8. test use case             fake repository, no HTTP, no DB

Repeat for every feature. Same steps every time.
*/

// =============================================================================
// 14. TESTING & COMMON MISTAKES
// =============================================================================

/*
TESTING:

    domain        → unit test pure functions
    use case      → fake Repository, test Exec() directly  (most tests here)
    repository    → integration test with real postgres (docker)
    handler       → httptest + mock UseCase


COMMON MISTAKES:

    ✗ SQL in handler              → use repository
    ✗ email exists in validator   → use case business rule
    ✗ domain imports infra        → interface in domain, impl in infra
    ✗ one struct for JSON+SQL     → separate DTO, contract, entity
    ✗ global var db               → inject in constructor
    ✗ empty folders, no example   → User feature fully implemented
    ✗ handler creates own repo    → composition root wires dependencies
*/

// =============================================================================
// 15. PHASE 10 VS 11
// =============================================================================

/*
Phase 10 — Template
    Skeleton + rules + User (register/login) + docs
    Deliverable: go-backend-template anyone can clone

Phase 11 — Capstone
    Products, orders, payments on TOP of template
    Do NOT redesign architecture — follow the checklist
*/

func main() {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║  01 — ARCHITECTURE DECISIONS                             ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Read this file section by section (comments above).")
	fmt.Println()
	fmt.Println("Core idea:")
	fmt.Println("  HTTP → use case → domain ← repository (postgres)")
	fmt.Println("  cmd/main wires everything")
	fmt.Println()
	fmt.Println("Layers:")
	fmt.Println("  domain/       entities + repo interfaces")
	fmt.Println("  application/  use cases + business rules")
	fmt.Println("  infra/        HTTP, DB, config, bootstrap")
	fmt.Println("  cmd/          entry point")
	fmt.Println()
	fmt.Println("Next: 02-project-layout")
	fmt.Println("Study: conning-configuration (CTO project)")
	fmt.Println()
}

/*
================================================================================
16. QUICK Q&A
================================================================================

Q: Why not one main.go?
A: Teams need boundaries — testable, maintainable, onboardable.

Q: Where does SQL live?
A: infra/adapters/repository/ only.

Q: Where does validate:"email" go?
A: infra/httpserver/<feature>/requests.go → 400.

Q: Where does "email already exists" go?
A: application/usecase/ → 409.

Q: Who wires repo → use case → handler?
A: cmd/api/main.go or infra/bootstrap/ (Wire).

Q: Can handler import postgres?
A: No. Handler → UseCase interface → Repository interface.

Q: Why 3 struct types (DTO, contract, entity)?
A: API, business, and persistence change independently.

Q: Repository interface in domain or application?
A: domain/ (CTO style).

Q: livez vs readyz?
A: livez = alive. readyz = DB up.

Q: Phase 10 vs 11?
A: 10 = template + User. 11 = e-commerce on template.


================================================================================
RUN:  go run ./10-production-template/01-architecture-decisions/lecture/
================================================================================
*/
