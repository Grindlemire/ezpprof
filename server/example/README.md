# Server Example

Run `go build` and run the server `./example`


If you hit the endpoint http://localhost:7777/fib?n=43 you will get the fibonacci result for 43.


If you hit any of the following `wget` you will get the profile
- http://localhost:7778/debug/pprof/profile?seconds=10 (default is 30 seconds)
- http://localhost:7778/debug/pprof/heap
- http://localhost:7778/debug/pprof/block
- http://localhost:7778/debug/pprof/mutex
- http://localhost:7778/debug/pprof/trace?seconds=10

Example: 
```
wget http://localhost:7778/debug/pprof/trace?seconds=10 -O trace.out
``` 
will output a trace file you can then run using `go tool trace`

