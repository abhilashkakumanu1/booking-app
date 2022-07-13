Source - [link](https://www.youtube.com/watch?v=yyUHQIec83I)

Why do we need another programming language?
Infrastructure has changed a lot - scalable & distributed (multi-core processors). Existing languages are not taking full advantage of it => 1 task at a time vs multiple tasks at once (multi-threading)

Go => designed to run on multiple cores. Concurrency in Go is cheap and easy

In go all the code is in packages

\_ => blank identifier

package level variables => variables shared among multiple functions

package is a collection of go files

You can't just do `go run main.go` when code is shared across files. You have to specify every file - `go run main.go helper.go`. Or you can also just specify the path to the package and it will run the `main func` - `go run .`

capitalizing the function is going to export the function - (reason why capital P in `fmt.Println`)

3 levels of variables => local, Package, Global (capital letter)

Cannot mix data types in maps. Slice also can have only one data type

struct => to save mixed data type values for an entity

type keyword => create custom type

goroutines => concurrency in Go

sleep function stops/blocks the current thread (goroutine) execution for the defined duration
go by default is single threaded and synchronous..but we can ask it to create a separate thread when a blocking code has to be executed. Once the job is done, the thread will be deleted

Just adding a `go` keyword in front of a blocking code will put it in a separate thread. A goroutine is a light thread managed by the Go routine

Problem => main thread can sometime finish executing and exit before the sendTicket thread has time to start and execute the code => By default main goroutine does not wait for other goroutines

WaitGroup => waits for the launched goroutine to finish. package `sync` provides basic synchronization functionality.
Add: Sets the number of goroutines to wait for (increases the counter by the provided number)
Wait => Blocked until the WaitGroup counter is 0
Done => Decrements the WaitGroup counter by 1. So, this is called by the goroutine to indicate that it's finished.

Goroutine
Go is using what's called a "Green Thread"
Abstraction of an actual thread
Managed by the go runtime, we are only interacting with these high level goroutines
You can run hundreds of thousands or millions of goroutines without affecting the performance
Channels (built-in) => functionality for goroutines to talk with one another

OS Threads
Managed by Kernel
Are hardware dependent
Cost of these threads are higher
Higher startup time
No easy communication between threads
