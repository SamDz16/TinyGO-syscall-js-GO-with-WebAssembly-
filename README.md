# TinyGo & WebAssembly with syscall/js
This little applicaton demonstrates how we can integrate Golang programming language to work along with WebAssembly. Two example were made, the first one adds two numbers from user input, the second one capitalize a user input string. 

## Run
The bellowed command is executed in the root folder in order to compile the `main.go` to get the wasm equivalent `main.wasm` in the `assets` folder.
Note that tinygo is used to get ***tiny*** wasm file size. In my case it tooks only **181KB** for the `main.wasm` file size wherease if we use the built in go way it generates wasm file size with almost **2MB** of size.

```shell
tinygo build -o ./assets/main.wasm -target wasm ./main.go
```

## Launch the server
In order to launch the server tto serve the assets folder (`index.html`), run the bellow command in the `./server` folder:
```shell
go run server.go
```
