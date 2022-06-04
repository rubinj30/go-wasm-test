## Messing around with WebAssembly, Go, and Web Workers (and maybe Partytown)

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
[ ] Figure out how to compare JS vs. Go sorting of a large number of items, and display the progress in the browser with WASM manipulated element with a progress bar or something
[ ] Cleanup styling and code
[ ] Use web worker to run JS sorting so it doesn't block main thread
    [x] setup a basic worker running JS script
    [ ] setup Partytown and test out JS script running in there