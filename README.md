# slog-demo

```
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
```