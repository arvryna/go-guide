# Third party dependencies that i like in Go

# Viper
- a lib that simplifies accesssing config across different environment 
- can parse config of file type = JSON, TOML, YAML, ENV, INI
- enviroment can be read automatically
- can read config from remote systems like etcd, or consul
- it can watch for config file changes and also notify the go-process
- you can marshal/unmarshal using mapstructure
- `go get github.com/spf13/viper`

# ORM
- gorm is a popular ORM
- to connect with postgres, driver is required additionally