# TinyGo to WASM Demo

Reference:

- [TinyGo WASM documentation](https://tinygo.org/docs/guides/webassembly/)
- ObservableHQ notebook [WASM from TinyGo - demo](https://observablehq.com/@oscar6echo/wasm-from-tinygo-demo)

## Build

```bash
# compile
tinygo build -o ./html/tinygo_demo.wasm -target wasm ./main/main.go
tinygo build -o ./html/tinygo_demo.wasm -target wasm -no-debug ./main/main.go

# boiler plate
cp $(tinygo env TINYGOROOT)/targets/wasm_exec.js ./html
mv ./html/wasm_exec.js ./html/wasm_exec_mod.js

# edit wasm_exec_mod.js
# tiny adjustment: change IIFE to module
# i.e. :
# change (() => {FUNC_BODY})() to FUNC_BODY
# change global.Go = class{...} to export const Go = class{...}
```

In order to export a Go function to WASM add comment above function in `main.go`:

```go
//export functionName
func functionName(){
    // blah
}
```

NOTE: Only int and float types are supported.

TinyGo .wasm generated files are small (even smaller with `-no-debug` flag):

```bash
❯ ll html
total 60K
-rw-rw-r-- 1 olivier olivier 1.1K Jul  4 00:24 index.html
-rwxrwxr-x 1 olivier olivier  14K Jul  4 00:24 tinygo_demo.wasm
-rw-r--r-- 1 olivier olivier  17K Jul  3 21:54 wasm_exec_mod.js
```

## Browser

Serve folder `/html`:

```bash
go run serve/serve.go
```

Open <http://localhost:8080>
