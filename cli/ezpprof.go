package ezpprof

import (
	"github.com/pkg/profile"
)

// Opts are options for which profiler you want to run.
// Note only one of these options should ever be set (aside from the profile path)
type Opts struct {
	Mem         bool   `long:"pprof_mem"                env:"pprof_mem"   description:"Memory profile. Mutually exlusive with all the other modes." `
	CPU         bool   `long:"pprof_cpu"                env:"pprof_cpu"   description:"CPU profile. Mutually exlusive with all the other modes."    `
	Trace       bool   `long:"pprof_trace"              env:"pprof_trace" description:"Trace profile. Mutually exlusive with all the other modes."  `
	Block       bool   `long:"pprof_block"              env:"pprof_block" description:"Block profile. Mutually exlusive with all the other modes."  `
	Mutex       bool   `long:"pprof_mutex"              env:"pprof_mutex" description:"Mutex profile. Mutually exlusive with all the other modes."  `
	ProfilePath string `long:"pprof_dir"   default:"./" env:"pprof_dir"   description:"The ouptut directory where ezpprof will write the file"`
}

type emptyProfile struct{}

func (e emptyProfile) Stop() {}

// RunProfiler starts a specific profiler based on the options passed to it.
// It will return a Stop function that can be deferred
func RunProfiler(opts Opts) interface{ Stop() } {
	p := getProfiler(opts)
	if p == nil {
		return emptyProfile{}
	}

	return profile.Start(p, profile.ProfilePath(opts.ProfilePath), profile.Quiet)
}

func getProfiler(opts Opts) func(p *profile.Profile) {
	if opts.Mutex {
		return profile.MutexProfile
	}

	if opts.Block {
		return profile.BlockProfile
	}

	if opts.Mem {
		return profile.MemProfile
	}

	if opts.CPU {
		return profile.CPUProfile
	}

	if opts.Trace {
		return profile.TraceProfile
	}

	return nil
}
