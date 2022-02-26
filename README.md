# Guide for beginners in golang
- Guide is organized with explanations and short snippets
- NOTE: If mistakes found, please make a PR or add comments so that i can fix them.

# Go Packages and its rules
- Naming of package: should not be camel/snakecase, same as the directory name
- Each directory can only one package, but many files (each bearing the same package name)
- Sources are organized in their packages
  - eg: jsman - json manipulator is residing in src/jsman as a seperate package
- Use small letters for private methods and vice versa (same rules applies to vars inside struct)
- Files residing in sub-folder has access to each other (only the public ones)

# JSON management
- JSON serialization and deserialization in go is analogously marshalling and unmarshalling
- NOTE:  "Unmarshal will only set exported fields of the struct.", first letter must be capital
- while marshalling you can provide rules on how marshalling needs to be done via, backsticks eg, `json:fieldName,fieldtype` 
- marshalling and unmarshalling code [here](https://github.com/arvryna/go-guide/blob/main/internal/jsman/jsman.go)

# Concurrency primitives
- **Waitgroups:** To wait for multiple goroutines to finish, we can use a wait group code [here](https://github.com/arvryna/go-guide/blob/main/internal/concur/sync.go)

- **Locks:** Mutex Lock = Mutually exclusive lock:
  - To avoid race among go-routines, ensure atomicity wherever required
  - Available in sync package, when used, the critical section of the code is allowed for only one go-routine
  - Either reader / writer will wait until lock is freed, they wait in a queue will be served FIFO
- RWMutex Lock = This lock allows access to critical section of code only one writer but any number of readers

# Go closure
- functions are first class in go. It can be stored in a variable
- A closure is a function that captures the state of the surrounding environment.

```
Example:

func(s string){fmt.print(s)}("Hello")
```

NOTE: capturing the state means that, the closure will somehow associate that specific closure with the variables passed to it. In the above example, this specific instance of a function is linked with the variable passed to it. It becomes particularly important when enclosing closures in loops. So each instance of closure will be given a unique reference to variable s.

# Pointers
- Map by default is reference in go
- Passing pointer by reference example

```
type fData struct {
  l sync.RWMutex
  data map[string]bool
}

func fetch(s *fData){
  // 
}

data := fData{}
fetch(&data)

```

# Context:
``` [client] -> | [middlewere] -> [app layer] -> [DB Layer] | ```

- Its like a kvstore, that one can propogate information across the layers of your app eg, req-id
- It is possible to set validity or timeout for a context, after timeout exceeded, context cease to exist
  and this context can be attached to a request and thereby control its flow, trace, debug etc.,
- invalidation of a context can be learnt by either checking its Err() func or listening for Done() channel

# Debugging:

> cannot use wg (type sync.WaitGroup) as type *sync.WaitGroup in argument to increment

- Problem: Passing value to a function that expects reference
- Fix: Use &wg to pass reference

> Deadlock

- Make sure concurrency primitives are used properly in code, make sure you call Unlock(), wg.Wait(), wg.Close().
  appropriate book-keeping is required to ensure we call the right constructs at the right time. usage of defer is encouraged. 

> Partial go-routine failures
- If all goroutines are deadlocked, go raises panic to kill program, but if there is partial failure of go-routines, the program 
  continues to function but may behave uncertain.

# Go Tools:
> Detecting race conditions ?
- Use race flag to detect race conditions in the program
``` go run -race main.go ```

How race flag works under the hood ? 
- In runtime, go allocates a small amount of memory (called as shadow memory) to each variables created by the program and checks the reads and writes by interwining code and there by identify the race. ref:[video](https://youtu.be/5erqWdlhQLA)

# Handling signals (SIGINT, SIGTERM) for graceful shutdown
You can add these lines at the end of the main function, this helps,
1) Wait for all your go routines to complete
2) To do last minute clean-up for the process, eg: to close DB connnections, flush data to disk etc.,

```
sigListen := make(chan os.Signal, 1)
signal.Notify(sigListen, syscall.SIGINT, syscall.SIGTERM)
<-sigListen
// Add code for graceful shutdown here
```

# Working with docker
An example of a simple hw app with a dockerfile can be found here `samples/greeter`

# Printing strings:
- ``` fmt.printf("%#v", struct) // print any struct ```

# Code style:

- [Uber go style guide](https://github.com/uber-go/guide/blob/master/style.md)
- [Google go's guide](https://github.com/golang/go/wiki/CodeReviewComments)
- [Effective go](https://go.dev/doc/effective_go)
- [Kubernets go coding convention](https://www.kubernetes.dev/docs/guide/coding-convention/)

# Best practices:
- Use defer whenever possible