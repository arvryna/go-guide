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

# Printing strings:
- ``` fmt.printf("%#v", struct) // print any struct ```
