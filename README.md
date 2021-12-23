## Basic clock using WebAssembly and Go

Followed [this](https://github.com/Yaoir/ClockExample-Go-WebAssembly) and another tutorial. 

### Get up and running 

Creates a build file (`lib.wasm`)
```
GOARCH=wasm GOOS=js go build -o lib.wasm main.go
```

That build file is served to browser when you start the go server on `localhost:8080`
```
go run server.go
```