# Guide for beginners in golang

- Guide is organized with explanations and short snippets

# Go Packages and its rules
- naming of package: should not be camel/snakecase, same as the directory name
- each directory can only one package, but many files (each bearing the same package name)
- sources are organized in their packages
  - eg: jsman - json manipulator is residing in src/jsman as a seperate package
- use small letters for private methods and vice versa (same rules applies to vars inside struct)
- files residing in sub-folder has access to each other (only the public ones)

# JSON management
- JSON serialization and deserialization in go is analogously marshalling and unmarshalling
- NOTE:  "Unmarshal will only set exported fields of the struct.", first letter must be capital
- while marshalling you can provide rules on how mashallelling needs to be done via, backsticks eg, `json:fieldName,fieldtype` 
- marshalling and unmarshalling code [here](https://github.com/arvryna/go-guide/blob/main/src/jsman/jsman.go)

# Context:
``` [client] -> | [middlewere] -> [app layer] -> [DB Layer] | ```

- Its like a kvstore, that one can propogate information across the layers of your app eg, req-id
- It is possible to set validity or timeout for a context, after timeout exceeded, context cease to exist
  and this context can be attached to a request and thereby control its flow, trace, debug etc.,
- invalidation of a context can be learnt by either checking its Err() func or listening for Done() channel

# Printing strings:
- ``` fmt.printf("%#v", struct) // print any struct ```

# Code style:

- [Uber go style guide](https://github.com/uber-go/guide/blob/master/style.md)
- [Google go's guide](https://github.com/golang/go/wiki/CodeReviewComments)
- [Effective go](https://go.dev/doc/effective_go)
- [Kubernets go coding convention](https://www.kubernetes.dev/docs/guide/coding-convention/)
