//go:build dev

package main

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/btcsuite/btclog/v2"
	"github.com/ellemouton/slog/util"
)

const dir = "/Users/elle/projects/slog/structured/log.log"

var log btclog.Logger

func main() {
	var err error
	log, err = util.SetupLogger(dir)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			userDoesThings(ctx, i)
		}()
	}

	wg.Wait()
}

func userDoesThings(ctx context.Context, userID int) {
	// At the API boarder, we add any request scoped values such as the
	// user ID to a derived child context.
	ctx = btclog.WithCtx(ctx, slog.Int("user_id", userID))

	// Basic log with no directives.
	log.InfoS(ctx, "This is a test log")

	// Log with more info.
	// 	1) Static search string. Easy to find both:
	// 		- Where in the _code_ this log line is.
	// 		- Where in the log file this log line is.
	// 		- Easy to count number of occurrences.
	// 	2) Includes user_id in the log line so:
	// 		- Easy to grep for all logs pertaining to a specific user.
	// 		- Don't need to remember to specify each time.
	// 		- Consistent format of the user_id= directive.
	log.InfoS(ctx, "Request made to GET /api/v1/users")

	// Log with amount.
	// 	 1) How much BTC did user X spend:
	// 		AGG "amount" WHERE "msg"="Channel open performed" AND "user_id" = X
	//
	// 	 2) How many times did this log line get called?
	// 		COUNT where "msg"="Channel open performed"
	//
	// 	 3) Total amount spent by all users?
	// 	        AGG "amount" WHERE msg="Channel open performed"
	log.InfoS(ctx, "Channel open performed",
		btclog.Fmt("amount", "%.8f", 0.00154))

	// Log with logs of directives:
	// 	1) Easy to write: no need to go count directives.
	// 	2) Easy to read.
	//
	// Proposed Style/suggestions:
	//   - Try to use helpers that return an `slog.Attr` to avoid BADKEY bugs (demo).
	//   - Unless all the attributes fit on the same line,
	//      go for: 1 attribute (kv pair) per line. I suggest we ignore log
	// 	lines for the `lll` linter.
	log.DebugS(ctx, "More interesting info",
		"key", "value",
		"key_2", 4,
		slog.String("key_3", "value3"),
		slog.String("key_4", "quoted string"),
		btclog.Hex("key_6", []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}),
		btclog.Hex6("key_7", []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}),
		btclog.Fmt("key_8", "%.12f", 3.241),
		"key_9", btclog.Sprintf("%d", 6),
	)

	// Slog helpers.
	log.DebugS(ctx, "slog helpers",
		slog.String("key_3", "value3"),
		slog.String("key_4", "quoted string"),
		slog.Int("key_5", 5))

	// Hex helpers.
	log.DebugS(ctx, "Hex helpers",
		btclog.Hex("key_6", []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}),
		btclog.Hex6("key_7", []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}))

	// Sprintf helpers.
	log.DebugS(ctx, "Sprintf helpers",
		btclog.Fmt("key_8", "%.12f", 3.241),
		"key_9", btclog.Sprintf("%d", 6))

	// Closure helpers
	log.DebugS(ctx, "Closure",
		btclog.ClosureAttr("key_10", func() string {
			return fmt.Sprintf("expensive string computation")
		}))

	// Closure helpers: skip expensive computation: show what happens if it does.
	log.TraceS(ctx, "Closure",
		btclog.ClosureAttr("key_10", func() string {
			panic("THIS SHOULD NOT GET COMPUTED")
		}))

	// Demonstrate error logs.
	// 	1) Discussion: structured errors .... dun dun duuuuun. See point 2 below.
	err := fmt.Errorf("oh no bad")
	log.ErrorS(ctx, "Static message here", err, "extra", "info")
}

/*
Discussion:

1) arguably every gRPC server should have an interceptor that adds a
    request ID to the incoming context.


	var (
		reqID   uint64
		reqIDMu sync.Mutex
	)

	func addRequestID(ctx context.Context) context.Context {
		var id uint64

		reqIDMu.Lock()
		id = reqID
		reqID++
		reqIDMu.Unlock()

		return btclog.WithCtx(ctx, slog.Uint64("request_id", id))
	}

     Then, in the grpc interceptor:

	ctx := addRequestID(reqctx)

   This at least works for synchronous code. For async code, we need to add the
   request ID to the context before the async operation is started.

2) Structured errors.  Same idea but applied to errors. That way we have context
   carrying structured values _down the call stack_ and structured errors
   carrying structured values _up the call stack. Will make error strings just
   as searchable/queryable.

*/
