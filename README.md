# go-wasm-simple-benchmark
This is a simple benchmark for the WASM new feature introduced in Go 1.11

# Why ??

Read [on my blog](https://blog.m4dfry.space/posts/go-wasm-benchmark/) a brief explanation.

# How to compile

If you already have Go installed, you can install 1.11 beta2 version by running:

	go get golang.org/dl/go1.11beta2

Then, use can use the go1.11beta2 command. Remember to add your go/bin folder to path:

	PATH="$PATH:$HOME/go/bin/"

Now, if you try to build with go1.11beta1 you got an error of missing sdk that you can fix with:

	go1.11beta2 download

This command download a new folder ~/sdk/go1.11beta2/ containing all you need to build.

Es.
	GOARCH=wasm GOOS=js go1.11beta2  build -o fibonacci.wasm fibonacci.go

If everything went as planned the compilation will end successfully.

# How to run

At this point it's necessary to start a small web server to serve our files.
You can use whatever you prefer as long as you can configure the download of .wasm file with its correct MIME.

*server_improved* is simple server that implement the correct file type and the gzip handler

Start with

	go run server_improved.go

Now you can connect to [localhost:3000/](localhost:3000/) and try yourself a simple WASM Fibonacci Benchmark.

# Makefile
Compile
> make wasm

Run server
> make run

## License

Used software (all software is licensed under appropriate license of authors)
*  Go Programming Language [https://golang.org/](https://golang.org/) BSD 3-clause "New" or "Revised" License

---

Have a request, suggestion, errata, critic or question? Open an Issue!

---
