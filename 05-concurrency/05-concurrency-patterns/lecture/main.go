package main

import "fmt"

func main() {

	// ============================================================
	// 1. Worker Pool ⭐⭐⭐
	// ============================================================

	fmt.Println("1. Worker Pool")

	/*
		A Worker Pool is a concurrency pattern where we create a
		fixed number of workers instead of creating an unlimited
		number of goroutines.

		Without a worker pool:

			1000 jobs
			    │
			    ├── goroutine
			    ├── goroutine
			    ├── goroutine
			    ├── ...
			    └── goroutine

		This can create too many goroutines and consume too many
		resources.

		With a Worker Pool:

			     Jobs
			      │
		       ┌──────┼──────┐
		       ↓      ↓      ↓
		    Worker  Worker  Worker
		       1       2      3
		       │       │      │
		       └───────┼──────┘
			       ↓
			    Results

		For example:

			1000 jobs
			   ↓
			3 workers
			   ↓
			workers process jobs from the queue

		The number of workers is fixed.

		Backend use cases:

			- Processing uploaded files
			- Sending emails
			- Processing database jobs
			- Image processing
			- Background jobs
			- CPU-intensive tasks

		Main benefit:

			Control the amount of concurrency.
	*/

	// ============================================================
	// 2. Fan-Out ⭐⭐⭐
	// ============================================================

	fmt.Println("2. Fan-Out")

	/*
		Fan-Out means distributing work from one source to
		multiple goroutines.

			      Jobs
				│
				↓
			    Channel
				│
		       ┌────────┼────────┐
		       ↓        ↓        ↓
		   Worker 1  Worker 2  Worker 3

		Multiple workers consume jobs from the same source.

		The goal is to increase parallel processing.

		For example:

			100 jobs
			   ↓
			Worker 1
			Worker 2
			Worker 3
			Worker 4

		Instead of one worker processing all 100 jobs,
		multiple workers process them concurrently.

		Main idea:

			ONE source → MANY workers
	*/

	// ============================================================
	// 3. Fan-In ⭐⭐⭐
	// ============================================================

	fmt.Println("3. Fan-In")

	/*
		Fan-In is the opposite direction.

		Multiple goroutines produce results into a single channel.

			 Worker 1 ──┐
			 Worker 2 ──┼──> Results Channel
			 Worker 3 ──┘

		Then one consumer receives all results.

		Main idea:

			MANY producers → ONE channel

		Fan-Out and Fan-In are often used together:

			     Jobs
			      │
			      ↓
		       ┌─────────────┐
		       │   Fan-Out   │
		       └─────────────┘
			 │    │    │
			 ↓    ↓    ↓
			W1   W2   W3
			 │    │    │
			 └────┼────┘
			      ↓
		       ┌─────────────┐
		       │   Fan-In    │
		       └─────────────┘
			      │
			      ↓
			   Results

		This pattern is useful when:

			- Work can be processed independently.
			- Multiple workers should process the work.
			- Results need to be collected in one place.
	*/

	// ============================================================
	// 4. Pipeline ⭐⭐⭐
	// ============================================================

	fmt.Println("4. Pipeline")

	/*
		A Pipeline divides processing into multiple stages.

		For example:

			Input
			  │
			  ↓
			Read
			  │
			  ↓
			Process
			  │
			  ↓
			Save
			  │
			  ↓
			Output

		Each stage performs a different job.

		Channels connect the stages:

			Generator
			    │
			    ↓
			 Channel
			    │
			    ↓
			Processor
			    │
			    ↓
			 Channel
			    │
			    ↓
			  Saver

		Each stage can run concurrently.

		For example:

			Stage 1:
				Read data

			Stage 2:
				Process data

			Stage 3:
				Save data

		While Stage 2 processes item 1,
		Stage 1 can already process item 2.

		This creates a flow of work through the system.

		Backend use cases:

			- Data processing
			- ETL systems
			- File processing
			- Message processing
			- Multi-stage background jobs

		Main idea:

			Stage 1 → Stage 2 → Stage 3
	*/

	// ============================================================
	// 5. Bounded Concurrency ⭐⭐⭐
	// ============================================================

	fmt.Println("5. Bounded Concurrency")

	/*
		Sometimes we don't want unlimited concurrent operations.

		For example:

			1000 database requests

		Creating 1000 simultaneous database operations may be
		a bad idea.

		Instead:

			1000 requests
			      │
			      ↓
			 max 10 at a time
			      │
			      ↓
			   database

		This is called Bounded Concurrency.

		The goal is to limit how many operations can execute
		concurrently.

		A common implementation uses a buffered channel as
		a semaphore:

			sem := make(chan struct{}, 10)

		Before starting work:

			sem <- struct{}{}

		After finishing work:

			<-sem

		The channel capacity represents the maximum number
		of concurrent operations.

		For example:

			make(chan struct{}, 10)

		means:

			Maximum 10 operations at the same time.

		Backend use cases:

			- Limit database queries
			- Limit HTTP requests
			- Limit external API calls
			- Limit file processing
			- Protect external services

		Main idea:

			Unlimited work
			      ↓
			Concurrency limit
			      ↓
			Safe number of active operations
	*/

	// ============================================================
	// 6. Graceful Shutdown ⭐⭐⭐
	// ============================================================

	fmt.Println("6. Graceful Shutdown")

	/*
		A real backend application must be able to stop safely.

		For example:

			Application receives shutdown signal
					│
					↓
			Stop accepting new work
					│
					↓
			Finish active work
					│
					↓
			Close resources
					│
					↓
				Exit

		We don't want the application to suddenly terminate
		while work is still running.

		A graceful shutdown can involve:

			- HTTP server
			- Goroutines
			- Worker pools
			- Database connections
			- Message queues
			- External resources

		For example:

			HTTP Server
			    │
			    ↓
			Shutdown signal
			    │
			    ↓
			Stop accepting requests
			    │
			    ↓
			Wait for active requests
			    │
			    ↓
			Close database
			    │
			    ↓
			   Exit

		This becomes especially important when we learn
		the context package.

		We will later use:

			context.Context

		to propagate cancellation and shutdown signals
		through goroutines and backend operations.

		Main idea:

			Stop safely
			    ↓
			Finish active work
			    ↓
			Cleanup resources
			    ↓
			Exit
	*/

	// ============================================================
	// 7. Important Mental Model ⭐⭐⭐
	// ============================================================

	fmt.Println("7. Important Mental Model")

	/*
		Worker Pool
		    ↓
		Fixed number of workers processing jobs

		Fan-Out
		    ↓
		Distribute work among multiple workers

		Fan-In
		    ↓
		Combine multiple result sources

		Pipeline
		    ↓
		Multiple processing stages

		Bounded Concurrency
		    ↓
		Limit how many operations run simultaneously

		Graceful Shutdown
		    ↓
		Stop the application safely
	*/

	// ============================================================
	// 8. Concurrency Patterns Together ⭐⭐⭐
	// ============================================================

	fmt.Println("8. Concurrency Patterns Together")

	/*
		In real backend systems, these patterns can be combined.

		For example:

			     Incoming Jobs
			          │
			          ↓
			       Fan-Out
			          │
			   ┌──────┼──────┐
			   ↓      ↓      ↓
			  W1     W2     W3
			   │      │      │
			   └──────┼──────┘
			          ↓
			       Fan-In
			          │
			          ↓
			       Results

		Each worker can also have bounded concurrency
		to protect a database or external API.

		The workers themselves can form a pipeline:

			Read → Process → Save

		And when the application receives a shutdown signal:

			Graceful Shutdown
			        ↓
			Stop accepting jobs
			        ↓
			Finish active jobs
			        ↓
			Close resources
			        ↓
			       Exit

		This is how individual concurrency concepts become
		real backend architecture.
	*/

	// ============================================================
	// Final Summary
	// ============================================================

	fmt.Println("41. concurrency patterns lecture finished")

	/*
		Final mental model:

		Worker Pool
		    → How many workers do I want?

		Fan-Out
		    → How do I distribute work?

		Fan-In
		    → How do I collect results?

		Pipeline
		    → How do I divide work into stages?

		Bounded Concurrency
		    → How many operations can run at once?

		Graceful Shutdown
		    → How do I stop safely?

		These patterns are not replacements for each other.

		They solve different concurrency problems and are often
		combined in real-world Go backend applications.
	*/
}
