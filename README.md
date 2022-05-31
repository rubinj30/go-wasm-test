## Messing around with WebAssembly and Go

Followed [this](https://github.com/Yaoir/ClockExample-Go-WebAssembly) as a starting point and then have veered off from there.

### Get up and running 

Creates a build file (`lib.wasm`)
```
GOARCH=wasm GOOS=js go build -o lib.wasm main.go
```

That build file is served to browser when you start the go server on `localhost:8080`
```
go run server.go
```

### Future ideas
- Figure out how to compare JS vs. Go sorting of a large number of items, and display the progress in the browser with WASM manipulated elements
- Cleanup styling
- Use web worker to run JS sorting so it doesn't block main thread
