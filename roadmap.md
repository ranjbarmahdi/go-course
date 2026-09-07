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
└── 10-capstone-project



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
PHASE 10 — FINAL PROJECT
====================================


10-capstone-project


Production E-Commerce Backend


Architecture:


cmd

internal

├── domain

├── application

├── infrastructure

└── presentation



Features:


Authentication

Users

Products

Orders

Payments

PostgreSQL

Redis

RabbitMQ

Docker

Testing

CI/CD



====================================
NEXT TOPIC TO LEARN
====================================

Phase 8 — 04-handling-transactions
