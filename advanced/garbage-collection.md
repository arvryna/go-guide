

# Basics:
- stack allocations are cleaned by themselves eg, function calls and its local vars (no need for GC)
- objects that are crated with new and structs that were created with `&mystruct{}` reside in heap

# Escape analysis:
- Compiler use this analysis to track lifetime of allocation/objects to check and see where they can possibly go. 
  - Eg, an allocation can be passed as param, return somewhere, or can be stored locally
  - escape is analogous to the familiar word "passing" as in passed as func 
  - By doing this analysis compiler can get information to optimize
- Compiler Optimizations:
  - Escape analysis provides valuable info for compiler to make decision on weather to keep a variable in stack or heap
  - moving allocations to stack from heap will greately reduce the load on Garbage collection
- usage: `go run -gcflags -m main.go`
- https://github.com/golang/go/wiki/CompilerOptimizations#escape-analysis

# References:
- Source code of Go's garbage collection here:
  https://github.com/golang/go/blob/master/src/runtime/mgc.go
