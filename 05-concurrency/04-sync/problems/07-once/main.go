/*
============================================================
Problem 7 — sync.Once
============================================================

Create a sync.Once.

Create a function that prints:

    "Initializing application"

Call once.Do() THREE times.

The initialization message must be printed ONLY ONCE.

Requirements:
- Import sync.
- Create a sync.Once.
- Call once.Do() exactly THREE times.
- Each call must provide the initialization function.
- The initialization function must print:
      "Initializing application"
- The message must appear exactly ONE time.
- Do NOT use if statements to control execution.
- Do NOT use channels.
- Do NOT use Mutex directly.
- Do NOT use time.Sleep().
- Do NOT use select.

Expected output:

    Initializing application

Important:

sync.Once guarantees that a function passed to:

    once.Do(func() {
        // initialization
    })

is executed only once.

Even if you call:

    once.Do(...)
    once.Do(...)
    once.Do(...)

the function runs only during the FIRST successful call.

Mental model:

    once.Do(init)
         │
         ↓
    Has init run?
       /     \
     NO       YES
     │         │
   run it    do nothing
     │
     ↓
   DONE

Backend examples:

    - Initialize configuration
    - Create a singleton resource
    - Initialize a database client
    - Lazy initialization
    - Load application configuration

Goal:
Understand that sync.Once is for:

    "This initialization must happen exactly once."
*/

package main

import (
	"fmt"
	"sync"
)

func initializationApp() {
	fmt.Println("Initializing application")
}

func main() {
	once := sync.Once{}
	once.Do(initializationApp)
	once.Do(initializationApp)
	once.Do(initializationApp)
}
