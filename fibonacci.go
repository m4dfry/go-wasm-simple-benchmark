package main

import (
  "fmt"
  "math/big"
  "syscall/js"
)

var done = make(chan struct{})

func main() {
  callback := js.NewCallback(start)
  setPrintMessage := js.Global().Get("setPrintResult")
  setPrintMessage.Invoke(callback)
  <-done
}

func start(args []js.Value) {
  num := args[0].Int()
  f := fibonacci()
  for i := 1; i < num; i++ {
    f()
  }
  fmt.Println(f())
  done <- struct{}{} // Notify printMessage has been called
}

func fibonacci() func() *big.Int {
        // Initialize two big ints with the first two numbers in the sequence.
	a := big.NewInt(0)
	b := big.NewInt(1)

        return func() *big.Int {
                defer func() {
			// Compute the next Fibonacci number, storing it in a.
			a.Add(a, b)
			// Swap a and b so that b is the next number in the sequence.
			a, b = b, a
                }()
                return a
        }
}
