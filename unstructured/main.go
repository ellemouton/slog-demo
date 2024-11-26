//go:build dev

package main

import (
	"github.com/btcsuite/btclog"
	"github.com/ellemouton/slog/util"
	"sync"
)

const dir = "/Users/elle/projects/slog/unstructured/log.log"

var log btclog.Logger

func main() {
	var err error
	log, err = util.SetupLogger(dir)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			userDoesThings(i)
		}()
	}

	wg.Wait()
}

func userDoesThings(userID int) {
	// Basic log with no directives.
	log.Info("This is a test log")

	// Log with directives.
	log.Infof("User %d made request to %s", userID, "GET /api/v1/users")

	// Log with many directives
	// 	Notice:
	// 		1) Hard to _write_ as you need to which param goes with
	// 		  which directive. Easy to make mistakes.
	// 		2) Log author must remember to add valuable info such as
	// 		   the user ID.
	// 		3) When grepping through log files, it's hard to find
	// 		   all the logs pertaining to this user since the logs
	// 		   don't use a consistent format for specifying user ID
	// 		   since it is manually specified each time.
	// 		4) Similarly: if we see this log line in the log file,
	// 		   it is hard to search for where this is in the code.
	// 		5) If we simply want to count the number of times this
	// 		   log line has been logged - it is difficult cause the
	// 		   variables are changing.
	// 		6) What if we just want to search the log file for where
	// 		   this log line was called? Right now this is hard to
	//	 	   do.
	log.Infof("(user_id=%d): Req: %s, Resp: %s", userID,
		"GET /api/v1/users", "200 OK")

	// Imagine we have a log like this. Ideally we'd like to be able to do
	// something like use our logs to query:
	// 	1) How much BTC did user X spend in the last 4 hours
	// 	2) How many times did this specific log line get called?
	//	   (regardless of user)
	// 	3) What is the total amount spent by all users in the last X
	// 	   time?
	log.Debugf("User %d just spent %.8f to open a channel", userID,
		0.0154)
}
