# ezpprof
ezpprof is one line configuration for profiling in either a cli or a webserver. All you need to do is embed an options struct in your struct (for a cli) or just call a single function for a server. See the examples directory in either subpackage to see how it works in more detail.


## Why?
I wanted to have as simple integration with profiling as possible. This is really a small wrapper around [github.com/pkg/profile](https://github.com/pkg/profile) that takes out the argument parsing you need to get it working.


## How do I run it?

### cli: 
```
defer ezpprof.RunProfiler(ezpprof.Opts{}).Stop()
```
The options are

```
type Opts struct {
	Mem         bool   `long:"pprof_mem"                env:"pprof_mem"   description:"Memory profile. Mutually exlusive with all the other modes." `
	CPU         bool   `long:"pprof_cpu"                env:"pprof_cpu"   description:"CPU profile. Mutually exlusive with all the other modes."    `
	Trace       bool   `long:"pprof_trace"              env:"pprof_trace" description:"Trace profile. Mutually exlusive with all the other modes."  `
	Block       bool   `long:"pprof_block"              env:"pprof_block" description:"Block profile. Mutually exlusive with all the other modes."  `
	Mutex       bool   `long:"pprof_mutex"              env:"pprof_mutex" description:"Mutex profile. Mutually exlusive with all the other modes."  `
	ProfilePath string `long:"pprof_dir"   default:"./" env:"pprof_dir"   description:"The ouptut directory where ezpprof will write the file"`
}
```

[Full Example](https://github.com/Grindlemire/ezpprof/tree/master/cli/example)

### server: 
```
ezpprof.ServeProfile(7778)
```

[Full Example](https://github.com/Grindlemire/ezpprof/tree/master/server/example)
