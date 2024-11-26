//go:build dev

package util

import (
	"flag"
	"github.com/btcsuite/btclog/v2"
	"github.com/lightningnetwork/lnd/build"
	"os"
)

var level = flag.String("level", "debug", "log level")

func SetupLogger(dir string) (btclog.Logger, error) {
	_ = os.Remove(dir)
	flag.Parse()

	logWriter := build.NewRotatingLogWriter()
	logCfg := build.DefaultLogConfig()
	logCfg.Console.Style = true
	logCfg.Console.LoggerConfig.NoTimestamps = true
	logCfg.File.LoggerConfig.NoTimestamps = true

	mgr := build.NewSubLoggerManager(
		build.NewDefaultLogHandlers(logCfg, logWriter)...,
	)

	err := logWriter.InitLogRotator(logCfg.File, dir)
	if err != nil {
		return nil, err
	}

	logger := mgr.GenSubLogger("DEMO", func() {})

	err = build.ParseAndSetDebugLevels(*level, mgr)
	if err != nil {
		return nil, err
	}

	return logger, nil
}
