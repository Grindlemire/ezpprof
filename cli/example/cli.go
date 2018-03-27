package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/grindlemire/ezpprof/cli"
	"github.com/jessevdk/go-flags"
)

// Opts represents your programs other options that you need for the cli
type Opts struct {
	// Other options that you use in your program
	Number    int64 `short:"n" long:"number"    default:"10" description:"Number to calculate the fibonacci for"`
	Efficient bool  `short:"e" long:"efficient"              description:"Calculate the fibonacci sequence efficiently or not"`

	// Embed the options directly in your options struct that you are parsing so it will be picked up
	ezpprof.ProfilingOpts
}

// Setup our parser and global options
var opts Opts
var parser = flags.NewParser(&opts, flags.Default)

func main() {
	// This parses the flags using go-flags
	_, err := parser.Parse()
	if nil != err {
		if isUsage(err) || isCommand(err) {
			os.Exit(1)
		}
		log.Fatalf("Error parsing arguments: %s", err)
		os.Exit(1)
	}

	// now just call our single line configuration for profiling. Defer stop it to end the profiling at the end of your program
	defer ezpprof.RunProfiler(opts.ProfilingOpts).Stop()

	// our actual logic here
	if !opts.Efficient {
		n := fib(opts.Number)
		fmt.Printf("Output from inefficient fibonacci: %d\n", n)
		return
	}

	n := fastfib(opts.Number)
	fmt.Printf("Output from efficient fibonacci: %d\n", n)
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

func isUsage(err error) bool {
	return strings.HasPrefix(err.Error(), "Usage:")
}

func isCommand(err error) bool {
	return strings.HasPrefix(err.Error(), "Please specify")
}
