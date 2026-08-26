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


├── 18-fmt

Topics:

- Print
- Println
- Printf
- Sprintf
- Formatting verbs


├── 19-strings

Topics:

- Contains
- Split
- Join
- Replace
- Trim
- Fields
- Builder


├── 20-strconv

Topics:

- Atoi
- Itoa
- ParseFloat
- FormatFloat
- ParseInt
- ParseBool


├── 21-time ⭐ CURRENT

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


├── 22-os

Topics:

- Environment variables
- Files
- Arguments

Examples:

- os.Getenv()
- os.ReadFile()
- os.WriteFile()


├── 23-io

Topics:

- Reader
- Writer
- Copy
- ReadAll

Used in:

- Files
- HTTP
- Network

├── 24-encoding-json ⭐⭐⭐

Topics:

- Marshal
- Unmarshal
- Struct tags
- JSON DTO



├── 25-bufio

Topics:

- Scanner
- Reader
- Writer
- Buffering


├── 26-filepath

Topics:

- Join
- Dir
- Base
- Ext
- Walk


├── 27-regexp

Topics:

- Match
- Find
- Replace
- Validation


├── 28-sort

Topics:

- Sort slices
- Custom sorting
- sort.Slice()


├── 29-flag

Topics:

- CLI arguments
- Command flags



====================================
PHASE 3 — TESTING
====================================


03-testing


├── 30-testing-package

Topics:

- _test.go
- Test functions
- Assertions
- Table driven tests
- Benchmarks
- Coverage


├── 31-mocking

Topics:

- Fake repository
- Interface mocks
- Dependency testing


└── 32-integration-testing

Topics:

- Database tests
- HTTP tests
- Docker test environment



====================================
PHASE 4 — ADVANCED GO
====================================


04-advanced-go


├── 33-advanced-structs

Topics:

- Embedding
- Composition
- Struct tags


├── 34-generics

Topics:

- Type parameters
- Constraints
- Generic functions
- Generic data structures


├── 35-reflection

Topics:

- reflect package
- Runtime inspection
- When not to use reflection


└── 36-memory-management

Topics:

- Garbage collector
- Stack vs Heap
- Escape analysis



====================================
PHASE 5 — CONCURRENCY ⭐⭐⭐
====================================


05-concurrency


├── 37-goroutines

Topics:

- Concurrency
- Goroutine lifecycle


├── 38-channels

Topics:

- Send
- Receive
- Buffered channels
- Closing channels


├── 39-select

Topics:

- Multiple channels
- Timeout
- Cancellation


├── 40-sync

Topics:

- Mutex
- RWMutex
- WaitGroup
- Once


└── 41-concurrency-patterns

Topics:

- Worker pool
- Fan-in
- Fan-out
- Pipeline



====================================
PHASE 6 — CONTEXT
====================================


06-context


└── 42-context


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


├── 43-http-server

Standard library:

net/http


Topics:

- Handler
- Router
- Request
- Response
- Middleware


├── 44-rest-api

Topics:

- DTO
- Validation
- Pagination
- Filtering
- Sorting


├── 45-json-api

Topics:

- Request models
- Response models
- Error responses


├── 46-middleware

Topics:

- Logging
- Recovery
- Authentication
- CORS


└── 47-authentication

Topics:

- JWT
- Refresh token
- Password hashing
- bcrypt



====================================
PHASE 8 — DATABASE
====================================


08-database


├── 48-database-sql

Topics:

- database/sql
- pgx
- Connection pool


├── 49-postgresql

Topics:

- Schema
- Migration
- Indexes
- Queries
- Transactions


└── 50-repository-pattern

Topics:

- Repository interface
- PostgreSQL implementation
- Unit of Work



====================================
PHASE 9 — PRODUCTION BACKEND
====================================


09-production


├── 51-configuration

Topics:

- Environment variables
- Config management


├── 52-logging

Topics:

- log/slog
- Structured logging


├── 53-caching

Topics:

- Redis
- TTL
- Cache strategies


├── 54-messaging

Topics:

- RabbitMQ
- Kafka basics
- Event driven design


├── 55-docker

Topics:

- Dockerfile
- Docker compose
- Multi stage builds


├── 56-deployment

Topics:

- Linux
- CI/CD
- Monitoring
- Health checks


├── 57-grpc

Topics:

- Protocol Buffers
- gRPC server
- gRPC client


├── 58-cli-tools

Topics:

- Cobra
- urfave/cli


├── 59-performance

Topics:

- pprof
- trace
- Benchmarking


└── 60-security

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
CURRENT POSITION
====================================


You are here:

02-standard-libraries

└── 23-io


Next:

24-encoding-json
