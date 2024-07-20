package retry

import (
	"fmt"
	"time"
)

// RetryFunc is a function that will be retried
// if it returns an error
type RetryFunc func() error

// Config is the configuration for the retry package
type Config struct {

	// MaxRetries is the maximum number of retries
	MaxRetries int

	// WaitTime is the time to wait between retries
	WaitTime time.Duration

	// Logger is the logger to use
	Logger func(log string)
}

// DefaultConfig is the default configuration for the retry package
var DefaultConfig = Config{
	MaxRetries: 3,
	WaitTime:   1 * time.Second,
}

// Run runs the given RetryFunc with the given Config
func Run(fn RetryFunc, cfg Config) error {
	if cfg.MaxRetries == 0 {
		cfg.MaxRetries = DefaultConfig.MaxRetries
	}
	if cfg.WaitTime == 0 {
		cfg.WaitTime = DefaultConfig.WaitTime
	}
	for {
		err := fn()
		if err == nil {
			break
		}
		cfg.MaxRetries--
		if cfg.MaxRetries == 0 {
			return err
		}
		if cfg.Logger != nil {
			cfg.Logger(fmt.Sprintf("retrying after error: %v", err))
		}
		time.Sleep(cfg.WaitTime)
	}
	return nil
}
