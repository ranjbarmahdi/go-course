# Go Backend Developer Roadmap

Goal:

Become a production-level Go Backend Developer.

Focus:

- Clean Architecture
- REST API
- PostgreSQL
- Redis
- RabbitMQ
- Docker
- Testing
- Production practices
- Production-ready project template


====================================
PROJECT STRUCTURE
====================================

go-crash

├── 01-Basics
│
├── 02-standard-libraries
│
├── 03-testing
│
├── 04-advanced-go
│
├── 05-concurrency
│
├── 06-context
│
├── 07-backend-development
│
├── 08-database
│
├── 09-production
│
├── 10-production-template
│
└── 11-capstone-project



====================================
PHASE 1 — GO FUNDAMENTALS ✅ DONE
====================================

Folder:

01-Basics


01-hello-go

02-variables-and-types

03-conditions-and-boolean-logic

04-switch

05-loops

06-arrays

07-slices

08-maps

09-functions

10-structs

11-methods

12-pointers

13-struct-pointers

14-interfaces

15-dependency-injection

16-packages

17-error-handling



====================================
PHASE 2 — STANDARD LIBRARIES ⭐ CURRENT
====================================


02-standard-libraries


├── 01-fmt

Topics:

- Print
- Println
- Printf
- Sprintf
- Formatting verbs


├── 02-strings

Topics:

- Contains
- Split
- Join
- Replace
- Trim
- Fields
- Builder


├── 03-strconv

Topics:

- Atoi
- Itoa
- ParseFloat
- FormatFloat
- ParseInt
- ParseBool


├── 04-time ⭐ CURRENT

Topics:

- time.Now()
- UTC
- Date extraction
- time.Date()
- Duration
- Sleep
- Add
- Sub
- Since
- Before
- After
- Equal
- Format
- Parse
- RFC3339
- Unix timestamp
- Timezone
- Timer
- Ticker


├── 05-os

Topics:

- Environment variables
- Files
- Arguments

Examples:

- os.Getenv()
- os.ReadFile()
- os.WriteFile()


├── 06-io

Topics:

- Reader
- Writer
- Copy
- ReadAll

Used in:

- Files
- HTTP
- Network

├── 07-encoding-json ⭐⭐⭐

Topics:

- Marshal
- Unmarshal
- Struct tags
- JSON DTO



├── 08-bufio

Topics:

- Scanner
- Reader
- Writer
- Buffering


├── 09-filepath

Topics:

- Join
- Dir
- Base
- Ext
- Walk


├── 10-regexp

Topics:

- Match
- Find
- Replace
- Validation


├── 11-sort

Topics:

- Sort slices
- Custom sorting
- sort.Slice()


├── 12-flag

Topics:

- CLI arguments
- Command flags



====================================
PHASE 3 — TESTING
====================================


03-testing


├── 01-testing-package

Topics:

- _test.go
- Test functions
- Assertions
- Table driven tests
- Benchmarks
- Coverage


├── 02-mocking

Topics:

- Fake repository
- Interface mocks
- Dependency testing


└── 03-integration-testing

Topics:

- Database tests
- HTTP tests
- Docker test environment



====================================
PHASE 4 — ADVANCED GO
====================================


04-advanced-go


├── 01-advanced-structs

Topics:

- Embedding
- Composition
- Struct tags


├── 02-generics

Topics:

- Type parameters
- Constraints
- Generic functions
- Generic data structures


├── 03-reflection

Topics:

- reflect package
- Runtime inspection
- When not to use reflection


└── 04-memory-management

Topics:

- Garbage collector
- Stack vs Heap
- Escape analysis



====================================
PHASE 5 — CONCURRENCY ⭐⭐⭐
====================================


05-concurrency


├── 01-goroutines

Topics:

- Concurrency
- Goroutine lifecycle


├── 02-channels

Topics:

- Send
- Receive
- Buffered channels
- Closing channels


├── 03-select

Topics:

- Multiple channels
- Timeout
- Cancellation


├── 04-sync

Topics:

- Mutex
- RWMutex
- WaitGroup
- Once


└── 05-concurrency-patterns

Topics:

- Worker pool
- Fan-in
- Fan-out
- Pipeline



====================================
PHASE 6 — CONTEXT
====================================


06-context


└── 01-context


Topics:

- context.Background()
- WithCancel()
- WithTimeout()
- WithDeadline()
- Request lifecycle
- Cancellation



====================================
PHASE 7 — BACKEND DEVELOPMENT
====================================


07-backend-development


├── 01-http-server

Standard library:

net/http


Topics:

- Handler
- Router
- Request
- Response
- Middleware


├── 02-rest-api

Topics:

- DTO
- Validation
- Pagination
- Filtering
- Sorting


├── 03-json-api

Topics:

- Request models
- Response models
- Error responses


├── 04-middleware

Topics:

- Logging
- Recovery
- Authentication
- CORS


├── 05-authentication

Topics:

- JWT
- Refresh token
- Password hashing
- bcrypt


└── 06-validation

Topics:

- validator/v10
- Struct tags
- Input validation
- Custom validators
- Validation vs business rules


====================================
PHASE 8 — DATABASE
====================================


08-database


├── 01-database-sql

Topics:

- database/sql
- pgx
- Connection pool


├── 02-postgresql

Topics:

- Schema
- Migration
- Indexes
- Queries
- Transactions


├── 03-repository-pattern

Topics:

- Repository interface
- PostgreSQL implementation
- Unit of Work


└── 04-handling-transactions

Topics:

- TransactionManager
- RunInTx
- Context-propagation pattern
- getExecutor



====================================
PHASE 9 — PRODUCTION BACKEND
====================================


09-production


├── 01-configuration

Topics:

- Environment variables
- Config management


├── 02-logging

Topics:

- log/slog
- Structured logging


├── 03-caching

Topics:

- Redis
- TTL
- Cache strategies


├── 04-messaging

Topics:

- RabbitMQ
- Kafka basics
- Event driven design


├── 05-docker

Topics:

- Dockerfile
- Docker compose
- Multi stage builds


├── 06-deployment

Topics:

- Linux
- CI/CD
- Monitoring
- Health checks


├── 07-grpc

Topics:

- Protocol Buffers
- gRPC server
- gRPC client


├── 08-cli-tools

Topics:

- Cobra
- urfave/cli


├── 09-performance

Topics:

- pprof
- trace
- Benchmarking


└── 10-security

Topics:

- Race detector
- govulncheck
- Secure coding



====================================
PHASE 10 — PRODUCTION TEMPLATE
====================================


Goal:

Learn to design and ship a Go backend template that other developers
can clone, run in 10 minutes, and extend with clear rules for every layer.

Reference architecture:

Study: conning-configuration (CTO project)
Build:  go-backend-template (used in Phase 11 capstone)


Folder:

10-production-template


├── 01-architecture-decisions

Topics:

- Layered vs vertical slice
- Domain / Application / Infra boundaries
- Import rules (who imports whom)
- Config vs input vs business validation
- When to add complexity (YAGNI)

Output:

ARCHITECTURE.md — "where do I put X?"


├── 02-project-layout

Topics:

- cmd/ for entry points
- domain/ for entities + repository interfaces
- application/usecase/ one folder per feature
- infra/ for HTTP, DB, messaging, config
- Standard Go Project Layout (practical subset)

Output:

Empty folder tree with README in each top folder


├── 03-domain-layer

Topics:

- Entities and aggregate roots
- Value objects
- Domain errors
- Repository interfaces IN domain
- Domain events (optional)

Practice:

User aggregate + UserRepository interface


├── 04-application-layer

Topics:

- UseCase interface + Exec(ctx, ...) per feature
- contract.go — use case input/output DTOs
- application/errors — Kind (InvalidInput, NotFound, Unauthorized, Internal)
- application/contracts — MessagePublisher, RunInTx
- Business rules live HERE (not in validator tags)

Practice:

RegisterUser use case depending on domain.Repository


├── 05-presentation-http

Topics:

- router.go — register all feature routes
- <feature>/routes.go — mux.Handle + middleware
- <feature>/handler.go — decode, validate, call use case
- <feature>/requests.go — JSON + validate tags
- <feature>/responses.go — API response mapping
- middlewares/chain.go — global + per-route
- utils — WriteSuccess, WriteAppError (Kind → status)

Practice:

POST /register end-to-end through layers


├── 06-infrastructure-adapters

Topics:

- infra/adapters/repository/ — PostgreSQL impl
- infra/database/ — connection, ping, cleanup
- infra/config/ — env struct tags + Validate
- Mapping domain ↔ persistence (no SQL in use case)

Practice:

PostgresUserRepository implements domain.UserRepository


├── 07-composition-root-manual

Topics:

- cmd/api/main.go — Load config, slog, signal context
- Manual wiring: db → repo → use case → handler → router → server
- Constructor injection (NewXxx dependencies)
- defer cleanup() for DB close

Practice:

Wire User feature manually in main (no Wire yet)


├── 08-runtime-runners

Topics:

- Runner interface (Start, Stop, Name)
- App runs []Runner with errgroup
- HTTP Server as Runner
- Graceful shutdown (Shutdown with timeout)
- Adding cmd/worker later (same internal/, new runner)

Reference:

conning-configuration/infra/runtime/

Practice:

Extract server into runtime.Runner; App.Run(ctx)


├── 09-health-indicators

Topics:

- GET /livez — process alive
- GET /readyz — dependencies ready
- Indicator interface (Name, Ready)
- PostgresIndicator pings DB

Reference:

conning-configuration/infra/httpserver/health.go

Practice:

/readyz fails when PostgreSQL is down


├── 10-error-mapping

Topics:

- Domain error → application error → HTTP status
- Single WriteAppError switch on Kind
- Never leak internal errors to client
- Validation 400 vs business 409 vs auth 401

Practice:

Central error map used by all handlers


├── 11-bootstrap-wire

Topics:

- Why Wire (compile-time DI)
- wire.NewSet — DatabaseSet, AdapterSet, UseCaseSet, HttpSet
- wire.Bind(interface, implementation)
- wire_gen.go — generated, do not edit
- make wire / go generate

Reference:

conning-configuration/infra/bootstrap/

Practice:

Convert manual main wiring to Google Wire


├── 12-messaging-contract

Topics:

- application/contracts MessagePublisher
- infra/adapters/kafka-producer or rabbitmq
- Use case publishes event after success
- Reserved structure — wire when needed

Practice:

Stub MessagePublisher + noop impl for tests


├── 13-developer-experience

Topics:

- .env.example
- docker-compose.yml (postgres + app)
- Makefile (run, test, migrate, wire, lint)
- README — quick start in 5 commands
- CONTRIBUTING — how to add new endpoint (checklist)

Output:

New dev clones → docker compose up → curl /health works


├── 14-example-vertical-slice

Topics:

- One complete feature in template: User (register + login)
- Every layer has real code (not empty folders)
- Copy this pattern for new features

Checklist for new endpoint:

    1. domain/<feature>/entity.go + repository.go
    2. application/usecase/<feature>/
    3. infra/adapters/repository/<feature>_repo.go
    4. infra/httpserver/<feature>/ routes, handler, requests, responses
    5. infra/bootstrap/sets.go — add to Wire sets
    6. Test use case with fake repository


└── 15-template-finalization

Topics:

- Finalize go-backend-template repository
- User feature as living example in every layer
- Version tag v1.0.0
- Optional: GitHub template repository

Final deliverable:

go-backend-template/

    clone → cp .env.example .env → make up → make run
    developer adds new features using ARCHITECTURE.md checklist


Reference comparison:


| Concern       | go-crash lectures | CTO project     | Your template       |
|---------------|-------------------|-----------------|---------------------|
| Repo interface| application port  | domain/         | domain/ (CTO style) |
| Use case      | one file          | usecase/<name>/ | usecase/<name>/     |
| DI            | manual main       | Google Wire     | manual → Wire       |
| Health        | /health           | /livez /readyz  | /livez /readyz      |
| Errors        | scattered         | application/    | Kind + WriteAppError|
| Runners       | single server     | runtime.App     | runtime.App         |



====================================
PHASE 11 — CAPSTONE PROJECT
====================================


11-capstone-project


Production E-Commerce Backend


Goal:

Build a full backend ON TOP of the Phase 10 production template.
Add e-commerce features using the template checklist — do NOT redesign architecture.


Architecture (from Phase 10 template):


cmd/
└── api/
    └── main.go              ← entry, config, signal, bootstrap

domain/
├── user/
├── product/
├── order/
└── ...

application/
├── contracts/               ← shared ports (MessagePublisher, RunInTx)
├── errors/                  ← typed app errors (Kind → HTTP)
└── usecase/
    ├── register-user/
    ├── create-product/
    └── create-order/

infra/
├── config/
├── bootstrap/               ← Google Wire sets
├── database/
├── adapters/
│   └── repository/          ← PostgreSQL implementations
├── httpserver/
│   ├── middlewares/
│   ├── router.go
│   ├── health.go
│   └── <feature>/           ← routes, handler, requests, responses
└── runtime/                 ← App, Runner, Indicator


Features (vertical slices):


01-users          register, login, profile
02-products       CRUD
03-orders         create order (RunInTx)
04-payments       stub or simple flow
05-auth           JWT middleware
06-health         /livez, /readyz


Stack:


PostgreSQL
Redis (cache)
RabbitMQ (events)
Docker
slog
validator/v10


Deliverable:

Working e-commerce API — proves the template scales to a real product.


Practice per feature:

    Follow Phase 10 checklist — one vertical slice at a time



====================================
LEARNING PATH
====================================

Phases 7–8  → HTTP, auth, validation, DB, repository, transactions
Phase 9     → config, logging, docker, health, deployment
Phase 10    → production template (learn architecture + build skeleton)
Phase 11    → capstone e-commerce (build ON the template)

Study CTO project (conning-configuration) while doing Phase 10.



====================================
NEXT TOPIC TO LEARN
====================================

Finish remaining Phase 9 topics, then Phase 10 template, then Phase 11 capstone.

Immediate next:

Phase 9 — 02-logging
