package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/grindlemire/ezpprof/server"
)

func main() {
	// setup our profiler server in one line. Just pass in the port you want to run on.
	// Note I only run it on localhost
	ezpprof.ServeProfile(7778)

	// Simpler router for running fibonacci calculations
	r := mux.NewRouter()
	r.Methods("GET").Path("/fib").HandlerFunc(FibHandler)
	r.Methods("GET").Path("/fastfib").HandlerFunc(FastFibHandler)
	err := http.ListenAndServe("localhost:7777", r)
	if err != nil {
		fmt.Printf("Error creating server: %v\n", err)
		return
	}
}

// FibHandler runs the fibonacci sequence on a number
func FibHandler(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	nstr := vals.Get("n")

	n, err := strconv.ParseInt(nstr, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Error converting argument to integer: %v\n", err)
	}
	output := fib(n)

	fmt.Fprintf(w, "Fibonacci output: %d\n", output)
	return
}

// FastFibHandler runs the fibonacci sequence on a number using memoization
func FastFibHandler(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	nstr := vals.Get("n")

	n, err := strconv.ParseInt(nstr, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Error converting argument to integer: %v\n", err)
	}
	output := fastfib(n)

	fmt.Fprintf(w, "Fast Fibonacci output: %d\n", output)
	return
}

func fib(n int64) int64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

var cache = map[int64]int64{}

func fastfib(n int64) int64 {
	cachedVal, found := cache[n]
	if found {
		return cachedVal
	}

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	result := fastfib(n-1) + fastfib(n-2)
	cache[n] = result
	return result
}
