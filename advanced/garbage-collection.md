
# Source code of Go's garbage collection here:
https://github.com/golang/go/blob/master/src/runtime/mgc.go

# Basics:
- stack allocations are cleaned by themselves eg, function calls and its local vars (no need for GC)
