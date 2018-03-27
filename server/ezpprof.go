package ezpprof

import (
	"fmt"
	"net/http"

	// This is how you import the pprof for serving over http requests
	_ "net/http/pprof"
)

// ServeProfile serves profiling information on the port you pass.
// It will launch a server in another go routine that will serve profiling information through the route /debug/pprof.
// An example of how you can use it: `wget http://localhost:8081/debug/pprof/trace?seconds=10`
// See https://golang.org/pkg/net/http/pprof/ for more information on how to use it.
func ServeProfile(port int) {
	go func() {
		pprofMux := http.DefaultServeMux
		http.DefaultServeMux = http.NewServeMux()
		pprofSrv := &http.Server{
			Addr:    fmt.Sprintf("localhost:%d", port),
			Handler: pprofMux,
		}
		pprofSrv.ListenAndServe()
	}()
}
