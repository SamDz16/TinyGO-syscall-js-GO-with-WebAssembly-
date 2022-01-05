# TinyGo & WebAssembly with syscall/js
This little applicaton demonstrates how we can integrate Golang programming language to work along with WebAssembly. Two example were made, the first one adds two numbers from user input, the second one capitalize a user input string. 

## run
The bellowed command is executed in the root folder in order to compile the `main.go` to get the wasm equivalent `main.wasm` in the `assets` folder.
Note that tinygo is used to get ***tiny*** wasm file size. In my case it tooks only 181KB wherease if we use the built in go way it generates wasm file size almost 2MB.

```shell
tinygo build -o ./assets/main.wasm -target wasm ./main.go
```

## launch the server
In order to launch the server tto serve the assets folder (`index.html`), run the bellow command in the `./server` folder:
```shell
go run server.go
```
