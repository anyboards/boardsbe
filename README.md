# Boards Backend

Real motherfucking backend

## Development Tools

Before start development install following tools

### Taskfile

Taskfile used instead of Makefile. Download here https://taskfile.dev

### Go

I recommend to install Go via Go Version Manager so you can easily upgrade Go version if needed

Install Go Version Manager (code taken from https://github.com/moovweb/gvm#installing)

```bash
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

Restart terminal

Install latest Go version. Versions can be taken from https://go.dev/dl. Latest version at the moment of writing this readme is `go1.21.5`

```bash
gvm install go1.21.5
```

You can see all the installed Go version with 

```bash
gvm list

### Output

gvm gos (installed)

=> go1.21
   go1.21.5
   system
```

Then select installed Go version

```bash
gvm use go1.21.5
```

Validate Go installed version using

```bash
go version

### Output

go version go1.21.5 darwin/arm64
```

## Development

Development process described for VS Code

### VS Code

Install Go plugin https://marketplace.visualstudio.com/items?itemName=golang.Go

#### `Dev` launch configuration

This launch configuration starts development environment and debug session

#### `App` launch configuration

Development environment can be run by hand
```bash
task -t Taskfile.dev.yaml env
```
In this case use `App` launch configuration. Development containers should be also stopped by hand
```bash
task -t Taskfile.dev.yaml stopenv
```

### `grpcui`

To run web ui for gRPC run 
```bash
task -t Taskfile.dev.yaml grpcui
```

