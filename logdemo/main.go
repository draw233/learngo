package main

import (
	"log/slog"
	"logdemo/config"
	"sync"

	"gopkg.in/lumberjack.v3"
)

var (
	cfg    = config.NewDefaultConfig()
	logger *slog.Logger
	wg     = sync.WaitGroup{}
)

func init() {
	fileHandler, err := lumberjack.New(
		lumberjack.WithBufferSize(cfg.BufferSize),
		lumberjack.WithFileName(cfg.Filename),
		lumberjack.WithMaxBytes(cfg.MaxSize),
		lumberjack.WithMaxBackups(cfg.BackUp),
		lumberjack.WithCompress(),
	)

	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
	logger = slog.New(slog.NewJSONHandler(fileHandler, nil))
	// slog.SetDefault(logger)
}

func main() {

	slog.Info("i", slog.Int("action", 9999))

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				logger.Info("gorutine", "id", i, "value", j)
			}
		}()
	}
	wg.Wait()

}
