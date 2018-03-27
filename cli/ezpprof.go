package ezpprof

import (
	"github.com/pkg/profile"
)

// ProfilingOpts options for which profiler you want to run.
// Note only one of these options should ever be set (aside from the profile path)
type ProfilingOpts struct {
	Mem         bool   `long:"mem" description:"memory profile"`
	CPU         bool   `long:"cpu" description:"cpu profile"`
	Trace       bool   `long:"trace" description:"trace profile"`
	Block       bool   `long:"block" description:"block profile"`
	Mutex       bool   `long:"mutex" description:"mutex profile"`
	ProfilePath string `long:"pprof-output" default:"./" description:"path to output the profile to"`
}

type emptyProfile struct{}

func (e emptyProfile) Stop() {}

// RunProfiler starts a specific profiler based on the options passed to it.
// It will return a Stop function that can be deferred
func RunProfiler(opts ProfilingOpts) interface{ Stop() } {
	p := getProfiler(opts)
	if p == nil {
		return emptyProfile{}
	}

	return profile.Start(p, profile.ProfilePath(opts.ProfilePath), profile.Quiet)
}

func getProfiler(opts ProfilingOpts) func(p *profile.Profile) {
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
