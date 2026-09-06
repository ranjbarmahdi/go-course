package main

/*
============================================================
42 — Context
============================================================

Context is used to control the lifetime of work.

It is especially useful when working with:

	- HTTP requests
	- Database queries
	- External APIs
	- Redis
	- Goroutines
	- Timeouts
	- Deadlines
	- Cancellation
	- Request-scoped values

Main idea:

	The caller creates and controls the Context.
	Functions receive the Context and should respect it.

============================================================
1. Why Context?
============================================================

Imagine an HTTP request:

	Client
	   |
	   v
	HTTP Handler
	   |
	   v
	Service
	   |
	   +----> Database
	   |
	   +----> External API
	   |
	   +----> Redis

If the client disconnects or the request times out,
we don't want all the work to continue unnecessarily.

Context allows us to signal:

	"Stop working because this request is no longer needed."

Context does NOT forcefully kill goroutines.

Instead, it provides a signal that cooperating
operations can observe and stop themselves.

============================================================
2. context.Context
============================================================

Context is an interface from the standard library:

	import "context"

A Context can carry:

	- Cancellation signals
	- Deadlines
	- Timeouts
	- Request-scoped values

Usually Context is passed as the first parameter:

	func doSomething(ctx context.Context) error {
		// ...
		return nil
	}

The Context should normally flow from the caller
down through the application layers.

Example:

	Handler
	   |
	   v
	Service
	   |
	   v
	Repository

Each layer receives and passes the same Context.

============================================================
3. context.Background()
============================================================

Background creates an empty root Context.

It is commonly used:

	- In main()
	- At application startup
	- As the root of a context tree
	- In tests

Example:

	ctx := context.Background()

Important:

	Background() is never cancelled,
	has no deadline,
	and contains no values.

It is usually the starting point for
a Context tree.

============================================================
4. context.TODO()
============================================================

TODO returns an empty Context like Background(),
but semantically means:

	"I don't know which Context should be used yet."

Example:

	ctx := context.TODO()

Use TODO when:

	- The correct Context source has not been decided yet.
	- You are temporarily working on code that will later
	  receive a proper Context.

Do not use TODO as a permanent replacement
for a proper Context.

============================================================
5. context.WithCancel()
============================================================

WithCancel creates:

	1. A child Context
	2. A cancel function

Example:

	ctx, cancel := context.WithCancel(
		context.Background(),
	)

	defer cancel()

The cancel function manually cancels the Context.

Example:

	ctx, cancel := context.WithCancel(
		context.Background(),
	)

	go worker(ctx)

	time.Sleep(1 * time.Second)

	cancel()

When cancel() is called:

	ctx.Done() becomes ready.

Cancellation also propagates to
all child Contexts.

============================================================
6. ctx.Done()
============================================================

Done() returns a channel.

When the Context is cancelled,
the channel is closed.

This allows us to use select:

	select {
	case <-ctx.Done():
		fmt.Println("Cancelled")

	case result := <-resultCh:
		fmt.Println("Result:", result)
	}

Mental model:

	ctx.Done()

means:

	"Should I stop?"

Important:

	Done() does not contain the cancellation reason.

	It only provides a signal that the Context
	has been cancelled or its deadline has expired.

============================================================
7. ctx.Err()
============================================================

Err() tells us why the Context finished.

Before cancellation:

	ctx.Err() == nil

After manual cancellation:

	ctx.Err() == context.Canceled

After timeout or deadline:

	ctx.Err() == context.DeadlineExceeded

Example:

	if err := ctx.Err(); err != nil {
		fmt.Println("Context ended:", err)
	}

Mental model:

	ctx.Done() → "Should I stop?"

	ctx.Err()  → "Why did I stop?"

============================================================
8. Cancellation Is Cooperative
============================================================

Context does NOT kill a goroutine.

This is very important.

Calling:

	cancel()

does NOT forcefully terminate:

	go worker(ctx)

The goroutine must check the Context
and stop itself.

Example:

	func worker(ctx context.Context) {

		for {
			select {

			case <-ctx.Done():
				fmt.Println("Worker stopped")
				return

			default:
				fmt.Println("Working...")
			}
		}
	}

The worker cooperates by listening to ctx.Done().

The important idea is:

	Context sends the signal.

	The goroutine decides how to respond.

============================================================
9. Context + Select
============================================================

A very common pattern is:

	select {
	case <-ctx.Done():
		return ctx.Err()

	case result := <-work:
		_ = result
		return nil
	}

This means:

	Wait for either:

		1. The work to finish
		2. The Context to be cancelled

Whichever case becomes ready is selected.

This pattern is extremely common in
concurrent Go programs.

============================================================
10. context.WithTimeout()
============================================================

WithTimeout automatically cancels a Context
after a specific duration.

Example:

	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)

	defer cancel()

The Context will automatically be cancelled
after 2 seconds.

Even though the timeout automatically cancels,
we still normally use:

	defer cancel()

This allows the resources associated with the
Context to be released as soon as the function
finishes instead of waiting for the timeout.

============================================================
11. Timeout + Select
============================================================

Example:

	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)

	defer cancel()

	select {

	case result := <-resultCh:
		fmt.Println("Result:", result)

	case <-ctx.Done():
		fmt.Println("Operation timed out")
	}

If resultCh produces a result within 1 second:

	Result is received.

If not:

	ctx.Done() becomes ready.

============================================================
12. Timeout Does Not Kill Goroutines
============================================================

Consider:

	func worker(
		ctx context.Context,
		resultCh chan<- string,
	) {
		select {

		case <-time.After(5 * time.Second):
			resultCh <- "Success"

		case <-ctx.Done():
			fmt.Println("Worker cancelled")
			return
		}
	}

And:

	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)

	defer cancel()

	go worker(ctx, resultCh)

The worker can perform work for up to 5 seconds.

But the Context expires after 1 second.

The worker receives:

	<-ctx.Done()

and returns.

The Context tells the worker:

	"Stop."

The worker decides to stop.

============================================================
13. context.WithDeadline()
============================================================

WithDeadline cancels a Context at a specific time.

Example:

	deadline := time.Now().Add(5 * time.Second)

	ctx, cancel := context.WithDeadline(
		context.Background(),
		deadline,
	)

	defer cancel()

WithTimeout:

	"Cancel after 5 seconds."

WithDeadline:

	"Cancel at this specific time."

Conceptually:

	context.WithTimeout(ctx, duration)

is equivalent to creating a deadline around:

	time.Now().Add(duration)

============================================================
14. WithCancel vs WithTimeout vs WithDeadline
============================================================

WithCancel:

	Manual cancellation.

	ctx, cancel := context.WithCancel(parent)

	cancel()

--------------------------------------------

WithTimeout:

	Automatic cancellation after a duration.

	ctx, cancel := context.WithTimeout(
		parent,
		5*time.Second,
	)

--------------------------------------------

WithDeadline:

	Automatic cancellation at a specific time.

	ctx, cancel := context.WithDeadline(
		parent,
		deadline,
	)

Mental model:

	WithCancel
		↓
	"I decide when to stop."

	WithTimeout
		↓
	"Stop after this amount of time."

	WithDeadline
		↓
	"Stop at this specific time."

============================================================
15. Parent → Child Context
============================================================

Contexts form a tree.

Example:

	Background
	    |
	    v
	Request Context
	    |
	    +--------> Database Context
	    |
	    +--------> API Context
	    |
	    +--------> Worker Context

If a parent Context is cancelled,
its children are also cancelled.

Example:

	parent, cancel := context.WithCancel(
		context.Background(),
	)

	child, childCancel := context.WithCancel(parent)

	defer cancel()
	defer childCancel()

If:

	cancel()

is called,

the cancellation propagates:

	parent
	   |
	   +----> child

Both Contexts become cancelled.

Important:

	A child cannot outlive its parent.

============================================================
16. Request Context
============================================================

In an HTTP server, every request already provides
a Context:

	func handler(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		service(ctx)
	}

The request Context represents the lifetime of
the HTTP request.

If the request is cancelled, the Context is cancelled.

The Context can then be passed through the application:

	HTTP Handler
	    |
	    v
	Service
	    |
	    v
	Repository
	    |
	    v
	Database

Example:

	func handler(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		err := service(ctx)

		if err != nil {
			// handle error
			return
		}
	}

	func service(ctx context.Context) error {

		return repository(ctx)
	}

	func repository(ctx context.Context) error {

		// Database operation using ctx

		return nil
	}

============================================================
17. Real Backend Example
============================================================

A request might need:

	- Database
	- External API
	- Redis

We can derive a timeout Context:

	ctx, cancel := context.WithTimeout(
		request.Context(),
		3*time.Second,
	)

	defer cancel()

Then pass it down:

	database.Query(ctx)

	externalAPI.Call(ctx)

	redis.Get(ctx)

Now all operations are connected to the
same request lifetime and timeout.

The flow becomes:

	HTTP Request
	    |
	    v
	Request Context
	    |
	    v
	3 Second Timeout
	    |
	    +------> Database
	    |
	    +------> External API
	    |
	    +------> Redis

If the request is cancelled or the timeout expires,
all cooperating operations can receive the
cancellation signal.

============================================================
18. Context-Aware Function
============================================================

A function that performs cancellable work
should usually accept Context:

	func operation(ctx context.Context) error {

		select {

		case <-ctx.Done():
			return ctx.Err()

		case result := <-work:
			_ = result

			return nil
		}
	}

This allows the caller to control
the lifetime of the operation.

============================================================
19. context.WithValue()
============================================================

WithValue allows a Context to carry
request-scoped values.

Example:

	type contextKey string

	const requestIDKey contextKey = "requestID"

	ctx := context.WithValue(
		context.Background(),
		requestIDKey,
		"abc-123",
	)

	requestID := ctx.Value(requestIDKey)

	fmt.Println(requestID)

WithValue creates a child Context
containing the value.

============================================================
20. Request-Scoped Values
============================================================

WithValue is useful for values that belong
to the lifetime of a request.

Examples:

	- Request ID
	- Trace ID
	- Correlation ID
	- Request-scoped authentication metadata

Example:

	ctx := context.WithValue(
		ctx,
		requestIDKey,
		"request-123",
	)

Then deeper functions can access it:

	requestID := ctx.Value(requestIDKey)

This can be useful when the value is
request-scoped metadata needed across
multiple layers.

============================================================
21. Context Values Are Inherited
============================================================

A child Context inherits values from its parent.

Example:

	parent := context.WithValue(
		context.Background(),
		requestIDKey,
		"abc-123",
	)

	child, cancel := context.WithCancel(parent)

	defer cancel()

	fmt.Println(child.Value(requestIDKey))

Output:

	abc-123

The value flows down the Context tree.

============================================================
22. Context Keys
============================================================

Avoid using plain strings as Context keys:

	context.WithValue(
		ctx,
		"requestID",
		value,
	)

Instead, use a custom key type:

	type contextKey string

	const requestIDKey contextKey = "requestID"

Why?

	Custom key types reduce the possibility
	of key collisions between packages.

Better:

	type contextKey string

	const requestIDKey contextKey = "requestID"

============================================================
23. Value Type Assertion
============================================================

Context.Value() returns any.

Therefore, when retrieving a value,
you can use a type assertion.

Example:

	requestID, ok := ctx.Value(requestIDKey).(string)

	if !ok {
		return
	}

	fmt.Println(requestID)

The ok value tells us whether:

	- The value exists
	- The value has the expected type

============================================================
24. Don't Abuse WithValue()
============================================================

Context values should be used for
request-scoped metadata.

Do NOT use Context as a general-purpose
storage container.

Bad:

	ctx = context.WithValue(ctx, "userName", "Mahdi")
	ctx = context.WithValue(ctx, "age", 27)
	ctx = context.WithValue(ctx, "limit", 100)
	ctx = context.WithValue(ctx, "page", 10)

If a function needs a normal business parameter,
pass it normally:

	func getUsers(
		ctx context.Context,
		page int,
		limit int,
	) {
		// ...
	}

Use Context mainly for:

	- Cancellation
	- Deadlines
	- Timeouts
	- Request-scoped values

============================================================
25. HTTP → Service → Repository
============================================================

A common backend architecture:

	HTTP Handler
	     |
	     | ctx
	     v
	Service
	     |
	     | ctx
	     v
	Repository
	     |
	     | ctx
	     v
	Database

Example:

	func handler(
		w http.ResponseWriter,
		r *http.Request,
	) {

		ctx := r.Context()

		err := userService(ctx)

		if err != nil {
			// handle error
			return
		}
	}

	func userService(ctx context.Context) error {

		return userRepository(ctx)
	}

	func userRepository(ctx context.Context) error {

		// db.QueryContext(ctx, ...)

		return nil
	}

The Context travels through the call chain.

============================================================
26. Context and Database
============================================================

Database libraries commonly provide
Context-aware operations.

For example:

	db.QueryContext(ctx, query)

instead of:

	db.Query(query)

The Context allows the database operation
to react to cancellation and deadlines.

Typical flow:

	HTTP Request
	    |
	    v
	Request Context
	    |
	    v
	Service
	    |
	    v
	Repository
	    |
	    v
	Database Query
	    |
	    v
	Database

============================================================
27. Context and External APIs
============================================================

HTTP clients can also use Context.

Example:

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)

	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)

If ctx is cancelled,
the HTTP request can be cancelled as well.

This prevents unnecessary work after
the original request is no longer needed.

============================================================
28. Context and Goroutines
============================================================

A goroutine should listen to the Context
when its lifetime is tied to a request.

Example:

	func worker(ctx context.Context) {

		for {
			select {

			case <-ctx.Done():
				fmt.Println("Worker stopped")
				return

			case <-time.After(500 * time.Millisecond):
				fmt.Println("Working...")
			}
		}
	}

The worker continues until:

	- Context is cancelled
	- Context times out
	- Context deadline is reached

The worker must cooperate with the Context.

============================================================
29. Common Context Pattern
============================================================

A very common backend pattern is:

	func operation(ctx context.Context) error {

		select {

		case <-ctx.Done():
			return ctx.Err()

		case result := <-work:
			_ = result

			return nil
		}
	}

This pattern combines:

	- Context
	- Cancellation
	- Channels
	- Select
	- Error handling

============================================================
30. Important Rules
============================================================

Rule 1:

	Pass Context explicitly.

	func service(ctx context.Context)

--------------------------------------------

Rule 2:

	Usually Context is the first parameter.

	func service(ctx context.Context, id int)

--------------------------------------------

Rule 3:

	Do not store Context inside structs.

	Pass it explicitly to functions instead.

--------------------------------------------

Rule 4:

	Do not pass a nil Context.

	Use context.Background(),
	context.TODO(),
	or another appropriate Context.

--------------------------------------------

Rule 5:

	Call cancel when you create a
	cancellable Context.

	ctx, cancel := context.WithTimeout(...)
	defer cancel()

--------------------------------------------

Rule 6:

	Functions performing cancellable work
	should respect ctx.Done().

--------------------------------------------

Rule 7:

	Do not use Context values for normal
	business parameters.

--------------------------------------------

Rule 8:

	Do not create a new Background Context
	deep inside your application when you
	already have a Context.

	Prefer:

	func service(ctx context.Context)

	instead of:

	func service() {
		ctx := context.Background()
	}

This preserves cancellation and deadlines
from the original caller.

============================================================
31. Complete Backend Mental Model
============================================================

Context is a lifetime-control mechanism.

The caller says:

	"Here is the Context for this work."

The lower-level function says:

	"I will respect this Context."

Example:

	HTTP Request
	    |
	    v
	Request Context
	    |
	    v
	Service
	    |
	    +------> Database
	    |
	    +------> External API
	    |
	    +------> Redis
	    |
	    +------> Goroutine

If the request is cancelled:

	Context
	    |
	    v
	Cancellation Signal
	    |
	    +------> Database
	    |
	    +------> External API
	    |
	    +------> Redis
	    |
	    +------> Goroutine

Each cooperating operation can stop.

============================================================
32. The Most Important APIs
============================================================

context.Background()

	Root Context.

--------------------------------------------

context.TODO()

	Temporary/unknown Context.

--------------------------------------------

context.WithCancel()

	Manual cancellation.

--------------------------------------------

context.WithTimeout()

	Cancel after a duration.

--------------------------------------------

context.WithDeadline()

	Cancel at a specific time.

--------------------------------------------

context.WithValue()

	Attach request-scoped values.

--------------------------------------------

ctx.Done()

	Channel that signals cancellation.

--------------------------------------------

ctx.Err()

	Reason why the Context ended.

--------------------------------------------

ctx.Value(key)

	Retrieve a request-scoped value.

============================================================
33. Final Mental Model
============================================================

Think about Context as information flowing
down the application:

	Caller
	  |
	  v
	Context
	  |
	  +--------> Handler
	  |             |
	  |             v
	  |          Service
	  |             |
	  |       +-----+-----+
	  |       |     |     |
	  |       v     v     v
	  |      DB   Redis   API
	  |
	  +--------> Goroutine

The Context carries:

	- Cancellation
	- Deadline
	- Timeout
	- Request-scoped values

The Context does NOT:

	- Kill goroutines
	- Store arbitrary application state
	- Replace normal function parameters
	- Replace configuration

The Context sends the signal:

	"Your work is no longer needed."

Each operation must cooperate and stop itself.

============================================================
Key Concepts to Remember
============================================================

	context.Background()
	context.TODO()

	context.WithCancel()
	context.WithTimeout()
	context.WithDeadline()
	context.WithValue()

	ctx.Done()
	ctx.Err()
	ctx.Value()

	Cancellation
	Timeout
	Deadline
	Request-scoped values
	Parent → Child propagation
	Cooperative cancellation

============================================================
End of 42 — Context
============================================================
*/

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type contextKey string

const requestIDKey contextKey = "requestID"

func main() {
	fmt.Println("42 — Context")

	// ========================================================
	// 1. Background
	// ========================================================

	ctx := context.Background()

	fmt.Println("Background:", ctx)

	// ========================================================
	// 2. TODO
	// ========================================================

	todoCtx := context.TODO()

	fmt.Println("TODO:", todoCtx)

	// ========================================================
	// 3. WithCancel
	// ========================================================

	cancelCtx, cancel := context.WithCancel(
		context.Background(),
	)

	fmt.Println("Before cancel:", cancelCtx.Err())

	cancel()

	fmt.Println("After cancel:", cancelCtx.Err())

	// ========================================================
	// 4. WithTimeout
	// ========================================================

	timeoutCtx, timeoutCancel := context.WithTimeout(
		context.Background(),
		100*time.Millisecond,
	)

	defer timeoutCancel()

	<-timeoutCtx.Done()

	fmt.Println("Timeout error:", timeoutCtx.Err())

	// ========================================================
	// 5. WithDeadline
	// ========================================================

	deadline := time.Now().Add(100 * time.Millisecond)

	deadlineCtx, deadlineCancel := context.WithDeadline(
		context.Background(),
		deadline,
	)

	defer deadlineCancel()

	<-deadlineCtx.Done()

	fmt.Println("Deadline error:", deadlineCtx.Err())

	// ========================================================
	// 6. WithValue
	// ========================================================

	valueCtx := context.WithValue(
		context.Background(),
		requestIDKey,
		"request-123",
	)

	requestID, ok := valueCtx.Value(requestIDKey).(string)

	if ok {
		fmt.Println("Request ID:", requestID)
	}

	// ========================================================
	// 7. HTTP Request Context
	// ========================================================

	req, err := http.NewRequest(
		http.MethodGet,
		"https://example.com",
		nil,
	)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	requestCtx := req.Context()

	fmt.Println("HTTP request context:", requestCtx)

	// ========================================================
	// 8. Context + Select
	// ========================================================

	work := make(chan string)

	ctx, cancel = context.WithTimeout(
		context.Background(),
		100*time.Millisecond,
	)

	defer cancel()

	select {
	case result := <-work:
		fmt.Println("Result:", result)

	case <-ctx.Done():
		fmt.Println("Work cancelled:", ctx.Err())
	}

	fmt.Println("Done")
}
