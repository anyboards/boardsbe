# Boards Backend

## Install Go

I recommend to install Go via `Go Version Manager` so you can easily upgrade Go version if needed

Install `Go Version Manager` (code taken from https://github.com/moovweb/gvm#installing)

```bash
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

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

Validate Go installed version in use

```bash
go version

### Output

go version go1.21.5 darwin/arm64
```

