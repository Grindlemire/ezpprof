# CLI Example

Run `go build` and run the cli with `./example`

I am using github.com/jessevdk/go-flags for easy flag parsing. This library is intended to be used and parsed by `go-flags` (it is a great library).

If you pass any of the following flags it will output a profile for that profiler:
- `--mem`
- `--cpu`
- `--block`
- `--mutext`
- `--trace`

You can then examine the output file with `go tool pprof FILENAME` (if you are using the trace it is through `go tool trace FILENAME`)

Example: 
```
./example -n 43 --trace
``` 
will output a trace file you can then run using


```
go tool trace trace.out
```

