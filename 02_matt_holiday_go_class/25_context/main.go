package main

import (
	"context"
	"log"
)

// Cancellation and timeouts
// 	The Context package offers common method to cancel requests
// 		Explicit cancellation
// 		Implicit cancellation
// 	A context may also carry request-specific values, such as a
// 	trace ID
// 	Many networks or database requests, for example, take a context
// 	for cancellation

// 	A context offers two controls:
// 		A channel that closes when cancellation occurs
// 		An error that's readable once the channel closes
// 	The error value tells you whether the request was cancelled
// 	or timed out
// 	We often use the channel from Done() in a select block

// 	Contexts form an immutable tree structure
//  (goroutine-safe; changes to a context do not affect its ancestors)
//  Cancellation/timeout applies to the current context and its subtree
// 	Same for a value
//  A subtree may be created with a shorter timeout (but not longer)
//  It's a tree of immutable nodes which can be extended
// ctx := context.Background()
// ctx = context.WithValue(ctx, "v", 7)
// ctx, canc := context.WithTimeout(ctx, t)
// req, _ := http.NewRequest(method, url, nil)
// req = req.WithContext(ctx)
// resp, err := http.DefaultClient(req)
// The Context value should always be the first param

// Values
// 	Context values should be data specific to a request, such as:
// 		A trace ID or start time (for latency calculation)
// 		Security or authorization data
// 	AVOID using the context to carry "optional" params
// 	Use a package-specific, private context key type (not str)
// 	to avoid collisions

type contextKey int

const TraceKey contextKey = 1

// ContextLog makes a log with the trace ID as a prefix
func ContextLog(ctx context.Context, f string, args ...interface{}) {
	// reflection -- tbd
	traceID, ok := ctx.Value(TraceKey).(string)

	if ok && traceID != "" {
		f = traceID + ": " + f
	}
	log.Printf(f, args...)
}

func main() {

}
