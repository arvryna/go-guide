# Guide for beginners in golang

- Guide is organized with explanations and short snippets

# Go Packages and its rules
- naming of package: should not be camel/snakecase, same as the directory name
- each directory can only one package, but many files (each bearing the same package name)
- sources are organized in their packages
  - eg: jsman - json manipulator is residing in src/jsman as a seperate package
- use small letters for private methods and vice versa (same rules applies to vars inside struct)
- files residing in sub-folder has access to each other (only the public ones)
