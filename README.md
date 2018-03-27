# ezpprof
ezpprof is one line configuration for profiling in either a cli or a webserver. All you need to do is embed an options struct in your struct (for a cli) or just call a single function for a server. See the examples directory in either subpackage to see how it works in more detail.


## Why?
I wanted to have as simple integration with profiling as possible. This is really a small wrapper around github.com/pkg/profile that takes out the argument parsing you need to get it working.


## How do I run it?
If you want to profile a cli use the [cli](https://github.com/Grindlemire/ezpprof/tree/master/cli) package (package name is still `ezpprof`). [example](https://github.com/Grindlemire/ezpprof/tree/master/cli/example)

If you want to profile a server use the [server](https://github.com/Grindlemire/ezpprof/tree/master/server) package (package name is still `ezpprof`). [Example](https://github.com/Grindlemire/ezpprof/tree/master/server/example)